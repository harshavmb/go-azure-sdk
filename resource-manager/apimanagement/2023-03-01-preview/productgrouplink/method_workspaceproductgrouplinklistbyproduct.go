package productgrouplink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceProductGroupLinkListByProductOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProductGroupLinkContract
}

type WorkspaceProductGroupLinkListByProductCompleteResult struct {
	Items []ProductGroupLinkContract
}

type WorkspaceProductGroupLinkListByProductOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultWorkspaceProductGroupLinkListByProductOperationOptions() WorkspaceProductGroupLinkListByProductOperationOptions {
	return WorkspaceProductGroupLinkListByProductOperationOptions{}
}

func (o WorkspaceProductGroupLinkListByProductOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WorkspaceProductGroupLinkListByProductOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o WorkspaceProductGroupLinkListByProductOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// WorkspaceProductGroupLinkListByProduct ...
func (c ProductGroupLinkClient) WorkspaceProductGroupLinkListByProduct(ctx context.Context, id WorkspaceProductId, options WorkspaceProductGroupLinkListByProductOperationOptions) (result WorkspaceProductGroupLinkListByProductOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/groupLinks", id.ID()),
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
		Values *[]ProductGroupLinkContract `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceProductGroupLinkListByProductComplete retrieves all the results into a single object
func (c ProductGroupLinkClient) WorkspaceProductGroupLinkListByProductComplete(ctx context.Context, id WorkspaceProductId, options WorkspaceProductGroupLinkListByProductOperationOptions) (WorkspaceProductGroupLinkListByProductCompleteResult, error) {
	return c.WorkspaceProductGroupLinkListByProductCompleteMatchingPredicate(ctx, id, options, ProductGroupLinkContractOperationPredicate{})
}

// WorkspaceProductGroupLinkListByProductCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProductGroupLinkClient) WorkspaceProductGroupLinkListByProductCompleteMatchingPredicate(ctx context.Context, id WorkspaceProductId, options WorkspaceProductGroupLinkListByProductOperationOptions, predicate ProductGroupLinkContractOperationPredicate) (result WorkspaceProductGroupLinkListByProductCompleteResult, err error) {
	items := make([]ProductGroupLinkContract, 0)

	resp, err := c.WorkspaceProductGroupLinkListByProduct(ctx, id, options)
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

	result = WorkspaceProductGroupLinkListByProductCompleteResult{
		Items: items,
	}
	return
}
