package types

type Realm struct {
	ID              string `json:"id" header:"ID"`
	Name            string `json:"name" header:"NAME"`
	LocationID      string `json:"location_id" header:"LOCATION_ID"`
	CloudProviderID string `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	ProviderName    string `json:"provider_name,omitempty" header:"PROVIDER_NAME"`
	Deprecated      bool   `json:"deprecated,omitempty" header:"DEPRECATED" show:"nolist,noshow"`
	ResourceType    string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
