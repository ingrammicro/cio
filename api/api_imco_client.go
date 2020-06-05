// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

const (
	pathAuditEvents       = "/audit/events"
	pathAuditSystemEvents = "/audit/system_events"

	pathBlueprintAttachment             = "/blueprint/attachments/%s"
	pathBlueprintAttachmentUploaded     = "/blueprint/attachments/%s/uploaded"
	pathBlueprintCookbookVersion        = "/blueprint/cookbook_versions/%s"
	pathBlueprintCookbookVersionProcess = "/blueprint/cookbook_versions/%s/process"
	pathBlueprintCookbookVersions       = "/blueprint/cookbook_versions"
	pathBlueprintScript                 = "/blueprint/scripts/%s"
	pathBlueprintScriptAttachments      = "/blueprint/scripts/%s/attachments"
	pathBlueprintScripts                = "/blueprint/scripts"
	pathBlueprintTemplate               = "/blueprint/templates/%s"
	pathBlueprintTemplateCompile        = "/blueprint/templates/%s/compile"
	pathBlueprintTemplates              = "/blueprint/templates"
	pathBlueprintTemplateScript         = "/blueprint/templates/%s/scripts/%s"
	pathBlueprintTemplateScripts        = "/blueprint/templates/%s/scripts"
	pathBlueprintTemplateScriptsReorder = "/blueprint/templates/%s/scripts/reorder"
	pathBlueprintTemplateScriptsType    = "/blueprint/templates/%s/scripts?type=%s"
	pathBlueprintTemplateServers        = "/blueprint/templates/%s/servers"

	pathBrownfieldCloudAccountImportFloatingIPs        = "/brownfield/cloud_accounts/%s/import_floating_ips"
	pathBrownfieldCloudAccountImportKubernetesClusters = "/brownfield/cloud_accounts/%s/import_kubernetes_clusters"
	pathBrownfieldCloudAccountImportPolicies           = "/brownfield/cloud_accounts/%s/import_policies"
	pathBrownfieldCloudAccountImportServers            = "/brownfield/cloud_accounts/%s/import_servers"
	pathBrownfieldCloudAccountImportVolumes            = "/brownfield/cloud_accounts/%s/import_volumes"
	pathBrownfieldCloudAccountImportVpcs               = "/brownfield/cloud_accounts/%s/import_vpcs"
	pathBrownfieldCloudAccount                         = "/brownfield/cloud_accounts/%s"
	pathBrownfieldCloudAccounts                        = "/brownfield/cloud_accounts"

	pathCloudCloudProviderClusterPlans       = "/cloud/cloud_providers/%s/cluster_plans"
	pathCloudCloudProviderLoadBalancerPlans  = "/cloud/cloud_providers/%s/load_balancer_plans"
	pathCloudCloudProviders                  = "/cloud/cloud_providers"
	pathCloudCloudProviderServerPlansByRealm = "/cloud/cloud_providers/%s/server_plans?realm_id=%s"
	pathCloudCloudProviderStoragePlans       = "/cloud/cloud_providers/%s/storage_plans"
	pathCloudGenericImages                   = "/cloud/generic_images"
	pathCloudProviderRealms                  = "/cloud/cloud_providers/%s/realms"
	pathCloudRealm                           = "/cloud/realms/%s"
	pathCloudRealmNodePoolPlans              = "/cloud/realms/%s/node_pool_plans"
	pathCloudServer                          = "/cloud/servers/%s"
	pathCloudServerArray                     = "/cloud/server_arrays/%s"
	pathCloudServerArrayBoot                 = "/cloud/server_arrays/%s/boot"
	pathCloudServerArrayEmpty                = "/cloud/server_arrays/%s/empty"
	pathCloudServerArrays                    = "/cloud/server_arrays"
	pathCloudServerArrayServers             = "/cloud/server_arrays/%s/servers"
	pathCloudServerArrayShutdown            = "/cloud/server_arrays/%s/shutdown"
	pathCloudServerBoot                     = "/cloud/servers/%s/boot"
	pathCloudServerEvents                   = "/cloud/servers/%s/events"
	pathCloudServerFloatingIps              = "/cloud/servers/%s/floating_ips"
	pathCloudServerOperationalScriptExecute = "/cloud/servers/%s/operational_scripts/%s/execute"
	pathCloudServerOperationalScripts       = "/cloud/servers/%s/operational_scripts"
	pathCloudServerOverride                 = "/cloud/servers/%s/override"
	pathCloudServerPlan                     = "/cloud/server_plans/%s"
	pathCloudServerReboot                   = "/cloud/servers/%s/reboot"
	pathCloudServers                        = "/cloud/servers"
	pathCloudServerShutdown                 = "/cloud/servers/%s/shutdown"
	pathCloudServerVolumes                  = "/cloud/servers/%s/volumes"
	pathCloudSshProfile                     = "/cloud/ssh_profiles/%s"
	pathCloudSshProfiles                    = "/cloud/ssh_profiles"

	pathCseDeployment          = "/cse/deployments/%s"
	pathCseDeployments         = "/cse/deployments"
	pathCseTemplate            = "/cse/templates/%s"
	pathCseTemplateDeployments = "/cse/templates/%s/deployments"
	pathCseTemplates           = "/cse/templates"

	pathKubernetesCluster          = "/kubernetes/clusters/%s"
	pathKubernetesClusterDiscard   = "/kubernetes/clusters/%s/discard"
	pathKubernetesClusterNodePools = "/kubernetes/clusters/%s/node_pools"
	pathKubernetesClusterPlan      = "/kubernetes/cluster_plans/%s"
	pathKubernetesClusterRetry     = "/kubernetes/clusters/%s/retry"
	pathKubernetesClusters         = "/kubernetes/clusters"
	pathKubernetesNodePool         = "/kubernetes/node_pools/%s"
	pathKubernetesNodePoolPlan     = "/kubernetes/node_pool_plans/%s"
	pathKubernetesNodePoolRetry    = "/kubernetes/node_pools/%s/retry"

	pathLabelResource  = "/labels/%s/resources/%s/%s"
	pathLabelResources = "/labels/%s/resources"
	pathLabels         = "/labels"

	pathNetworkDnsDomain                = "/network/dns/domains/%s"
	pathNetworkDnsDomainRecords         = "/network/dns/domains/%s/records"
	pathNetworkDnsDomainRetry           = "/network/dns/domains/%s/retry"
	pathNetworkDnsDomains               = "/network/dns/domains"
	pathNetworkDnsRecord                = "/network/dns/records/%s"
	pathNetworkDnsRecordRetry           = "/network/dns/records/%s/retry"
	pathNetworkFirewallProfile          = "/network/firewall_profiles/%s"
	pathNetworkFirewallProfiles         = "/network/firewall_profiles"
	pathNetworkFloatingIp               = "/network/floating_ips/%s"
	pathNetworkFloatingIpAttachedServer = "/network/floating_ips/%s/attached_server"
	pathNetworkFloatingIpDiscard        = "/network/floating_ips/%s/discard"
	pathNetworkFloatingIps              = "/network/floating_ips"
	pathNetworkListener                 = "/network/listeners/%s"
	pathNetworkListenerRetry            = "/network/listeners/%s/retry"
	pathNetworkListenerRule             = "/network/listeners/%s/rules/%s"
	pathNetworkListenerRules            = "/network/listeners/%s/rules"
	pathNetworkLoadBalancer             = "/network/load_balancers/%s"
	pathNetworkLoadBalancerCertificate  = "/network/load_balancers/%s/certificates/%s"
	pathNetworkLoadBalancerCertificates = "/network/load_balancers/%s/certificates"
	pathNetworkLoadBalancerListeners    = "/network/load_balancers/%s/listeners"
	pathNetworkLoadBalancerPlan         = "/network/load_balancer_plans/%s"
	pathNetworkLoadBalancerRetry        = "/network/load_balancers/%s/retry"
	pathNetworkLoadBalancers            = "/network/load_balancers"
	pathNetworkLoadBalancerTargetGroups = "/network/load_balancers/%s/target_groups"
	pathNetworkSubnet                   = "/network/subnets/%s"
	pathNetworkSubnetServerArrays       = "/network/subnets/%s/server_arrays"
	pathNetworkSubnetServers            = "/network/subnets/%s/servers"
	pathNetworkTargetGroup              = "/network/target_groups/%s"
	pathNetworkTargetGroupRetry         = "/network/target_groups/%s/retry"
	pathNetworkTargetGroupTarget        = "/network/target_groups/%s/targets/%s/%s"
	pathNetworkTargetGroupTargets       = "/network/target_groups/%s/targets"
	pathNetworkVpc                      = "/network/vpcs/%s"
	pathNetworkVpcDiscard               = "/network/vpcs/%s/discard"
	pathNetworkVpcs                     = "/network/vpcs"
	pathNetworkVpcSubnets               = "/network/vpcs/%s/subnets"
	pathNetworkVpcVpn                   = "/network/vpcs/%s/vpn"
	pathNetworkVpcVpnPlans              = "/network/vpcs/%s/vpn_plans"

	pathPluginsToscaCat                     = "/plugins/tosca/cats/%s"
	pathPluginsToscaCatDeploymentTask       = "/plugins/tosca/cats/%s/deployment_tasks/%s"
	pathPluginsToscaCatDeploymentTasks      = "/plugins/tosca/cats/%s/deployment_tasks"
	pathPluginsToscaCatParseMetadata        = "/plugins/tosca/cats/%s/parse_metadata"
	pathPluginsToscaCats                    = "/plugins/tosca/cats"
	pathPluginsToscaDeployment              = "/plugins/tosca/deployments/%s"
	pathPluginsToscaTemporaryArchiveExport  = "/plugins/tosca/temporary_archives/%s/export"
	pathPluginsToscaTemporaryArchiveImport  = "/plugins/tosca/temporary_archives/%s/import"
	pathPluginsToscaTemporaryArchives       = "/plugins/tosca/temporary_archives"
	pathPluginsToscaTemporaryArchivesExport = "/plugins/tosca/temporary_archives/export"

	pathPolicyAssignment            = "/policy/assignments/%s"
	pathPolicyDefinition            = "/policy/definitions/%s"
	pathPolicyDefinitionAssignments = "/policy/definitions/%s/assignments"
	pathPolicyDefinitions           = "/policy/definitions"

	pathSettingsCloudAccount                  = "/settings/cloud_accounts/%s"
	pathSettingsCloudAccountPolicyAssignments = "/settings/cloud_accounts/%s/policy_assignments"
	pathSettingsCloudAccounts                 = "/settings/cloud_accounts"

	pathStoragePlan                 = "/storage/plans/%s"
	pathStorageVolume               = "/storage/volumes/%s"
	pathStorageVolumeAttachedServer = "/storage/volumes/%s/attached_server"
	pathStorageVolumeDiscard        = "/storage/volumes/%s/discard"
	pathStorageVolumes              = "/storage/volumes"

	pathWizardAppDeploy      = "/wizard/apps/%s/deploy"
	pathWizardApps           = "/wizard/apps"
	pathWizardCloudProviders = "/wizard/cloud_providers?app_id=%s&location_id=%s"
	pathWizardLocations      = "/wizard/locations"
	pathWizardServerPlans    = "/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s"
)

