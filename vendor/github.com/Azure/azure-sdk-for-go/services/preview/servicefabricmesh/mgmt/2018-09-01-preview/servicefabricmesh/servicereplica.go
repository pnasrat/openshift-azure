package servicefabricmesh

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ServiceReplicaClient is the service Fabric Mesh Management Client
type ServiceReplicaClient struct {
	BaseClient
}

// NewServiceReplicaClient creates an instance of the ServiceReplicaClient client.
func NewServiceReplicaClient(subscriptionID string) ServiceReplicaClient {
	return NewServiceReplicaClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceReplicaClientWithBaseURI creates an instance of the ServiceReplicaClient client.
func NewServiceReplicaClientWithBaseURI(baseURI string, subscriptionID string) ServiceReplicaClient {
	return ServiceReplicaClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the information about the service replica with the given name. The information include the description and
// other properties of the service replica.
// Parameters:
// resourceGroupName - azure resource group name
// applicationResourceName - the identity of the application.
// serviceResourceName - the identity of the service.
// replicaName - service Fabric replica name.
func (client ServiceReplicaClient) Get(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string) (result ServiceReplicaDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceReplicaClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, applicationResourceName, serviceResourceName, replicaName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceReplicaClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
		"replicaName":             replicaName,
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serviceResourceName":     serviceResourceName,
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}/replicas/{replicaName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceReplicaClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceReplicaClient) GetResponder(resp *http.Response) (result ServiceReplicaDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the information about all replicas of a given service of an application. The information includes the
// runtime properties of the replica instance.
// Parameters:
// resourceGroupName - azure resource group name
// applicationResourceName - the identity of the application.
// serviceResourceName - the identity of the service.
func (client ServiceReplicaClient) List(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string) (result ServiceReplicaDescriptionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceReplicaClient.List")
		defer func() {
			sc := -1
			if result.srdl.Response.Response != nil {
				sc = result.srdl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, applicationResourceName, serviceResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.srdl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "List", resp, "Failure sending request")
		return
	}

	result.srdl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceReplicaClient) ListPreparer(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serviceResourceName":     serviceResourceName,
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}/replicas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceReplicaClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceReplicaClient) ListResponder(resp *http.Response) (result ServiceReplicaDescriptionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ServiceReplicaClient) listNextResults(ctx context.Context, lastResults ServiceReplicaDescriptionList) (result ServiceReplicaDescriptionList, err error) {
	req, err := lastResults.serviceReplicaDescriptionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceReplicaClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceReplicaClient) ListComplete(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string) (result ServiceReplicaDescriptionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceReplicaClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, applicationResourceName, serviceResourceName)
	return
}
