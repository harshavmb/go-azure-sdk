package cosmosdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GremlinGraphResource struct {
	ConflictResolutionPolicy *ConflictResolutionPolicy `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int64                    `json:"defaultTtl,omitempty"`
	Id                       string                    `json:"id"`
	IndexingPolicy           *IndexingPolicy           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey    `json:"partitionKey,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `json:"uniqueKeyPolicy,omitempty"`
}
