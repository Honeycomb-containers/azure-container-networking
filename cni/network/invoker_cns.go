package network

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/Azure/azure-container-networking/cni"
	"github.com/Azure/azure-container-networking/cns"
	"github.com/Azure/azure-container-networking/iptables"
	"github.com/Azure/azure-container-networking/log"
	"github.com/Azure/azure-container-networking/network"
	cniSkel "github.com/containernetworking/cni/pkg/skel"
	cniTypes "github.com/containernetworking/cni/pkg/types"
	cniTypesCurr "github.com/containernetworking/cni/pkg/types/current"
)

var (
	errEmtpyHostSubnetPrefix = errors.New("empty host subnet prefix not allowed")
	errEmptyCNIArgs          = errors.New("empty CNI cmd args not allowed")
)

const (
	cnsPort = 10090
)

type CNSIPAMInvoker struct {
	podName      string
	podNamespace string
	cnsClient    cnsclient
}

type IPv4ResultInfo struct {
	podIPAddress       string
	ncSubnetPrefix     uint8
	ncPrimaryIP        string
	ncGatewayIPAddress string
	hostSubnet         string
	hostPrimaryIP      string
	hostGateway        string
}

func NewCNSInvoker(podName, namespace string, cnsClient cnsclient) *CNSIPAMInvoker {
	return &CNSIPAMInvoker{
		podName:      podName,
		podNamespace: namespace,
		cnsClient:    cnsClient,
	}
}

// Add uses the requestipconfig API in cns, and returns ipv4 and a nil ipv6 as CNS doesn't support IPv6 yet
func (invoker *CNSIPAMInvoker) Add( //nolint don't consider unnamedResult
	_ *cni.NetworkConfig,
	args *cniSkel.CmdArgs,
	hostSubnetPrefix *net.IPNet,
	options map[string]interface{}) (*cniTypesCurr.Result, *cniTypesCurr.Result, error) {
	// Parse Pod arguments.
	podInfo := cns.KubernetesPodInfo{
		PodName:      invoker.podName,
		PodNamespace: invoker.podNamespace,
	}

	log.Printf(podInfo.PodName)
	orchestratorContext, err := json.Marshal(podInfo)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to unmarshal orchestrator context during add: %w", err)
	}

	if args == nil {
		return nil, nil, errEmptyCNIArgs
	}

	ipconfig := cns.IPConfigRequest{
		OrchestratorContext: orchestratorContext,
		PodInterfaceID:      GetEndpointID(args),
		InfraContainerID:    args.ContainerID,
	}

	log.Printf("Requesting IP for pod %+v using ipconfig %+v", podInfo, ipconfig)
	response, err := invoker.cnsClient.RequestIPAddress(context.TODO(), ipconfig)
	if err != nil {
		log.Printf("Failed to get IP address from CNS with error %v, response: %v", err, response)
		return nil, nil, err
	}

	info := IPv4ResultInfo{
		podIPAddress:       response.PodIpInfo.PodIPConfig.IPAddress,
		ncSubnetPrefix:     response.PodIpInfo.NetworkContainerPrimaryIPConfig.IPSubnet.PrefixLength,
		ncPrimaryIP:        response.PodIpInfo.NetworkContainerPrimaryIPConfig.IPSubnet.IPAddress,
		ncGatewayIPAddress: response.PodIpInfo.NetworkContainerPrimaryIPConfig.GatewayIPAddress,
		hostSubnet:         response.PodIpInfo.HostPrimaryIPInfo.Subnet,
		hostPrimaryIP:      response.PodIpInfo.HostPrimaryIPInfo.PrimaryIP,
		hostGateway:        response.PodIpInfo.HostPrimaryIPInfo.Gateway,
	}

	// set the NC Primary IP in options
	options[network.SNATIPKey] = info.ncPrimaryIP

	log.Printf("[cni-invoker-cns] Received info %+v for pod %v", info, podInfo)

	ncgw := net.ParseIP(info.ncGatewayIPAddress)
	if ncgw == nil {
		return nil, nil, fmt.Errorf("Gateway address %v from response is invalid", info.ncGatewayIPAddress)
	}

	// set result ipconfigArgument from CNS Response Body
	ip, ncipnet, err := net.ParseCIDR(info.podIPAddress + "/" + fmt.Sprint(info.ncSubnetPrefix))
	if ip == nil {
		return nil, nil, fmt.Errorf("Unable to parse IP from response: %v with err %v", info.podIPAddress, err)
	}

	// construct ipnet for result
	resultIPnet := net.IPNet{
		IP:   ip,
		Mask: ncipnet.Mask,
	}

	result := &cniTypesCurr.Result{
		IPs: []*cniTypesCurr.IPConfig{
			{
				Version: "4",
				Address: resultIPnet,
				Gateway: ncgw,
			},
		},
		Routes: []*cniTypes.Route{
			{
				Dst: network.Ipv4DefaultRouteDstPrefix,
				GW:  ncgw,
			},
		},
	}

	// set subnet prefix for host vm
	err = setHostOptions(hostSubnetPrefix, ncipnet, options, &info)
	if err != nil {
		return nil, nil, err
	}

	// first result is ipv4, second is ipv6, SWIFT doesn't currently support IPv6
	return result, nil, nil
}

