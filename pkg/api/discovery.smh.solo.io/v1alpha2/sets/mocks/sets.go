// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha2sets is a generated GoMock package.
package mock_v1alpha2sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockTrafficTargetSet is a mock of TrafficTargetSet interface
type MockTrafficTargetSet struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetSetMockRecorder
}

// MockTrafficTargetSetMockRecorder is the mock recorder for MockTrafficTargetSet
type MockTrafficTargetSetMockRecorder struct {
	mock *MockTrafficTargetSet
}

// NewMockTrafficTargetSet creates a new mock instance
func NewMockTrafficTargetSet(ctrl *gomock.Controller) *MockTrafficTargetSet {
	mock := &MockTrafficTargetSet{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrafficTargetSet) EXPECT() *MockTrafficTargetSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockTrafficTargetSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockTrafficTargetSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockTrafficTargetSet)(nil).Keys))
}

// List mocks base method
func (m *MockTrafficTargetSet) List(filterResource ...func(*v1alpha2.TrafficTarget) bool) []*v1alpha2.TrafficTarget {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha2.TrafficTarget)
	return ret0
}

// List indicates an expected call of List
func (mr *MockTrafficTargetSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTrafficTargetSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockTrafficTargetSet) Map() map[string]*v1alpha2.TrafficTarget {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha2.TrafficTarget)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockTrafficTargetSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockTrafficTargetSet)(nil).Map))
}

// Insert mocks base method
func (m *MockTrafficTargetSet) Insert(trafficTarget ...*v1alpha2.TrafficTarget) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficTarget {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockTrafficTargetSetMockRecorder) Insert(trafficTarget ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTrafficTargetSet)(nil).Insert), trafficTarget...)
}

// Equal mocks base method
func (m *MockTrafficTargetSet) Equal(trafficTargetSet v1alpha2sets.TrafficTargetSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", trafficTargetSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockTrafficTargetSetMockRecorder) Equal(trafficTargetSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockTrafficTargetSet)(nil).Equal), trafficTargetSet)
}

// Has mocks base method
func (m *MockTrafficTargetSet) Has(trafficTarget ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", trafficTarget)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockTrafficTargetSetMockRecorder) Has(trafficTarget interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockTrafficTargetSet)(nil).Has), trafficTarget)
}

// Delete mocks base method
func (m *MockTrafficTargetSet) Delete(trafficTarget ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", trafficTarget)
}

// Delete indicates an expected call of Delete
func (mr *MockTrafficTargetSetMockRecorder) Delete(trafficTarget interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrafficTargetSet)(nil).Delete), trafficTarget)
}

// Union mocks base method
func (m *MockTrafficTargetSet) Union(set v1alpha2sets.TrafficTargetSet) v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockTrafficTargetSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockTrafficTargetSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockTrafficTargetSet) Difference(set v1alpha2sets.TrafficTargetSet) v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockTrafficTargetSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockTrafficTargetSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockTrafficTargetSet) Intersection(set v1alpha2sets.TrafficTargetSet) v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockTrafficTargetSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockTrafficTargetSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockTrafficTargetSet) Find(id ezkube.ResourceId) (*v1alpha2.TrafficTarget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha2.TrafficTarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockTrafficTargetSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTrafficTargetSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockTrafficTargetSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockTrafficTargetSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockTrafficTargetSet)(nil).Length))
}

// MockWorkloadSet is a mock of WorkloadSet interface
type MockWorkloadSet struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadSetMockRecorder
}

// MockWorkloadSetMockRecorder is the mock recorder for MockWorkloadSet
type MockWorkloadSetMockRecorder struct {
	mock *MockWorkloadSet
}

// NewMockWorkloadSet creates a new mock instance
func NewMockWorkloadSet(ctrl *gomock.Controller) *MockWorkloadSet {
	mock := &MockWorkloadSet{ctrl: ctrl}
	mock.recorder = &MockWorkloadSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkloadSet) EXPECT() *MockWorkloadSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockWorkloadSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockWorkloadSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockWorkloadSet)(nil).Keys))
}

// List mocks base method
func (m *MockWorkloadSet) List(filterResource ...func(*v1alpha2.Workload) bool) []*v1alpha2.Workload {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha2.Workload)
	return ret0
}

// List indicates an expected call of List
func (mr *MockWorkloadSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWorkloadSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockWorkloadSet) Map() map[string]*v1alpha2.Workload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha2.Workload)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockWorkloadSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockWorkloadSet)(nil).Map))
}

