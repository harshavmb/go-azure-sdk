package modelversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CodeConfiguration struct {
	CodeId        *string `json:"codeId,omitempty"`
	ScoringScript string  `json:"scoringScript"`
}
