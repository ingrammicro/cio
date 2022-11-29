// Copyright (c) 2017-2022 Ingram Micro Inc.

package types

// CloudSpecificExtensionTemplate CSE Template
type CloudSpecificExtensionTemplate struct {
	ID           string                 `json:"id"            header:"ID"`
	Name         string                 `json:"name"          header:"NAME"`
	Global       bool                   `json:"global"        header:"GLOBAL"`
	Definition   string                 `json:"definition"    header:"DEFINITION"    show:"nolist"`
	Parameters   map[string]interface{} `json:"parameters"    header:"PARAMETERS"    show:"nolist"`
	Syntax       string                 `json:"syntax"        header:"SYNTAX"`
	ResourceType string                 `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
}

// CloudSpecificExtensionDeployment CSE Deployment
type CloudSpecificExtensionDeployment struct {
	ID              string                 `json:"id"               header:"ID"`
	Name            string                 `json:"name"             header:"NAME"`
	TemplateID      string                 `json:"template_id"      header:"TEMPLATE_ID"`
	State           string                 `json:"state"            header:"STATE"`
	RemoteID        string                 `json:"remote_id"        header:"REMOTE_ID"`
	CloudAccountID  string                 `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	RealmID         string                 `json:"realm_id"         header:"REALM_ID"`
	RealmIDs        []string               `json:"realm_ids"        header:"REALM_IDS"`
	ErrorEventID    string                 `json:"error_event_id"   header:"ERROR_EVENT_ID"   show:"nolist,noshow"`
	ParameterValues map[string]interface{} `json:"parameter_values" header:"PARAMETER_VALUES" show:"nolist"`
	Outputs         map[string]interface{} `json:"outputs"          header:"OUTPUTS"          show:"nolist"`
	ResourceType    string                 `json:"resource_type"    header:"RESOURCE_TYPE"    show:"nolist"`
	LabelableFields
}
