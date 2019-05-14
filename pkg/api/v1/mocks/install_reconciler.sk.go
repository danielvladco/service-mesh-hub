// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/install_reconciler.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

// MockInstallReconciler is a mock of InstallReconciler interface
type MockInstallReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockInstallReconcilerMockRecorder
}

// MockInstallReconcilerMockRecorder is the mock recorder for MockInstallReconciler
type MockInstallReconcilerMockRecorder struct {
	mock *MockInstallReconciler
}

// NewMockInstallReconciler creates a new mock instance
func NewMockInstallReconciler(ctrl *gomock.Controller) *MockInstallReconciler {
	mock := &MockInstallReconciler{ctrl: ctrl}
	mock.recorder = &MockInstallReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstallReconciler) EXPECT() *MockInstallReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockInstallReconciler) Reconcile(namespace string, desiredResources v1.InstallList, transition v1.TransitionInstallFunc, opts clients.ListOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", namespace, desiredResources, transition, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockInstallReconcilerMockRecorder) Reconcile(namespace, desiredResources, transition, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockInstallReconciler)(nil).Reconcile), namespace, desiredResources, transition, opts)
}