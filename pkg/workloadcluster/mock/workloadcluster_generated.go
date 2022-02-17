// Code generated by MockGen. DO NOT EDIT.
// Source: ./workloadcluster.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	context "sigs.k8s.io/cluster-api-provider-kubevirt/pkg/context"
)

// MockWorkloadCluster is a mock of WorkloadCluster interface.
type MockWorkloadCluster struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadClusterMockRecorder
}

// MockWorkloadClusterMockRecorder is the mock recorder for MockWorkloadCluster.
type MockWorkloadClusterMockRecorder struct {
	mock *MockWorkloadCluster
}

// NewMockWorkloadCluster creates a new mock instance.
func NewMockWorkloadCluster(ctrl *gomock.Controller) *MockWorkloadCluster {
	mock := &MockWorkloadCluster{ctrl: ctrl}
	mock.recorder = &MockWorkloadClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadCluster) EXPECT() *MockWorkloadClusterMockRecorder {
	return m.recorder
}

// GenerateWorkloadClusterClient mocks base method.
func (m *MockWorkloadCluster) GenerateWorkloadClusterClient(ctx *context.MachineContext) (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkloadClusterClient", ctx)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateWorkloadClusterClient indicates an expected call of GenerateWorkloadClusterClient.
func (mr *MockWorkloadClusterMockRecorder) GenerateWorkloadClusterClient(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkloadClusterClient", reflect.TypeOf((*MockWorkloadCluster)(nil).GenerateWorkloadClusterClient), ctx)
}
