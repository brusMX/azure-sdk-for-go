package servicefabric

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
)

// ServicesClient is the client for the Services methods of the Servicefabric
// service.
type ServicesClient struct {
    ManagementClient
}
// NewServicesClient creates an instance of the ServicesClient client.
func NewServicesClient(timeout *int32) ServicesClient {
        return NewServicesClientWithBaseURI(DefaultBaseURI, timeout)
        }

// NewServicesClientWithBaseURI creates an instance of the ServicesClient
// client.
    func NewServicesClientWithBaseURI(baseURI string, timeout *int32) ServicesClient {
        return ServicesClient{ NewWithBaseURI(baseURI, timeout)}
    }

// Create create services
//
// applicationName is the name of the application createServiceDescription is
// the description of the service
func (client ServicesClient) Create(applicationName string, createServiceDescription CreateServiceDescription) (result String, err error) {
    req, err := client.CreatePreparer(applicationName, createServiceDescription)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Create", nil , "Failure preparing request")
        return
    }

    resp, err := client.CreateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Create", resp, "Failure sending request")
        return
    }

    result, err = client.CreateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Create", resp, "Failure responding to request")
    }

    return
}

// CreatePreparer prepares the Create request.
func (client ServicesClient) CreatePreparer(applicationName string, createServiceDescription CreateServiceDescription) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "applicationName": applicationName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Applications/{applicationName}/$/GetServices/$/Create",pathParameters),
                        autorest.WithJSON(createServiceDescription),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) CreateSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServicesClient) CreateResponder(resp *http.Response) (result String, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Get get services
//
// applicationName is the name of the application serviceName is the name of
// the service
func (client ServicesClient) Get(applicationName string, serviceName string) (result Service, err error) {
    req, err := client.GetPreparer(applicationName, serviceName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Get", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Get", resp, "Failure sending request")
        return
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client ServicesClient) GetPreparer(applicationName string, serviceName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "applicationName": applicationName,
    "serviceName": serviceName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Applications/{applicationName}/$/GetServices/{serviceName}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServicesClient) GetResponder(resp *http.Response) (result Service, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// List list services
//
// applicationName is the name of the application
func (client ServicesClient) List(applicationName string) (result ServiceList, err error) {
    req, err := client.ListPreparer(applicationName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "List", nil , "Failure preparing request")
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "List", resp, "Failure sending request")
        return
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "List", resp, "Failure responding to request")
    }

    return
}

// ListPreparer prepares the List request.
func (client ServicesClient) ListPreparer(applicationName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "applicationName": applicationName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Applications/{applicationName}/$/GetServices",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) ListSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServicesClient) ListResponder(resp *http.Response) (result ServiceList, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Remove remove services
//
// serviceName is the name of the service
func (client ServicesClient) Remove(serviceName string) (result String, err error) {
    req, err := client.RemovePreparer(serviceName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Remove", nil , "Failure preparing request")
        return
    }

    resp, err := client.RemoveSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Remove", resp, "Failure sending request")
        return
    }

    result, err = client.RemoveResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Remove", resp, "Failure responding to request")
    }

    return
}

// RemovePreparer prepares the Remove request.
func (client ServicesClient) RemovePreparer(serviceName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "serviceName": serviceName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Services/{serviceName}/$/Delete",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) RemoveSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client ServicesClient) RemoveResponder(resp *http.Response) (result String, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Resolve resolve services
//
// serviceName is the name of the service partitionKeyType is the type of the
// partition key partitionKeyValue is the value of the partition key
// previousRspVersion is the version of the previous rsp
func (client ServicesClient) Resolve(serviceName string, partitionKeyType *int32, partitionKeyValue string, previousRspVersion string) (result ResolvedServicePartition, err error) {
    req, err := client.ResolvePreparer(serviceName, partitionKeyType, partitionKeyValue, previousRspVersion)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Resolve", nil , "Failure preparing request")
        return
    }

    resp, err := client.ResolveSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Resolve", resp, "Failure sending request")
        return
    }

    result, err = client.ResolveResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Resolve", resp, "Failure responding to request")
    }

    return
}

// ResolvePreparer prepares the Resolve request.
func (client ServicesClient) ResolvePreparer(serviceName string, partitionKeyType *int32, partitionKeyValue string, previousRspVersion string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "serviceName": serviceName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if partitionKeyType != nil {
        queryParameters["PartitionKeyType"] = autorest.Encode("query",*partitionKeyType)
    }
    if len(partitionKeyValue) > 0 {
        queryParameters["PartitionKeyValue"] = autorest.Encode("query",partitionKeyValue)
    }
    if len(previousRspVersion) > 0 {
        queryParameters["PreviousRspVersion"] = autorest.Encode("query",previousRspVersion)
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Services/{serviceName}/$/ResolvePartition",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// ResolveSender sends the Resolve request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) ResolveSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ResolveResponder handles the response to the Resolve request. The method always
// closes the http.Response Body.
func (client ServicesClient) ResolveResponder(resp *http.Response) (result ResolvedServicePartition, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Update update services
//
// serviceName is the name of the service updateServiceDescription is the
// description of the service update
func (client ServicesClient) Update(serviceName string, updateServiceDescription UpdateServiceDescription) (result String, err error) {
    req, err := client.UpdatePreparer(serviceName, updateServiceDescription)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Update", nil , "Failure preparing request")
        return
    }

    resp, err := client.UpdateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Update", resp, "Failure sending request")
        return
    }

    result, err = client.UpdateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Update", resp, "Failure responding to request")
    }

    return
}

// UpdatePreparer prepares the Update request.
func (client ServicesClient) UpdatePreparer(serviceName string, updateServiceDescription UpdateServiceDescription) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "serviceName": serviceName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Services/{serviceName}/$/Update",pathParameters),
                        autorest.WithJSON(updateServiceDescription),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) UpdateSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServicesClient) UpdateResponder(resp *http.Response) (result String, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

