// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

// Add test
// Add context
// Separate by contexts: Client / Server?
// OpenAPI

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

type API interface {
	// UploadFile uploads a file to target url
	UploadFile(sourceFilePath, targetURL string) error

	// ListEvents returns the list of events as an array of Event
	ListEvents() (events []*types.Event, err error)

	// ListSysEvents returns the list of events as an array of Event
	ListSysEvents() (events []*types.Event, err error)

	// ListApps returns the list of apps as an array of App
	ListApps() (apps []*types.WizardApp, err error)

	// DeployApp deploys a app
	DeployApp(appID string, appParams *map[string]interface{}) (app *types.Server, err error)

	// ListLocations returns the list of locations as an array of Location
	ListLocations() (locations []*types.Location, err error)

	// ListWizardCloudProviders returns the list of cloud providers as an array of CloudProvider
	ListWizardCloudProviders(appID string, locationID string) (cloudProviders []*types.CloudProvider, err error)

	// ListWizardServerPlans returns the list of server plans as an array of ServerPlan
	ListWizardServerPlans(appID string, locationID string, cloudProviderID string,
	) (serverPlans []*types.ServerPlan, err error)

	// ListLabels returns the list of labels as an array of Label
	ListLabels() (labels []*types.Label, err error)

	// CreateLabel creates a label
	CreateLabel(labelParams *map[string]interface{}) (label *types.Label, err error)

	// AddLabel assigns a single label from a single labelable resource
	AddLabel(labelID string, labelParams *map[string]interface{}) (labeledResources []*types.LabeledResource, err error)

	// RemoveLabel de-assigns a single label from a single labelable resource
	RemoveLabel(labelID string, resourceType string, resourceID string) error

	// ListCloudApplicationDeployments returns the list of cloud application deployments as an array of
	// CloudApplicationDeployment
	ListCloudApplicationDeployments() (deployments []*types.CloudApplicationDeployment, err error)

	// GetCloudApplicationDeployment returns a cloud application deployment by its ID
	GetCloudApplicationDeployment(deploymentID string,
	) (deployment *types.CloudApplicationDeployment, status int, err error)

	// DeleteCloudApplicationDeployment deletes a cloud application deployment by its ID
	DeleteCloudApplicationDeployment(deploymentID string) (deployment *types.CloudApplicationDeployment, err error)

	// CreateCloudApplicationDeploymentTask creates a cloud application deployment task by a given CAT ID
	CreateCloudApplicationDeploymentTask(catID string, deploymentParams *map[string]interface{},
	) (deploymentTask *types.CloudApplicationDeploymentTask, err error)

	// GetCloudApplicationDeploymentTask gets a cloud application deployment task by its ID and given CAT ID
	GetCloudApplicationDeploymentTask(catID string, deploymentTaskID string,
	) (deploymentTask *types.CloudApplicationDeploymentTask, err error)

	// ListCloudAccounts returns the list of cloudAccounts as an array of CloudAccount
	ListCloudAccounts() (cloudAccounts []*types.CloudAccount, err error)

	// GetCloudAccount returns a cloudAccount by its ID
	GetCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error)

	// ListCloudSpecificExtensionDeployments returns the list of cloud specific extension deployments as an array of
	// CloudSpecificExtensionDeployment
	ListCloudSpecificExtensionDeployments() (deployments []*types.CloudSpecificExtensionDeployment, err error)

	// GetCloudSpecificExtensionDeployment returns a cloud specific extension deployment by its ID
	GetCloudSpecificExtensionDeployment(deploymentID string,
	) (deployment *types.CloudSpecificExtensionDeployment, err error)

	// CreateCloudSpecificExtensionDeployment creates a cloud specific extension deployment
	CreateCloudSpecificExtensionDeployment(templateID string, deploymentParams *map[string]interface{},
	) (deployment *types.CloudSpecificExtensionDeployment, err error)

	// UpdateCloudSpecificExtensionDeployment updates a cloud specific extension deployment by its ID
	UpdateCloudSpecificExtensionDeployment(deploymentID string, deploymentParams *map[string]interface{},
	) (deployment *types.CloudSpecificExtensionDeployment, err error)

	// DeleteCloudSpecificExtensionDeployment deletes a cloud specific extension deployment by its ID
	DeleteCloudSpecificExtensionDeployment(deploymentID string,
	) (deployment *types.CloudSpecificExtensionDeployment, err error)

	// GetStoragePlan returns a storage plan by its ID
	GetStoragePlan(storagePlanID string) (storagePlan *types.StoragePlan, err error)

	// ListCloudApplicationTemplates returns the list of cloud application templates as an array of
	// CloudApplicationTemplate
	ListCloudApplicationTemplates() (templates []*types.CloudApplicationTemplate, err error)

	// GetCloudApplicationTemplate returns a cloud application template by its ID
	GetCloudApplicationTemplate(templateID string) (template *types.CloudApplicationTemplate, err error)

	// CreateCloudApplicationTemplate creates a cloud application template
	CreateCloudApplicationTemplate(catParams *map[string]interface{},
	) (template *types.CloudApplicationTemplate, err error)

	// ParseMetadataCloudApplicationTemplate process cloud application template metadata
	ParseMetadataCloudApplicationTemplate(templateID string) (template *types.CloudApplicationTemplate, err error)

	// DeleteCloudApplicationTemplate deletes a cloud application template by its ID
	DeleteCloudApplicationTemplate(templateID string) (err error)

	// ListCloudSpecificExtensionTemplates returns the list of cloud specific extension templates as an array of
	// CloudSpecificExtensionTemplate
	ListCloudSpecificExtensionTemplates() (templates []*types.CloudSpecificExtensionTemplate, err error)

	// GetCloudSpecificExtensionTemplate returns a cloud specific extension template by its ID
	GetCloudSpecificExtensionTemplate(templateID string) (template *types.CloudSpecificExtensionTemplate, err error)

	// CreateCloudSpecificExtensionTemplate creates a cloud specific extension template
	CreateCloudSpecificExtensionTemplate(templateParams *map[string]interface{},
	) (template *types.CloudSpecificExtensionTemplate, err error)

	// UpdateCloudSpecificExtensionTemplate updates a cloud specific extension template by its ID
	UpdateCloudSpecificExtensionTemplate(templateID string, templateParams *map[string]interface{},
	) (template *types.CloudSpecificExtensionTemplate, err error)

	// ListCloudSpecificExtensionDeployment returns the list of cloud specific extension deployments for a CSE template
	// as an array of CloudSpecificExtensionDeployment
	ListCloudSpecificExtensionDeployment(templateID string,
	) (deployments []*types.CloudSpecificExtensionDeployment, err error)

	// DeleteCloudSpecificExtensionTemplate deletes a cloud specific extension template by its ID
	DeleteCloudSpecificExtensionTemplate(templateID string) (err error)

	// GetAttachment returns a attachment by its ID
	GetAttachment(attachmentID string) (attachment *types.Attachment, err error)

	// DeleteAttachment deletes a attachment by its ID
	DeleteAttachment(attachmentID string) (err error)

	// ListStorageVolumes returns the list of Volumes as an array of Volume
	ListStorageVolumes(serverID string) (volumes []*types.Volume, err error)

	// GetStorageVolume returns a Volume by its ID
	GetStorageVolume(volumeID string) (volume *types.Volume, err error)

	// CreateStorageVolume creates a Volume
	CreateStorageVolume(volumeParams *map[string]interface{}) (volume *types.Volume, err error)

	// UpdateStorageVolume updates a Volume by its ID
	UpdateStorageVolume(volumeID string, volumeParams *map[string]interface{}) (volume *types.Volume, err error)

	// AttachStorageVolume attaches a Volume by its ID
	AttachStorageVolume(volumeID string, volumeParams *map[string]interface{}) (server *types.Server, err error)

	// DetachStorageVolume detaches a Volume by its ID
	DetachStorageVolume(volumeID string) (err error)

	// DeleteStorageVolume deletes a Volume by its ID
	DeleteStorageVolume(volumeID string) (err error)

	// DiscardStorageVolume discards a Volume by its ID
	DiscardStorageVolume(volumeID string) (err error)

	// ListBrownfieldCloudAccounts returns the list of Brownfield Cloud Accounts as an array of CloudAccount
	ListBrownfieldCloudAccounts() (cloudAccounts []*types.CloudAccount, err error)

	// GetBrownfieldCloudAccount returns a Brownfield Cloud Account by its ID
	GetBrownfieldCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error)

	// ImportServer imports a brownfield server import candidate
	ImportServer(serverID string, serverParams *map[string]interface{}) (server *types.Server, err error)

	// ImportVPC imports a brownfield vpc import candidate
	ImportVPC(vpcID string, vpcParams *map[string]interface{}) (vpc *types.Vpc, err error)

	// ImportFloatingIP imports a brownfield floating ip import candidate
	ImportFloatingIP(floatingIPID string, floatingIPParams *map[string]interface{},
	) (floatingIP *types.FloatingIP, err error)

	// ImportVolume imports a brownfield volume import candidate
	ImportVolume(volumeID string, volumeParams *map[string]interface{}) (volume *types.Volume, err error)

	// ListCloudProviders returns the list of cloudProviders as an array of CloudProvider
	ListCloudProviders() (cloudProviders []*types.CloudProvider, err error)

	// ListServerStoragePlans returns the list of storage plans as an array of StoragePlan
	ListServerStoragePlans(providerID string) (storagePlans []*types.StoragePlan, err error)

	// ListLoadBalancerPlans returns the list of load balancer plans as an array of LoadBalancerPlan
	ListLoadBalancerPlans(providerID string) (loadBalancerPlans []*types.LoadBalancerPlan, err error)

	// ListGenericImages returns the list of generic images as an array of GenericImage
	ListGenericImages() (genericImages []*types.GenericImage, err error)

	// ListCookbookVersions returns the list of cookbook versions as an array of CookbookVersion
	ListCookbookVersions() (cookbookVersions []*types.CookbookVersion, err error)

	// GetCookbookVersion returns a cookbook version by its ID
	GetCookbookVersion(cookbookVersionID string) (cookbookVersion *types.CookbookVersion, err error)

	// CreateCookbookVersion creates a new cookbook version
	CreateCookbookVersion(cookbookVersionParams *map[string]interface{},
	) (cookbookVersion *types.CookbookVersion, err error)

	// ProcessCookbookVersion process a cookbook version by its ID
	ProcessCookbookVersion(cookbookVersionID string, cookbookVersionParams *map[string]interface{},
	) (cookbookVersion *types.CookbookVersion, err error)

	// DeleteCookbookVersion deletes a cookbook version by its ID
	DeleteCookbookVersion(cookbookVersionID string) (err error)

	// ListScripts returns the list of scripts as an array of Scripts
	ListScripts() (scripts []*types.Script, err error)

	// GetScript returns a script by its ID
	GetScript(scriptID string) (script *types.Script, err error)

	// CreateScript creates a script
	CreateScript(scriptParams *map[string]interface{}) (script *types.Script, err error)

	// UpdateScript updates a script by its ID
	UpdateScript(scriptID string, scriptParams *map[string]interface{}) (script *types.Script, err error)

	// DeleteScript deletes a script by its ID
	DeleteScript(scriptID string) (err error)

	// AddScriptAttachment adds an attachment to script by its ID
	AddScriptAttachment(scriptID string, attachmentIn *map[string]interface{}) (script *types.Attachment, err error)

	// UploadedScriptAttachment sets "uploaded" status to the attachment by its ID
	UploadedScriptAttachment(attachmentID string, attachmentParams *map[string]interface{},
	) (attachment *types.Attachment, err error)

	// ListScriptAttachments returns the list of Attachments for a given script ID
	ListScriptAttachments(scriptID string) (attachments []*types.Attachment, err error)

	// ListServerArrays returns the list of server arrays as an array of ServerArray
	ListServerArrays() (serverArrays []*types.ServerArray, err error)

	// GetServerArray returns a server array by its ID
	GetServerArray(serverArrayID string) (serverArray *types.ServerArray, err error)

	// CreateServerArray creates a server array
	CreateServerArray(serverArrayParams *map[string]interface{}) (serverArray *types.ServerArray, err error)

	// UpdateServerArray updates a server array by its ID
	UpdateServerArray(serverArrayID string, serverArrayParams *map[string]interface{},
	) (serverArray *types.ServerArray, err error)

	// BootServerArray boots a server array by its ID
	BootServerArray(serverArrayID string, serverArrayParams *map[string]interface{},
	) (serverArray *types.ServerArray, err error)

	// ShutdownServerArray shuts down a server array by its ID
	ShutdownServerArray(serverArrayID string, serverArrayParams *map[string]interface{},
	) (serverArray *types.ServerArray, err error)

	// EmptyServerArray empties a server array by its ID
	EmptyServerArray(serverArrayID string, serverArrayParams *map[string]interface{},
	) (serverArray *types.ServerArray, err error)

	// EnlargeServerArray enlarges a server array by its ID
	EnlargeServerArray(serverArrayID string, serverArrayParams *map[string]interface{},
	) (serverArray *types.ServerArray, err error)

	// ListServerArrayServers returns the list of servers in a server array as an array of server
	ListServerArrayServers(serverArrayID string) (servers []*types.Server, err error)

	// DeleteServerArray deletes a server array by its ID
	DeleteServerArray(serverArrayID string) (err error)

	// ListTemplates returns the list of templates as an array of Template
	ListTemplates() (templates []*types.Template, err error)

	// GetTemplate returns a template by its ID
	GetTemplate(templateID string) (template *types.Template, err error)

	// CreateTemplate creates a template
	CreateTemplate(templateParams *map[string]interface{}) (template *types.Template, err error)

	// UpdateTemplate updates a template by its ID
	UpdateTemplate(templateID string, templateParams *map[string]interface{}) (template *types.Template, err error)

	// CompileTemplate requests compile for a given template by its ID
	CompileTemplate(templateID string, payload *map[string]interface{}) (template *types.Template, err error)

	// DeleteTemplate deletes a template by its ID
	DeleteTemplate(templateID string) (err error)

	// ListTemplateScripts returns a list of templateScript by template ID
	ListTemplateScripts(templateID string, scriptType string) (templateScript []*types.TemplateScript, err error)

	// GetTemplateScript returns a templateScript
	GetTemplateScript(templateID string, templateScriptID string) (templateScript *types.TemplateScript, err error)

	// CreateTemplateScript creates a templateScript
	CreateTemplateScript(templateID string, templateScriptParams *map[string]interface{},
	) (templateScript *types.TemplateScript, err error)

	// UpdateTemplateScript updates a templateScript
	UpdateTemplateScript(templateID string, templateScriptID string, templateScriptParams *map[string]interface{},
	) (templateScript *types.TemplateScript, err error)

	// DeleteTemplateScript deletes a template record
	DeleteTemplateScript(templateID string, templateScriptID string) (err error)

	// ReorderTemplateScript returns a list of templateScript
	ReorderTemplateScript(templateID string, templateScriptParams *map[string]interface{},
	) (templateScript []*types.TemplateScript, err error)

	// ListTemplateServers returns a list of templateServers by template ID
	ListTemplateServers(templateID string) (templateServer []*types.TemplateServer, err error)

	// ListServerPlans returns the list of serverPlans as an array of ServerPlan
	ListServerPlans(providerID string) (serverPlans []*types.ServerPlan, err error)

	// GetServerPlan returns a serverPlan by its ID
	GetServerPlan(planID string) (serverPlan *types.ServerPlan, err error)

	// ListServers returns the list of servers as an array of server
	ListServers() (servers []*types.Server, err error)

	// GetServer returns a server by its ID
	GetServer(serverID string) (server *types.Server, err error)

	// CreateServer creates a server
	CreateServer(serverParams *map[string]interface{}) (server *types.Server, err error)

	// UpdateServer updates a server by its ID
	UpdateServer(serverID string, serverParams *map[string]interface{}) (server *types.Server, err error)

	// BootServer boots a server by its ID
	BootServer(serverID string, serverParams *map[string]interface{}) (server *types.Server, err error)

	// RebootServer reboots a server by its ID
	RebootServer(serverID string, serverParams *map[string]interface{}) (server *types.Server, err error)

	// ShutdownServer shuts down a server by its ID
	ShutdownServer(serverID string, serverParams *map[string]interface{}) (server *types.Server, err error)

	// OverrideServer overrides a server by its ID
	OverrideServer(serverID string, serverParams *map[string]interface{}) (server *types.Server, err error)

	// DeleteServer deletes a server by its ID
	DeleteServer(serverID string) (err error)

	// ListServerFloatingIPs returns the list of floating IPs as an array of FloatingIP
	ListServerFloatingIPs(serverID string) (floatingIPs []*types.FloatingIP, err error)

	// ListServerVolumes returns the list of volumes as an array of Volume
	ListServerVolumes(serverID string) (volumes []*types.Volume, err error)

	// ListServerEvents returns a list of events by server ID
	ListServerEvents(serverID string) (events []*types.Event, err error)

	// ListOperationalScripts returns a list of scripts by server ID
	ListOperationalScripts(serverID string) (scripts []*types.ScriptChar, err error)

	// ExecuteOperationalScript executes an operational script by its server ID and the script id
	ExecuteOperationalScript(serverID string, scriptID string, serverParams *map[string]interface{},
	) (script *types.Event, err error)

	// ListSSHProfiles returns the list of sshProfiles as an array of SSHProfile
	ListSSHProfiles() (sshProfiles []*types.SSHProfile, err error)

	// GetSSHProfile returns a sshProfile by its ID
	GetSSHProfile(sshProfileID string) (sshProfile *types.SSHProfile, err error)

	// CreateSSHProfile creates a sshProfile
	CreateSSHProfile(sshProfileParams *map[string]interface{}) (sshProfile *types.SSHProfile, err error)

	// UpdateSSHProfile updates a sshProfile by its ID
	UpdateSSHProfile(sshProfileID string, sshProfileParams *map[string]interface{},
	) (sshProfile *types.SSHProfile, err error)

	// DeleteSSHProfile deletes a sshProfile by its ID
	DeleteSSHProfile(sshProfileID string) (err error)

	// CreateTemporaryArchive creates a temporary archive
	CreateTemporaryArchive(temporaryArchiveParams *map[string]interface{},
	) (temporaryArchive *types.TemporaryArchive, err error)

	// CreateTemporaryArchiveImport creates a temporary archive import
	CreateTemporaryArchiveImport(temporaryArchiveID string, temporaryArchiveImportParams *map[string]interface{},
	) (temporaryArchiveImport *types.TemporaryArchiveImport, err error)

	// GetTemporaryArchiveImport returns a temporary archive import by its ID
	GetTemporaryArchiveImport(temporaryArchiveImportID string,
	) (temporaryArchiveImport *types.TemporaryArchiveImport, err error)

	// CreateTemporaryArchiveExport creates a temporary archive export
	CreateTemporaryArchiveExport(temporaryArchiveExportParams *map[string]interface{},
	) (temporaryArchiveExport *types.TemporaryArchiveExport, err error)

	// GetTemporaryArchiveExportTask returns a temporary archive export task by its ID
	GetTemporaryArchiveExportTask(temporaryArchiveID string,
	) (temporaryArchiveExportTask *types.TemporaryArchiveExportTask, err error)

	// DownloadTemporaryArchiveExport gets a file from given url saving file into given file path
	DownloadTemporaryArchiveExport(url string, filepath string) (realFileName string, status int, err error)

	// ListCertificates returns the list of certificates in a load balancer by its ID, as an array of Certificate
	ListCertificates(loadBalancerID string) (certificates []*types.Certificate, err error)

	// GetCertificate returns a certificate by its ID
	GetCertificate(loadBalancerID string, certificateID string) (certificate *types.Certificate, err error)

	// CreateCertificate creates a certificate in a load balancer by its ID
	CreateCertificate(loadBalancerID string, certificateParams *map[string]interface{},
	) (certificate *types.Certificate, err error)

	// UpdateCertificate updates a certificate by its ID
	UpdateCertificate(loadBalancerID string, certificateID string, certificateParams *map[string]interface{},
	) (certificate *types.Certificate, err error)

	// DeleteCertificate deletes a certificate by its ID
	DeleteCertificate(loadBalancerID string, certificateID string) (err error)

	// ListFirewallProfiles returns the list of firewallProfiles as an array of FirewallProfile
	ListFirewallProfiles() (firewallProfiles []*types.FirewallProfile, err error)

	// GetFirewallProfile returns a firewallProfile by its ID
	GetFirewallProfile(firewallProfileID string) (firewallProfile *types.FirewallProfile, err error)

	// CreateFirewallProfile creates a firewallProfile
	CreateFirewallProfile(firewallProfileParams *map[string]interface{},
	) (firewallProfile *types.FirewallProfile, err error)

	// UpdateFirewallProfile updates a firewallProfile by its ID
	UpdateFirewallProfile(firewallProfileID string, firewallProfileParams *map[string]interface{},
	) (firewallProfile *types.FirewallProfile, err error)

	// DeleteFirewallProfile deletes a firewallProfile by its ID
	DeleteFirewallProfile(firewallProfileID string) (err error)

	// ListFloatingIPs returns the list of FloatingIPs as an array of FloatingIP
	ListFloatingIPs(serverID string) (floatingIPs []*types.FloatingIP, err error)

	// GetFloatingIP returns a FloatingIP by its ID
	GetFloatingIP(floatingIPID string) (floatingIP *types.FloatingIP, err error)

	// CreateFloatingIP creates a FloatingIP
	CreateFloatingIP(floatingIPParams *map[string]interface{}) (floatingIP *types.FloatingIP, err error)

	// UpdateFloatingIP updates a FloatingIP by its ID
	UpdateFloatingIP(floatingIPID string, floatingIPParams *map[string]interface{},
	) (floatingIP *types.FloatingIP, err error)

	// AttachFloatingIP attaches a FloatingIP by its ID
	AttachFloatingIP(floatingIPID string, floatingIPParams *map[string]interface{}) (server *types.Server, err error)

	// DetachFloatingIP detaches a FloatingIP by its ID
	DetachFloatingIP(floatingIPID string) (err error)

	// DeleteFloatingIP deletes a FloatingIP by its ID
	DeleteFloatingIP(floatingIPID string) (err error)

	// DiscardFloatingIP discards a FloatingIP by its ID
	DiscardFloatingIP(floatingIPID string) (err error)

	// ListListeners returns the list of listeners in a load balancer by its ID, as an array of Listener
	ListListeners(loadBalancerID string) (listeners []*types.Listener, err error)

	// GetListener returns a listener by its ID
	GetListener(listenerID string) (listener *types.Listener, err error)

	// CreateListener creates a listener in a load balancer by its ID
	CreateListener(loadBalancerID string, listenerParams *map[string]interface{}) (listener *types.Listener, err error)

	// UpdateListener updates a listener by its ID
	UpdateListener(listenerID string, listenerParams *map[string]interface{}) (listener *types.Listener, err error)

	// DeleteListener deletes a listener by its ID
	DeleteListener(listenerID string) (listener *types.Listener, err error)

	// RetryListener retries a listener by its ID
	RetryListener(listenerID string, listenerParams *map[string]interface{}) (listener *types.Listener, err error)

	// ListRules returns the list of rules in a listener by its ID, as an array of ListenerRule
	ListRules(listenerID string) (listenerRules []*types.ListenerRule, err error)

	// CreateRule creates a rule in a listener by its ID
	CreateRule(listenerID string, listenerRuleParams *map[string]interface{},
	) (listenerRule *types.ListenerRule, err error)

	// UpdateRule updates a rule in a listener by its ID
	UpdateRule(listenerID string, listenerRuleID string, listenerRuleParams *map[string]interface{},
	) (listenerRule *types.ListenerRule, err error)

	// DeleteRule deletes a rule in a listener by given IDs
	DeleteRule(listenerID string, listenerRuleID string) (err error)

	// ListLoadBalancers returns the list of load balancers as an array of LoadBalancer
	ListLoadBalancers() (loadBalancers []*types.LoadBalancer, err error)

	// GetLoadBalancer returns a load balancer by its ID
	GetLoadBalancer(loadBalancerID string) (loadBalancer *types.LoadBalancer, err error)

	// CreateLoadBalancer creates a load balancer
	CreateLoadBalancer(loadBalancerParams *map[string]interface{}) (loadBalancer *types.LoadBalancer, err error)

	// UpdateLoadBalancer updates a load balancer by its ID
	UpdateLoadBalancer(loadBalancerID string, loadBalancerParams *map[string]interface{},
	) (loadBalancer *types.LoadBalancer, err error)

	// DeleteLoadBalancer deletes a load balancer by its ID
	DeleteLoadBalancer(loadBalancerID string) (loadBalancer *types.LoadBalancer, err error)

	// RetryLoadBalancer retries a load balancer by its ID
	RetryLoadBalancer(loadBalancerID string, loadBalancerParams *map[string]interface{},
	) (loadBalancer *types.LoadBalancer, err error)

	// GetLoadBalancerPlan returns a load balancer plan by its ID
	GetLoadBalancerPlan(loadBalancerPlanID string) (loadBalancerPlan *types.LoadBalancerPlan, err error)

	// ListSubnets returns the list of Subnets of a VPC as an array of Subnet
	ListSubnets(vpcID string) (subnets []*types.Subnet, err error)

	// GetSubnet returns a Subnet by its ID
	GetSubnet(subnetID string) (subnet *types.Subnet, err error)

	// CreateSubnet creates a Subnet
	CreateSubnet(vpcID string, subnetParams *map[string]interface{}) (subnet *types.Subnet, err error)

	// UpdateSubnet updates a Subnet by its ID
	UpdateSubnet(subnetID string, subnetParams *map[string]interface{}) (subnet *types.Subnet, err error)

	// DeleteSubnet deletes a Subnet by its ID
	DeleteSubnet(subnetID string) (err error)

	// ListSubnetServers returns the list of Servers of a Subnet as an array of server
	ListSubnetServers(subnetID string) (servers []*types.Server, err error)

	// ListSubnetServerArrays returns the list of server arrays of a Subnet as an array of ServerArray
	ListSubnetServerArrays(subnetID string) (serverArrays []*types.ServerArray, err error)

	// ListTargetGroups returns the list of target groups in a load balancer by its ID, as an array of TargetGroup
	ListTargetGroups(loadBalancerID string) (targetGroups []*types.TargetGroup, err error)

	// GetTargetGroup returns a target group by its ID
	GetTargetGroup(targetGroupID string) (targetGroup *types.TargetGroup, err error)

	// CreateTargetGroup creates a target group in a load balancer by its ID
	CreateTargetGroup(loadBalancerID string, targetGroupParams *map[string]interface{},
	) (targetGroup *types.TargetGroup, err error)

	// UpdateTargetGroup updates a target group by its ID
	UpdateTargetGroup(targetGroupID string, targetGroupParams *map[string]interface{},
	) (targetGroup *types.TargetGroup, err error)

	// DeleteTargetGroup deletes a target group by its ID
	DeleteTargetGroup(targetGroupID string) (targetGroup *types.TargetGroup, err error)

	// RetryTargetGroup retries a target group by its ID
	RetryTargetGroup(targetGroupID string, targetGroupParams *map[string]interface{},
	) (targetGroup *types.TargetGroup, err error)

	// ListTargets returns the list of targets in a target group by its ID, as an array of Target
	ListTargets(targetGroupID string) (targets []*types.Target, err error)

	// CreateTarget creates a target in a target group by its ID
	CreateTarget(targetGroupID string, targetParams *map[string]interface{}) (target *types.Target, err error)

	// DeleteTarget deletes a target in a target group by given IDs and resource type
	DeleteTarget(targetGroupID string, targetResourceType string, targetResourceID string) (err error)

	// ListVPCs returns the list of VPCs as an array of VPC
	ListVPCs() (vpcs []*types.Vpc, err error)

	// GetVPC returns a VPC by its ID
	GetVPC(vpcID string) (vpc *types.Vpc, err error)

	// CreateVPC creates a VPC
	CreateVPC(vpcParams *map[string]interface{}) (vpc *types.Vpc, err error)

	// UpdateVPC updates a VPC by its ID
	UpdateVPC(vpcID string, vpcParams *map[string]interface{}) (vpc *types.Vpc, err error)

	// DeleteVPC deletes a VPC by its ID
	DeleteVPC(vpcID string) (err error)

	// DiscardVPC discards a VPC by its ID
	DiscardVPC(vpcID string) (err error)

	// GetVPN returns a VPN by VPC ID
	GetVPN(vpcID string) (vpn *types.Vpn, err error)

	// CreateVPN creates a VPN
	CreateVPN(vpcID string, vpnParams *map[string]interface{}) (vpn *types.Vpn, err error)

	// DeleteVPN deletes VPN by VPC ID
	DeleteVPN(vpcID string) (err error)

	// ListVPNPlans returns the list of VPN plans for a given VPC ID
	ListVPNPlans(vpcID string) (vpnPlans []*types.VpnPlan, err error)

	// SERVER_CMDS

	// GetDispatcherScriptCharacterizationsByType returns script characterizations list for a given phase
	GetDispatcherScriptCharacterizationsByType(phase string,
	) (scriptCharacterizations []*types.ScriptCharacterization, err error)

	// GetDispatcherScriptCharacterizationByUUID returns script characterizations list for a given UUID
	GetDispatcherScriptCharacterizationByUUID(scriptCharacterizationUUID string) (*types.ScriptCharacterization, error)

	// ReportScriptConclusions reports a result
	ReportScriptConclusions(scriptConclusions *map[string]interface{},
	) (command *types.ScriptConclusion, status int, err error)

	// GetBootstrappingConfiguration returns the list of policy files as a JSON response with the desired configuration
	// changes
	GetBootstrappingConfiguration() (bootstrappingConfigurations *types.BootstrappingConfiguration, status int, err error)

	// ReportBootstrappingAppliedConfiguration informs the platform of applied changes
	ReportBootstrappingAppliedConfiguration(bootstrappingAppliedConfigurationParams *map[string]interface{}) (err error)

	// ReportBootstrappingLog reports a policy files application result
	ReportBootstrappingLog(bootstrappingContinuousReportParams *map[string]interface{},
	) (command *types.BootstrappingContinuousReport, status int, err error)

	// GetPolicy returns firewall policy
	GetPolicy() (policy *types.Policy, err error)

	// AddPolicyRule adds a new firewall policy rule
	AddPolicyRule(ruleParams *map[string]interface{}) (policyRule *types.PolicyRule, err error)

	// UpdatePolicy update firewall policy
	UpdatePolicy(policyParams *map[string]interface{}) (policy *types.Policy, err error)

	// Ping resolves if new command is waiting for execution
	Ping() (ping *types.PollingPing, status int, err error)

	// GetNextCommand returns the command to be executed
	GetNextCommand() (command *types.PollingCommand, status int, err error)

	// UpdateCommand updates a command by its ID
	UpdateCommand(commandID string, pollingCommandParams *map[string]interface{},
	) (command *types.PollingCommand, status int, err error)

	// ReportBootstrapLog reports a command result
	ReportBootstrapLog(pollingContinuousReportParams *map[string]interface{},
	) (command *types.PollingContinuousReport, status int, err error)
}

