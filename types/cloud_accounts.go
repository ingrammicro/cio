// Copyright (c) 2017-2021 Ingram Micro Inc.

package types

type CloudAccount struct {
	ID                                  string `json:"id" header:"ID"`
	Name                                string `json:"name" header:"NAME"`
	SubscriptionID                      string `json:"subscription_id" header:"SUBSCRIPTION_D"`
	RemoteID                            string `json:"remote_id" header:"REMOTE_ID"`
	CloudProviderID                     string `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName                   string `header:"CLOUD_PROVIDER_NAME"`
	SupportsImporting                   bool   `json:"supports_importing" header:"SUPPORTS_IMPORTING" show:"nolist"`
	SupportsImportingVPCs bool `json:"supports_importing_vpcs" header:"SUPPORTS_IMPORTING_VPCS" show:"nolist"`
	SupportsImportingFloatingIPs bool `json:"supports_importing_floating_ips" header:"SUPPORTS_IMPORTING_FLOATING_IPS" show:"nolist"`
	SupportsImportingVolumes bool `json:"supports_importing_volumes" header:"SUPPORTS_IMPORTING_VOLUMES" show:"nolist"`
	SupportsImportingPolicies bool `json:"supports_importing_policies" header:"SUPPORTS_IMPORTING_POLICIES" show:"nolist"`
	SupportsImportingKubernetesClusters bool `json:"supports_importing_kubernetes_clusters" header:"SUPPORTS_IMPORTING_KUBERNETES_CLUSTERS" show:"nolist"`
	State                               string `json:"state" header:"STATE"`
	ErrorEventID                        string `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType                        string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type RequiredCredentials interface{}
