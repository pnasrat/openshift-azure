// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient (interfaces: Client,VirtualMachineScaleSetsClient,VirtualMachineScaleSetVMsClient,VirtualMachineScaleSetExtensionsClient,ApplicationsClient,MarketPlaceAgreementsClient,DeploymentsClient,AccountsClient,KeyVaultClient,VaultMgmtClient,ZonesClient,RecordSetsClient)

// Package mock_azureclient is a generated GoMock package.
package mock_azureclient

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	dns "github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2017-10-01/dns"
	keyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	keyvault0 "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"
	marketplaceordering "github.com/Azure/azure-sdk-for-go/services/marketplaceordering/mgmt/2015-06-01/marketplaceordering"
	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	managedapplications "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-06-01/managedapplications"
	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockClient)(nil).Client))
}

// MockVirtualMachineScaleSetsClient is a mock of VirtualMachineScaleSetsClient interface
type MockVirtualMachineScaleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetsClientMockRecorder
}

// MockVirtualMachineScaleSetsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetsClient
type MockVirtualMachineScaleSetsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetsClient
}

// NewMockVirtualMachineScaleSetsClient creates a new mock instance
func NewMockVirtualMachineScaleSetsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetsClient {
	mock := &MockVirtualMachineScaleSetsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualMachineScaleSetsClient) EXPECT() *MockVirtualMachineScaleSetsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockVirtualMachineScaleSetsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Client))
}

// CreateOrUpdate mocks base method
func (m *MockVirtualMachineScaleSetsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockVirtualMachineScaleSetsClient) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockVirtualMachineScaleSetsClient) List(arg0 context.Context, arg1 string) ([]compute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).List), arg0, arg1)
}

// Update mocks base method
func (m *MockVirtualMachineScaleSetsClient) Update(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSetUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Update), arg0, arg1, arg2, arg3)
}

// UpdateInstances mocks base method
func (m *MockVirtualMachineScaleSetsClient) UpdateInstances(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSetVMInstanceRequiredIDs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstances", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstances indicates an expected call of UpdateInstances
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) UpdateInstances(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstances", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).UpdateInstances), arg0, arg1, arg2, arg3)
}

// MockVirtualMachineScaleSetVMsClient is a mock of VirtualMachineScaleSetVMsClient interface
type MockVirtualMachineScaleSetVMsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetVMsClientMockRecorder
}

// MockVirtualMachineScaleSetVMsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetVMsClient
type MockVirtualMachineScaleSetVMsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetVMsClient
}

// NewMockVirtualMachineScaleSetVMsClient creates a new mock instance
func NewMockVirtualMachineScaleSetVMsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetVMsClient {
	mock := &MockVirtualMachineScaleSetVMsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetVMsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualMachineScaleSetVMsClient) EXPECT() *MockVirtualMachineScaleSetVMsClientMockRecorder {
	return m.recorder
}

// Deallocate mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Deallocate(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deallocate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deallocate indicates an expected call of Deallocate
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Deallocate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deallocate", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Deallocate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Delete(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Delete), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) List(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) ([]compute.VirtualMachineScaleSetVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSetVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Reimage mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Reimage(arg0 context.Context, arg1, arg2, arg3 string, arg4 *compute.VirtualMachineScaleSetVMReimageParameters) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reimage", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reimage indicates an expected call of Reimage
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Reimage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reimage", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Reimage), arg0, arg1, arg2, arg3, arg4)
}

// ReimageAll mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) ReimageAll(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineScaleSetVMsReimageAllFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReimageAll", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMsReimageAllFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReimageAll indicates an expected call of ReimageAll
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) ReimageAll(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReimageAll", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).ReimageAll), arg0, arg1, arg2, arg3)
}

// Restart mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Restart(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restart", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restart indicates an expected call of Restart
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Restart(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restart", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Restart), arg0, arg1, arg2, arg3)
}

// Start mocks base method
func (m *MockVirtualMachineScaleSetVMsClient) Start(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Start(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Start), arg0, arg1, arg2, arg3)
}

