// Copyright (c) 2017-2021 Ingram Micro Inc.

package kubernetes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// ListNodePoolsMocked test mocked function
func ListNodePoolsMocked(t *testing.T, clusterID string, nodePoolsIn []*types.NodePool) []*types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolsIn)
	assert.Nil(err, "NodePools test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID)).Return(dIn, 200, nil)
	nodePoolsOut, err := ds.ListNodePools(clusterID)

	assert.Nil(err, "Error getting node pools")
	assert.Equal(nodePoolsIn, nodePoolsOut, "ListNodePools returned different node pools")

	return nodePoolsOut
}

// ListNodePoolsFailErrMocked test mocked function
func ListNodePoolsFailErrMocked(t *testing.T, clusterID string, nodePoolsIn []*types.NodePool) []*types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolsIn)
	assert.Nil(err, "NodePools test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	nodePoolsOut, err := ds.ListNodePools(clusterID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolsOut
}

// ListNodePoolsFailStatusMocked test mocked function
func ListNodePoolsFailStatusMocked(t *testing.T, clusterID string, nodePoolsIn []*types.NodePool) []*types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolsIn)
	assert.Nil(err, "NodePools test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID)).Return(dIn, 499, nil)
	nodePoolsOut, err := ds.ListNodePools(clusterID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolsOut
}

// ListNodePoolsFailJSONMocked test mocked function
func ListNodePoolsFailJSONMocked(t *testing.T, clusterID string, nodePoolsIn []*types.NodePool) []*types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID)).Return(dIn, 200, nil)
	nodePoolsOut, err := ds.ListNodePools(clusterID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolsOut
}

// GetNodePoolMocked test mocked function
func GetNodePoolMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).Return(dIn, 200, nil)
	nodePoolOut, err := ds.GetNodePool(nodePoolIn.ID)

	assert.Nil(err, "Error getting node pool")
	assert.Equal(*nodePoolIn, *nodePoolOut, "GetNodePool returned different node pool")

	return nodePoolOut
}

// GetNodePoolFailErrMocked test mocked function
func GetNodePoolFailErrMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	nodePoolOut, err := ds.GetNodePool(nodePoolIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolOut
}

// GetNodePoolFailStatusMocked test mocked function
func GetNodePoolFailStatusMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).Return(dIn, 499, nil)
	nodePoolOut, err := ds.GetNodePool(nodePoolIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolOut
}

// GetNodePoolFailJSONMocked test mocked function
func GetNodePoolFailJSONMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).Return(dIn, 200, nil)
	nodePoolOut, err := ds.GetNodePool(nodePoolIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolOut
}

// CreateNodePoolMocked test mocked function
func CreateNodePoolMocked(t *testing.T, clusterID string, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID), mapIn).Return(dOut, 200, nil)
	nodePoolOut, err := ds.CreateNodePool(clusterID, mapIn)

	assert.Nil(err, "Error creating node pool")
	assert.Equal(nodePoolIn, nodePoolOut, "CreateNodePool returned different node pool")

	return nodePoolOut
}

// CreateNodePoolFailErrMocked test mocked function
func CreateNodePoolFailErrMocked(t *testing.T, clusterID string, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	nodePoolOut, err := ds.CreateNodePool(clusterID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolOut
}

// CreateNodePoolFailStatusMocked test mocked function
func CreateNodePoolFailStatusMocked(t *testing.T, clusterID string, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID), mapIn).Return(dOut, 499, nil)
	nodePoolOut, err := ds.CreateNodePool(clusterID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolOut
}

// CreateNodePoolFailJSONMocked test mocked function
func CreateNodePoolFailJSONMocked(t *testing.T, clusterID string, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathKubernetesClusterNodePools, clusterID), mapIn).Return(dIn, 200, nil)
	nodePoolOut, err := ds.CreateNodePool(clusterID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolOut
}

// UpdateNodePoolMocked test mocked function
func UpdateNodePoolMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID), mapIn).Return(dOut, 200, nil)
	nodePoolOut, err := ds.UpdateNodePool(nodePoolIn.ID, mapIn)

	assert.Nil(err, "Error updating node pool")
	assert.Equal(nodePoolIn, nodePoolOut, "UpdateNodePool returned different node pool")

	return nodePoolOut
}

// UpdateNodePoolFailErrMocked test mocked function
func UpdateNodePoolFailErrMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	nodePoolOut, err := ds.UpdateNodePool(nodePoolIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolOut
}

// UpdateNodePoolFailStatusMocked test mocked function
func UpdateNodePoolFailStatusMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID), mapIn).Return(dOut, 499, nil)
	nodePoolOut, err := ds.UpdateNodePool(nodePoolIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolOut
}

// UpdateNodePoolFailJSONMocked test mocked function
func UpdateNodePoolFailJSONMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID), mapIn).Return(dIn, 200, nil)
	nodePoolOut, err := ds.UpdateNodePool(nodePoolIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolOut
}

