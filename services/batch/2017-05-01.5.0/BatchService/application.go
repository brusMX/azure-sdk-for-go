package batchservice

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
    "github.com/Azure/go-autorest/autorest/date"
    "github.com/Azure/go-autorest/autorest/validation"
)

// ApplicationClient is the a client for issuing REST requests to the Azure
// Batch service.
type ApplicationClient struct {
    ManagementClient
}
// NewApplicationClient creates an instance of the ApplicationClient client.
func NewApplicationClient() ApplicationClient {
        return NewApplicationClientWithBaseURI(DefaultBaseURI, )
        }

// NewApplicationClientWithBaseURI creates an instance of the ApplicationClient
// client.
    func NewApplicationClientWithBaseURI(baseURI string, ) ApplicationClient {
        return ApplicationClient{ NewWithBaseURI(baseURI, )}
    }

// Get sends the get request.
//
// applicationID is the ID of the application. timeout is the maximum time that
// the server can spend processing the request, in seconds. The default is 30
// seconds. clientRequestID is the caller-generated request identity, in the
// form of a GUID with no decoration such as curly braces, e.g.
// 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0. returnClientRequestID is whether the
// server should return the client-request-id in the response. ocpDate is the
// time the request was issued. Client libraries typically set this to the
// current system clock time; set it explicitly if you are calling the REST API
// directly.
func (client ApplicationClient) Get(applicationID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result ApplicationSummary, err error) {
    req, err := client.GetPreparer(applicationID, timeout, clientRequestID, returnClientRequestID, ocpDate)
    if err != nil {
        err = autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "Get", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "Get", resp, "Failure sending request")
        return
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client ApplicationClient) GetPreparer(applicationID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "applicationId": autorest.Encode("path",applicationID),
    }

        const APIVersion = "2017-05-01.5.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/applications/{applicationId}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    if returnClientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("return-client-request-id",autorest.String(returnClientRequestID)))
    }
    if ocpDate != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("ocp-date",autorest.String(ocpDate)))
    }
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationClient) GetResponder(resp *http.Response) (result ApplicationSummary, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// List this operation returns only applications and versions that are
// available for use on compute nodes; that is, that can be used in an
// application package reference. For administrator information about
// applications and versions that are not yet available to compute nodes, use
// the Azure portal or the Azure Resource Manager API.
//
// maxResults is the maximum number of items to return in the response. A
// maximum of 1000 applications can be returned. timeout is the maximum time
// that the server can spend processing the request, in seconds. The default is
// 30 seconds. clientRequestID is the caller-generated request identity, in the
// form of a GUID with no decoration such as curly braces, e.g.
// 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0. returnClientRequestID is whether the
// server should return the client-request-id in the response. ocpDate is the
// time the request was issued. Client libraries typically set this to the
// current system clock time; set it explicitly if you are calling the REST API
// directly.
func (client ApplicationClient) List(maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result ApplicationListResult, err error) {
    if err := validation.Validate([]validation.Validation{
    { TargetValue: maxResults,
     Constraints: []validation.Constraint{	{Target: "maxResults", Name: validation.Null, Rule: false ,
    Chain: []validation.Constraint{	{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil },
    	{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil },
    }}}}}); err != nil {
    return result, validation.NewErrorWithValidationError(err, "batchservice.ApplicationClient","List")
}

    req, err := client.ListPreparer(maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
    if err != nil {
        err = autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "List", nil , "Failure preparing request")
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "List", resp, "Failure sending request")
        return
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "List", resp, "Failure responding to request")
    }

    return
}

// ListPreparer prepares the List request.
func (client ApplicationClient) ListPreparer(maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
        const APIVersion = "2017-05-01.5.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if maxResults != nil {
        queryParameters["maxresults"] = autorest.Encode("query",*maxResults)
    }
    if timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/applications"),
                        autorest.WithQueryParameters(queryParameters))
    if clientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("client-request-id",autorest.String(clientRequestID)))
    }
    if returnClientRequestID != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("return-client-request-id",autorest.String(returnClientRequestID)))
    }
    if ocpDate != nil {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithHeader("ocp-date",autorest.String(ocpDate)))
    }
    return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) ListSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationClient) ListResponder(resp *http.Response) (result ApplicationListResult, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// ListNextResults retrieves the next set of results, if any.
func (client ApplicationClient) ListNextResults(lastResults ApplicationListResult) (result ApplicationListResult, err error) {
    req, err := lastResults.ApplicationListResultPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "List", nil , "Failure preparing next results request")
    }
    if req == nil {
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "List", resp, "Failure sending next results request")
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "batchservice.ApplicationClient", "List", resp, "Failure responding to next results request")
    }

    return
}