// SERVER CMDs

const (
	pathBlueprintScriptCharacterizationsType = "/blueprint/script_characterizations?type=%s"
	pathBlueprintScriptCharacterization      = "/blueprint/script_characterizations/%s"
	pathBlueprintScriptConclusions           = "/blueprint/script_conclusions"
	pathBlueprintConfiguration               = "/blueprint/configuration"
	pathBlueprintAppliedConfiguration        = "/blueprint/applied_configuration"
	pathBlueprintBootstrapLogs               = "/blueprint/bootstrap_logs"
	pathCloudFirewallProfile                 = "/cloud/firewall_profile"
	pathCloudFirewallProfileRules            = "/cloud/firewall_profile/rules"
	pathCommandPollingPings                  = "/command_polling/pings"
	pathCommandPollingNextCommand            = "/command_polling/command"
	pathCommandPollingCommand                = "/command_polling/commands/%s"
	pathCommandPollingBootstrapLogs          = "/command_polling/bootstrap_logs"
)

// GetDispatcherScriptCharacterizationsByType returns script characterizations list for a given phase
func (imco *IMCOClient) GetDispatcherScriptCharacterizationsByType(phase string,
) (scriptCharacterizations []*types.ScriptCharacterization, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathBlueprintScriptCharacterizationsType, phase),
		true,
		&scriptCharacterizations,
	)
	if err != nil {
		return nil, err
	}
	return scriptCharacterizations, nil
}

