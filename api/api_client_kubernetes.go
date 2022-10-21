// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListClusters returns the list of clusters as an array of cluster
func (imco *ClientAPI) ListClusters(ctx context.Context) (clusters []*types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathKubernetesClusters, true, &clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

// GetCluster returns a cluster by its ID
func (imco *ClientAPI) GetCluster(ctx context.Context, clusterID string) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathKubernetesCluster, clusterID), true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// CreateCluster creates a cluster
func (imco *ClientAPI) CreateCluster(ctx context.Context, clusterParams *map[string]interface{},
) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathKubernetesClusters, clusterParams, true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// UpdateCluster updates a cluster by its ID
func (imco *ClientAPI) UpdateCluster(ctx context.Context, clusterID string, clusterParams *map[string]interface{},
) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathKubernetesCluster, clusterID), clusterParams, true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// DeleteCluster deletes a cluster by its ID
func (imco *ClientAPI) DeleteCluster(ctx context.Context, clusterID string) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathKubernetesCluster, clusterID), true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// RetryCluster retries a cluster by its ID
func (imco *ClientAPI) RetryCluster(ctx context.Context, clusterID string, clusterParams *map[string]interface{},
) (cluster *types.Cluster, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathKubernetesClusterRetry, clusterID), clusterParams, true, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// DiscardCluster discards a cluster by its ID
func (imco *ClientAPI) DiscardCluster(ctx context.Context, clusterID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathKubernetesClusterDiscard, clusterID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetClusterPlan returns a cluster plan by its ID
func (imco *ClientAPI) GetClusterPlan(ctx context.Context, clusterPlanID string,
) (clusterPlan *types.ClusterPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathKubernetesClusterPlan, clusterPlanID), true, &clusterPlan)
	if err != nil {
		return nil, err
	}
	return clusterPlan, err
}

// ListNodePools returns the list of node pools as an array of node pool for a given cluster ID
func (imco *ClientAPI) ListNodePools(ctx context.Context, clusterID string) (nodePools []*types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathKubernetesClusterNodePools, clusterID), true, &nodePools)
	if err != nil {
		return nil, err
	}
	return nodePools, nil
}

// GetNodePool returns a node pool by its ID
func (imco *ClientAPI) GetNodePool(ctx context.Context, nodePoolID string) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathKubernetesNodePool, nodePoolID), true, &nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// CreateNodePool creates a node pool
func (imco *ClientAPI) CreateNodePool(ctx context.Context, clusterID string, nodePoolParams *map[string]interface{},
) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(PathKubernetesClusterNodePools, clusterID),
		nodePoolParams,
		true,
		&nodePool,
	)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// UpdateNodePool updates a node pool by its ID
func (imco *ClientAPI) UpdateNodePool(ctx context.Context, nodePoolID string, nodePoolParams *map[string]interface{},
) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathKubernetesNodePool, nodePoolID), nodePoolParams, true, &nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// DeleteNodePool deletes a node pool by its ID
func (imco *ClientAPI) DeleteNodePool(ctx context.Context, nodePoolID string) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathKubernetesNodePool, nodePoolID), true, &nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// RetryNodePool retries a node pool by its ID
func (imco *ClientAPI) RetryNodePool(ctx context.Context, nodePoolID string, nodePoolParams *map[string]interface{},
) (nodePool *types.NodePool, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathKubernetesNodePoolRetry, nodePoolID),
		nodePoolParams,
		true,
		&nodePool,
	)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

// GetNodePoolPlan returns a node pool plan by its ID
func (imco *ClientAPI) GetNodePoolPlan(ctx context.Context, nodePoolPlanID string,
) (nodePoolPlan *types.NodePoolPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathKubernetesNodePoolPlan, nodePoolPlanID), true, &nodePoolPlan)
	if err != nil {
		return nil, err
	}
	return nodePoolPlan, err
}
