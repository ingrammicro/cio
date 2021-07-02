// Copyright (c) 2017-2021 Ingram Micro Inc.

package kubernetes

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathKubernetesClusters = "/kubernetes/clusters"
const APIPathKubernetesCluster = "/kubernetes/clusters/%s"
const APIPathKubernetesClusterRetry = "/kubernetes/clusters/%s/retry"
const APIPathKubernetesClusterDiscard = "/kubernetes/clusters/%s/discard"
const APIPathKubernetesClusterPlan = "/kubernetes/cluster_plans/%s"

// ClusterService manages cluster operations
type ClusterService struct {
	concertoService utils.ConcertoService
}

// NewClusterService returns a Concerto Cluster service
func NewClusterService(concertoService utils.ConcertoService) (*ClusterService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ClusterService{
		concertoService: concertoService,
	}, nil
}

// ListClusters returns the list of clusters as an array of cluster
func (cs *ClusterService) ListClusters() (clusters []*types.Cluster, err error) {
	log.Debug("ListClusters")

	data, status, err := cs.concertoService.Get(APIPathKubernetesClusters)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &clusters); err != nil {
		return nil, err
	}

	return clusters, nil
}

// GetCluster returns a cluster by its ID
func (cs *ClusterService) GetCluster(clusterID string) (cluster *types.Cluster, err error) {
	log.Debug("GetCluster")

	data, status, err := cs.concertoService.Get(fmt.Sprintf(APIPathKubernetesCluster, clusterID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cluster); err != nil {
		return nil, err
	}

	return cluster, nil
}

// CreateCluster creates a cluster
func (cs *ClusterService) CreateCluster(clusterParams *map[string]interface{}) (cluster *types.Cluster, err error) {
	log.Debug("CreateCluster")

	data, status, err := cs.concertoService.Post(APIPathKubernetesClusters, clusterParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cluster); err != nil {
		return nil, err
	}

	return cluster, nil
}

// UpdateCluster updates a cluster by its ID
func (cs *ClusterService) UpdateCluster(
	clusterID string, clusterParams *map[string]interface{},
) (cluster *types.Cluster, err error) {
	log.Debug("UpdateCluster")

	data, status, err := cs.concertoService.Put(fmt.Sprintf(APIPathKubernetesCluster, clusterID), clusterParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cluster); err != nil {
		return nil, err
	}

	return cluster, nil
}

// DeleteCluster deletes a cluster by its ID
func (cs *ClusterService) DeleteCluster(clusterID string) (cluster *types.Cluster, err error) {
	log.Debug("DeleteCluster")

	data, status, err := cs.concertoService.Delete(fmt.Sprintf(APIPathKubernetesCluster, clusterID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cluster); err != nil {
		return nil, err
	}

	return cluster, nil
}

// RetryCluster retries a cluster by its ID
func (cs *ClusterService) RetryCluster(clusterID string, clusterParams *map[string]interface{}) (
	cluster *types.Cluster, err error,
) {
	log.Debug("RetryCluster")

	data, status, err := cs.concertoService.Put(fmt.Sprintf(APIPathKubernetesClusterRetry, clusterID), clusterParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cluster); err != nil {
		return nil, err
	}

	return cluster, nil
}

// DiscardCluster discards a cluster by its ID
func (cs *ClusterService) DiscardCluster(clusterID string) (err error) {
	log.Debug("DiscardCluster")

	data, status, err := cs.concertoService.Delete(fmt.Sprintf(APIPathKubernetesClusterDiscard, clusterID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// GetClusterPlan returns a cluster plan by its ID
func (cs *ClusterService) GetClusterPlan(clusterPlanID string) (clusterPlan *types.ClusterPlan, err error) {
	log.Debug("GetClusterPlan")

	data, status, err := cs.concertoService.Get(fmt.Sprintf(APIPathKubernetesClusterPlan, clusterPlanID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &clusterPlan); err != nil {
		return nil, err
	}

	return clusterPlan, err
}
