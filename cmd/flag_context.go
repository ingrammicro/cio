// Copyright (c) 2017-2022 Ingram Micro Inc.

package cmd

type (
	FlagType = int
)

const (
	Bool FlagType = iota + 1
	String
	Int
	Int64
)

// command line flag names
const (
	AdminPassword                   = "admin-password"
	AppId                           = "app-id"
	CertificateId                   = "certificate-id"
	Chain                           = "chain"
	Cidr                            = "cidr"
	CloudAccountId                  = "cloud-account-id"
	CloudProviderId                 = "cloud-provider-id"
	ClusterPlanId                   = "cluster-plan-id"
	Code                            = "code"
	ConfigurationAttributes         = "configuration-attributes"
	ConfigurationAttributesFromFile = "configuration-attributes-from-file"
	ConfigurationManagementSystem   = "configuration-management-system"
	Content                         = "content"
	CookbookVersions                = "cookbook-versions"
	CpuType                         = "cpu-type"
	DefaultVpcCidr                  = "default-vpc-cidr"
	DefaultVpcCreation              = "default-vpc-creation"
	Definition                      = "definition"
	DefinitionFromFile              = "definition-from-file"
	DefinitionId                    = "definition-id"
	Description                     = "description"
	DesiredNodes                    = "desired-nodes"
	DiskSize                        = "disk-size"
	DomainId                        = "domain-id"
	ExposedCidrs                    = "exposed-cidrs"
	Field                           = "field"
	Filepath                        = "filepath"
	FirewallProfileId               = "firewall-profile-id"
	FloatingIpId                    = "floating-ip-id"
	GenericImageId                  = "generic-image-id"
	HealthCheckInterval             = "health-check-interval"
	HealthCheckPath                 = "health-check-path"
	HealthCheckPort                 = "health-check-port"
	HealthCheckProtocol             = "health-check-protocol"
	HealthCheckThresholdCount       = "health-check-threshold-count"
	Hostname                        = "hostname"
	Id                              = "id"
	Inputs                          = "inputs"
	InputsFromFile                  = "inputs-from-file"
	IpProtocol                      = "ip-protocol"
	IsStackset                      = "is-stackset"
	Label                           = "label"
	Labels                          = "labels"
	Lines                           = "lines"
	ListenerId                      = "listener-id"
	LoadBalancerId                  = "load-balancer-id"
	LocationId                      = "location-id"
	LongTime                        = "longTime"
	MaxNodes                        = "max-nodes"
	MaxPort                         = "max-port"
	MinNodes                        = "min-nodes"
	MinPort                         = "min-port"
	Name                            = "name"
	NodePoolPlanId                  = "node-pool-plan-id"
	Parameters                      = "parameters"
	ParametersFromFile              = "parameters-from-file"
	ParameterValues                 = "parameter-values"
	ParameterValuesFromFile         = "parameter-values-from-file"
	PlanId                          = "plan-id"
	PodsPerNode                     = "pods-per-node"
	Port                            = "port"
	Priority                        = "priority"
	PrivateKey                      = "private-key"
	Privateness                     = "privateness"
	Protocol                        = "protocol"
	Psk                             = "psk"
	PublicAccessIpAddresses         = "public-access-ip-addresses"
	PublicIp                        = "public-ip"
	PublicKey                       = "public-key"
	RealmId                         = "realm-id"
	RealmIds                        = "realm-ids"
	RealmProviderName               = "realm-provider-name"
	ResourceId                      = "resource-id"
	ResourceType                    = "resource-type"
	Rules                           = "rules"
	RunList                         = "run-list"
	ScriptId                        = "script-id"
	ScriptIds                       = "script-ids"
	ServerArrayIds                  = "server-array-ids"
	ServerId                        = "server-id"
	ServerIds                       = "server-ids"
	ServerPlanId                    = "server-plan-id"
	ShortTime                       = "shortTime"
	Size                            = "size"
	SSHProfileId                    = "ssh-profile-id"
	SSHProfileIds                   = "ssh-profile-ids"
	Stickiness                      = "stickiness"
	StoragePlanId                   = "storage-plan-id"
	SubnetId                        = "subnet-id"
	Syntax                          = "syntax"
	TargetGroupId                   = "target-group-id"
	TemplateId                      = "template-id"
	Time                            = "time"
	Ttl                             = "ttl"
	Type                            = "type"
	Values                          = "values"
	Version                         = "version"
	VpcId                           = "vpc-id"
	VpnPlanId                       = "vpn-plan-id"
	Weight                          = "weight"
)

type FlagContext struct {
	cmd          *CommandContext
	Name         string
	Shorthand    string
	Usage        string
	Type         FlagType
	DefaultValue interface{}
	Required     bool
	Hidden       bool
}