//type ClientAPI interface {
//	ListEvents(ctx context.Context) (events []*types.Event, err error)
//	ListSysEvents(ctx context.Context) (events []*types.Event, err error)
//
//	GetAttachment(ctx context.Context, attachmentID string) (attachment *types.Attachment, err error)
//	DeleteAttachment(ctx context.Context, attachmentID string) (err error)
//
//	ListCookbookVersions(ctx context.Context) (cookbookVersions []*types.CookbookVersion, err error)
//	GetCookbookVersion(ctx context.Context, cookbookVersionID string) (cookbookVersion *types.CookbookVersion, err error)
//	CreateCookbookVersion(ctx context.Context, cookbookVersionParams *map[string]interface{}) (
//		cookbookVersion *types.CookbookVersion, err error)
//	ProcessCookbookVersion(ctx context.Context, cookbookVersionID string,
//		cookbookVersionParams *map[string]interface{}) (cookbookVersion *types.CookbookVersion, err error)
//	DeleteCookbookVersion(ctx context.Context, cookbookVersionID string) (err error)
//
//	ListScripts(ctx context.Context) (scripts []*types.Script, err error)
//	GetScript(ctx context.Context, scriptID string) (script *types.Script, err error)
//	CreateScript(ctx context.Context, scriptParams *map[string]interface{}) (script *types.Script, err error)
//	UpdateScript(ctx context.Context, scriptID string, scriptParams *map[string]interface{}) (
//		script *types.Script, err error)
//	DeleteScript(ctx context.Context, scriptID string) (err error)
//
//	AddScriptAttachment(ctx context.Context, scriptID string, attachmentIn *map[string]interface{}) (
//		script *types.Attachment, err error)
//	UploadedScriptAttachment(ctx context.Context, attachmentID string, attachmentParams *map[string]interface{}) (
//		attachment *types.Attachment, err error)
//	ListScriptAttachments(ctx context.Context, scriptID string) (attachments []*types.Attachment, err error)
//
//	ListTemplates(ctx context.Context) (templates []*types.Template, err error)
//	GetTemplate(ctx context.Context, templateID string) (template *types.Template, err error)
//	CreateTemplate(ctx context.Context, templateParams *map[string]interface{}) (template *types.Template, err error)
//	UpdateTemplate(ctx context.Context, templateID string, templateParams *map[string]interface{}) (
//		template *types.Template, err error)
//	CompileTemplate(ctx context.Context, templateID string, payload *map[string]interface{}) (
//		template *types.Template, err error)
//	DeleteTemplate(ctx context.Context, templateID string) (err error)
//
//	ListTemplateScripts(ctx context.Context, templateID string, scriptType string) (
//		templateScript []*types.TemplateScript, err error)
//	GetTemplateScript(ctx context.Context, templateID string, templateScriptID string) (
//		templateScript *types.TemplateScript, err error)
//	CreateTemplateScript(ctx context.Context, templateID string, templateScriptParams *map[string]interface{}) (
//		templateScript *types.TemplateScript, err error)
//	UpdateTemplateScript(ctx context.Context, templateID string, templateScriptID string,
//		templateScriptParams *map[string]interface{}) (templateScript *types.TemplateScript, err error)
//	DeleteTemplateScript(ctx context.Context, templateID string, templateScriptID string) (err error)
//	ReorderTemplateScript(ctx context.Context, templateID string, templateScriptParams *map[string]interface{}) (
//		templateScript []*types.TemplateScript, err error)
//	ListTemplateServers(ctx context.Context, templateID string) (templateServer []*types.TemplateServer, err error)
//
//	ListBrownfieldCloudAccounts(ctx context.Context) (cloudAccounts []*types.CloudAccount, err error)
//	GetBrownfieldCloudAccount(ctx context.Context, cloudAccountID string) (cloudAccount *types.CloudAccount, err error)
//	ImportServers(ctx context.Context, cloudAccountID string, params *map[string]interface{}) (
//		cloudAccount *types.CloudAccount, err error)
//	ImportVPCs(ctx context.Context, cloudAccountID string, params *map[string]interface{}) (
//		cloudAccount *types.CloudAccount, err error)
//	ImportFloatingIPs(ctx context.Context, cloudAccountID string, params *map[string]interface{}) (
//		cloudAccount *types.CloudAccount, err error)
//	ImportVolumes(ctx context.Context, cloudAccountID string, params *map[string]interface{}) (
//		cloudAccount *types.CloudAccount, err error)
//	ImportKubernetesClusters(ctx context.Context, cloudAccountID string, params *map[string]interface{}) (
//		cloudAccount *types.CloudAccount, err error)
//	ImportPolicies(ctx context.Context, cloudAccountID string, params *map[string]interface{}) (
//		cloudAccount *types.CloudAccount, err error)
//
//	ListStorageVolumes(ctx context.Context, serverID string) (volumes []*types.Volume, err error)
//	ListCloudProviders(ctx context.Context) (cloudProviders []*types.CloudProvider, err error)
//	ListServerStoragePlans(ctx context.Context, providerID string) (storagePlans []*types.StoragePlan, err error)
//	ListLoadBalancerPlans(ctx context.Context, providerID string) (
//		loadBalancerPlans []*types.LoadBalancerPlan, err error)
//	ListClusterPlans(ctx context.Context, providerID string) (clusterPlans []*types.ClusterPlan, err error)
//	ListGenericImages(ctx context.Context) (genericImages []*types.GenericImage, err error)
//	ListServerArrays(ctx context.Context) (serverArrays []*types.ServerArray, err error)
//	GetServerArray(ctx context.Context, serverArrayID string) (serverArray *types.ServerArray, err error)
//	CreateServerArray(ctx context.Context, serverArrayParams *map[string]interface{}) (
//		serverArray *types.ServerArray, err error)
//	UpdateServerArray(ctx context.Context, serverArrayID string, serverArrayParams *map[string]interface{}) (
//		serverArray *types.ServerArray, err error)
//
//	BootServerArray(ctx context.Context, serverArrayID string) (serverArray *types.ServerArray, err error)
//	ShutdownServerArray(ctx context.Context, serverArrayID string) (serverArray *types.ServerArray, err error)
//	EmptyServerArray(ctx context.Context, serverArrayID string) (serverArray *types.ServerArray, err error)
//	EnlargeServerArray(ctx context.Context, serverArrayID string, serverArrayParams *map[string]interface{}) (
//		serverArray *types.ServerArray, err error)
//
//	ListServerArrayServers(ctx context.Context, serverArrayID string) (servers []*types.Server, err error)
//	DeleteServerArray(ctx context.Context, serverArrayID string) (err error)
//	ListServerPlans(ctx context.Context, providerID string) (serverPlans []*types.ServerPlan, err error)
//	GetServerPlan(ctx context.Context, planID string) (serverPlan *types.ServerPlan, err error)
//	ListServers(ctx context.Context) (servers []*types.Server, err error)
//	GetServer(ctx context.Context, serverID string) (server *types.Server, err error)
//	CreateServer(ctx context.Context, serverParams *map[string]interface{}) (server *types.Server, err error)
//	UpdateServer(ctx context.Context, serverID string, serverParams *map[string]interface{}) (
//		server *types.Server, err error)
//	BootServer(ctx context.Context, serverID string, serverParams *map[string]interface{}) (
//		server *types.Server, err error)
//	RebootServer(ctx context.Context, serverID string, serverParams *map[string]interface{}) (
//		server *types.Server, err error)
//	ShutdownServer(ctx context.Context, serverID string, serverParams *map[string]interface{}) (
//		server *types.Server, err error)
//	OverrideServer(ctx context.Context, serverID string, serverParams *map[string]interface{}) (
//		server *types.Server, err error)
//	DeleteServer(ctx context.Context, serverID string) (err error)
//	ListServerFloatingIPs(ctx context.Context, serverID string) (floatingIPs []*types.FloatingIP, err error)
//	ListServerVolumes(ctx context.Context, serverID string) (volumes []*types.Volume, err error)
//	ListServerEvents(ctx context.Context, serverID string) (events []*types.Event, err error)
//	ListOperationalScripts(ctx context.Context, serverID string) (scripts []*types.ScriptChar, err error)
//	ExecuteOperationalScript(ctx context.Context, serverID string, scriptID string,
//		serverParams *map[string]interface{}) (script *types.Event, err error)
//	ListSSHProfiles(ctx context.Context) (sshProfiles []*types.SSHProfile, err error)
//	GetSSHProfile(ctx context.Context, sshProfileID string) (sshProfile *types.SSHProfile, err error)
//	CreateSSHProfile(ctx context.Context, sshProfileParams *map[string]interface{}) (
//		sshProfile *types.SSHProfile, err error)
//	UpdateSSHProfile(ctx context.Context, sshProfileID string, sshProfileParams *map[string]interface{}) (
//		sshProfile *types.SSHProfile, err error)
//	DeleteSSHProfile(ctx context.Context, sshProfileID string) (err error)
//	ListFloatingIPs(ctx context.Context, serverID string) (floatingIPs []*types.FloatingIP, err error)
//	ListRealms(ctx context.Context, providerID string) (realms []*types.Realm, err error)
//	GetRealm(ctx context.Context, realmID string) (realm *types.Realm, err error)
//	ListRealmNodePoolPlans(ctx context.Context, realmID string) (nodePoolPlans []*types.NodePoolPlan, err error)
//
//	ListCloudSpecificExtensionDeployments(ctx context.Context) (
//		deployments []*types.CloudSpecificExtensionDeployment, err error)
//	GetCloudSpecificExtensionDeployment(ctx context.Context, deploymentID string) (
//		deployment *types.CloudSpecificExtensionDeployment, err error)
//	CreateCloudSpecificExtensionDeployment(ctx context.Context, templateID string,
//		deploymentParams *map[string]interface{}) (deployment *types.CloudSpecificExtensionDeployment, err error)
//	UpdateCloudSpecificExtensionDeployment(ctx context.Context, deploymentID string,
//		deploymentParams *map[string]interface{}) (deployment *types.CloudSpecificExtensionDeployment, err error)
//	DeleteCloudSpecificExtensionDeployment(ctx context.Context, deploymentID string) (
//		deployment *types.CloudSpecificExtensionDeployment, err error)
//	ListCloudSpecificExtensionTemplates(ctx context.Context) (
//		templates []*types.CloudSpecificExtensionTemplate, err error)
//	GetCloudSpecificExtensionTemplate(ctx context.Context, templateID string) (
//		template *types.CloudSpecificExtensionTemplate, err error)
//	CreateCloudSpecificExtensionTemplate(ctx context.Context, templateParams *map[string]interface{}) (
//		template *types.CloudSpecificExtensionTemplate, err error)
//	UpdateCloudSpecificExtensionTemplate(ctx context.Context, templateID string,
//		templateParams *map[string]interface{}) (template *types.CloudSpecificExtensionTemplate, err error)
//	ListCloudSpecificExtensionTemplateDeployments(ctx context.Context, templateID string) (
//		deployments []*types.CloudSpecificExtensionDeployment, err error)
//	DeleteCloudSpecificExtensionTemplate(ctx context.Context, templateID string) (err error)
//
//	ListClusters(ctx context.Context) (clusters []*types.Cluster, err error)
//	GetCluster(ctx context.Context, clusterID string) (cluster *types.Cluster, err error)
//	CreateCluster(ctx context.Context, clusterParams *map[string]interface{}) (cluster *types.Cluster, err error)
//	UpdateCluster(ctx context.Context, clusterID string, clusterParams *map[string]interface{}) (
//		cluster *types.Cluster, err error)
//	DeleteCluster(ctx context.Context, clusterID string) (cluster *types.Cluster, err error)
//	RetryCluster(ctx context.Context, clusterID string, clusterParams *map[string]interface{}) (
//		cluster *types.Cluster, err error)
//	DiscardCluster(ctx context.Context, clusterID string) (err error)
//	GetClusterPlan(ctx context.Context, clusterPlanID string) (clusterPlan *types.ClusterPlan, err error)
//	ListNodePools(ctx context.Context, clusterID string) (nodePools []*types.NodePool, err error)
//	GetNodePool(ctx context.Context, nodePoolID string) (nodePool *types.NodePool, err error)
//	CreateNodePool(ctx context.Context, clusterID string, nodePoolParams *map[string]interface{}) (
//		nodePool *types.NodePool, err error)
//	UpdateNodePool(ctx context.Context, nodePoolID string, nodePoolParams *map[string]interface{}) (
//		nodePool *types.NodePool, err error)
//	DeleteNodePool(ctx context.Context, nodePoolID string) (nodePool *types.NodePool, err error)
//	RetryNodePool(ctx context.Context, nodePoolID string, nodePoolParams *map[string]interface{}) (
//		nodePool *types.NodePool, err error)
//	GetNodePoolPlan(ctx context.Context, nodePoolPlanID string) (nodePoolPlan *types.NodePoolPlan, err error)
//
//	ListLabels(ctx context.Context) (labels []*types.Label, err error)
//	CreateLabel(ctx context.Context, labelParams *map[string]interface{}) (label *types.Label, err error)
//	AddLabel(ctx context.Context, labelID string, labelParams *map[string]interface{}) (
//		labeledResources []*types.LabeledResource, err error)
//	RemoveLabel(ctx context.Context, labelID string, resourceType string, resourceID string) (err error)
//	ListCloudApplicationDeployments(ctx context.Context) (deployments []*types.CloudApplicationDeployment, err error)
//
//	ListCertificates(ctx context.Context, loadBalancerID string) (certificates []*types.Certificate, err error)
//	GetCertificate(ctx context.Context, loadBalancerID string, certificateID string) (
//		certificate *types.Certificate, err error)
//	CreateCertificate(ctx context.Context, loadBalancerID string, certificateParams *map[string]interface{}) (
//		certificate *types.Certificate, err error)
//	UpdateCertificate(ctx context.Context, loadBalancerID string, certificateID string,
//		certificateParams *map[string]interface{}) (certificate *types.Certificate, err error)
//	DeleteCertificate(ctx context.Context, loadBalancerID string, certificateID string) (err error)
//	ListFirewallProfiles(ctx context.Context) (firewallProfiles []*types.FirewallProfile, err error)
//	GetFirewallProfile(ctx context.Context, firewallProfileID string) (firewallProfile *types.FirewallProfile, err error)
//	CreateFirewallProfile(ctx context.Context, firewallProfileParams *map[string]interface{}) (
//		firewallProfile *types.FirewallProfile, err error)
//	UpdateFirewallProfile(ctx context.Context, firewallProfileID string, firewallProfileParams *map[string]interface{}) (
//		firewallProfile *types.FirewallProfile, err error)
//	DeleteFirewallProfile(ctx context.Context, firewallProfileID string) (err error)
//	GetFloatingIP(ctx context.Context, floatingIPID string) (floatingIP *types.FloatingIP, err error)
//	CreateFloatingIP(ctx context.Context, floatingIPParams *map[string]interface{}) (
//		floatingIP *types.FloatingIP, err error)
//	UpdateFloatingIP(ctx context.Context, floatingIPID string, floatingIPParams *map[string]interface{}) (
//		floatingIP *types.FloatingIP, err error)
//	AttachFloatingIP(ctx context.Context, floatingIPID string, floatingIPParams *map[string]interface{}) (
//		server *types.Server, err error)
//	DetachFloatingIP(ctx context.Context, floatingIPID string) (err error)
//	DeleteFloatingIP(ctx context.Context, floatingIPID string) (err error)
//	DiscardFloatingIP(ctx context.Context, floatingIPID string) (err error)
//	ListListeners(ctx context.Context, loadBalancerID string) (listeners []*types.Listener, err error)
//	GetListener(ctx context.Context, listenerID string) (listener *types.Listener, err error)
//	CreateListener(ctx context.Context, loadBalancerID string, listenerParams *map[string]interface{}) (
//		listener *types.Listener, err error)
//	UpdateListener(ctx context.Context, listenerID string, listenerParams *map[string]interface{}) (
//		listener *types.Listener, err error)
//	DeleteListener(ctx context.Context, listenerID string) (listener *types.Listener, err error)
//	RetryListener(ctx context.Context, listenerID string, listenerParams *map[string]interface{}) (
//		listener *types.Listener, err error)
//	ListRules(ctx context.Context, listenerID string) (listenerRules []*types.ListenerRule, err error)
//	CreateRule(ctx context.Context, listenerID string, listenerRuleParams *map[string]interface{}) (
//		listenerRule *types.ListenerRule, err error)
//	UpdateRule(ctx context.Context, listenerID string, listenerRuleID string,
//		listenerRuleParams *map[string]interface{}) (listenerRule *types.ListenerRule, err error)
//	DeleteRule(ctx context.Context, listenerID string, listenerRuleID string) (err error)
//	ListLoadBalancers(ctx context.Context) (loadBalancers []*types.LoadBalancer, err error)
//	GetLoadBalancer(ctx context.Context, loadBalancerID string) (loadBalancer *types.LoadBalancer, err error)
//	CreateLoadBalancer(ctx context.Context, loadBalancerParams *map[string]interface{}) (
//		loadBalancer *types.LoadBalancer, err error)
//	UpdateLoadBalancer(ctx context.Context, loadBalancerID string, loadBalancerParams *map[string]interface{}) (
//		loadBalancer *types.LoadBalancer, err error)
//	DeleteLoadBalancer(ctx context.Context, loadBalancerID string) (loadBalancer *types.LoadBalancer, err error)
//	RetryLoadBalancer(ctx context.Context, loadBalancerID string, loadBalancerParams *map[string]interface{}) (
//		loadBalancer *types.LoadBalancer, err error)
//	GetLoadBalancerPlan(ctx context.Context, loadBalancerPlanID string) (
//		loadBalancerPlan *types.LoadBalancerPlan, err error)
//	ListSubnets(ctx context.Context, vpcID string) (subnets []*types.Subnet, err error)
//	GetSubnet(ctx context.Context, subnetID string) (subnet *types.Subnet, err error)
//	CreateSubnet(ctx context.Context, vpcID string, subnetParams *map[string]interface{}) (
//		subnet *types.Subnet, err error)
//	UpdateSubnet(ctx context.Context, subnetID string, subnetParams *map[string]interface{}) (
//		subnet *types.Subnet, err error)
//	DeleteSubnet(ctx context.Context, subnetID string) (err error)
//	ListSubnetServers(ctx context.Context, subnetID string) (servers []*types.Server, err error)
//	ListSubnetServerArrays(ctx context.Context, subnetID string) (serverArrays []*types.ServerArray, err error)
//	ListTargetGroups(ctx context.Context, loadBalancerID string) (targetGroups []*types.TargetGroup, err error)
//	GetTargetGroup(ctx context.Context, targetGroupID string) (targetGroup *types.TargetGroup, err error)
//	CreateTargetGroup(ctx context.Context, loadBalancerID string, targetGroupParams *map[string]interface{}) (
//		targetGroup *types.TargetGroup, err error)
//	UpdateTargetGroup(ctx context.Context, targetGroupID string, targetGroupParams *map[string]interface{}) (
//		targetGroup *types.TargetGroup, err error)
//	DeleteTargetGroup(ctx context.Context, targetGroupID string) (targetGroup *types.TargetGroup, err error)
//	RetryTargetGroup(ctx context.Context, targetGroupID string, targetGroupParams *map[string]interface{}) (
//		targetGroup *types.TargetGroup, err error)
//	ListTargets(ctx context.Context, targetGroupID string) (targets []*types.Target, err error)
//	CreateTarget(ctx context.Context, targetGroupID string, targetParams *map[string]interface{}) (
//		target *types.Target, err error)
//	DeleteTarget(ctx context.Context, targetGroupID string, targetResourceType string, targetResourceID string) (
//		err error)
//	ListVPCs(ctx context.Context) (vpcs []*types.Vpc, err error)
//	GetVPC(ctx context.Context, vpcID string) (vpc *types.Vpc, err error)
//	CreateVPC(ctx context.Context, vpcParams *map[string]interface{}) (vpc *types.Vpc, err error)
//	UpdateVPC(ctx context.Context, vpcID string, vpcParams *map[string]interface{}) (vpc *types.Vpc, err error)
//	DeleteVPC(ctx context.Context, vpcID string) (err error)
//	DiscardVPC(ctx context.Context, vpcID string) (err error)
//	GetVPN(ctx context.Context, vpcID string) (vpn *types.Vpn, err error)
//	CreateVPN(ctx context.Context, vpcID string, vpnParams *map[string]interface{}) (vpn *types.Vpn, err error)
//	DeleteVPN(ctx context.Context, vpcID string) (err error)
//	ListVPNPlans(ctx context.Context, vpcID string) (vpnPlans []*types.VpnPlan, err error)
//	ListDomains(ctx context.Context) (domains []*types.Domain, err error)
//	GetDomain(ctx context.Context, domainID string) (domain *types.Domain, err error)
//	CreateDomain(ctx context.Context, domainParams *map[string]interface{}) (domain *types.Domain, err error)
//	DeleteDomain(ctx context.Context, domainID string) (domain *types.Domain, err error)
//	RetryDomain(ctx context.Context, domainID string) (domain *types.Domain, err error)
//	ListRecords(ctx context.Context, domainID string) (records []*types.Record, err error)
//	GetRecord(ctx context.Context, recordID string) (record *types.Record, err error)
//	CreateRecord(ctx context.Context, domainID string, recordParams *map[string]interface{}) (
//		record *types.Record, err error)
//	UpdateRecord(ctx context.Context, recordID string, recordParams *map[string]interface{}) (
//		record *types.Record, err error)
//	DeleteRecord(ctx context.Context, recordID string) (record *types.Record, err error)
//	RetryRecord(ctx context.Context, recordID string) (record *types.Record, err error)
//
//	GetCloudApplicationDeployment(ctx context.Context, deploymentID string) (
//		deployment *types.CloudApplicationDeployment, status int, err error)
//	DeleteCloudApplicationDeployment(ctx context.Context, deploymentID string) (
//		deployment *types.CloudApplicationDeployment, err error)
//	CreateCloudApplicationDeploymentTask(ctx context.Context, catID string, deploymentParams *map[string]interface{}) (
//		deploymentTask *types.CloudApplicationDeploymentTask, err error)
//	GetCloudApplicationDeploymentTask(ctx context.Context, catID string, deploymentTaskID string) (
//		deploymentTask *types.CloudApplicationDeploymentTask, err error)
//	ListCloudApplicationTemplates(ctx context.Context) (templates []*types.CloudApplicationTemplate, err error)
//	GetCloudApplicationTemplate(ctx context.Context, templateID string) (
//		template *types.CloudApplicationTemplate, err error)
//	CreateCloudApplicationTemplate(ctx context.Context, catParams *map[string]interface{}) (
//		template *types.CloudApplicationTemplate, err error)
//	ParseMetadataCloudApplicationTemplate(ctx context.Context, templateID string) (
//		template *types.CloudApplicationTemplate, err error)
//	DeleteCloudApplicationTemplate(ctx context.Context, templateID string) (err error)
//	CreateTemporaryArchive(ctx context.Context, temporaryArchiveParams *map[string]interface{}) (
//		temporaryArchive *types.TemporaryArchive, err error)
//	CreateTemporaryArchiveImport(ctx context.Context, temporaryArchiveID string,
//		temporaryArchiveImportParams *map[string]interface{}) (
//		temporaryArchiveImport *types.TemporaryArchiveImport, err error)
//	GetTemporaryArchiveImport(ctx context.Context, temporaryArchiveImportID string) (
//		temporaryArchiveImport *types.TemporaryArchiveImport, err error)
//	CreateTemporaryArchiveExport(ctx context.Context, temporaryArchiveExportParams *map[string]interface{}) (
//		temporaryArchiveExport *types.TemporaryArchiveExport, err error)
//	GetTemporaryArchiveExportTask(ctx context.Context, temporaryArchiveID string) (
//		temporaryArchiveExportTask *types.TemporaryArchiveExportTask, err error)
//
//	ListPolicyDefinitions(ctx context.Context) (definitions []*types.PolicyDefinition, err error)
//	GetPolicyDefinition(ctx context.Context, definitionID string) (definition *types.PolicyDefinition, err error)
//	CreatePolicyDefinition(ctx context.Context, definitionParams *map[string]interface{}) (
//		definition *types.PolicyDefinition, err error)
//	UpdatePolicyDefinition(ctx context.Context, definitionID string, definitionParams *map[string]interface{}) (
//		definition *types.PolicyDefinition, err error)
//	ListPolicyDefinitionAssignments(ctx context.Context, definitionID string) (
//		assignments []*types.PolicyAssignment, err error)
//	GetPolicyAssignment(ctx context.Context, assignmentID string) (assignment *types.PolicyAssignment, err error)
//	CreatePolicyAssignment(ctx context.Context, definitionID string, assignmentParams *map[string]interface{}) (
//		assignment *types.PolicyAssignment, err error)
//	UpdatePolicyAssignment(ctx context.Context, assignmentID string, assignmentParams *map[string]interface{}) (
//		assignment *types.PolicyAssignment, err error)
//	DeletePolicyAssignment(ctx context.Context, assignmentID string) (assignment *types.PolicyAssignment, err error)
//
//	ListCloudAccounts(ctx context.Context) (cloudAccounts []*types.CloudAccount, err error)
//	GetCloudAccount(ctx context.Context, cloudAccountID string) (cloudAccount *types.CloudAccount, err error)
//	ListPolicyAssignments(ctx context.Context, cloudAccountID string) (
//		assignments []*types.PolicyAssignment, err error)
//
//	GetStoragePlan(ctx context.Context, storagePlanID string) (storagePlan *types.StoragePlan, err error)
//	GetStorageVolume(ctx context.Context, volumeID string) (volume *types.Volume, err error)
//	CreateStorageVolume(ctx context.Context, volumeParams *map[string]interface{}) (volume *types.Volume, err error)
//	UpdateStorageVolume(ctx context.Context, volumeID string, volumeParams *map[string]interface{}) (
//		volume *types.Volume, err error)
//	AttachStorageVolume(ctx context.Context, volumeID string, volumeParams *map[string]interface{}) (
//		server *types.Server, err error)
//	DetachStorageVolume(ctx context.Context, volumeID string) (err error)
//	DeleteStorageVolume(ctx context.Context, volumeID string) (err error)
//	DiscardStorageVolume(ctx context.Context, volumeID string) (err error)
//
//	ListApps(ctx context.Context) (apps []*types.WizardApp, err error)
//	DeployApp(ctx context.Context, appID string, appParams *map[string]interface{}) (app *types.Server, err error)
//	ListLocations(ctx context.Context) (locations []*types.Location, err error)
//	ListWizardCloudProviders(ctx context.Context, appID string, locationID string) (
//		cloudProviders []*types.CloudProvider, err error)
//	ListWizardServerPlans(ctx context.Context, appID string, locationID string, cloudProviderID string) (
//		serverPlans []*types.ServerPlan, err error)
//}

// ClientAPI web service manager
type ClientAPI struct {
	HTTPClient
}
