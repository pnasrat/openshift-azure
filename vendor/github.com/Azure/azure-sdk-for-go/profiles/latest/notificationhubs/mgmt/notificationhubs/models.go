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

package notificationhubs

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/notificationhubs/mgmt/2017-04-01/notificationhubs"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessRights = original.AccessRights

const (
	Listen AccessRights = original.Listen
	Manage AccessRights = original.Manage
	Send   AccessRights = original.Send
)

type NamespaceType = original.NamespaceType

const (
	Messaging       NamespaceType = original.Messaging
	NotificationHub NamespaceType = original.NotificationHub
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Free     SkuName = original.Free
	Standard SkuName = original.Standard
)

type AdmCredential = original.AdmCredential
type AdmCredentialProperties = original.AdmCredentialProperties
type ApnsCredential = original.ApnsCredential
type ApnsCredentialProperties = original.ApnsCredentialProperties
type BaiduCredential = original.BaiduCredential
type BaiduCredentialProperties = original.BaiduCredentialProperties
type BaseClient = original.BaseClient
type CheckAvailabilityParameters = original.CheckAvailabilityParameters
type CheckAvailabilityResult = original.CheckAvailabilityResult
type Client = original.Client
type CreateOrUpdateParameters = original.CreateOrUpdateParameters
type DebugSendResponse = original.DebugSendResponse
type DebugSendResult = original.DebugSendResult
type ErrorResponse = original.ErrorResponse
type GcmCredential = original.GcmCredential
type GcmCredentialProperties = original.GcmCredentialProperties
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type MpnsCredential = original.MpnsCredential
type MpnsCredentialProperties = original.MpnsCredentialProperties
type NamespaceCreateOrUpdateParameters = original.NamespaceCreateOrUpdateParameters
type NamespaceListResult = original.NamespaceListResult
type NamespaceListResultIterator = original.NamespaceListResultIterator
type NamespaceListResultPage = original.NamespaceListResultPage
type NamespacePatchParameters = original.NamespacePatchParameters
type NamespaceProperties = original.NamespaceProperties
type NamespaceResource = original.NamespaceResource
type NamespacesClient = original.NamespacesClient
type NamespacesDeleteFuture = original.NamespacesDeleteFuture
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PatchParameters = original.PatchParameters
type PnsCredentialsProperties = original.PnsCredentialsProperties
type PnsCredentialsResource = original.PnsCredentialsResource
type PolicykeyResource = original.PolicykeyResource
type Properties = original.Properties
type Resource = original.Resource
type ResourceListKeys = original.ResourceListKeys
type ResourceType = original.ResourceType
type SharedAccessAuthorizationRuleCreateOrUpdateParameters = original.SharedAccessAuthorizationRuleCreateOrUpdateParameters
type SharedAccessAuthorizationRuleListResult = original.SharedAccessAuthorizationRuleListResult
type SharedAccessAuthorizationRuleListResultIterator = original.SharedAccessAuthorizationRuleListResultIterator
type SharedAccessAuthorizationRuleListResultPage = original.SharedAccessAuthorizationRuleListResultPage
type SharedAccessAuthorizationRuleProperties = original.SharedAccessAuthorizationRuleProperties
type SharedAccessAuthorizationRuleResource = original.SharedAccessAuthorizationRuleResource
type Sku = original.Sku
type SubResource = original.SubResource
type WnsCredential = original.WnsCredential
type WnsCredentialProperties = original.WnsCredentialProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(getNextPage)
}
func NewNamespaceListResultIterator(page NamespaceListResultPage) NamespaceListResultIterator {
	return original.NewNamespaceListResultIterator(page)
}
func NewNamespaceListResultPage(getNextPage func(context.Context, NamespaceListResult) (NamespaceListResult, error)) NamespaceListResultPage {
	return original.NewNamespaceListResultPage(getNextPage)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSharedAccessAuthorizationRuleListResultIterator(page SharedAccessAuthorizationRuleListResultPage) SharedAccessAuthorizationRuleListResultIterator {
	return original.NewSharedAccessAuthorizationRuleListResultIterator(page)
}
func NewSharedAccessAuthorizationRuleListResultPage(getNextPage func(context.Context, SharedAccessAuthorizationRuleListResult) (SharedAccessAuthorizationRuleListResult, error)) SharedAccessAuthorizationRuleListResultPage {
	return original.NewSharedAccessAuthorizationRuleListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleNamespaceTypeValues() []NamespaceType {
	return original.PossibleNamespaceTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