// GetDispatcherScriptCharacterizationByUUID returns script characterizations list for a given UUID
func (imco *IMCOClient) GetDispatcherScriptCharacterizationByUUID(scriptCharacterizationUUID string,
) (scriptCharacterization *types.ScriptCharacterization, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathBlueprintScriptCharacterization, scriptCharacterizationUUID),
		true,
		&scriptCharacterization,
	)
	if err != nil {
		return nil, err
	}
	return scriptCharacterization, nil
}

// ReportScriptConclusions reports a result
func (imco *IMCOClient) ReportScriptConclusions(scriptConclusions *map[string]interface{},
) (command *types.ScriptConclusion, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.postAndCheck(pathBlueprintScriptConclusions, scriptConclusions, true, &command)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// GetBootstrappingConfiguration returns the list of policy files as a JSON response with the desired configuration
// changes
func (imco *IMCOClient) GetBootstrappingConfiguration() (
	bootstrappingConfigurations *types.BootstrappingConfiguration, status int, err error,
) {
	logger.DebugFuncInfo()

	status, err = imco.getAndCheck(pathBlueprintConfiguration, true, &bootstrappingConfigurations)
	if err != nil {
		return nil, status, err
	}
	return bootstrappingConfigurations, status, nil
}

// ReportBootstrappingAppliedConfiguration informs the platform of applied changes
func (imco *IMCOClient) ReportBootstrappingAppliedConfiguration(
	bootstrappingAppliedConfigurationParams *map[string]interface{},
) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(pathBlueprintAppliedConfiguration, bootstrappingAppliedConfigurationParams, true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ReportBootstrappingLog reports a policy files application result
func (imco *IMCOClient) ReportBootstrappingLog(bootstrappingContinuousReportParams *map[string]interface{},
) (command *types.BootstrappingContinuousReport, status int, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathBlueprintBootstrapLogs, bootstrappingContinuousReportParams, false, &command)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// GetPolicy returns firewall policy
func (imco *IMCOClient) GetPolicy() (policy *types.Policy, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCloudFirewallProfile, true, &policy)
	if err != nil {
		return nil, err
	}

	var data []byte
	if data, err = json.Marshal(policy); err != nil {
		return nil, err
	}
	policy.Md5 = fmt.Sprintf("%x", md5.Sum(data))
	return policy, nil
}

// AddPolicyRule adds a new firewall policy rule
func (imco *IMCOClient) AddPolicyRule(ruleParams *map[string]interface{}) (policyRule *types.PolicyRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathCloudFirewallProfileRules, ruleParams, true, &policyRule)
	if err != nil {
		return nil, err
	}
	return policyRule, nil
}

