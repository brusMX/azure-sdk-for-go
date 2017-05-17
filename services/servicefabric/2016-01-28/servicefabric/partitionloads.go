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

// PartitionLoadsClient is the client for the PartitionLoads methods of the
// Servicefabric service.
type PartitionLoadsClient struct {
    ManagementClient
}
// NewPartitionLoadsClient creates an instance of the PartitionLoadsClient
// client.
func NewPartitionLoadsClient(timeout *int32) PartitionLoadsClient {
        return NewPartitionLoadsClientWithBaseURI(DefaultBaseURI, timeout)
        }

// NewPartitionLoadsClientWithBaseURI creates an instance of the
// PartitionLoadsClient client.
    func NewPartitionLoadsClientWithBaseURI(baseURI string, timeout *int32) PartitionLoadsClient {
        return PartitionLoadsClient{ NewWithBaseURI(baseURI, timeout)}
    }

// Reset reset partition loads
//
// partitionID is the id of the partition
func (client PartitionLoadsClient) Reset(partitionID string) (result String, err error) {
    req, err := client.ResetPreparer(partitionID)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.PartitionLoadsClient", "Reset", nil , "Failure preparing request")
        return
    }

    resp, err := client.ResetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.PartitionLoadsClient", "Reset", resp, "Failure sending request")
        return
    }

    result, err = client.ResetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.PartitionLoadsClient", "Reset", resp, "Failure responding to request")
    }

    return
}

// ResetPreparer prepares the Reset request.
func (client PartitionLoadsClient) ResetPreparer(partitionID string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "partitionId": partitionID,
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
                        autorest.WithPathParameters("/Partitions/{partitionId}/$/ResetLoad",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client PartitionLoadsClient) ResetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client PartitionLoadsClient) ResetResponder(resp *http.Response) (result String, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

