// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: AlertsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAlertsService is a mock of AlertsService interface
type MockAlertsService struct {
	ctrl     *gomock.Controller
	recorder *MockAlertsServiceMockRecorder
}

// MockAlertsServiceMockRecorder is the mock recorder for MockAlertsService
type MockAlertsServiceMockRecorder struct {
	mock *MockAlertsService
}

// NewMockAlertsService creates a new mock instance
func NewMockAlertsService(ctrl *gomock.Controller) *MockAlertsService {
	mock := &MockAlertsService{ctrl: ctrl}
	mock.recorder = &MockAlertsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlertsService) EXPECT() *MockAlertsServiceMockRecorder {
	return m.recorder
}

// AlertAdd mocks base method
func (m *MockAlertsService) AlertAdd(arg0 string, arg1 go_scalingo.AlertAddParams) (*go_scalingo.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertAdd", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertAdd indicates an expected call of AlertAdd
func (mr *MockAlertsServiceMockRecorder) AlertAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertAdd", reflect.TypeOf((*MockAlertsService)(nil).AlertAdd), arg0, arg1)
}

// AlertRemove mocks base method
func (m *MockAlertsService) AlertRemove(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertRemove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlertRemove indicates an expected call of AlertRemove
func (mr *MockAlertsServiceMockRecorder) AlertRemove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertRemove", reflect.TypeOf((*MockAlertsService)(nil).AlertRemove), arg0, arg1)
}

// AlertShow mocks base method
func (m *MockAlertsService) AlertShow(arg0, arg1 string) (*go_scalingo.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertShow", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertShow indicates an expected call of AlertShow
func (mr *MockAlertsServiceMockRecorder) AlertShow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertShow", reflect.TypeOf((*MockAlertsService)(nil).AlertShow), arg0, arg1)
}

// AlertUpdate mocks base method
func (m *MockAlertsService) AlertUpdate(arg0, arg1 string, arg2 go_scalingo.AlertUpdateParams) (*go_scalingo.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertUpdate indicates an expected call of AlertUpdate
func (mr *MockAlertsServiceMockRecorder) AlertUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertUpdate", reflect.TypeOf((*MockAlertsService)(nil).AlertUpdate), arg0, arg1, arg2)
}

// AlertsList mocks base method
func (m *MockAlertsService) AlertsList(arg0 string) ([]*go_scalingo.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertsList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertsList indicates an expected call of AlertsList
func (mr *MockAlertsServiceMockRecorder) AlertsList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertsList", reflect.TypeOf((*MockAlertsService)(nil).AlertsList), arg0)
}
