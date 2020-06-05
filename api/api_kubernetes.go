// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListClusters returns the list of clusters as an array of cluster
func (imco *IMCOClient) ListClusters() (clusters []*types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathKubernetesClusters, true, &clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

// GetCluster returns a cluster by its ID
func (imco *IMCOClient) GetCluster(clusterID string) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathKubernetesCluster, clusterID), true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// CreateCluster creates a cluster
func (imco *IMCOClient) CreateCluster(clusterParams *map[string]interface{}) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathKubernetesClusters, clusterParams, true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// UpdateCluster updates a cluster by its ID
func (imco *IMCOClient) UpdateCluster(clusterID string, clusterParams *map[string]interface{},
) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathKubernetesCluster, clusterID), clusterParams, true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// DeleteCluster deletes a cluster by its ID
func (imco *IMCOClient) DeleteCluster(clusterID string) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathKubernetesCluster, clusterID), true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// RetryCluster retries a cluster by its ID
func (imco *IMCOClient) RetryCluster(clusterID string, clusterParams *map[string]interface{},
) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathKubernetesClusterRetry, clusterID), clusterParams, true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// DiscardCluster discards a cluster by its ID
func (imco *IMCOClient) DiscardCluster(clusterID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathKubernetesClusterDiscard, clusterID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetClusterPlan returns a cluster plan by its ID
func (imco *IMCOClient) GetClusterPlan(clusterPlanID string) (clusterPlan *types.ClusterPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathKubernetesClusterPlan, clusterPlanID), true, &clusterPlan)
	if err != nil {
		return nil, err
	}
	return clusterPlan, err
}

// ListNodePools returns the list of node pools as an array of node pool for a given cluster ID
func (imco *IMCOClient) ListNodePools(clusterID string) (nodePools []*types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathKubernetesClusterNodePools, clusterID), true, &nodePools)
	if err != nil {
		return nil, err
	}
	return nodePools, nil
}

// GetNodePool returns a node pool by its ID
func (imco *IMCOClient) GetNodePool(nodePoolID string) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathKubernetesNodePool, nodePoolID), true, &nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// CreateNodePool creates a node pool
func (imco *IMCOClient) CreateNodePool(clusterID string, nodePoolParams *map[string]interface{},
) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathKubernetesClusterNodePools, clusterID), nodePoolParams, true, &nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// UpdateNodePool updates a node pool by its ID
func (imco *IMCOClient) UpdateNodePool(nodePoolID string, nodePoolParams *map[string]interface{},
) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathKubernetesNodePool, nodePoolID), nodePoolParams, true, &nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// DeleteNodePool deletes a node pool by its ID
func (imco *IMCOClient) DeleteNodePool(nodePoolID string) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathKubernetesNodePool, nodePoolID), true, nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// RetryNodePool retries a node pool by its ID
func (imco *IMCOClient) RetryNodePool(nodePoolID string, nodePoolParams *map[string]interface{},
) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathKubernetesNodePoolRetry, nodePoolID), nodePoolParams, true, &nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// GetNodePoolPlan returns a node pool plan by its ID
func (imco *IMCOClient) GetNodePoolPlan(nodePoolPlanID string) (nodePoolPlan *types.NodePoolPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathKubernetesNodePoolPlan, nodePoolPlanID), true, &nodePoolPlan)
	if err != nil {
		return nil, err
	}
	return nodePoolPlan, err
}
