package kubernetes

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// NodePoolService manages node pool operations
type NodePoolService struct {
	concertoService utils.ConcertoService
}

// NewNodePoolService returns a Concerto node pool service
func NewNodePoolService(concertoService utils.ConcertoService) (*NodePoolService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &NodePoolService{
		concertoService: concertoService,
	}, nil
}

// ListNodePools returns the list of node pools as an array of node pool for a given cluster ID
func (nps *NodePoolService) ListNodePools(clusterID string) (nodePools []*types.NodePool, err error) {
	log.Debug("ListNodePools")

	data, status, err := nps.concertoService.Get(fmt.Sprintf("/kubernetes/clusters/%s/node_pools", clusterID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePools); err != nil {
		return nil, err
	}

	return nodePools, nil
}

// GetNodePool returns a node pool by its ID
func (nps *NodePoolService) GetNodePool(nodePoolID string) (nodePool *types.NodePool, err error) {
	log.Debug("GetNodePool")

	data, status, err := nps.concertoService.Get(fmt.Sprintf("/kubernetes/node_pools/%s", nodePoolID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePool); err != nil {
		return nil, err
	}

	return nodePool, nil
}

// CreateNodePool creates a node pool
func (nps *NodePoolService) CreateNodePool(clusterID string, nodePoolParams *map[string]interface{}) (nodePool *types.NodePool, err error) {
	log.Debug("CreateNodePool")

	data, status, err := nps.concertoService.Post(fmt.Sprintf("/kubernetes/clusters/%s/node_pools", clusterID), nodePoolParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePool); err != nil {
		return nil, err
	}

	return nodePool, nil
}

// UpdateNodePool updates a node pool by its ID
func (nps *NodePoolService) UpdateNodePool(nodePoolID string, nodePoolParams *map[string]interface{}) (nodePool *types.NodePool, err error) {
	log.Debug("UpdateNodePool")

	data, status, err := nps.concertoService.Put(fmt.Sprintf("/kubernetes/node_pools/%s", nodePoolID), nodePoolParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePool); err != nil {
		return nil, err
	}

	return nodePool, nil
}

// DeleteNodePool deletes a node pool by its ID
func (nps *NodePoolService) DeleteNodePool(nodePoolID string) (nodePool *types.NodePool, err error) {
	log.Debug("DeleteNodePool")

	data, status, err := nps.concertoService.Delete(fmt.Sprintf("/kubernetes/node_pools/%s", nodePoolID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePool); err != nil {
		return nil, err
	}

	return nodePool, nil
}

// RetryNodePool retries a node pool by its ID
func (nps *NodePoolService) RetryNodePool(nodePoolID string, nodePoolParams *map[string]interface{}) (nodePool *types.NodePool, err error) {
	log.Debug("RetryNodePool")

	data, status, err := nps.concertoService.Put(fmt.Sprintf("/kubernetes/node_pools/%s/retry", nodePoolID), nodePoolParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePool); err != nil {
		return nil, err
	}

	return nodePool, nil
}

// GetNodePoolPlan returns a node pool plan by its ID
func (nps *NodePoolService) GetNodePoolPlan(nodePoolPlanID string) (nodePoolPlan *types.NodePoolPlan, err error) {
	log.Debug("GetNodePoolPlan")

	data, status, err := nps.concertoService.Get(fmt.Sprintf("/kubernetes/node_pool_plans/%s", nodePoolPlanID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePoolPlan); err != nil {
		return nil, err
	}

	return nodePoolPlan, err
}
