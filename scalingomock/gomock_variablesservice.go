// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: VariablesService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVariablesService is a mock of VariablesService interface
type MockVariablesService struct {
	ctrl     *gomock.Controller
	recorder *MockVariablesServiceMockRecorder
}

// MockVariablesServiceMockRecorder is the mock recorder for MockVariablesService
type MockVariablesServiceMockRecorder struct {
	mock *MockVariablesService
}

// NewMockVariablesService creates a new mock instance
func NewMockVariablesService(ctrl *gomock.Controller) *MockVariablesService {
	mock := &MockVariablesService{ctrl: ctrl}
	mock.recorder = &MockVariablesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVariablesService) EXPECT() *MockVariablesServiceMockRecorder {
	return m.recorder
}

// VariableMultipleSet mocks base method
func (m *MockVariablesService) VariableMultipleSet(arg0 string, arg1 go_scalingo.Variables) (go_scalingo.Variables, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariableMultipleSet", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VariableMultipleSet indicates an expected call of VariableMultipleSet
func (mr *MockVariablesServiceMockRecorder) VariableMultipleSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableMultipleSet", reflect.TypeOf((*MockVariablesService)(nil).VariableMultipleSet), arg0, arg1)
}

// VariableSet mocks base method
func (m *MockVariablesService) VariableSet(arg0, arg1, arg2 string) (*go_scalingo.Variable, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariableSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_scalingo.Variable)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VariableSet indicates an expected call of VariableSet
func (mr *MockVariablesServiceMockRecorder) VariableSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableSet", reflect.TypeOf((*MockVariablesService)(nil).VariableSet), arg0, arg1, arg2)
}

// VariableUnset mocks base method
func (m *MockVariablesService) VariableUnset(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariableUnset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VariableUnset indicates an expected call of VariableUnset
func (mr *MockVariablesServiceMockRecorder) VariableUnset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableUnset", reflect.TypeOf((*MockVariablesService)(nil).VariableUnset), arg0, arg1)
}

// VariablesList mocks base method
func (m *MockVariablesService) VariablesList(arg0 string) (go_scalingo.Variables, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariablesList", arg0)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariablesList indicates an expected call of VariablesList
func (mr *MockVariablesServiceMockRecorder) VariablesList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariablesList", reflect.TypeOf((*MockVariablesService)(nil).VariablesList), arg0)
}

// VariablesListWithoutAlias mocks base method
func (m *MockVariablesService) VariablesListWithoutAlias(arg0 string) (go_scalingo.Variables, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VariablesListWithoutAlias", arg0)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariablesListWithoutAlias indicates an expected call of VariablesListWithoutAlias
func (mr *MockVariablesServiceMockRecorder) VariablesListWithoutAlias(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariablesListWithoutAlias", reflect.TypeOf((*MockVariablesService)(nil).VariablesListWithoutAlias), arg0)
}
