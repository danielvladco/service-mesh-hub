// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/external/istio/rbac/v1alpha1/service_role_binding_reconciler.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/rbac/v1alpha1"
)

// MockServiceRoleBindingReconciler is a mock of ServiceRoleBindingReconciler interface
type MockServiceRoleBindingReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceRoleBindingReconcilerMockRecorder
}

// MockServiceRoleBindingReconcilerMockRecorder is the mock recorder for MockServiceRoleBindingReconciler
type MockServiceRoleBindingReconcilerMockRecorder struct {
	mock *MockServiceRoleBindingReconciler
}

// NewMockServiceRoleBindingReconciler creates a new mock instance
func NewMockServiceRoleBindingReconciler(ctrl *gomock.Controller) *MockServiceRoleBindingReconciler {
	mock := &MockServiceRoleBindingReconciler{ctrl: ctrl}
	mock.recorder = &MockServiceRoleBindingReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceRoleBindingReconciler) EXPECT() *MockServiceRoleBindingReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockServiceRoleBindingReconciler) Reconcile(namespace string, desiredResources v1alpha1.ServiceRoleBindingList, transition v1alpha1.TransitionServiceRoleBindingFunc, opts clients.ListOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", namespace, desiredResources, transition, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockServiceRoleBindingReconcilerMockRecorder) Reconcile(namespace, desiredResources, transition, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockServiceRoleBindingReconciler)(nil).Reconcile), namespace, desiredResources, transition, opts)
}