// UpdatePolicy update firewall policy
func (imco *IMCOClient) UpdatePolicy(policyParams *map[string]interface{}) (policy *types.Policy, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(pathCloudFirewallProfile, policyParams, true, &policy)
	if err != nil {
		return nil, err
	}
	return policy, nil
}

// Ping resolves if new command is waiting for execution
func (imco *IMCOClient) Ping() (ping *types.PollingPing, status int, err error) {
	logger.DebugFuncInfo()

	payload := make(map[string]interface{})
	status, err = imco.postAndCheck(pathCommandPollingPings, &payload, false, &ping)
	if err != nil {
		return nil, status, err
	}
	return ping, status, nil
}

// GetNextCommand returns the command to be executed
func (imco *IMCOClient) GetNextCommand() (command *types.PollingCommand, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.getAndCheck(pathCommandPollingNextCommand, false, &command)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// UpdateCommand updates a command by its ID
func (imco *IMCOClient) UpdateCommand(commandID string, pollingCommandParams *map[string]interface{},
) (command *types.PollingCommand, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.putAndCheck(
		fmt.Sprintf(pathCommandPollingCommand, commandID),
		pollingCommandParams,
		false,
		&command,
	)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// ReportBootstrapLog reports a command result
func (imco *IMCOClient) ReportBootstrapLog(pollingContinuousReportParams *map[string]interface{},
) (command *types.PollingContinuousReport, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.postAndCheck(pathCommandPollingBootstrapLogs, pollingContinuousReportParams, false, &command)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}
