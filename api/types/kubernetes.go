package types

type Cluster struct {
	ID                      string   `json:"id" header:"ID"`
	Name                    string   `json:"name" header:"NAME"`
	State                   string   `json:"state" header:"STATE"`
	RemoteID                string   `json:"remote_id" header:"REMOTE_ID"`
	CloudAccountID          string   `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	RealmID                 string   `json:"realm_id" header:"REALM_ID"`
	VpcID                   string   `json:"vpc_id,omitempty" header:"VPC_ID"`
	Brownfield              bool     `json:"brownfield,omitempty" header:"BROWNFIELD" show:"nolist,noshow"`
	Version                 string   `json:"version" header:"VERSION"`
	Endpoint                string   `json:"endpoint" header:"ENDPOINT" show:"nolist"`
	ClusterPlanID           string   `json:"cluster_plan_id" header:"CLUSTER_PLAN_ID" show:"nolist"`
	PublicAccessIPAddresses []string `json:"public_access_ip_addresses" header:"PUBLIC_ACCESS_IP_ADDRESSES" show:"nolist"`
	ErrorEventID            string   `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType            string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
}

type NodePool struct {
	ID             string `json:"id" header:"ID"`
	Name           string `json:"name" header:"NAME"`
	State          string `json:"state" header:"STATE"`
	RemoteID       string `json:"remote_id" header:"REMOTE_ID"`
	ClusterID      string `json:"cluster_id" header:"CLUSTER_ID" show:"nolist"`
	SubnetID       string `json:"subnet_id" header:"SUBNET_ID" show:"nolist"`
	NodePoolPlanID string `json:"node_pool_plan_id" header:"NODE_POOL_PLAN_ID" show:"nolist"`
	DiskSize       int    `json:"disk_size" header:"DISK_SIZE"`
	OSType         string `json:"os_type" header:"OS_TYPE"`
	CpuType        string `json:"cpu_type" header:"CPU_TYPE"`
	MinNodes       int    `json:"min_nodes" header:"MIN_NODES"`
	MaxNodes       int    `json:"max_nodes" header:"MAX_NODES"`
	DesiredNodes   int    `json:"desired_nodes" header:"DESIRED_NODES" show:"nolist"`
	PodsPerNode    int    `json:"pods_per_node" header:"PODS_PER_NODE" show:"nolist"`
	Autoscaling    bool   `json:"autoscaling" header:"AUTOSCALING" show:"nolist"`
	Brownfield     bool   `json:"brownfield,omitempty" header:"BROWNFIELD" show:"nolist,noshow"`
	ErrorEventID   string `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType   string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type ClusterPlan struct {
	ID                  string   `json:"id" header:"ID"`
	Name                string   `json:"name" header:"NAME"`
	AvailableVersions   []string `json:"available_versions" header:"AVAILABLE_VERSIONS"`
	DefaultVersion      string   `json:"default_version" header:"DEFAULT_VERSION"`
	MaxPodsPerNode      int      `json:"max_pods_per_node" header:"MAX_PODS_PER_NODE"`
	MaxNodesPerNodePool int      `json:"max_nodes_per_node_pool" header:"MAX_NODES_PER_NODE_POOL"`
	CloudProviderID     string   `json:"cloud_provider_id,omitempty" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName   string   `header:"CLOUD_PROVIDER_NAME"`
	RealmID             string   `json:"realm_id,omitempty" header:"REALM_ID"`
	RealmProviderName   string   `json:"realm_provider_name,omitempty" header:"REALM_PROVIDER_NAME"`
	FlavourProviderName string   `json:"flavour_provider_name,omitempty" header:"FLAVOUR_PROVIDER_NAME"`
	Deprecated          bool     `json:"deprecated,omitempty" header:"DEPRECATED" show:"nolist,noshow"`
	ResourceType        string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type NodePoolPlan struct {
	ID                   string   `json:"id" header:"ID"`
	Name                 string   `json:"name" header:"NAME"`
	RemoteID             string   `json:"remote_id" header:"REMOTE_ID"`
	CPUTypes             []string `json:"cpu_types" header:"CPU_TYPES"`
	CPUs                 int      `json:"cpus" header:"CPUS"`
	Memory               int      `json:"memory" header:"MEMORY"`
	CloudProviderID      string   `json:"cloud_provider_id,omitempty" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName    string   `header:"CLOUD_PROVIDER_NAME"`
	RealmID              string   `json:"realm_id,omitempty" header:"REALM_ID"`
	ServerPlanID         string   `json:"server_plan_id,omitempty" header:"SERVER_PLAN_ID"`
	AutoscalingCapable   bool     `json:"autoscaling_capable" header:"AUTOSCALING_CAPABLE"`
	PodsPerNodePresence  bool     `json:"pods_per_node_presence" header:"PODS_PER_NODE_PRESENCE"`
	FirstNodePoolCapable bool     `json:"first_node_pool_capable" header:"FIRST_NODE_POOL_CAPABLE"`
	Deprecated           bool     `json:"deprecated,omitempty" header:"DEPRECATED" show:"nolist,noshow"`
	ResourceType         string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
