package dataversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoDeleteCondition string

const (
	AutoDeleteConditionCreatedGreaterThan      AutoDeleteCondition = "CreatedGreaterThan"
	AutoDeleteConditionLastAccessedGreaterThan AutoDeleteCondition = "LastAccessedGreaterThan"
)

func PossibleValuesForAutoDeleteCondition() []string {
	return []string{
		string(AutoDeleteConditionCreatedGreaterThan),
		string(AutoDeleteConditionLastAccessedGreaterThan),
	}
}

func (s *AutoDeleteCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoDeleteCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoDeleteCondition(input string) (*AutoDeleteCondition, error) {
	vals := map[string]AutoDeleteCondition{
		"createdgreaterthan":      AutoDeleteConditionCreatedGreaterThan,
		"lastaccessedgreaterthan": AutoDeleteConditionLastAccessedGreaterThan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoDeleteCondition(input)
	return &out, nil
}

type DataImportSourceType string

const (
	DataImportSourceTypeDatabase   DataImportSourceType = "database"
	DataImportSourceTypeFileSystem DataImportSourceType = "file_system"
)

func PossibleValuesForDataImportSourceType() []string {
	return []string{
		string(DataImportSourceTypeDatabase),
		string(DataImportSourceTypeFileSystem),
	}
}

func (s *DataImportSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataImportSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataImportSourceType(input string) (*DataImportSourceType, error) {
	vals := map[string]DataImportSourceType{
		"database":    DataImportSourceTypeDatabase,
		"file_system": DataImportSourceTypeFileSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataImportSourceType(input)
	return &out, nil
}

type DataType string

const (
	DataTypeMltable   DataType = "mltable"
	DataTypeUriFile   DataType = "uri_file"
	DataTypeUriFolder DataType = "uri_folder"
)

func PossibleValuesForDataType() []string {
	return []string{
		string(DataTypeMltable),
		string(DataTypeUriFile),
		string(DataTypeUriFolder),
	}
}

func (s *DataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataType(input string) (*DataType, error) {
	vals := map[string]DataType{
		"mltable":    DataTypeMltable,
		"uri_file":   DataTypeUriFile,
		"uri_folder": DataTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataType(input)
	return &out, nil
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

func (s *ListViewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseListViewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
}

type ProtectionLevel string

const (
	ProtectionLevelAll  ProtectionLevel = "All"
	ProtectionLevelNone ProtectionLevel = "None"
)

func PossibleValuesForProtectionLevel() []string {
	return []string{
		string(ProtectionLevelAll),
		string(ProtectionLevelNone),
	}
}

func (s *ProtectionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionLevel(input string) (*ProtectionLevel, error) {
	vals := map[string]ProtectionLevel{
		"all":  ProtectionLevelAll,
		"none": ProtectionLevelNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionLevel(input)
	return &out, nil
}
