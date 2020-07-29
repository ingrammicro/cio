package types

type LoadBalancer struct {
	ID                 string `json:"id" header:"ID"`
	Name               string `json:"name" header:"NAME"`
	State              string `json:"state" header:"STATE"`
	RemoteID           string `json:"remote_id" header:"REMOTE_ID"`
	CloudAccountID     string `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	RealmID            string `json:"realm_id" header:"REALM_ID"`
	VpcID              string `json:"vpc_id,omitempty" header:"VPC_ID"`
	LoadBalancerPlanID string `json:"load_balancer_plan_id" header:"LOAD_BALANCER_PLAN_ID"`
	DnsName            string `json:"dns_name" header:"DNS_NAME"`
	GlobalState        string `json:"global_state" header:"GLOBAL_STATE"`
	ErrorEventID       string `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType       string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
}

type LoadBalancerPlan struct {
	ID                                   string                 `json:"id" header:"ID"`
	Name                                 string                 `json:"name" header:"NAME"`
	CloudProviderID                      string                 `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName                    string                 `header:"CLOUD_PROVIDER_NAME"`
	RealmID                              string                 `json:"realm_id" header:"REALM_ID"`
	RealmProviderName                    string                 `json:"realm_provider_name,omitempty" header:"REALM_PROVIDER_NAME"`
	FlavourProviderName                  string                 `json:"flavour_provider_name,omitempty" header:"FLAVOUR_PROVIDER_NAME"`
	Protocols                            []string               `json:"protocols" header:"PROTOCOLS"`
	HealthCheckProtocols                 []string               `json:"health_check_protocols" header:"HEALTH_CHECK_PROTOCOLS"`
	RuleFields                           []string               `json:"rule_fields" header:"RULE_FIELDS" show:"nolist"`
	HealthCheckIntervalValidValues       map[string]interface{} `json:"health_check_interval_valid_values" header:"HEALTH_CHECK_INTERVAL_VALID_VALUES" show:"nolist"`
	HealthCheckTimeoutValidValues        map[string]interface{} `json:"health_check_timeout_valid_values" header:"HEALTH_CHECK_TIMEOUT_VALID_VALUES" show:"nolist"`
	HealthCheckThresholdCountValidValues map[string]interface{} `json:"health_check_threshold_count_valid_values" header:"HEALTH_CHECK_THRESHOLD_COUNT_VALID_VALUES" show:"nolist"`
	Deprecated                           bool                   `json:"deprecated,omitempty" header:"DEPRECATED" show:"nolist,noshow"`
	ResourceType                         string                 `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type TargetGroup struct {
	ID                        string `json:"id" header:"ID"`
	Name                      string `json:"name" header:"NAME"`
	State                     string `json:"state" header:"STATE"`
	RemoteID                  string `json:"remote_id" header:"REMOTE_ID"`
	Protocol                  string `json:"protocol" header:"PROTOCOL"`
	Port                      int    `json:"port" header:"PORT"`
	Stickiness                bool   `json:"stickiness" header:"STICKINESS" show:"nolist"`
	HealthCheckProtocol       string `json:"health_check_protocol" header:"HEALTH_CHECK_PROTOCOL" show:"nolist"`
	HealthCheckPort           int    `json:"health_check_port" header:"HEALTH_CHECK_PORT" show:"nolist"`
	HealthCheckInterval       int    `json:"health_check_interval" header:"HEALTH_CHECK_INTERVAL" show:"nolist"`
	HealthCheckTimeout        int    `json:"health_check_timeout" header:"HEALTH_CHECK_TIMEOUT" show:"nolist"`
	HealthCheckThresholdCount int    `json:"health_check_threshold_count" header:"HEALTH_CHECK_THRESHOLD_COUNT" show:"nolist"`
	HealthCheckPath           string `json:"health_check_path" header:"health_check_path" show:"nolist"`
	LoadBalancerID            string `json:"load_balancer_id" header:"LOAD_BALANCER_ID"`
	ErrorEventID              string `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType              string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type Target struct {
	ID           string `json:"id" header:"ID"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE"`
}

type Listener struct {
	ID                   string `json:"id" header:"ID"`
	State                string `json:"state" header:"STATE"`
	RemoteID             string `json:"remote_id" header:"REMOTE_ID"`
	Protocol             string `json:"protocol" header:"PROTOCOL"`
	Port                 int    `json:"port" header:"PORT"`
	LoadBalancerID       string `json:"load_balancer_id" header:"LOAD_BALANCER_ID"`
	CertificateID        string `json:"certificate_id" header:"CERTIFICATE_ID"`
	DefaultTargetGroupID string `json:"default_target_group_id" header:"DEFAULT_TARGET_GROUP_ID"`
	ErrorEventID         string `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType         string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

// Rule
type ListenerRule struct {
	ID            string   `json:"id" header:"ID"`
	Field         string   `json:"field" header:"FIELD"`
	Values        []string `json:"values" header:"VALUES"`
	ListenerID    string   `json:"listener_id" header:"LISTENER_ID"`
	TargetGroupID string   `json:"target_group_id" header:"TARGET_GROUP_ID"`
	ResourceType  string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type Certificate struct {
	ID             string `json:"id" header:"ID"`
	Name           string `json:"name" header:"NAME"`
	PublicKey      string `json:"public_key" header:"PUBLIC_KEY"`
	Chain          string `json:"chain" header:"CHAIN"`
	PrivateKey     string `json:"private_key" header:"PRIVATE_KEY"`
	LoadBalancerID string `json:"load_balancer_id" header:"LOAD_BALANCER_ID"`
	ResourceType   string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
