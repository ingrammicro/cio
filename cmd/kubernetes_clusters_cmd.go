package cmd

import (
	"fmt"
	"github.com/ingrammicro/cio/api/kubernetes"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
	"strings"
)

// WireUpCluster prepares common resources to send request to Concerto API
func WireUpCluster(c *cli.Context) (ds *kubernetes.ClusterService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = kubernetes.NewClusterService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up cluster service", err)
	}

	return ds, f
}

// ClusterList subcommand function
func ClusterList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	clusters, err := clusterSvc.ListClusters()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cluster data", err)
	}

	labelables := make([]types.Labelable, len(clusters))
	for i := 0; i < len(clusters); i++ {
		labelables[i] = types.Labelable(clusters[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	clusters = make([]*types.Cluster, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.Cluster)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.Cluster, got a %T", labelable))
		}
		clusters[i] = v
	}
	if err = formatter.PrintList(clusters); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterShow subcommand function
func ClusterShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cluster, err := clusterSvc.GetCluster(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cluster data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterCreate subcommand function
func ClusterCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"name", "version", "cloud-account-id", "cluster-plan-id"}, formatter)
	clusterIn := map[string]interface{}{
		"name":             c.String("name"),
		"version":          c.String("version"),
		"cloud_account_id": c.String("cloud-account-id"),
		"cluster_plan_id":  c.String("cluster-plan-id"),
	}

	// If 'default_vpc_creation' is set as true, 'default_vpc_cidr' must be provided.
	// If 'default_vpc_creation' is provided as false or not provided, 'vpc_id' must be provided.
	if c.IsSet("default-vpc-creation") {
		clusterIn["default_vpc_creation"] = c.Bool("default-vpc-creation")
		if !c.IsSet("default-vpc-cidr") {
			return fmt.Errorf("invalid parameters detected. Please provide 'default-vpc-cidr'")
		}
		clusterIn["default_vpc_cidr"] = c.String("default-vpc-cidr")
	} else {
		if !c.IsSet("vpc-id") {
			return fmt.Errorf("invalid parameters detected. Please provide 'vpc-id'")
		}
		clusterIn["vpc_id"] = c.String("vpc-id")
	}

	if c.IsSet("public-access-ip-addresses") {
		clusterIn["public_access_ip_addresses"] = strings.Split(c.String("public-access-ip-addresses"), ",")
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		clusterIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	cluster, err := clusterSvc.CreateCluster(&clusterIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create cluster", err)
	}

	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterUpdate subcommand function
func ClusterUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)

	clusterIn := map[string]interface{}{
		"name": c.String("name"),
	}

	if c.IsSet("version") {
		clusterIn["version"] = c.String("version")
	}
	if c.IsSet("public-access-ip-addresses") {
		clusterIn["public_access_ip_addresses"] = strings.Split(c.String("public-access-ip-addresses"), ",")
	}

	cluster, err := clusterSvc.UpdateCluster(c.String("id"), &clusterIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update cluster", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterDelete subcommand function
func ClusterDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cluster, err := clusterSvc.DeleteCluster(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cluster", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterRetry subcommand function
func ClusterRetry(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cluster, err := clusterSvc.RetryCluster(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't retry cluster", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	cluster.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterDiscard subcommand function
func ClusterDiscard(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := clusterSvc.DiscardCluster(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't discard cluster", err)
	}
	return nil
}

// ClusterPlanShow subcommand function
func ClusterPlanShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	clusterPlan, err := clusterSvc.GetClusterPlan(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't show cluster plan", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	clusterPlan.CloudProviderName = cloudProvidersMap[clusterPlan.CloudProviderID]

	if err = formatter.PrintItem(*clusterPlan); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