// DeleteNodePoolMocked test mocked function
func DeleteNodePoolMocked(t *testing.T, nodePoolIn *types.NodePool) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).Return(dIn, 200, nil)
	nodePoolOut, err := ds.DeleteNodePool(nodePoolIn.ID)

	assert.Nil(err, "Error deleting node pool")
	assert.Equal(nodePoolIn, nodePoolOut, "DeleteNodePool returned different node pool")

}

// DeleteNodePoolFailErrMocked test mocked function
func DeleteNodePoolFailErrMocked(t *testing.T, nodePoolIn *types.NodePool) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	nodePoolOut, err := ds.DeleteNodePool(nodePoolIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteNodePoolFailStatusMocked test mocked function
func DeleteNodePoolFailStatusMocked(t *testing.T, nodePoolIn *types.NodePool) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).Return(dIn, 499, nil)
	nodePoolOut, err := ds.DeleteNodePool(nodePoolIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteNodePoolFailJSONMocked test mocked function
func DeleteNodePoolFailJSONMocked(t *testing.T, nodePoolIn *types.NodePool) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesNodePool, nodePoolIn.ID)).Return(dIn, 200, nil)
	nodePoolOut, err := ds.DeleteNodePool(nodePoolIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")
}

// RetryNodePoolMocked test mocked function
func RetryNodePoolMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePoolRetry, nodePoolIn.ID), mapIn).Return(dOut, 200, nil)
	nodePoolOut, err := ds.RetryNodePool(nodePoolIn.ID, mapIn)

	assert.Nil(err, "Error retrying node pool")
	assert.Equal(nodePoolIn, nodePoolOut, "RetryNodePool returned different node pool")

	return nodePoolOut
}

// RetryNodePoolFailErrMocked test mocked function
func RetryNodePoolFailErrMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePoolRetry, nodePoolIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	nodePoolOut, err := ds.RetryNodePool(nodePoolIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolOut
}

// RetryNodePoolFailStatusMocked test mocked function
func RetryNodePoolFailStatusMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// to json
	dOut, err := json.Marshal(nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePoolRetry, nodePoolIn.ID), mapIn).Return(dOut, 499, nil)
	nodePoolOut, err := ds.RetryNodePool(nodePoolIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolOut
}

// RetryNodePoolFailJSONMocked test mocked function
func RetryNodePoolFailJSONMocked(t *testing.T, nodePoolIn *types.NodePool) *types.NodePool {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodePoolIn)
	assert.Nil(err, "NodePool test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesNodePoolRetry, nodePoolIn.ID), mapIn).Return(dIn, 200, nil)
	nodePoolOut, err := ds.RetryNodePool(nodePoolIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolOut
}

// GetNodePoolPlanMocked test mocked function
func GetNodePoolPlanMocked(
	t *testing.T,
	nodePoolPlanID string,
	nodePoolPlanIn *types.NodePoolPlan,
) *types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlanIn)
	assert.Nil(err, "NodePoolPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePoolPlan, nodePoolPlanID)).Return(dIn, 200, nil)
	nodePoolPlanOut, err := ds.GetNodePoolPlan(nodePoolPlanID)

	assert.Nil(err, "Error getting node pool plan")
	assert.Equal(nodePoolPlanIn, nodePoolPlanOut, "GetNodePoolPlan returned different node pool plan")

	return nodePoolPlanOut
}

// GetNodePoolPlanFailErrMocked test mocked function
func GetNodePoolPlanFailErrMocked(
	t *testing.T,
	nodePoolPlanID string,
	nodePoolPlanIn *types.NodePoolPlan,
) *types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlanIn)
	assert.Nil(err, "NodePoolPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePoolPlan, nodePoolPlanID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	nodePoolPlanOut, err := ds.GetNodePoolPlan(nodePoolPlanID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolPlanOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolPlanOut
}

// GetNodePoolPlanFailStatusMocked test mocked function
func GetNodePoolPlanFailStatusMocked(
	t *testing.T,
	nodePoolPlanID string,
	nodePoolPlanIn *types.NodePoolPlan,
) *types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlanIn)
	assert.Nil(err, "NodePoolPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePoolPlan, nodePoolPlanID)).Return(dIn, 499, nil)
	nodePoolPlanOut, err := ds.GetNodePoolPlan(nodePoolPlanID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolPlanOut
}

// GetNodePoolPlanFailJSONMocked test mocked function
func GetNodePoolPlanFailJSONMocked(
	t *testing.T,
	nodePoolPlanID string,
	nodePoolPlanIn *types.NodePoolPlan,
) *types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodePoolService(cs)
	assert.Nil(err, "Couldn't load node pool service")
	assert.NotNil(ds, "NodePool service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesNodePoolPlan, nodePoolPlanID)).Return(dIn, 200, nil)
	nodePoolPlanOut, err := ds.GetNodePoolPlan(nodePoolPlanID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolPlanOut
}
