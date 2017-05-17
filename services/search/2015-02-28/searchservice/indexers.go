package searchservice

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
    "github.com/satori/uuid"
    "github.com/Azure/go-autorest/autorest/validation"
)

// IndexersClient is the client that can be used to manage and query indexes
// and documents, as well as manage other resources, on an Azure Search
// service.
type IndexersClient struct {
    ManagementClient
}
// NewIndexersClient creates an instance of the IndexersClient client.
func NewIndexersClient() IndexersClient {
        return NewIndexersClientWithBaseURI(DefaultBaseURI, )
        }

// NewIndexersClientWithBaseURI creates an instance of the IndexersClient
// client.
    func NewIndexersClientWithBaseURI(baseURI string, ) IndexersClient {
        return IndexersClient{ NewWithBaseURI(baseURI, )}
    }

// Create creates a new Azure Search indexer.
//
// indexer is the definition of the indexer to create. clientRequestID is
// tracking ID sent with the request to help with debugging.
func (client IndexersClient) Create(indexer Indexer, clientRequestID *uuid.UUID) (result Indexer, err error) {
    if err := validation.Validate([]validation.Validation{
    { TargetValue: indexer,
     Constraints: []validation.Constraint{	{Target: "indexer.Name", Name: validation.Null, Rule: true, Chain: nil },
    	{Target: "indexer.DataSourceName", Name: validation.Null, Rule: true, Chain: nil },
    	{Target: "indexer.TargetIndexName", Name: validation.Null, Rule: true, Chain: nil },
    	{Target: "indexer.Schedule", Name: validation.Null, Rule: false ,
    Chain: []validation.Constraint{	{Target: "indexer.Schedule.Interval", Name: validation.Null, Rule: true, Chain: nil },
    }}}}}); err != nil {
    return result, validation.NewErrorWithValidationError(err, "searchservice.IndexersClient","Create")
}

    req, err := client.CreatePreparer(indexer, clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Create", nil , "Failure preparing request")
        return
    }

    resp, err := client.CreateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Create", resp, "Failure sending request")
        return
    }

    result, err = client.CreateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Create", resp, "Failure responding to request")
    }

    return
}

// CreatePreparer prepares the Create request.
func (client IndexersClient) CreatePreparer(indexer Indexer, clientRequestID *uuid.UUID) (*http.Request, error) {
        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/indexers"),
                        autorest.WithJSON(indexer),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) CreateSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client IndexersClient) CreateResponder(resp *http.Response) (result Indexer, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// CreateOrUpdate creates a new Azure Search indexer or updates an indexer if
// it already exists.
//
// indexerName is the name of the indexer to create or update. indexer is the
// definition of the indexer to create or update. clientRequestID is tracking
// ID sent with the request to help with debugging.
func (client IndexersClient) CreateOrUpdate(indexerName string, indexer Indexer, clientRequestID *uuid.UUID) (result Indexer, err error) {
    if err := validation.Validate([]validation.Validation{
    { TargetValue: indexer,
     Constraints: []validation.Constraint{	{Target: "indexer.Name", Name: validation.Null, Rule: true, Chain: nil },
    	{Target: "indexer.DataSourceName", Name: validation.Null, Rule: true, Chain: nil },
    	{Target: "indexer.TargetIndexName", Name: validation.Null, Rule: true, Chain: nil },
    	{Target: "indexer.Schedule", Name: validation.Null, Rule: false ,
    Chain: []validation.Constraint{	{Target: "indexer.Schedule.Interval", Name: validation.Null, Rule: true, Chain: nil },
    }}}}}); err != nil {
    return result, validation.NewErrorWithValidationError(err, "searchservice.IndexersClient","CreateOrUpdate")
}

    req, err := client.CreateOrUpdatePreparer(indexerName, indexer, clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "CreateOrUpdate", nil , "Failure preparing request")
        return
    }

    resp, err := client.CreateOrUpdateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "CreateOrUpdate", resp, "Failure sending request")
        return
    }

    result, err = client.CreateOrUpdateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "CreateOrUpdate", resp, "Failure responding to request")
    }

    return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IndexersClient) CreateOrUpdatePreparer(indexerName string, indexer Indexer, clientRequestID *uuid.UUID) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "indexerName": autorest.Encode("path",indexerName),
    }

        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/indexers('{indexerName}')",pathParameters),
                        autorest.WithJSON(indexer),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IndexersClient) CreateOrUpdateResponder(resp *http.Response) (result Indexer, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusCreated,http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Delete deletes an Azure Search indexer.
//
// indexerName is the name of the indexer to delete. clientRequestID is
// tracking ID sent with the request to help with debugging.
func (client IndexersClient) Delete(indexerName string, clientRequestID *uuid.UUID) (result autorest.Response, err error) {
    req, err := client.DeletePreparer(indexerName, clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Delete", nil , "Failure preparing request")
        return
    }

    resp, err := client.DeleteSender(req)
    if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Delete", resp, "Failure sending request")
        return
    }

    result, err = client.DeleteResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Delete", resp, "Failure responding to request")
    }

    return
}

