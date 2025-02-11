package nodereports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetContentOperationResponse struct {
	HttpResponse *http.Response
	Model        *interface{}
}

// GetContent ...
func (c NodeReportsClient) GetContent(ctx context.Context, id ReportId) (result GetContentOperationResponse, err error) {
	req, err := c.preparerForGetContent(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "GetContent", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "GetContent", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetContent(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "GetContent", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetContent prepares the GetContent request.
func (c NodeReportsClient) preparerForGetContent(ctx context.Context, id ReportId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/content", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetContent handles the response to the GetContent request. The method always
// closes the http.Response Body.
func (c NodeReportsClient) responderForGetContent(resp *http.Response) (result GetContentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
