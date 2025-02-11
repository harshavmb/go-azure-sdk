package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListWebJobsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebJob
}

type ListWebJobsSlotCompleteResult struct {
	Items []WebJob
}

// ListWebJobsSlot ...
func (c WebAppsClient) ListWebJobsSlot(ctx context.Context, id SlotId) (result ListWebJobsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/webJobs", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]WebJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWebJobsSlotComplete retrieves all the results into a single object
func (c WebAppsClient) ListWebJobsSlotComplete(ctx context.Context, id SlotId) (ListWebJobsSlotCompleteResult, error) {
	return c.ListWebJobsSlotCompleteMatchingPredicate(ctx, id, WebJobOperationPredicate{})
}

// ListWebJobsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListWebJobsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate WebJobOperationPredicate) (result ListWebJobsSlotCompleteResult, err error) {
	items := make([]WebJob, 0)

	resp, err := c.ListWebJobsSlot(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListWebJobsSlotCompleteResult{
		Items: items,
	}
	return
}