// MockVirtualMachineScaleSetExtensionsClient is a mock of VirtualMachineScaleSetExtensionsClient interface
type MockVirtualMachineScaleSetExtensionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetExtensionsClientMockRecorder
}

// MockVirtualMachineScaleSetExtensionsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetExtensionsClient
type MockVirtualMachineScaleSetExtensionsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetExtensionsClient
}

// NewMockVirtualMachineScaleSetExtensionsClient creates a new mock instance
func NewMockVirtualMachineScaleSetExtensionsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetExtensionsClient {
	mock := &MockVirtualMachineScaleSetExtensionsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetExtensionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualMachineScaleSetExtensionsClient) EXPECT() *MockVirtualMachineScaleSetExtensionsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).Client))
}

// CreateOrUpdate mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2, arg3 string, arg4 compute.VirtualMachineScaleSetExtension) (compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) Get(arg0 context.Context, arg1, arg2, arg3, arg4 string) (compute.VirtualMachineScaleSetExtension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// List mocks base method
func (m *MockVirtualMachineScaleSetExtensionsClient) List(arg0 context.Context, arg1, arg2 string) (compute.VirtualMachineScaleSetExtensionListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtensionListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).List), arg0, arg1, arg2)
}

// MockApplicationsClient is a mock of ApplicationsClient interface
type MockApplicationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationsClientMockRecorder
}

// MockApplicationsClientMockRecorder is the mock recorder for MockApplicationsClient
type MockApplicationsClientMockRecorder struct {
	mock *MockApplicationsClient
}

// NewMockApplicationsClient creates a new mock instance
func NewMockApplicationsClient(ctrl *gomock.Controller) *MockApplicationsClient {
	mock := &MockApplicationsClient{ctrl: ctrl}
	mock.recorder = &MockApplicationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationsClient) EXPECT() *MockApplicationsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockApplicationsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockApplicationsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockApplicationsClient)(nil).Client))
}

// Get mocks base method
func (m *MockApplicationsClient) Get(arg0 context.Context, arg1, arg2 string) (managedapplications.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(managedapplications.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockApplicationsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationsClient)(nil).Get), arg0, arg1, arg2)
}

// GetByID mocks base method
func (m *MockApplicationsClient) GetByID(arg0 context.Context, arg1 string) (managedapplications.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(managedapplications.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockApplicationsClientMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockApplicationsClient)(nil).GetByID), arg0, arg1)
}

// ListByResourceGroup mocks base method
func (m *MockApplicationsClient) ListByResourceGroup(arg0 context.Context, arg1 string) (managedapplications.ApplicationListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(managedapplications.ApplicationListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockApplicationsClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockApplicationsClient)(nil).ListByResourceGroup), arg0, arg1)
}

// MockMarketPlaceAgreementsClient is a mock of MarketPlaceAgreementsClient interface
type MockMarketPlaceAgreementsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMarketPlaceAgreementsClientMockRecorder
}

// MockMarketPlaceAgreementsClientMockRecorder is the mock recorder for MockMarketPlaceAgreementsClient
type MockMarketPlaceAgreementsClientMockRecorder struct {
	mock *MockMarketPlaceAgreementsClient
}

// NewMockMarketPlaceAgreementsClient creates a new mock instance
func NewMockMarketPlaceAgreementsClient(ctrl *gomock.Controller) *MockMarketPlaceAgreementsClient {
	mock := &MockMarketPlaceAgreementsClient{ctrl: ctrl}
	mock.recorder = &MockMarketPlaceAgreementsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketPlaceAgreementsClient) EXPECT() *MockMarketPlaceAgreementsClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMarketPlaceAgreementsClient) Create(arg0 context.Context, arg1, arg2, arg3 string, arg4 marketplaceordering.AgreementTerms) (marketplaceordering.AgreementTerms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(marketplaceordering.AgreementTerms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMarketPlaceAgreementsClientMockRecorder) Create(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMarketPlaceAgreementsClient)(nil).Create), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method
func (m *MockMarketPlaceAgreementsClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (marketplaceordering.AgreementTerms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(marketplaceordering.AgreementTerms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMarketPlaceAgreementsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMarketPlaceAgreementsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// MockDeploymentsClient is a mock of DeploymentsClient interface
type MockDeploymentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentsClientMockRecorder
}

// MockDeploymentsClientMockRecorder is the mock recorder for MockDeploymentsClient
type MockDeploymentsClientMockRecorder struct {
	mock *MockDeploymentsClient
}

// NewMockDeploymentsClient creates a new mock instance
func NewMockDeploymentsClient(ctrl *gomock.Controller) *MockDeploymentsClient {
	mock := &MockDeploymentsClient{ctrl: ctrl}
	mock.recorder = &MockDeploymentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeploymentsClient) EXPECT() *MockDeploymentsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockDeploymentsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockDeploymentsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockDeploymentsClient)(nil).Client))
}

// CreateOrUpdate mocks base method
func (m *MockDeploymentsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 resources.Deployment) (resources.DeploymentsCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(resources.DeploymentsCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockDeploymentsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockDeploymentsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// DeploymentClient mocks base method
func (m *MockDeploymentsClient) DeploymentClient() resources.DeploymentsClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentClient")
	ret0, _ := ret[0].(resources.DeploymentsClient)
	return ret0
}

// DeploymentClient indicates an expected call of DeploymentClient
func (mr *MockDeploymentsClientMockRecorder) DeploymentClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentClient", reflect.TypeOf((*MockDeploymentsClient)(nil).DeploymentClient))
}

// MockAccountsClient is a mock of AccountsClient interface
type MockAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsClientMockRecorder
}

// MockAccountsClientMockRecorder is the mock recorder for MockAccountsClient
type MockAccountsClientMockRecorder struct {
	mock *MockAccountsClient
}

// NewMockAccountsClient creates a new mock instance
func NewMockAccountsClient(ctrl *gomock.Controller) *MockAccountsClient {
	mock := &MockAccountsClient{ctrl: ctrl}
	mock.recorder = &MockAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountsClient) EXPECT() *MockAccountsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockAccountsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockAccountsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockAccountsClient)(nil).Client))
}

// Create mocks base method
func (m *MockAccountsClient) Create(arg0 context.Context, arg1, arg2 string, arg3 storage.AccountCreateParameters) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockAccountsClientMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountsClient)(nil).Create), arg0, arg1, arg2, arg3)
}

// ListByResourceGroup mocks base method
func (m *MockAccountsClient) ListByResourceGroup(arg0 context.Context, arg1 string) (storage.AccountListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(storage.AccountListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockAccountsClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockAccountsClient)(nil).ListByResourceGroup), arg0, arg1)
}

// ListKeys mocks base method
func (m *MockAccountsClient) ListKeys(arg0 context.Context, arg1, arg2 string) (storage.AccountListKeysResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].(storage.AccountListKeysResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockAccountsClientMockRecorder) ListKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockAccountsClient)(nil).ListKeys), arg0, arg1, arg2)
}

// MockKeyVaultClient is a mock of KeyVaultClient interface
type MockKeyVaultClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultClientMockRecorder
}

// MockKeyVaultClientMockRecorder is the mock recorder for MockKeyVaultClient
type MockKeyVaultClientMockRecorder struct {
	mock *MockKeyVaultClient
}

// NewMockKeyVaultClient creates a new mock instance
func NewMockKeyVaultClient(ctrl *gomock.Controller) *MockKeyVaultClient {
	mock := &MockKeyVaultClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyVaultClient) EXPECT() *MockKeyVaultClientMockRecorder {
	return m.recorder
}

// GetSecret mocks base method
func (m *MockKeyVaultClient) GetSecret(arg0 context.Context, arg1, arg2, arg3 string) (keyvault.SecretBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(keyvault.SecretBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockKeyVaultClientMockRecorder) GetSecret(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockKeyVaultClient)(nil).GetSecret), arg0, arg1, arg2, arg3)
}

