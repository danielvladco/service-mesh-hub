// Code generated by MockGen. DO NOT EDIT.
// Source: ./access_policy_translator.go

// Package mock_access is a generated GoMock package.
package mock_access

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	istio "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/output/istio"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method
func (m *MockTranslator) Translate(mesh *v1alpha2.Mesh, virtualMesh *v1alpha2.MeshStatus_AppliedVirtualMesh, outputs istio.Builder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Translate", mesh, virtualMesh, outputs)
}

// Translate indicates an expected call of Translate
func (mr *MockTranslatorMockRecorder) Translate(mesh, virtualMesh, outputs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslator)(nil).Translate), mesh, virtualMesh, outputs)
}
