package containerappsrevisionreplicas

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ReplicaId{}

// ReplicaId is a struct representing the Resource ID for a Replica
type ReplicaId struct {
	SubscriptionId    string
	ResourceGroupName string
	ContainerAppName  string
	RevisionName      string
	ReplicaName       string
}

// NewReplicaID returns a new ReplicaId struct
func NewReplicaID(subscriptionId string, resourceGroupName string, containerAppName string, revisionName string, replicaName string) ReplicaId {
	return ReplicaId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ContainerAppName:  containerAppName,
		RevisionName:      revisionName,
		ReplicaName:       replicaName,
	}
}

// ParseReplicaID parses 'input' into a ReplicaId
func ParseReplicaID(input string) (*ReplicaId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicaId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReplicaIDInsensitively parses 'input' case-insensitively into a ReplicaId
// note: this method should only be used for API response data and not user input
func ParseReplicaIDInsensitively(input string) (*ReplicaId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReplicaId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicaId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ContainerAppName, ok = input.Parsed["containerAppName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "containerAppName", input)
	}

	if id.RevisionName, ok = input.Parsed["revisionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "revisionName", input)
	}

	if id.ReplicaName, ok = input.Parsed["replicaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicaName", input)
	}

	return nil
}

// ValidateReplicaID checks that 'input' can be parsed as a Replica ID
func ValidateReplicaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replica ID
func (id ReplicaId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/containerApps/%s/revisions/%s/replicas/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ContainerAppName, id.RevisionName, id.ReplicaName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replica ID
func (id ReplicaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticContainerApps", "containerApps", "containerApps"),
		resourceids.UserSpecifiedSegment("containerAppName", "containerAppValue"),
		resourceids.StaticSegment("staticRevisions", "revisions", "revisions"),
		resourceids.UserSpecifiedSegment("revisionName", "revisionValue"),
		resourceids.StaticSegment("staticReplicas", "replicas", "replicas"),
		resourceids.UserSpecifiedSegment("replicaName", "replicaValue"),
	}
}

// String returns a human-readable description of this Replica ID
func (id ReplicaId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Container App Name: %q", id.ContainerAppName),
		fmt.Sprintf("Revision Name: %q", id.RevisionName),
		fmt.Sprintf("Replica Name: %q", id.ReplicaName),
	}
	return fmt.Sprintf("Replica (%s)", strings.Join(components, "\n"))
}
