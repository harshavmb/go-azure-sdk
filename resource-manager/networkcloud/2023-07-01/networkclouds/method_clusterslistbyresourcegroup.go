package networkclouds

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

type ClustersListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Cluster
}

type ClustersListByResourceGroupCompleteResult struct {
	Items []Cluster
}

// ClustersListByResourceGroup ...
func (c NetworkcloudsClient) ClustersListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ClustersListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/clusters", id.ID()),
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
		Values *[]Cluster `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ClustersListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) ClustersListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ClustersListByResourceGroupCompleteResult, error) {
	return c.ClustersListByResourceGroupCompleteMatchingPredicate(ctx, id, ClusterOperationPredicate{})
}

// ClustersListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) ClustersListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate ClusterOperationPredicate) (result ClustersListByResourceGroupCompleteResult, err error) {
	items := make([]Cluster, 0)

	resp, err := c.ClustersListByResourceGroup(ctx, id)
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

	result = ClustersListByResourceGroupCompleteResult{
		Items: items,
	}
	return
}
