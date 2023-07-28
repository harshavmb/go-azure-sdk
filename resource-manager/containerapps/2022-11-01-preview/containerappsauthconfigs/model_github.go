package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitHub struct {
	Enabled      *bool               `json:"enabled,omitempty"`
	Login        *LoginScopes        `json:"login,omitempty"`
	Registration *ClientRegistration `json:"registration,omitempty"`
}
