package updateruns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StopOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type StopOperationOptions struct {
	IfMatch *string
}

func DefaultStopOperationOptions() StopOperationOptions {
	return StopOperationOptions{}
}

func (o StopOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o StopOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// Stop ...
func (c UpdateRunsClient) Stop(ctx context.Context, id UpdateRunId, options StopOperationOptions) (result StopOperationResponse, err error) {
	req, err := c.preparerForStop(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "updateruns.UpdateRunsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStop(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "updateruns.UpdateRunsClient", "Stop", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StopThenPoll performs Stop then polls until it's completed
func (c UpdateRunsClient) StopThenPoll(ctx context.Context, id UpdateRunId, options StopOperationOptions) error {
	result, err := c.Stop(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing Stop: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Stop: %+v", err)
	}

	return nil
}

// preparerForStop prepares the Stop request.
func (c UpdateRunsClient) preparerForStop(ctx context.Context, id UpdateRunId, options StopOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForStop sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (c UpdateRunsClient) senderForStop(ctx context.Context, req *http.Request) (future StopOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
