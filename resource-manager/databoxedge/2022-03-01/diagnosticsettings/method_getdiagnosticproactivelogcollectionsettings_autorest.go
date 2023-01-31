package diagnosticsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDiagnosticProactiveLogCollectionSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticProactiveLogCollectionSettings
}

// GetDiagnosticProactiveLogCollectionSettings ...
func (c DiagnosticSettingsClient) GetDiagnosticProactiveLogCollectionSettings(ctx context.Context, id DataBoxEdgeDeviceId) (result GetDiagnosticProactiveLogCollectionSettingsOperationResponse, err error) {
	req, err := c.preparerForGetDiagnosticProactiveLogCollectionSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "GetDiagnosticProactiveLogCollectionSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "GetDiagnosticProactiveLogCollectionSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDiagnosticProactiveLogCollectionSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "GetDiagnosticProactiveLogCollectionSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDiagnosticProactiveLogCollectionSettings prepares the GetDiagnosticProactiveLogCollectionSettings request.
func (c DiagnosticSettingsClient) preparerForGetDiagnosticProactiveLogCollectionSettings(ctx context.Context, id DataBoxEdgeDeviceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/diagnosticProactiveLogCollectionSettings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDiagnosticProactiveLogCollectionSettings handles the response to the GetDiagnosticProactiveLogCollectionSettings request. The method always
// closes the http.Response Body.
func (c DiagnosticSettingsClient) responderForGetDiagnosticProactiveLogCollectionSettings(resp *http.Response) (result GetDiagnosticProactiveLogCollectionSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
