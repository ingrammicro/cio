package types

type Domain struct {
	ID             string   `json:"id" header:"ID"`
	Name           string   `json:"name" header:"NAME"`
	State          string   `json:"state" header:"STATE"`
	RemoteID       string   `json:"remote_id" header:"REMOTE_ID" show:"nolist"`
	CloudAccountID string   `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	Nameservers    []string `json:"name_servers" header:"NAMESERVERS"`
	GlobalState    string   `json:"global_state" header:"GLOBAL_STATE" show:"nolist"`
	ErrorEventID   string   `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType   string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
	//CloudAccountName
}

type Record struct {
	ID             string `json:"id" header:"ID"`
	Name           string `json:"name" header:"NAME"`
	State          string `json:"state" header:"STATE"`
	Content        string `json:"content" header:"CONTENT"`
	RemoteID       string `json:"remote_id" header:"REMOTE_ID" show:"nolist"`
	Type           string `json:"type" header:"TYPE"`
	TTL            int    `json:"ttl" header:"TTL"`
	DomainID       string `json:"domain_id" header:"DOMAIN_ID" show:"nolist"`
	InstanceID     string `json:"instance_id" header:"INSTANCE_ID" show:"nolist"`           // Only valid for records of type 'a'
	FloatingIpID   string `json:"floating_ip_id" header:"FLOATING_IP_ID" show:"nolist"`     // Only valid for records of type 'a'
	LoadBalancerID string `json:"load_balancer_id" header:"LOAD_BALANCER_ID" show:"nolist"` // Only valid for records of type 'cname'
	Priority       int    `json:"priority" header:"PRIORITY" show:"nolist"`                 // Only valid for records of types 'mx' and 'srv'
	Weight         int    `json:"weight" header:"WEIGHT" show:"nolist"`                     // Only valid for records of type 'srv'
	Port           int    `json:"port" header:"PORT" show:"nolist"`                         // Only valid for records of type 'srv'
	ErrorEventID   string `json:"error_event_id" header:"ERROR_EVENT_ID" show:"nolist"`
	ResourceType   string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
