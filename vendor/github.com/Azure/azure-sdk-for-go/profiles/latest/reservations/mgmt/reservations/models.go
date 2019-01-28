// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package reservations

import original "github.com/Azure/azure-sdk-for-go/services/reservations/mgmt/2018-06-01/reservations"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppliedScopeType = original.AppliedScopeType

const (
	Shared AppliedScopeType = original.Shared
	Single AppliedScopeType = original.Single
)

type ErrorResponseCode = original.ErrorResponseCode

const (
	ActivateQuoteFailed                           ErrorResponseCode = original.ActivateQuoteFailed
	AppliedScopesNotAssociatedWithCommerceAccount ErrorResponseCode = original.AppliedScopesNotAssociatedWithCommerceAccount
	AuthorizationFailed                           ErrorResponseCode = original.AuthorizationFailed
	BadRequest                                    ErrorResponseCode = original.BadRequest
	BillingCustomerInputError                     ErrorResponseCode = original.BillingCustomerInputError
	BillingError                                  ErrorResponseCode = original.BillingError
	BillingPaymentInstrumentHardError             ErrorResponseCode = original.BillingPaymentInstrumentHardError
	BillingPaymentInstrumentSoftError             ErrorResponseCode = original.BillingPaymentInstrumentSoftError
	BillingScopeIDCannotBeChanged                 ErrorResponseCode = original.BillingScopeIDCannotBeChanged
	BillingTransientError                         ErrorResponseCode = original.BillingTransientError
	CalculatePriceFailed                          ErrorResponseCode = original.CalculatePriceFailed
	CapacityUpdateScopesFailed                    ErrorResponseCode = original.CapacityUpdateScopesFailed
	ClientCertificateThumbprintNotSet             ErrorResponseCode = original.ClientCertificateThumbprintNotSet
	CreateQuoteFailed                             ErrorResponseCode = original.CreateQuoteFailed
	Forbidden                                     ErrorResponseCode = original.Forbidden
	FulfillmentConfigurationError                 ErrorResponseCode = original.FulfillmentConfigurationError
	FulfillmentError                              ErrorResponseCode = original.FulfillmentError
	FulfillmentOutOfStockError                    ErrorResponseCode = original.FulfillmentOutOfStockError
	FulfillmentTransientError                     ErrorResponseCode = original.FulfillmentTransientError
	HTTPMethodNotSupported                        ErrorResponseCode = original.HTTPMethodNotSupported
	InternalServerError                           ErrorResponseCode = original.InternalServerError
	InvalidAccessToken                            ErrorResponseCode = original.InvalidAccessToken
	InvalidFulfillmentRequestParameters           ErrorResponseCode = original.InvalidFulfillmentRequestParameters
	InvalidHealthCheckType                        ErrorResponseCode = original.InvalidHealthCheckType
	InvalidLocationID                             ErrorResponseCode = original.InvalidLocationID
	InvalidRefundQuantity                         ErrorResponseCode = original.InvalidRefundQuantity
	InvalidRequestContent                         ErrorResponseCode = original.InvalidRequestContent
	InvalidRequestURI                             ErrorResponseCode = original.InvalidRequestURI
	InvalidReservationID                          ErrorResponseCode = original.InvalidReservationID
	InvalidReservationOrderID                     ErrorResponseCode = original.InvalidReservationOrderID
	InvalidSingleAppliedScopesCount               ErrorResponseCode = original.InvalidSingleAppliedScopesCount
	InvalidSubscriptionID                         ErrorResponseCode = original.InvalidSubscriptionID
	InvalidTenantID                               ErrorResponseCode = original.InvalidTenantID
	MissingAppliedScopesForSingle                 ErrorResponseCode = original.MissingAppliedScopesForSingle
	MissingTenantID                               ErrorResponseCode = original.MissingTenantID
	NonsupportedAccountID                         ErrorResponseCode = original.NonsupportedAccountID
	NotSpecified                                  ErrorResponseCode = original.NotSpecified
	NotSupportedCountry                           ErrorResponseCode = original.NotSupportedCountry
	NoValidReservationsToReRate                   ErrorResponseCode = original.NoValidReservationsToReRate
	OperationCannotBePerformedInCurrentState      ErrorResponseCode = original.OperationCannotBePerformedInCurrentState
	OperationFailed                               ErrorResponseCode = original.OperationFailed
	PatchValuesSameAsExisting                     ErrorResponseCode = original.PatchValuesSameAsExisting
	PaymentInstrumentNotFound                     ErrorResponseCode = original.PaymentInstrumentNotFound
	PurchaseError                                 ErrorResponseCode = original.PurchaseError
	ReRateOnlyAllowedForEA                        ErrorResponseCode = original.ReRateOnlyAllowedForEA
	ReservationIDNotInReservationOrder            ErrorResponseCode = original.ReservationIDNotInReservationOrder
	ReservationOrderCreationFailed                ErrorResponseCode = original.ReservationOrderCreationFailed
	ReservationOrderIDAlreadyExists               ErrorResponseCode = original.ReservationOrderIDAlreadyExists
	ReservationOrderNotEnabled                    ErrorResponseCode = original.ReservationOrderNotEnabled
	ReservationOrderNotFound                      ErrorResponseCode = original.ReservationOrderNotFound
	RiskCheckFailed                               ErrorResponseCode = original.RiskCheckFailed
	RoleAssignmentCreationFailed                  ErrorResponseCode = original.RoleAssignmentCreationFailed
	ServerTimeout                                 ErrorResponseCode = original.ServerTimeout
	UnauthenticatedRequestsThrottled              ErrorResponseCode = original.UnauthenticatedRequestsThrottled
	UnsupportedReservationTerm                    ErrorResponseCode = original.UnsupportedReservationTerm
)

