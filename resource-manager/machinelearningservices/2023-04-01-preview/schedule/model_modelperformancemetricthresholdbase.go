package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelPerformanceMetricThresholdBase interface {
}

// RawModelPerformanceMetricThresholdBaseImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawModelPerformanceMetricThresholdBaseImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalModelPerformanceMetricThresholdBaseImplementation(input []byte) (ModelPerformanceMetricThresholdBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ModelPerformanceMetricThresholdBase into map[string]interface: %+v", err)
	}

	value, ok := temp["modelType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Classification") {
		var out ClassificationModelPerformanceMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClassificationModelPerformanceMetricThreshold: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Regression") {
		var out RegressionModelPerformanceMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegressionModelPerformanceMetricThreshold: %+v", err)
		}
		return out, nil
	}

	out := RawModelPerformanceMetricThresholdBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
