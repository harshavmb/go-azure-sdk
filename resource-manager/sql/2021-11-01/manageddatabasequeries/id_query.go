package manageddatabasequeries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = QueryId{}

// QueryId is a struct representing the Resource ID for a Query
type QueryId struct {
	SubscriptionId      string
	ResourceGroupName   string
	ManagedInstanceName string
	DatabaseName        string
	QueryId             string
}

// NewQueryID returns a new QueryId struct
func NewQueryID(subscriptionId string, resourceGroupName string, managedInstanceName string, databaseName string, queryId string) QueryId {
	return QueryId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		ManagedInstanceName: managedInstanceName,
		DatabaseName:        databaseName,
		QueryId:             queryId,
	}
}

// ParseQueryID parses 'input' into a QueryId
func ParseQueryID(input string) (*QueryId, error) {
	parser := resourceids.NewParserFromResourceIdType(QueryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := QueryId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseQueryIDInsensitively parses 'input' case-insensitively into a QueryId
// note: this method should only be used for API response data and not user input
func ParseQueryIDInsensitively(input string) (*QueryId, error) {
	parser := resourceids.NewParserFromResourceIdType(QueryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := QueryId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *QueryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedInstanceName, ok = input.Parsed["managedInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", input)
	}

	if id.DatabaseName, ok = input.Parsed["databaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseName", input)
	}

	if id.QueryId, ok = input.Parsed["queryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "queryId", input)
	}

	return nil
}

// ValidateQueryID checks that 'input' can be parsed as a Query ID
func ValidateQueryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQueryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Query ID
func (id QueryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/databases/%s/queries/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.DatabaseName, id.QueryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Query ID
func (id QueryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticQueries", "queries", "queries"),
		resourceids.UserSpecifiedSegment("queryId", "queryIdValue"),
	}
}

// String returns a human-readable description of this Query ID
func (id QueryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Query: %q", id.QueryId),
	}
	return fmt.Sprintf("Query (%s)", strings.Join(components, "\n"))
}
