package policies

import "testing"

func TestAddPolicy(t *testing.T) {
	pMgr := NewPolicyManager()

	netpol := NPMNetworkPolicy{}

	err := pMgr.AddPolicy(&netpol, nil)
	if err != nil {
		t.Errorf("AddPolicy() returned error %s", err.Error())
	}
}

func TestRemovePolicy(t *testing.T) {
	pMgr := NewPolicyManager()

	err := pMgr.RemovePolicy("test", nil)
	if err != nil {
		t.Errorf("RemovePolicy() returned error %s", err.Error())
	}
}