func setHostOptions(hostSubnetPrefix, ncSubnetPrefix *net.IPNet, options map[string]interface{}, info *IPv4ResultInfo) error {
	// get the name of the primary IP address
	_, hostIPNet, err := net.ParseCIDR(info.hostSubnet)
	if err != nil {
		return err
	}

	if hostSubnetPrefix == nil {
		return errEmtpyHostSubnetPrefix
	}

	*hostSubnetPrefix = *hostIPNet

	// get the host ip
	hostIP := net.ParseIP(info.hostPrimaryIP)
	if hostIP == nil {
		return fmt.Errorf("Host IP address %v from response is invalid", info.hostPrimaryIP)
	}

	// get host gateway
	hostGateway := net.ParseIP(info.hostGateway)
	if hostGateway == nil {
		return fmt.Errorf("Host Gateway %v from response is invalid", info.hostGateway)
	}

	// this route is needed when the vm on subnet A needs to send traffic to a pod in subnet B on a different vm
	options[network.RoutesKey] = []network.RouteInfo{
		{
			Dst: *ncSubnetPrefix,
			Gw:  hostGateway,
		},
	}

	azureDNSMatch := fmt.Sprintf(" -m addrtype ! --dst-type local -s %s -d %s -p %s --dport %d", ncSubnetPrefix.String(), iptables.AzureDNS, iptables.UDP, iptables.DNSPort)
	azureIMDSMatch := fmt.Sprintf(" -m addrtype ! --dst-type local -s %s -d %s -p %s --dport %d", ncSubnetPrefix.String(), iptables.AzureIMDS, iptables.TCP, iptables.HTTPPort)

	snatPrimaryIPJump := fmt.Sprintf("%s --to %s", iptables.Snat, info.ncPrimaryIP)
	// we need to snat IMDS traffic to node IP, this sets up snat '--to'
	snatHostIPJump := fmt.Sprintf("%s --to %s", iptables.Snat, info.hostPrimaryIP)
	options[network.IPTablesKey] = []iptables.IPTableEntry{
		iptables.GetCreateChainCmd(iptables.V4, iptables.Nat, iptables.Swift),
		iptables.GetAppendIptableRuleCmd(iptables.V4, iptables.Nat, iptables.Postrouting, "", iptables.Swift),
		// add a snat rule to primary NC IP for DNS
		iptables.GetInsertIptableRuleCmd(iptables.V4, iptables.Nat, iptables.Swift, azureDNSMatch, snatPrimaryIPJump),
		// add a snat rule to node IP for IMDS http traffic
		iptables.GetInsertIptableRuleCmd(iptables.V4, iptables.Nat, iptables.Swift, azureIMDSMatch, snatHostIPJump),
	}

	return nil
}

// Delete calls into the releaseipconfiguration API in CNS
func (invoker *CNSIPAMInvoker) Delete(address *net.IPNet, _ *cni.NetworkConfig, args *cniSkel.CmdArgs, _ map[string]interface{}) error {
	// Parse Pod arguments.
	podInfo := cns.KubernetesPodInfo{
		PodName:      invoker.podName,
		PodNamespace: invoker.podNamespace,
	}

	orchestratorContext, err := json.Marshal(podInfo)
	if err != nil {
		return err
	}

	if args == nil {
		return errEmptyCNIArgs
	}

	req := cns.IPConfigRequest{
		OrchestratorContext: orchestratorContext,
		PodInterfaceID:      GetEndpointID(args),
		InfraContainerID:    args.ContainerID,
	}

	if address != nil {
		req.DesiredIPAddress = address.IP.String()
	} else {
		log.Printf("CNS invoker called with empty IP address")
	}

	if err := invoker.cnsClient.ReleaseIPAddress(context.TODO(), req); err != nil {
		return fmt.Errorf("failed to release IP %v with err %w", address, err)
	}

	return nil
}