// Insert mocks base method
func (m *MockWorkloadSet) Insert(workload ...*v1alpha2.Workload) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range workload {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockWorkloadSetMockRecorder) Insert(workload ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockWorkloadSet)(nil).Insert), workload...)
}

// Equal mocks base method
func (m *MockWorkloadSet) Equal(workloadSet v1alpha2sets.WorkloadSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", workloadSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockWorkloadSetMockRecorder) Equal(workloadSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockWorkloadSet)(nil).Equal), workloadSet)
}

// Has mocks base method
func (m *MockWorkloadSet) Has(workload ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", workload)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockWorkloadSetMockRecorder) Has(workload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockWorkloadSet)(nil).Has), workload)
}

// Delete mocks base method
func (m *MockWorkloadSet) Delete(workload ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", workload)
}

// Delete indicates an expected call of Delete
func (mr *MockWorkloadSetMockRecorder) Delete(workload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWorkloadSet)(nil).Delete), workload)
}

// Union mocks base method
func (m *MockWorkloadSet) Union(set v1alpha2sets.WorkloadSet) v1alpha2sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha2sets.WorkloadSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockWorkloadSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockWorkloadSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockWorkloadSet) Difference(set v1alpha2sets.WorkloadSet) v1alpha2sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha2sets.WorkloadSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockWorkloadSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockWorkloadSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockWorkloadSet) Intersection(set v1alpha2sets.WorkloadSet) v1alpha2sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha2sets.WorkloadSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockWorkloadSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockWorkloadSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockWorkloadSet) Find(id ezkube.ResourceId) (*v1alpha2.Workload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha2.Workload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockWorkloadSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockWorkloadSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockWorkloadSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockWorkloadSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockWorkloadSet)(nil).Length))
}

// MockMeshSet is a mock of MeshSet interface
type MockMeshSet struct {
	ctrl     *gomock.Controller
	recorder *MockMeshSetMockRecorder
}

// MockMeshSetMockRecorder is the mock recorder for MockMeshSet
type MockMeshSetMockRecorder struct {
	mock *MockMeshSet
}

// NewMockMeshSet creates a new mock instance
func NewMockMeshSet(ctrl *gomock.Controller) *MockMeshSet {
	mock := &MockMeshSet{ctrl: ctrl}
	mock.recorder = &MockMeshSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshSet) EXPECT() *MockMeshSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockMeshSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockMeshSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockMeshSet)(nil).Keys))
}

// List mocks base method
func (m *MockMeshSet) List(filterResource ...func(*v1alpha2.Mesh) bool) []*v1alpha2.Mesh {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha2.Mesh)
	return ret0
}

// List indicates an expected call of List
func (mr *MockMeshSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMeshSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockMeshSet) Map() map[string]*v1alpha2.Mesh {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha2.Mesh)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockMeshSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockMeshSet)(nil).Map))
}

// Insert mocks base method
func (m *MockMeshSet) Insert(mesh ...*v1alpha2.Mesh) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range mesh {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockMeshSetMockRecorder) Insert(mesh ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMeshSet)(nil).Insert), mesh...)
}

// Equal mocks base method
func (m *MockMeshSet) Equal(meshSet v1alpha2sets.MeshSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", meshSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockMeshSetMockRecorder) Equal(meshSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockMeshSet)(nil).Equal), meshSet)
}

// Has mocks base method
func (m *MockMeshSet) Has(mesh ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", mesh)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockMeshSetMockRecorder) Has(mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMeshSet)(nil).Has), mesh)
}

// Delete mocks base method
func (m *MockMeshSet) Delete(mesh ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", mesh)
}

// Delete indicates an expected call of Delete
func (mr *MockMeshSetMockRecorder) Delete(mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshSet)(nil).Delete), mesh)
}

// Union mocks base method
func (m *MockMeshSet) Union(set v1alpha2sets.MeshSet) v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockMeshSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockMeshSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockMeshSet) Difference(set v1alpha2sets.MeshSet) v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockMeshSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockMeshSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockMeshSet) Intersection(set v1alpha2sets.MeshSet) v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockMeshSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockMeshSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockMeshSet) Find(id ezkube.ResourceId) (*v1alpha2.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha2.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockMeshSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMeshSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockMeshSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockMeshSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockMeshSet)(nil).Length))
}
