package v2019_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/account"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/consumerinvitation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/dataset"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/datasetmapping"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/emailregistration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/invitation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/share"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/sharesubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/synchronizationsetting"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/trigger"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Account                *account.AccountClient
	ConsumerInvitation     *consumerinvitation.ConsumerInvitationClient
	DataSet                *dataset.DataSetClient
	DataSetMapping         *datasetmapping.DataSetMappingClient
	EmailRegistration      *emailregistration.EmailRegistrationClient
	Invitation             *invitation.InvitationClient
	Share                  *share.ShareClient
	ShareSubscription      *sharesubscription.ShareSubscriptionClient
	SynchronizationSetting *synchronizationsetting.SynchronizationSettingClient
	Trigger                *trigger.TriggerClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountClient, err := account.NewAccountClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Account client: %+v", err)
	}
	configureFunc(accountClient.Client)

	consumerInvitationClient, err := consumerinvitation.NewConsumerInvitationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ConsumerInvitation client: %+v", err)
	}
	configureFunc(consumerInvitationClient.Client)

	dataSetClient, err := dataset.NewDataSetClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataSet client: %+v", err)
	}
	configureFunc(dataSetClient.Client)

	dataSetMappingClient, err := datasetmapping.NewDataSetMappingClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataSetMapping client: %+v", err)
	}
	configureFunc(dataSetMappingClient.Client)

	emailRegistrationClient, err := emailregistration.NewEmailRegistrationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EmailRegistration client: %+v", err)
	}
	configureFunc(emailRegistrationClient.Client)

	invitationClient, err := invitation.NewInvitationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Invitation client: %+v", err)
	}
	configureFunc(invitationClient.Client)

	shareClient, err := share.NewShareClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Share client: %+v", err)
	}
	configureFunc(shareClient.Client)

	shareSubscriptionClient, err := sharesubscription.NewShareSubscriptionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ShareSubscription client: %+v", err)
	}
	configureFunc(shareSubscriptionClient.Client)

	synchronizationSettingClient, err := synchronizationsetting.NewSynchronizationSettingClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationSetting client: %+v", err)
	}
	configureFunc(synchronizationSettingClient.Client)

	triggerClient, err := trigger.NewTriggerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Trigger client: %+v", err)
	}
	configureFunc(triggerClient.Client)

	return &Client{
		Account:                accountClient,
		ConsumerInvitation:     consumerInvitationClient,
		DataSet:                dataSetClient,
		DataSetMapping:         dataSetMappingClient,
		EmailRegistration:      emailRegistrationClient,
		Invitation:             invitationClient,
		Share:                  shareClient,
		ShareSubscription:      shareSubscriptionClient,
		SynchronizationSetting: synchronizationSettingClient,
		Trigger:                triggerClient,
	}, nil
}
