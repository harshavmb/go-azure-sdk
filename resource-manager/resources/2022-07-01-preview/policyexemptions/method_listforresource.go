package policyexemptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListForResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicyExemption
}

type ListForResourceCompleteResult struct {
	Items []PolicyExemption
}

type ListForResourceOperationOptions struct {
	Filter *string
}

func DefaultListForResourceOperationOptions() ListForResourceOperationOptions {
	return ListForResourceOperationOptions{}
}

func (o ListForResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListForResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListForResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListForResource ...
func (c PolicyExemptionsClient) ListForResource(ctx context.Context, id commonids.ScopeId, options ListForResourceOperationOptions) (result ListForResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Authorization/policyExemptions", id.ID()),
		OptionsObject: options,
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
		Values *[]PolicyExemption `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListForResourceComplete retrieves all the results into a single object
func (c PolicyExemptionsClient) ListForResourceComplete(ctx context.Context, id commonids.ScopeId, options ListForResourceOperationOptions) (ListForResourceCompleteResult, error) {
	return c.ListForResourceCompleteMatchingPredicate(ctx, id, options, PolicyExemptionOperationPredicate{})
}

// ListForResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyExemptionsClient) ListForResourceCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options ListForResourceOperationOptions, predicate PolicyExemptionOperationPredicate) (result ListForResourceCompleteResult, err error) {
	items := make([]PolicyExemption, 0)

	resp, err := c.ListForResource(ctx, id, options)
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

	result = ListForResourceCompleteResult{
		Items: items,
	}
	return
}
