// Copyright (c) 2017-2022 Ingram Micro Inc.

package types

type ServerPlan struct {
	ID                  string  `json:"id"                    header:"ID"`
	Name                string  `json:"name"                  header:"NAME"`
	Memory              int     `json:"memory"                header:"MEMORY"`
	CPUs                float32 `json:"cpus"                  header:"CPUS"`
	Storage             int     `json:"storage"               header:"STORAGE"`
	LocationID          string  `json:"location_id"           header:"LOCATION_ID"`
	LocationName        string  `header:"LOCATION_NAME"`
	RealmID             string  `json:"realm_id"              header:"REALM_ID"`
	RealmProviderName   string  `json:"realm_provider_name"   header:"REALM_PROVIDER_NAME"`
	FlavourProviderName string  `json:"flavour_provider_name" header:"FLAVOUR_PROVIDER_NAME"`
	CloudProviderID     string  `json:"cloud_provider_id,omitempty"   header:"CLOUD_PROVIDER_ID" show:"noshow"`
	CloudProviderName   string  `header:"CLOUD_PROVIDER_NAME" show:"noshow"`
	ResourceType        string  `json:"resource_type"         header:"RESOURCE_TYPE" show:"nolist"`
}
