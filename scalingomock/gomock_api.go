// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: API)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	http "net/http"
	reflect "reflect"

	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	websocket "golang.org/x/net/websocket"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// APIVersion mocks base method
func (m *MockAPI) APIVersion() string {
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// APIVersion indicates an expected call of APIVersion
func (mr *MockAPIMockRecorder) APIVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockAPI)(nil).APIVersion))
}

// AddonDestroy mocks base method
func (m *MockAPI) AddonDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AddonDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddonDestroy indicates an expected call of AddonDestroy
func (mr *MockAPIMockRecorder) AddonDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonDestroy", reflect.TypeOf((*MockAPI)(nil).AddonDestroy), arg0, arg1)
}

// AddonProviderPlansList mocks base method
func (m *MockAPI) AddonProviderPlansList(arg0 string) ([]*go_scalingo.Plan, error) {
	ret := m.ctrl.Call(m, "AddonProviderPlansList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProviderPlansList indicates an expected call of AddonProviderPlansList
func (mr *MockAPIMockRecorder) AddonProviderPlansList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProviderPlansList", reflect.TypeOf((*MockAPI)(nil).AddonProviderPlansList), arg0)
}

// AddonProvidersList mocks base method
func (m *MockAPI) AddonProvidersList() ([]*go_scalingo.AddonProvider, error) {
	ret := m.ctrl.Call(m, "AddonProvidersList")
	ret0, _ := ret[0].([]*go_scalingo.AddonProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProvidersList indicates an expected call of AddonProvidersList
func (mr *MockAPIMockRecorder) AddonProvidersList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProvidersList", reflect.TypeOf((*MockAPI)(nil).AddonProvidersList))
}

// AddonProvision mocks base method
func (m *MockAPI) AddonProvision(arg0, arg1, arg2 string) (go_scalingo.AddonRes, error) {
	ret := m.ctrl.Call(m, "AddonProvision", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.AddonRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProvision indicates an expected call of AddonProvision
func (mr *MockAPIMockRecorder) AddonProvision(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProvision", reflect.TypeOf((*MockAPI)(nil).AddonProvision), arg0, arg1, arg2)
}

// AddonUpgrade mocks base method
func (m *MockAPI) AddonUpgrade(arg0, arg1, arg2 string) (go_scalingo.AddonRes, error) {
	ret := m.ctrl.Call(m, "AddonUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.AddonRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonUpgrade indicates an expected call of AddonUpgrade
func (mr *MockAPIMockRecorder) AddonUpgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonUpgrade", reflect.TypeOf((*MockAPI)(nil).AddonUpgrade), arg0, arg1, arg2)
}

// AddonsList mocks base method
func (m *MockAPI) AddonsList(arg0 string) ([]*go_scalingo.Addon, error) {
	ret := m.ctrl.Call(m, "AddonsList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Addon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonsList indicates an expected call of AddonsList
func (mr *MockAPIMockRecorder) AddonsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonsList", reflect.TypeOf((*MockAPI)(nil).AddonsList), arg0)
}

// AlertAdd mocks base method
func (m *MockAPI) AlertAdd(arg0 string, arg1 go_scalingo.AlertAddParams) (*go_scalingo.Alert, error) {
	ret := m.ctrl.Call(m, "AlertAdd", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertAdd indicates an expected call of AlertAdd
func (mr *MockAPIMockRecorder) AlertAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertAdd", reflect.TypeOf((*MockAPI)(nil).AlertAdd), arg0, arg1)
}

// AlertRemove mocks base method
func (m *MockAPI) AlertRemove(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AlertRemove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlertRemove indicates an expected call of AlertRemove
func (mr *MockAPIMockRecorder) AlertRemove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertRemove", reflect.TypeOf((*MockAPI)(nil).AlertRemove), arg0, arg1)
}

// AlertShow mocks base method
func (m *MockAPI) AlertShow(arg0, arg1 string) (*go_scalingo.Alert, error) {
	ret := m.ctrl.Call(m, "AlertShow", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertShow indicates an expected call of AlertShow
func (mr *MockAPIMockRecorder) AlertShow(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertShow", reflect.TypeOf((*MockAPI)(nil).AlertShow), arg0, arg1)
}

// AlertUpdate mocks base method
func (m *MockAPI) AlertUpdate(arg0, arg1 string, arg2 go_scalingo.AlertUpdateParams) (*go_scalingo.Alert, error) {
	ret := m.ctrl.Call(m, "AlertUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertUpdate indicates an expected call of AlertUpdate
func (mr *MockAPIMockRecorder) AlertUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertUpdate", reflect.TypeOf((*MockAPI)(nil).AlertUpdate), arg0, arg1, arg2)
}

// AlertsList mocks base method
func (m *MockAPI) AlertsList(arg0 string) ([]*go_scalingo.Alert, error) {
	ret := m.ctrl.Call(m, "AlertsList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertsList indicates an expected call of AlertsList
func (mr *MockAPIMockRecorder) AlertsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertsList", reflect.TypeOf((*MockAPI)(nil).AlertsList), arg0)
}

// AppsCreate mocks base method
func (m *MockAPI) AppsCreate(arg0 go_scalingo.AppsCreateOpts) (*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsCreate", arg0)
	ret0, _ := ret[0].(*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsCreate indicates an expected call of AppsCreate
func (mr *MockAPIMockRecorder) AppsCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsCreate", reflect.TypeOf((*MockAPI)(nil).AppsCreate), arg0)
}

// AppsDestroy mocks base method
func (m *MockAPI) AppsDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AppsDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppsDestroy indicates an expected call of AppsDestroy
func (mr *MockAPIMockRecorder) AppsDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsDestroy", reflect.TypeOf((*MockAPI)(nil).AppsDestroy), arg0, arg1)
}

// AppsList mocks base method
func (m *MockAPI) AppsList() ([]*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsList")
	ret0, _ := ret[0].([]*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsList indicates an expected call of AppsList
func (mr *MockAPIMockRecorder) AppsList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsList", reflect.TypeOf((*MockAPI)(nil).AppsList))
}

// AppsPs mocks base method
func (m *MockAPI) AppsPs(arg0 string) ([]go_scalingo.ContainerType, error) {
	ret := m.ctrl.Call(m, "AppsPs", arg0)
	ret0, _ := ret[0].([]go_scalingo.ContainerType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsPs indicates an expected call of AppsPs
func (mr *MockAPIMockRecorder) AppsPs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsPs", reflect.TypeOf((*MockAPI)(nil).AppsPs), arg0)
}

// AppsRename mocks base method
func (m *MockAPI) AppsRename(arg0, arg1 string) (*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsRename", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsRename indicates an expected call of AppsRename
func (mr *MockAPIMockRecorder) AppsRename(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsRename", reflect.TypeOf((*MockAPI)(nil).AppsRename), arg0, arg1)
}

// AppsRestart mocks base method
func (m *MockAPI) AppsRestart(arg0 string, arg1 *go_scalingo.AppsRestartParams) (*http.Response, error) {
	ret := m.ctrl.Call(m, "AppsRestart", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsRestart indicates an expected call of AppsRestart
func (mr *MockAPIMockRecorder) AppsRestart(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsRestart", reflect.TypeOf((*MockAPI)(nil).AppsRestart), arg0, arg1)
}

// AppsScale mocks base method
func (m *MockAPI) AppsScale(arg0 string, arg1 *go_scalingo.AppsScaleParams) (*http.Response, error) {
	ret := m.ctrl.Call(m, "AppsScale", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsScale indicates an expected call of AppsScale
func (mr *MockAPIMockRecorder) AppsScale(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsScale", reflect.TypeOf((*MockAPI)(nil).AppsScale), arg0, arg1)
}

// AppsShow mocks base method
func (m *MockAPI) AppsShow(arg0 string) (*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsShow", arg0)
	ret0, _ := ret[0].(*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsShow indicates an expected call of AppsShow
func (mr *MockAPIMockRecorder) AppsShow(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsShow", reflect.TypeOf((*MockAPI)(nil).AppsShow), arg0)
}

// AppsStats mocks base method
func (m *MockAPI) AppsStats(arg0 string) (*go_scalingo.AppStatsRes, error) {
	ret := m.ctrl.Call(m, "AppsStats", arg0)
	ret0, _ := ret[0].(*go_scalingo.AppStatsRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsStats indicates an expected call of AppsStats
func (mr *MockAPIMockRecorder) AppsStats(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsStats", reflect.TypeOf((*MockAPI)(nil).AppsStats), arg0)
}

// AppsTransfer mocks base method
func (m *MockAPI) AppsTransfer(arg0, arg1 string) (*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsTransfer", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsTransfer indicates an expected call of AppsTransfer
func (mr *MockAPIMockRecorder) AppsTransfer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsTransfer", reflect.TypeOf((*MockAPI)(nil).AppsTransfer), arg0, arg1)
}

// CollaboratorAdd mocks base method
func (m *MockAPI) CollaboratorAdd(arg0, arg1 string) (go_scalingo.Collaborator, error) {
	ret := m.ctrl.Call(m, "CollaboratorAdd", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Collaborator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollaboratorAdd indicates an expected call of CollaboratorAdd
func (mr *MockAPIMockRecorder) CollaboratorAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollaboratorAdd", reflect.TypeOf((*MockAPI)(nil).CollaboratorAdd), arg0, arg1)
}

// CollaboratorRemove mocks base method
func (m *MockAPI) CollaboratorRemove(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "CollaboratorRemove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CollaboratorRemove indicates an expected call of CollaboratorRemove
func (mr *MockAPIMockRecorder) CollaboratorRemove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollaboratorRemove", reflect.TypeOf((*MockAPI)(nil).CollaboratorRemove), arg0, arg1)
}

// CollaboratorsList mocks base method
func (m *MockAPI) CollaboratorsList(arg0 string) ([]go_scalingo.Collaborator, error) {
	ret := m.ctrl.Call(m, "CollaboratorsList", arg0)
	ret0, _ := ret[0].([]go_scalingo.Collaborator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollaboratorsList indicates an expected call of CollaboratorsList
func (mr *MockAPIMockRecorder) CollaboratorsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollaboratorsList", reflect.TypeOf((*MockAPI)(nil).CollaboratorsList), arg0)
}

// Deployment mocks base method
func (m *MockAPI) Deployment(arg0, arg1 string) (*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "Deployment", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deployment indicates an expected call of Deployment
func (mr *MockAPIMockRecorder) Deployment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockAPI)(nil).Deployment), arg0, arg1)
}

// DeploymentList mocks base method
func (m *MockAPI) DeploymentList(arg0 string) ([]*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "DeploymentList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentList indicates an expected call of DeploymentList
func (mr *MockAPIMockRecorder) DeploymentList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentList", reflect.TypeOf((*MockAPI)(nil).DeploymentList), arg0)
}

// DeploymentLogs mocks base method
func (m *MockAPI) DeploymentLogs(arg0 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "DeploymentLogs", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentLogs indicates an expected call of DeploymentLogs
func (mr *MockAPIMockRecorder) DeploymentLogs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentLogs", reflect.TypeOf((*MockAPI)(nil).DeploymentLogs), arg0)
}

// DeploymentStream mocks base method
func (m *MockAPI) DeploymentStream(arg0 string) (*websocket.Conn, error) {
	ret := m.ctrl.Call(m, "DeploymentStream", arg0)
	ret0, _ := ret[0].(*websocket.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentStream indicates an expected call of DeploymentStream
func (mr *MockAPIMockRecorder) DeploymentStream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentStream", reflect.TypeOf((*MockAPI)(nil).DeploymentStream), arg0)
}

// DeploymentsCreate mocks base method
func (m *MockAPI) DeploymentsCreate(arg0 string, arg1 *go_scalingo.DeploymentsCreateParams) (*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "DeploymentsCreate", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentsCreate indicates an expected call of DeploymentsCreate
func (mr *MockAPIMockRecorder) DeploymentsCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentsCreate", reflect.TypeOf((*MockAPI)(nil).DeploymentsCreate), arg0, arg1)
}

// DomainsAdd mocks base method
func (m *MockAPI) DomainsAdd(arg0 string, arg1 go_scalingo.Domain) (go_scalingo.Domain, error) {
	ret := m.ctrl.Call(m, "DomainsAdd", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainsAdd indicates an expected call of DomainsAdd
func (mr *MockAPIMockRecorder) DomainsAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsAdd", reflect.TypeOf((*MockAPI)(nil).DomainsAdd), arg0, arg1)
}

// DomainsList mocks base method
func (m *MockAPI) DomainsList(arg0 string) ([]go_scalingo.Domain, error) {
	ret := m.ctrl.Call(m, "DomainsList", arg0)
	ret0, _ := ret[0].([]go_scalingo.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainsList indicates an expected call of DomainsList
func (mr *MockAPIMockRecorder) DomainsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsList", reflect.TypeOf((*MockAPI)(nil).DomainsList), arg0)
}

// DomainsRemove mocks base method
func (m *MockAPI) DomainsRemove(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "DomainsRemove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DomainsRemove indicates an expected call of DomainsRemove
func (mr *MockAPIMockRecorder) DomainsRemove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsRemove", reflect.TypeOf((*MockAPI)(nil).DomainsRemove), arg0, arg1)
}

// DomainsUpdate mocks base method
func (m *MockAPI) DomainsUpdate(arg0, arg1, arg2, arg3 string) (go_scalingo.Domain, error) {
	ret := m.ctrl.Call(m, "DomainsUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(go_scalingo.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainsUpdate indicates an expected call of DomainsUpdate
func (mr *MockAPIMockRecorder) DomainsUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsUpdate", reflect.TypeOf((*MockAPI)(nil).DomainsUpdate), arg0, arg1, arg2, arg3)
}

// Endpoint mocks base method
func (m *MockAPI) Endpoint() string {
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// Endpoint indicates an expected call of Endpoint
func (mr *MockAPIMockRecorder) Endpoint() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockAPI)(nil).Endpoint))
}

// EventsList mocks base method
func (m *MockAPI) EventsList(arg0 string, arg1 go_scalingo.PaginationOpts) (go_scalingo.Events, go_scalingo.PaginationMeta, error) {
	ret := m.ctrl.Call(m, "EventsList", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Events)
	ret1, _ := ret[1].(go_scalingo.PaginationMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EventsList indicates an expected call of EventsList
func (mr *MockAPIMockRecorder) EventsList(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsList", reflect.TypeOf((*MockAPI)(nil).EventsList), arg0, arg1)
}

// GetAccessToken mocks base method
func (m *MockAPI) GetAccessToken() (string, error) {
	ret := m.ctrl.Call(m, "GetAccessToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessToken indicates an expected call of GetAccessToken
func (mr *MockAPIMockRecorder) GetAccessToken() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockAPI)(nil).GetAccessToken))
}

// HTTPClient mocks base method
func (m *MockAPI) HTTPClient() go_scalingo.HTTPClient {
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(go_scalingo.HTTPClient)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient
func (mr *MockAPIMockRecorder) HTTPClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockAPI)(nil).HTTPClient))
}

// KeysAdd mocks base method
func (m *MockAPI) KeysAdd(arg0, arg1 string) (*go_scalingo.Key, error) {
	ret := m.ctrl.Call(m, "KeysAdd", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeysAdd indicates an expected call of KeysAdd
func (mr *MockAPIMockRecorder) KeysAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysAdd", reflect.TypeOf((*MockAPI)(nil).KeysAdd), arg0, arg1)
}

// KeysDelete mocks base method
func (m *MockAPI) KeysDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "KeysDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// KeysDelete indicates an expected call of KeysDelete
func (mr *MockAPIMockRecorder) KeysDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysDelete", reflect.TypeOf((*MockAPI)(nil).KeysDelete), arg0)
}

// KeysList mocks base method
func (m *MockAPI) KeysList() ([]go_scalingo.Key, error) {
	ret := m.ctrl.Call(m, "KeysList")
	ret0, _ := ret[0].([]go_scalingo.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeysList indicates an expected call of KeysList
func (mr *MockAPIMockRecorder) KeysList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysList", reflect.TypeOf((*MockAPI)(nil).KeysList))
}

// Login mocks base method
func (m *MockAPI) Login(arg0, arg1 string) (*go_scalingo.LoginResponse, error) {
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAPIMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAPI)(nil).Login), arg0, arg1)
}

// Logs mocks base method
func (m *MockAPI) Logs(arg0 string, arg1 int, arg2 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Logs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs
func (mr *MockAPIMockRecorder) Logs(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockAPI)(nil).Logs), arg0, arg1, arg2)
}

// LogsArchives mocks base method
func (m *MockAPI) LogsArchives(arg0 string, arg1 int) (*go_scalingo.LogsArchivesResponse, error) {
	ret := m.ctrl.Call(m, "LogsArchives", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.LogsArchivesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsArchives indicates an expected call of LogsArchives
func (mr *MockAPIMockRecorder) LogsArchives(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsArchives", reflect.TypeOf((*MockAPI)(nil).LogsArchives), arg0, arg1)
}

// LogsArchivesByCursor mocks base method
func (m *MockAPI) LogsArchivesByCursor(arg0, arg1 string) (*go_scalingo.LogsArchivesResponse, error) {
	ret := m.ctrl.Call(m, "LogsArchivesByCursor", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.LogsArchivesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsArchivesByCursor indicates an expected call of LogsArchivesByCursor
func (mr *MockAPIMockRecorder) LogsArchivesByCursor(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsArchivesByCursor", reflect.TypeOf((*MockAPI)(nil).LogsArchivesByCursor), arg0, arg1)
}

// LogsURL mocks base method
func (m *MockAPI) LogsURL(arg0 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "LogsURL", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsURL indicates an expected call of LogsURL
func (mr *MockAPIMockRecorder) LogsURL(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsURL", reflect.TypeOf((*MockAPI)(nil).LogsURL), arg0)
}

// NotificationDestroy mocks base method
func (m *MockAPI) NotificationDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "NotificationDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotificationDestroy indicates an expected call of NotificationDestroy
func (mr *MockAPIMockRecorder) NotificationDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationDestroy", reflect.TypeOf((*MockAPI)(nil).NotificationDestroy), arg0, arg1)
}

// NotificationPlatformByName mocks base method
func (m *MockAPI) NotificationPlatformByName(arg0 string) ([]*go_scalingo.NotificationPlatform, error) {
	ret := m.ctrl.Call(m, "NotificationPlatformByName", arg0)
	ret0, _ := ret[0].([]*go_scalingo.NotificationPlatform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationPlatformByName indicates an expected call of NotificationPlatformByName
func (mr *MockAPIMockRecorder) NotificationPlatformByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationPlatformByName", reflect.TypeOf((*MockAPI)(nil).NotificationPlatformByName), arg0)
}

// NotificationPlatformsList mocks base method
func (m *MockAPI) NotificationPlatformsList() ([]*go_scalingo.NotificationPlatform, error) {
	ret := m.ctrl.Call(m, "NotificationPlatformsList")
	ret0, _ := ret[0].([]*go_scalingo.NotificationPlatform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationPlatformsList indicates an expected call of NotificationPlatformsList
func (mr *MockAPIMockRecorder) NotificationPlatformsList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationPlatformsList", reflect.TypeOf((*MockAPI)(nil).NotificationPlatformsList))
}

// NotificationProvision mocks base method
func (m *MockAPI) NotificationProvision(arg0, arg1 string) (go_scalingo.NotificationRes, error) {
	ret := m.ctrl.Call(m, "NotificationProvision", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.NotificationRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationProvision indicates an expected call of NotificationProvision
func (mr *MockAPIMockRecorder) NotificationProvision(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationProvision", reflect.TypeOf((*MockAPI)(nil).NotificationProvision), arg0, arg1)
}

// NotificationUpdate mocks base method
func (m *MockAPI) NotificationUpdate(arg0, arg1, arg2 string) (go_scalingo.NotificationRes, error) {
	ret := m.ctrl.Call(m, "NotificationUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.NotificationRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationUpdate indicates an expected call of NotificationUpdate
func (mr *MockAPIMockRecorder) NotificationUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationUpdate", reflect.TypeOf((*MockAPI)(nil).NotificationUpdate), arg0, arg1, arg2)
}

// NotificationsList mocks base method
func (m *MockAPI) NotificationsList(arg0 string) ([]*go_scalingo.Notification, error) {
	ret := m.ctrl.Call(m, "NotificationsList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationsList indicates an expected call of NotificationsList
func (mr *MockAPIMockRecorder) NotificationsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationsList", reflect.TypeOf((*MockAPI)(nil).NotificationsList), arg0)
}

// NotifierByID mocks base method
func (m *MockAPI) NotifierByID(arg0, arg1 string) (*go_scalingo.Notifier, error) {
	ret := m.ctrl.Call(m, "NotifierByID", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifierByID indicates an expected call of NotifierByID
func (mr *MockAPIMockRecorder) NotifierByID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierByID", reflect.TypeOf((*MockAPI)(nil).NotifierByID), arg0, arg1)
}

// NotifierDestroy mocks base method
func (m *MockAPI) NotifierDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "NotifierDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifierDestroy indicates an expected call of NotifierDestroy
func (mr *MockAPIMockRecorder) NotifierDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierDestroy", reflect.TypeOf((*MockAPI)(nil).NotifierDestroy), arg0, arg1)
}

// NotifierProvision mocks base method
func (m *MockAPI) NotifierProvision(arg0, arg1 string, arg2 go_scalingo.NotifierParams) (*go_scalingo.Notifier, error) {
	ret := m.ctrl.Call(m, "NotifierProvision", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_scalingo.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifierProvision indicates an expected call of NotifierProvision
func (mr *MockAPIMockRecorder) NotifierProvision(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierProvision", reflect.TypeOf((*MockAPI)(nil).NotifierProvision), arg0, arg1, arg2)
}

// NotifierUpdate mocks base method
func (m *MockAPI) NotifierUpdate(arg0, arg1, arg2 string, arg3 go_scalingo.NotifierParams) (*go_scalingo.Notifier, error) {
	ret := m.ctrl.Call(m, "NotifierUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*go_scalingo.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifierUpdate indicates an expected call of NotifierUpdate
func (mr *MockAPIMockRecorder) NotifierUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierUpdate", reflect.TypeOf((*MockAPI)(nil).NotifierUpdate), arg0, arg1, arg2, arg3)
}

// NotifiersList mocks base method
func (m *MockAPI) NotifiersList(arg0 string) (go_scalingo.Notifiers, error) {
	ret := m.ctrl.Call(m, "NotifiersList", arg0)
	ret0, _ := ret[0].(go_scalingo.Notifiers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifiersList indicates an expected call of NotifiersList
func (mr *MockAPIMockRecorder) NotifiersList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifiersList", reflect.TypeOf((*MockAPI)(nil).NotifiersList), arg0)
}

// OperationsShow mocks base method
func (m *MockAPI) OperationsShow(arg0, arg1 string) (*go_scalingo.Operation, error) {
	ret := m.ctrl.Call(m, "OperationsShow", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperationsShow indicates an expected call of OperationsShow
func (mr *MockAPIMockRecorder) OperationsShow(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationsShow", reflect.TypeOf((*MockAPI)(nil).OperationsShow), arg0, arg1)
}

// Run mocks base method
func (m *MockAPI) Run(arg0 go_scalingo.RunOpts) (*go_scalingo.RunRes, error) {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(*go_scalingo.RunRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockAPIMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockAPI)(nil).Run), arg0)
}

// Self mocks base method
func (m *MockAPI) Self() (*go_scalingo.User, error) {
	ret := m.ctrl.Call(m, "Self")
	ret0, _ := ret[0].(*go_scalingo.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Self indicates an expected call of Self
func (mr *MockAPIMockRecorder) Self() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Self", reflect.TypeOf((*MockAPI)(nil).Self))
}

// SignUp mocks base method
func (m *MockAPI) SignUp(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SignUp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp
func (mr *MockAPIMockRecorder) SignUp(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAPI)(nil).SignUp), arg0, arg1)
}

// SourcesCreate mocks base method
func (m *MockAPI) SourcesCreate() (*go_scalingo.Source, error) {
	ret := m.ctrl.Call(m, "SourcesCreate")
	ret0, _ := ret[0].(*go_scalingo.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SourcesCreate indicates an expected call of SourcesCreate
func (mr *MockAPIMockRecorder) SourcesCreate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourcesCreate", reflect.TypeOf((*MockAPI)(nil).SourcesCreate))
}

// TokenCreate mocks base method
func (m *MockAPI) TokenCreate(arg0 go_scalingo.TokenCreateParams) (go_scalingo.Token, error) {
	ret := m.ctrl.Call(m, "TokenCreate", arg0)
	ret0, _ := ret[0].(go_scalingo.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenCreate indicates an expected call of TokenCreate
func (mr *MockAPIMockRecorder) TokenCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenCreate", reflect.TypeOf((*MockAPI)(nil).TokenCreate), arg0)
}

// TokenExchange mocks base method
func (m *MockAPI) TokenExchange(arg0 go_scalingo.TokenExchangeParams) (string, error) {
	ret := m.ctrl.Call(m, "TokenExchange", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenExchange indicates an expected call of TokenExchange
func (mr *MockAPIMockRecorder) TokenExchange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenExchange", reflect.TypeOf((*MockAPI)(nil).TokenExchange), arg0)
}

// TokenShow mocks base method
func (m *MockAPI) TokenShow(arg0 int) (go_scalingo.Token, error) {
	ret := m.ctrl.Call(m, "TokenShow", arg0)
	ret0, _ := ret[0].(go_scalingo.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenShow indicates an expected call of TokenShow
func (mr *MockAPIMockRecorder) TokenShow(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenShow", reflect.TypeOf((*MockAPI)(nil).TokenShow), arg0)
}

// TokensList mocks base method
func (m *MockAPI) TokensList() (go_scalingo.Tokens, error) {
	ret := m.ctrl.Call(m, "TokensList")
	ret0, _ := ret[0].(go_scalingo.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokensList indicates an expected call of TokensList
func (mr *MockAPIMockRecorder) TokensList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokensList", reflect.TypeOf((*MockAPI)(nil).TokensList))
}

// UpdateUser mocks base method
func (m *MockAPI) UpdateUser(arg0 go_scalingo.UpdateUserParams) (*go_scalingo.User, error) {
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(*go_scalingo.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockAPIMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAPI)(nil).UpdateUser), arg0)
}

// UserEventsList mocks base method
func (m *MockAPI) UserEventsList(arg0 go_scalingo.PaginationOpts) (go_scalingo.Events, go_scalingo.PaginationMeta, error) {
	ret := m.ctrl.Call(m, "UserEventsList", arg0)
	ret0, _ := ret[0].(go_scalingo.Events)
	ret1, _ := ret[1].(go_scalingo.PaginationMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserEventsList indicates an expected call of UserEventsList
func (mr *MockAPIMockRecorder) UserEventsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEventsList", reflect.TypeOf((*MockAPI)(nil).UserEventsList), arg0)
}

// VariableMultipleSet mocks base method
func (m *MockAPI) VariableMultipleSet(arg0 string, arg1 go_scalingo.Variables) (go_scalingo.Variables, int, error) {
	ret := m.ctrl.Call(m, "VariableMultipleSet", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VariableMultipleSet indicates an expected call of VariableMultipleSet
func (mr *MockAPIMockRecorder) VariableMultipleSet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableMultipleSet", reflect.TypeOf((*MockAPI)(nil).VariableMultipleSet), arg0, arg1)
}

// VariableSet mocks base method
func (m *MockAPI) VariableSet(arg0, arg1, arg2 string) (*go_scalingo.Variable, int, error) {
	ret := m.ctrl.Call(m, "VariableSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_scalingo.Variable)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VariableSet indicates an expected call of VariableSet
func (mr *MockAPIMockRecorder) VariableSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableSet", reflect.TypeOf((*MockAPI)(nil).VariableSet), arg0, arg1, arg2)
}

// VariableUnset mocks base method
func (m *MockAPI) VariableUnset(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "VariableUnset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VariableUnset indicates an expected call of VariableUnset
func (mr *MockAPIMockRecorder) VariableUnset(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableUnset", reflect.TypeOf((*MockAPI)(nil).VariableUnset), arg0, arg1)
}

// VariablesList mocks base method
func (m *MockAPI) VariablesList(arg0 string) (go_scalingo.Variables, error) {
	ret := m.ctrl.Call(m, "VariablesList", arg0)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariablesList indicates an expected call of VariablesList
func (mr *MockAPIMockRecorder) VariablesList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariablesList", reflect.TypeOf((*MockAPI)(nil).VariablesList), arg0)
}

// VariablesListWithoutAlias mocks base method
func (m *MockAPI) VariablesListWithoutAlias(arg0 string) (go_scalingo.Variables, error) {
	ret := m.ctrl.Call(m, "VariablesListWithoutAlias", arg0)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariablesListWithoutAlias indicates an expected call of VariablesListWithoutAlias
func (mr *MockAPIMockRecorder) VariablesListWithoutAlias(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariablesListWithoutAlias", reflect.TypeOf((*MockAPI)(nil).VariablesListWithoutAlias), arg0)
}
