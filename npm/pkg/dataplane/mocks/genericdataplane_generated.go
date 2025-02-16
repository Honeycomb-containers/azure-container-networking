// autogenerated
//

// Code generated by MockGen. DO NOT EDIT.
// Source: /mnt/c/Users/vamsi/Desktop/Microsoft_ws/azure-container-networking/npm/pkg/dataplane/types.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dataplane "github.com/Azure/azure-container-networking/npm/pkg/dataplane"
	ipsets "github.com/Azure/azure-container-networking/npm/pkg/dataplane/ipsets"
	policies "github.com/Azure/azure-container-networking/npm/pkg/dataplane/policies"
	gomock "github.com/golang/mock/gomock"
)

// MockGenericDataplane is a mock of GenericDataplane interface.
type MockGenericDataplane struct {
	ctrl     *gomock.Controller
	recorder *MockGenericDataplaneMockRecorder
}

// MockGenericDataplaneMockRecorder is the mock recorder for MockGenericDataplane.
type MockGenericDataplaneMockRecorder struct {
	mock *MockGenericDataplane
}

// NewMockGenericDataplane creates a new mock instance.
func NewMockGenericDataplane(ctrl *gomock.Controller) *MockGenericDataplane {
	mock := &MockGenericDataplane{ctrl: ctrl}
	mock.recorder = &MockGenericDataplaneMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenericDataplane) EXPECT() *MockGenericDataplaneMockRecorder {
	return m.recorder
}

// AddPolicy mocks base method.
func (m *MockGenericDataplane) AddPolicy(policies *policies.NPMNetworkPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPolicy", policies)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPolicy indicates an expected call of AddPolicy.
func (mr *MockGenericDataplaneMockRecorder) AddPolicy(policies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPolicy", reflect.TypeOf((*MockGenericDataplane)(nil).AddPolicy), policies)
}

// AddToList mocks base method.
func (m *MockGenericDataplane) AddToList(listName *ipsets.IPSetMetadata, setNames []*ipsets.IPSetMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToList", listName, setNames)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToList indicates an expected call of AddToList.
func (mr *MockGenericDataplaneMockRecorder) AddToList(listName, setNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToList", reflect.TypeOf((*MockGenericDataplane)(nil).AddToList), listName, setNames)
}

// AddToSet mocks base method.
func (m *MockGenericDataplane) AddToSet(setNames []*ipsets.IPSetMetadata, ip, podKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToSet", setNames, ip, podKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToSet indicates an expected call of AddToSet.
func (mr *MockGenericDataplaneMockRecorder) AddToSet(setNames, ip, podKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToSet", reflect.TypeOf((*MockGenericDataplane)(nil).AddToSet), setNames, ip, podKey)
}

// ApplyDataPlane mocks base method.
func (m *MockGenericDataplane) ApplyDataPlane() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyDataPlane")
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyDataPlane indicates an expected call of ApplyDataPlane.
func (mr *MockGenericDataplaneMockRecorder) ApplyDataPlane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyDataPlane", reflect.TypeOf((*MockGenericDataplane)(nil).ApplyDataPlane))
}

// CreateIPSet mocks base method.
func (m *MockGenericDataplane) CreateIPSet(setMetadata *ipsets.IPSetMetadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateIPSet", setMetadata)
}

// CreateIPSet indicates an expected call of CreateIPSet.
func (mr *MockGenericDataplaneMockRecorder) CreateIPSet(setMetadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIPSet", reflect.TypeOf((*MockGenericDataplane)(nil).CreateIPSet), setMetadata)
}

// DeleteIPSet mocks base method.
func (m *MockGenericDataplane) DeleteIPSet(setMetadata *ipsets.IPSetMetadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteIPSet", setMetadata)
}

// DeleteIPSet indicates an expected call of DeleteIPSet.
func (mr *MockGenericDataplaneMockRecorder) DeleteIPSet(setMetadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIPSet", reflect.TypeOf((*MockGenericDataplane)(nil).DeleteIPSet), setMetadata)
}

// InitializeDataPlane mocks base method.
func (m *MockGenericDataplane) InitializeDataPlane() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeDataPlane")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeDataPlane indicates an expected call of InitializeDataPlane.
func (mr *MockGenericDataplaneMockRecorder) InitializeDataPlane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeDataPlane", reflect.TypeOf((*MockGenericDataplane)(nil).InitializeDataPlane))
}

// RemoveFromList mocks base method.
func (m *MockGenericDataplane) RemoveFromList(listName *ipsets.IPSetMetadata, setNames []*ipsets.IPSetMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromList", listName, setNames)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromList indicates an expected call of RemoveFromList.
func (mr *MockGenericDataplaneMockRecorder) RemoveFromList(listName, setNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromList", reflect.TypeOf((*MockGenericDataplane)(nil).RemoveFromList), listName, setNames)
}

// RemoveFromSet mocks base method.
func (m *MockGenericDataplane) RemoveFromSet(setNames []*ipsets.IPSetMetadata, ip, podKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromSet", setNames, ip, podKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromSet indicates an expected call of RemoveFromSet.
func (mr *MockGenericDataplaneMockRecorder) RemoveFromSet(setNames, ip, podKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromSet", reflect.TypeOf((*MockGenericDataplane)(nil).RemoveFromSet), setNames, ip, podKey)
}

// RemovePolicy mocks base method.
func (m *MockGenericDataplane) RemovePolicy(policyName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePolicy", policyName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePolicy indicates an expected call of RemovePolicy.
func (mr *MockGenericDataplaneMockRecorder) RemovePolicy(policyName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePolicy", reflect.TypeOf((*MockGenericDataplane)(nil).RemovePolicy), policyName)
}

// ResetDataPlane mocks base method.
func (m *MockGenericDataplane) ResetDataPlane() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetDataPlane")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetDataPlane indicates an expected call of ResetDataPlane.
func (mr *MockGenericDataplaneMockRecorder) ResetDataPlane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetDataPlane", reflect.TypeOf((*MockGenericDataplane)(nil).ResetDataPlane))
}

// UpdatePod mocks base method.
func (m *MockGenericDataplane) UpdatePod(pod *dataplane.UpdateNPMPod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePod", pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePod indicates an expected call of UpdatePod.
func (mr *MockGenericDataplaneMockRecorder) UpdatePod(pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePod", reflect.TypeOf((*MockGenericDataplane)(nil).UpdatePod), pod)
}

// UpdatePolicy mocks base method.
func (m *MockGenericDataplane) UpdatePolicy(policies *policies.NPMNetworkPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePolicy", policies)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePolicy indicates an expected call of UpdatePolicy.
func (mr *MockGenericDataplaneMockRecorder) UpdatePolicy(policies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePolicy", reflect.TypeOf((*MockGenericDataplane)(nil).UpdatePolicy), policies)
}
