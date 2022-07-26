package topicsauthorizationrule

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopicsDeleteAuthorizationRuleOperationResponse struct {
	HttpResponse *http.Response
}

// TopicsDeleteAuthorizationRule ...
func (c TopicsAuthorizationRuleClient) TopicsDeleteAuthorizationRule(ctx context.Context, id TopicAuthorizationRuleId) (result TopicsDeleteAuthorizationRuleOperationResponse, err error) {
	req, err := c.preparerForTopicsDeleteAuthorizationRule(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "topicsauthorizationrule.TopicsAuthorizationRuleClient", "TopicsDeleteAuthorizationRule", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "topicsauthorizationrule.TopicsAuthorizationRuleClient", "TopicsDeleteAuthorizationRule", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTopicsDeleteAuthorizationRule(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "topicsauthorizationrule.TopicsAuthorizationRuleClient", "TopicsDeleteAuthorizationRule", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTopicsDeleteAuthorizationRule prepares the TopicsDeleteAuthorizationRule request.
func (c TopicsAuthorizationRuleClient) preparerForTopicsDeleteAuthorizationRule(ctx context.Context, id TopicAuthorizationRuleId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTopicsDeleteAuthorizationRule handles the response to the TopicsDeleteAuthorizationRule request. The method always
// closes the http.Response Body.
func (c TopicsAuthorizationRuleClient) responderForTopicsDeleteAuthorizationRule(resp *http.Response) (result TopicsDeleteAuthorizationRuleOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
