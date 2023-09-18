package replicationprotecteditems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwitchProviderProviderSpecificInput interface {
}

// RawSwitchProviderProviderSpecificInputImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSwitchProviderProviderSpecificInputImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalSwitchProviderProviderSpecificInputImplementation(input []byte) (SwitchProviderProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SwitchProviderProviderSpecificInput into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "InMageAzureV2") {
		var out InMageAzureV2SwitchProviderProviderInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageAzureV2SwitchProviderProviderInput: %+v", err)
		}
		return out, nil
	}

	out := RawSwitchProviderProviderSpecificInputImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