type InstanceFlexibility = original.InstanceFlexibility

const (
	NotSupported InstanceFlexibility = original.NotSupported
	Off          InstanceFlexibility = original.Off
	On           InstanceFlexibility = original.On
)

type ReservationTerm = original.ReservationTerm

const (
	P1Y ReservationTerm = original.P1Y
	P3Y ReservationTerm = original.P3Y
)

type ReservedResourceType = original.ReservedResourceType

const (
	CosmosDb        ReservedResourceType = original.CosmosDb
	SQLDatabases    ReservedResourceType = original.SQLDatabases
	SuseLinux       ReservedResourceType = original.SuseLinux
	VirtualMachines ReservedResourceType = original.VirtualMachines
)

type StatusCode = original.StatusCode

const (
	StatusCodeActive                 StatusCode = original.StatusCodeActive
	StatusCodeExpired                StatusCode = original.StatusCodeExpired
	StatusCodeMerged                 StatusCode = original.StatusCodeMerged
	StatusCodeNone                   StatusCode = original.StatusCodeNone
	StatusCodePaymentInstrumentError StatusCode = original.StatusCodePaymentInstrumentError
	StatusCodePending                StatusCode = original.StatusCodePending
	StatusCodePurchaseError          StatusCode = original.StatusCodePurchaseError
	StatusCodeSplit                  StatusCode = original.StatusCodeSplit
	StatusCodeSucceeded              StatusCode = original.StatusCodeSucceeded
)

type AppliedReservationList = original.AppliedReservationList
type AppliedReservations = original.AppliedReservations
type AppliedReservationsProperties = original.AppliedReservationsProperties
type BaseClient = original.BaseClient
type Catalog = original.Catalog
type Client = original.Client
type Error = original.Error
type ExtendedErrorInfo = original.ExtendedErrorInfo
type ExtendedStatusInfo = original.ExtendedStatusInfo
type List = original.List
type ListCatalog = original.ListCatalog
type ListIterator = original.ListIterator
type ListPage = original.ListPage
type ListResponse = original.ListResponse
type MergeProperties = original.MergeProperties
type MergePropertiesType = original.MergePropertiesType
type MergeRequest = original.MergeRequest
type OperationClient = original.OperationClient
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationResponse = original.OperationResponse
type OrderClient = original.OrderClient
type OrderList = original.OrderList
type OrderListIterator = original.OrderListIterator
type OrderListPage = original.OrderListPage
type OrderProperties = original.OrderProperties
type OrderResponse = original.OrderResponse
type Patch = original.Patch
type PatchProperties = original.PatchProperties
type Properties = original.Properties
type ReservationMergeFuture = original.ReservationMergeFuture
type ReservationUpdateFuture = original.ReservationUpdateFuture
type Response = original.Response
type SkuName = original.SkuName
type SkuProperty = original.SkuProperty
type SkuRestriction = original.SkuRestriction
type SplitFuture = original.SplitFuture
type SplitProperties = original.SplitProperties
type SplitPropertiesType = original.SplitPropertiesType
type SplitRequest = original.SplitRequest

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewOperationClient() OperationClient {
	return original.NewOperationClient()
}
func NewOperationClientWithBaseURI(baseURI string) OperationClient {
	return original.NewOperationClientWithBaseURI(baseURI)
}
func NewOrderClient() OrderClient {
	return original.NewOrderClient()
}
func NewOrderClientWithBaseURI(baseURI string) OrderClient {
	return original.NewOrderClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return original.PossibleAppliedScopeTypeValues()
}
func PossibleErrorResponseCodeValues() []ErrorResponseCode {
	return original.PossibleErrorResponseCodeValues()
}
func PossibleInstanceFlexibilityValues() []InstanceFlexibility {
	return original.PossibleInstanceFlexibilityValues()
}
func PossibleReservationTermValues() []ReservationTerm {
	return original.PossibleReservationTermValues()
}
func PossibleReservedResourceTypeValues() []ReservedResourceType {
	return original.PossibleReservedResourceTypeValues()
}
func PossibleStatusCodeValues() []StatusCode {
	return original.PossibleStatusCodeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