// ImportCertificate mocks base method
func (m *MockKeyVaultClient) ImportCertificate(arg0 context.Context, arg1, arg2 string, arg3 keyvault.CertificateImportParameters) (keyvault.CertificateBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportCertificate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(keyvault.CertificateBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportCertificate indicates an expected call of ImportCertificate
func (mr *MockKeyVaultClientMockRecorder) ImportCertificate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCertificate", reflect.TypeOf((*MockKeyVaultClient)(nil).ImportCertificate), arg0, arg1, arg2, arg3)
}

// MockVaultMgmtClient is a mock of VaultMgmtClient interface
type MockVaultMgmtClient struct {
	ctrl     *gomock.Controller
	recorder *MockVaultMgmtClientMockRecorder
}

// MockVaultMgmtClientMockRecorder is the mock recorder for MockVaultMgmtClient
type MockVaultMgmtClientMockRecorder struct {
	mock *MockVaultMgmtClient
}

// NewMockVaultMgmtClient creates a new mock instance
func NewMockVaultMgmtClient(ctrl *gomock.Controller) *MockVaultMgmtClient {
	mock := &MockVaultMgmtClient{ctrl: ctrl}
	mock.recorder = &MockVaultMgmtClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVaultMgmtClient) EXPECT() *MockVaultMgmtClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockVaultMgmtClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 keyvault0.VaultCreateOrUpdateParameters) (keyvault0.Vault, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(keyvault0.Vault)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockVaultMgmtClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVaultMgmtClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// MockZonesClient is a mock of ZonesClient interface
type MockZonesClient struct {
	ctrl     *gomock.Controller
	recorder *MockZonesClientMockRecorder
}

// MockZonesClientMockRecorder is the mock recorder for MockZonesClient
type MockZonesClientMockRecorder struct {
	mock *MockZonesClient
}

// NewMockZonesClient creates a new mock instance
func NewMockZonesClient(ctrl *gomock.Controller) *MockZonesClient {
	mock := &MockZonesClient{ctrl: ctrl}
	mock.recorder = &MockZonesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockZonesClient) EXPECT() *MockZonesClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockZonesClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 dns.Zone, arg4, arg5 string) (dns.Zone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(dns.Zone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockZonesClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockZonesClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Delete mocks base method
func (m *MockZonesClient) Delete(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockZonesClientMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockZonesClient)(nil).Delete), arg0, arg1, arg2, arg3)
}

// MockRecordSetsClient is a mock of RecordSetsClient interface
type MockRecordSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRecordSetsClientMockRecorder
}

// MockRecordSetsClientMockRecorder is the mock recorder for MockRecordSetsClient
type MockRecordSetsClientMockRecorder struct {
	mock *MockRecordSetsClient
}

// NewMockRecordSetsClient creates a new mock instance
func NewMockRecordSetsClient(ctrl *gomock.Controller) *MockRecordSetsClient {
	mock := &MockRecordSetsClient{ctrl: ctrl}
	mock.recorder = &MockRecordSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRecordSetsClient) EXPECT() *MockRecordSetsClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockRecordSetsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2, arg3 string, arg4 dns.RecordType, arg5 dns.RecordSet, arg6, arg7 string) (dns.RecordSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(dns.RecordSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockRecordSetsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockRecordSetsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// Delete mocks base method
func (m *MockRecordSetsClient) Delete(arg0 context.Context, arg1, arg2, arg3 string, arg4 dns.RecordType, arg5 string) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockRecordSetsClientMockRecorder) Delete(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRecordSetsClient)(nil).Delete), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Get mocks base method
func (m *MockRecordSetsClient) Get(arg0 context.Context, arg1, arg2, arg3 string, arg4 dns.RecordType) (dns.RecordSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(dns.RecordSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRecordSetsClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRecordSetsClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}
