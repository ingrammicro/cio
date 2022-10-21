// Copyright (c) 2017-2022 Ingram Micro Inc.

package kubernetes

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"
	"strings"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Cluster Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Logical name of the cluster"}
	fNameReq := fName
	fNameReq.Required = true

	fVersion := cmd.FlagContext{Type: cmd.String, Name: cmd.Version, Usage: "Kubernetes version of the cluster"}
	fVersionReq := fVersion
	fVersionReq.Required = true

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account the cluster will be deployed"}

	fClusterPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.ClusterPlanId, Required: true,
		Usage: "Identifier of the cluster plan that will use the cluster to be created"}

	fPublicAccessIpAddresses := cmd.FlagContext{Type: cmd.String, Name: cmd.PublicAccessIpAddresses,
		Usage: "A list of comma separated CIDR blocks the cluster will allow to receive requests"}

	fDefaultVpcCreation := cmd.FlagContext{Type: cmd.Bool, Name: cmd.DefaultVpcCreation,
		Usage: "Flag indicating if the cluster must create a VPC first"}

	fDefaultVpcCidr := cmd.FlagContext{Type: cmd.String, Name: cmd.DefaultVpcCidr,
		Usage: "CIDR block where the default VPC will have when created"}

	fVpcId := cmd.FlagContext{Type: cmd.String, Name: cmd.VpcId,
		Usage: "Identifier of the VPC where the cluster will be deployed"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with cluster"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "cluster", Hidden: true,
		Usage: "Resource Type"}

	clustersCmd := cmd.NewCommand(kubernetesCmd, &cmd.CommandContext{
		Use:   "clusters",
		Short: "Provides information on kubernetes clusters"},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists all existing clusters",
		RunMethod: ClusterList},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the cluster identified by the given id",
		RunMethod:    ClusterShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates a new cluster",
		RunMethod: ClusterCreate,
		FlagContexts: []cmd.FlagContext{
			fNameReq,
			fVersionReq,
			fCloudAccountId,
			fClusterPlanId,
			fPublicAccessIpAddresses,
			fDefaultVpcCreation,
			fDefaultVpcCidr,
			fVpcId,
			fLabels}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing cluster identified by the given id",
		RunMethod:    ClusterUpdate,
		FlagContexts: []cmd.FlagContext{fId, fNameReq, fVersion, fPublicAccessIpAddresses}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a cluster",
		RunMethod:    ClusterDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "retry",
		Short:        "Retries the application of cluster identified by the given id",
		RunMethod:    ClusterRetry,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "discard",
		Short:        "Discards a cluster but does not delete it from the cloud provider",
		RunMethod:    ClusterDiscard,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "show-plan",
		Short:        "Shows information about a specific cluster plan identified by the given id",
		RunMethod:    ClusterPlanShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(clustersCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// ClusterList subcommand function
func ClusterList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	clusters, err := svc.ListClusters(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive cluster data", err)
		return err
	}

	labelables := make([]types.Labelable, len(clusters))
	for i := 0; i < len(clusters); i++ {
		labelables[i] = types.Labelable(clusters[i])
	}
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	filteredLabelables, err := labels.LabelFiltering(labelables, labelIDsByName)
	if err != nil {
		return err
	}
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	clusters = make([]*types.Cluster, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.Cluster)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.Cluster, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		clusters[i] = v
	}
	if err = formatter.PrintList(clusters); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ClusterShow subcommand function
func ClusterShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cluster, err := svc.GetCluster(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive cluster data", err)
		return err
	}
	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ClusterCreate subcommand function
func ClusterCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	clusterIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"version":          viper.GetString(cmd.Version),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
		"cluster_plan_id":  viper.GetString(cmd.ClusterPlanId),
	}

	// If 'default_vpc_creation' is set as true, 'default_vpc_cidr' must be provided.
	// If 'default_vpc_creation' is provided as false or not provided, 'vpc_id' must be provided.
	if viper.IsSet(cmd.DefaultVpcCreation) {
		clusterIn["default_vpc_creation"] = viper.GetBool(cmd.DefaultVpcCreation)
		if !viper.IsSet(cmd.DefaultVpcCidr) {
			return fmt.Errorf("invalid parameters detected. Please provide 'default-vpc-cidr'")
		}
		clusterIn["default_vpc_cidr"] = viper.GetString(cmd.DefaultVpcCidr)
	} else {
		if !viper.IsSet(cmd.VpcId) {
			return fmt.Errorf("invalid parameters detected. Please provide 'vpc-id'")
		}
		clusterIn["vpc_id"] = viper.GetString(cmd.VpcId)
	}

	if viper.IsSet(cmd.PublicAccessIpAddresses) {
		clusterIn["public_access_ip_addresses"] = strings.Split(viper.GetString(cmd.PublicAccessIpAddresses), ",")
	}

	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		clusterIn["label_ids"], err = labels.LabelResolution(
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	cluster, err := svc.CreateCluster(cmd.GetContext(), &clusterIn)
	if err != nil {
		formatter.PrintError("Couldn't create cluster", err)
		return err
	}

	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ClusterUpdate subcommand function
func ClusterUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	clusterIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	if viper.IsSet(cmd.Version) {
		clusterIn["version"] = viper.GetString(cmd.Version)
	}
	if viper.IsSet(cmd.PublicAccessIpAddresses) {
		clusterIn["public_access_ip_addresses"] = strings.Split(viper.GetString(cmd.PublicAccessIpAddresses), ",")
	}

	cluster, err := svc.UpdateCluster(cmd.GetContext(), viper.GetString(cmd.Id), &clusterIn)
	if err != nil {
		formatter.PrintError("Couldn't update cluster", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ClusterDelete subcommand function
func ClusterDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cluster, err := svc.DeleteCluster(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete cluster", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ClusterRetry subcommand function
func ClusterRetry() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cluster, err := svc.RetryCluster(cmd.GetContext(), viper.GetString(cmd.Id), &map[string]interface{}{})
	if err != nil {
		formatter.PrintError("Couldn't retry cluster", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ClusterDiscard subcommand function
func ClusterDiscard() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DiscardCluster(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't discard cluster", err)
		return err
	}
	return nil
}

// ClusterPlanShow subcommand function
func ClusterPlanShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	clusterPlan, err := svc.GetClusterPlan(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't show cluster plan", err)
		return err
	}

	if err = formatter.PrintItem(*clusterPlan); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
