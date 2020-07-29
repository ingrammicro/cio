package types

import "time"

type ServerImportCandidate struct {
	ID                          string                      `json:"id" header:"ID"`
	Name                        string                      `json:"name" header:"NAME"`
	Fqdn                        string                      `json:"fqdn" header:"FQDN"`
	State                       string                      `json:"state" header:"STATE"`
	RemoteID                    string                      `json:"remote_id" header:"REMOTE_ID"`
	Image                       Image                       `json:"image"`
	ServerPlanID                string                      `json:"server_plan_id" header:"SERVER_PLAN_ID"`
	PublicIP                    string                      `json:"public_ip" header:"PUBLIC_IP"`
	CloudAccountID              string                      `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	FloatingIPsImportCandidates []FloatingIPImportCandidate `json:"floating_ip_import_candidates" header:"FLOATING_IPS" show:"nolist"`
	VolumesImportCandidates     []VolumeImportCandidate     `json:"volume_import_candidates" header:"VOLUMES" show:"nolist"`
	ResourceType                string                      `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type Image struct {
	Name     string `json:"name" header:"IMAGE_NAME"`
	RemoteID string `json:"remote_id" header:"IMAGE_REMOTE_ID"`
}

type FloatingIPImportCandidate struct {
	ID               string `json:"id" header:"ID"`
	Name             string `json:"name" header:"NAME"`
	Address          string `json:"address" header:"ADDRESS"`
	RemoteID         string `json:"remote_id" header:"REMOTE_ID"`
	CloudAccountID   string `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	RealmID          string `json:"realm_id" header:"REALM_ID"`
	AttachedServerID string `json:"attached_server_id" header:"ATTACHED_SERVER_ID"`
	ResourceType     string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type VolumeImportCandidate struct {
	ID               string `json:"id" header:"ID"`
	Name             string `json:"name" header:"NAME"`
	Size             int    `json:"size" header:"SIZE"`
	RemoteID         string `json:"remote_id" header:"REMOTE_ID"`
	CloudAccountID   string `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	StoragePlanID    string `json:"storage_plan_id" header:"STORAGE_PLAN_ID"`
	AttachedServerID string `json:"attached_server_id" header:"ATTACHED_SERVER_ID"`
	ResourceType     string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type VpcImportCandidate struct {
	ID                string            `json:"id" header:"ID"`
	Name              string            `json:"name" header:"NAME"`
	Cidr              string            `json:"cidr" header:"CIDR"`
	RemoteID          string            `json:"remote_id" header:"REMOTE_ID"`
	CloudAccountID    string            `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	RealmID           string            `json:"realm_id" header:"REALM_ID"`
	SubnetsCandidates []SubnetCandidate `json:"subnet_candidates" header:"SUBNETS" show:"nolist"`
	ResourceType      string            `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type SubnetCandidate struct {
	ID                     string    `json:"_id" header:"ID"`
	Name                   string    `json:"name" header:"NAME"`
	Cidr                   string    `json:"cidr" header:"CIDR"`
	State                  string    `json:"state" header:"STATE"`
	RemoteID               string    `json:"remote_id" header:"REMOTE_ID"`
	Type                   string    `json:"type" header:"TYPE"`
	CreatedAt              time.Time `json:"created_at,omitempty" header:"CREATED_AT" show:"nolist"`
	UpdateAt               time.Time `json:"updated_at,omitempty" header:"UPDATED_AT" show:"nolist"`
	ServerCreationDisabled bool      `json:"server_creation_disabled" header:"SERVER_CREATION_DISABLED" show:"nolist"`
}
