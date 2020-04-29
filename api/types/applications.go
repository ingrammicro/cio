package types

// CAT
type CloudApplicationTemplate struct {
	ID           string  `json:"id" header:"ID"`
	Name         string  `json:"name" header:"NAME"`
	Version      string  `json:"version" header:"VERSION"`
	Global       bool    `json:"global" header:"GLOBAL"`
	UploadURL    string  `json:"upload_url,omitempty" header:"UPLOAD_URL" show:"noshow,nolist"`
	VendorID     string  `json:"vendor_id" header:"VENDOR_ID"`
	Inputs       []Input `json:"inputs" header:"INPUTS" show:"nolist"`
	IsMock       bool    `json:"is_mock" header:"IS_MOCK" show:"nolist"`
	ResourceType string  `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type Input struct {
	Name         string                 `json:"name" header:"NAME"`
	Description  string                 `json:"description" header:"DESCRIPTION"`
	Type         string                 `json:"type" header:"TYPE"`
	Required     bool                   `json:"required" header:"REQUIRED"`
	Dependencies map[string]interface{} `json:"dependencies" header:"DEPENDENCIES"`
}

// CAD
type CloudApplicationDeployment struct {
	ID           string `json:"id" header:"ID"`
	Name         string `json:"name" header:"NAME"`
	Namespace    string `json:"namespace" header:"NAMESPACE" show:"nolist"`
	Value        string `json:"value" header:"VALUE"`
	CatID        string `json:"cat_id" header:"CAT_ID" show:"nolist"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

// tosca task
type CloudApplicationDeploymentTask struct {
	ID           string `json:"id" header:"ID"`
	Type         string `json:"type" header:"TYPE"`
	LabelName    string `json:"label_name" header:"LABEL_NAME"`
	LabelID      string `json:"label_id" header:"LABEL_ID"`
	State        string `json:"state" header:"STATE"`
	ErrorMessage string `json:"error_message" header:"ERROR_MESSAGE" show:"nolist"`
	Outputs      string `json:"outputs" header:"OUTPUTS" show:"nolist"`
	UserID       string `json:"user_id" header:"USER_ID"`
	ArchiveID    string `json:"archive_id" header:"ARCHIVE_ID"`
	DeploymentID string `json:"deployment_id" header:"DEPLOYMENT_ID"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
