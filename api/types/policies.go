package types

type PolicyDefinition struct {
	ID           string                 `json:"id" header:"ID"`
	Name         string                 `json:"name" header:"NAME"`
	Description  string                 `json:"description" header:"DESCRIPTION"`
	Definition   string                 `json:"definition" header:"DEFINITION" show:"nolist"`
	Parameters   map[string]interface{} `json:"parameters" header:"PARAMETERS" show:"nolist"`
	Builtin      bool                   `json:"builtin" header:"BUILTIN" show:"nolist"`
	ResourceType string                 `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type PolicyAssignment struct {
	ID              string                 `json:"id" header:"ID"`
	Name            string                 `json:"name" header:"NAME"`
	Description     string                 `json:"description" header:"DESCRIPTION"`
	State           string                 `json:"state" header:"STATE"`
	RemoteID        string                 `json:"remote_id" header:"REMOTE_ID"`
	CloudAccountID  string                 `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	ErrorEventID    string                 `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist,noshow"`
	ParameterValues map[string]interface{} `json:"parameter_values" header:"PARAMETER_VALUES" show:"nolist"`
	DefinitionID    string                 `json:"definition_id" header:"DEFINITION_ID"`
	ResellerApplied bool                   `json:"reseller_applied" header:"RESELLER_APPLIED" show:"nolist"`
	ResourceType    string                 `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