// DeletePreparer prepares the Delete request.
func (client IndexersClient) DeletePreparer(indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "indexerName": autorest.Encode("path",indexerName),
    }

        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsDelete(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/indexers('{indexerName}')",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) DeleteSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IndexersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNotFound,http.StatusNoContent),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Get retrieves an indexer definition from Azure Search.
//
// indexerName is the name of the indexer to retrieve. clientRequestID is
// tracking ID sent with the request to help with debugging.
func (client IndexersClient) Get(indexerName string, clientRequestID *uuid.UUID) (result Indexer, err error) {
    req, err := client.GetPreparer(indexerName, clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Get", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Get", resp, "Failure sending request")
        return
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client IndexersClient) GetPreparer(indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "indexerName": autorest.Encode("path",indexerName),
    }

        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/indexers('{indexerName}')",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IndexersClient) GetResponder(resp *http.Response) (result Indexer, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// GetStatus returns the current status and execution history of an indexer.
//
// indexerName is the name of the indexer for which to retrieve status.
// clientRequestID is tracking ID sent with the request to help with debugging.
func (client IndexersClient) GetStatus(indexerName string, clientRequestID *uuid.UUID) (result IndexerExecutionInfo, err error) {
    req, err := client.GetStatusPreparer(indexerName, clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "GetStatus", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetStatusSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "GetStatus", resp, "Failure sending request")
        return
    }

    result, err = client.GetStatusResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "GetStatus", resp, "Failure responding to request")
    }

    return
}

// GetStatusPreparer prepares the GetStatus request.
func (client IndexersClient) GetStatusPreparer(indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "indexerName": autorest.Encode("path",indexerName),
    }

        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/indexers('{indexerName}')/search.status",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// GetStatusSender sends the GetStatus request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) GetStatusSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetStatusResponder handles the response to the GetStatus request. The method always
// closes the http.Response Body.
func (client IndexersClient) GetStatusResponder(resp *http.Response) (result IndexerExecutionInfo, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// List lists all indexers available for an Azure Search service.
//
// clientRequestID is tracking ID sent with the request to help with debugging.
func (client IndexersClient) List(clientRequestID *uuid.UUID) (result IndexerListResult, err error) {
    req, err := client.ListPreparer(clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "List", nil , "Failure preparing request")
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "List", resp, "Failure sending request")
        return
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "List", resp, "Failure responding to request")
    }

    return
}

// ListPreparer prepares the List request.
func (client IndexersClient) ListPreparer(clientRequestID *uuid.UUID) (*http.Request, error) {
        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/indexers"),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) ListSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IndexersClient) ListResponder(resp *http.Response) (result IndexerListResult, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Reset resets the change tracking state associated with an Azure Search
// indexer.
//
// indexerName is the name of the indexer to reset. clientRequestID is tracking
// ID sent with the request to help with debugging.
func (client IndexersClient) Reset(indexerName string, clientRequestID *uuid.UUID) (result autorest.Response, err error) {
    req, err := client.ResetPreparer(indexerName, clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Reset", nil , "Failure preparing request")
        return
    }

    resp, err := client.ResetSender(req)
    if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Reset", resp, "Failure sending request")
        return
    }

    result, err = client.ResetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Reset", resp, "Failure responding to request")
    }

    return
}

// ResetPreparer prepares the Reset request.
func (client IndexersClient) ResetPreparer(indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "indexerName": autorest.Encode("path",indexerName),
    }

        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/indexers('{indexerName}')/search.reset",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) ResetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client IndexersClient) ResetResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Run runs an Azure Search indexer on-demand.
//
// indexerName is the name of the indexer to run. clientRequestID is tracking
// ID sent with the request to help with debugging.
func (client IndexersClient) Run(indexerName string, clientRequestID *uuid.UUID) (result autorest.Response, err error) {
    req, err := client.RunPreparer(indexerName, clientRequestID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Run", nil , "Failure preparing request")
        return
    }

    resp, err := client.RunSender(req)
    if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Run", resp, "Failure sending request")
        return
    }

    result, err = client.RunResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Run", resp, "Failure responding to request")
    }

    return
}

// RunPreparer prepares the Run request.
func (client IndexersClient) RunPreparer(indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "indexerName": autorest.Encode("path",indexerName),
    }

        const APIVersion = "2015-02-28"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/indexers('{indexerName}')/search.run",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    return preparer.Prepare(&http.Request{})
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) RunSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client IndexersClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
    result.Response = resp
    return
}

