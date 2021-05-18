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

// ListClustersMocked test mocked function
func ListClustersMocked(t *testing.T, clustersIn []*types.Cluster) []*types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clustersIn)
	assert.Nil(err, "Clusters test data corrupted")

	// call service
	cs.On("Get", APIPathKubernetesClusters).Return(dIn, 200, nil)
	clustersOut, err := ds.ListClusters()

	assert.Nil(err, "Error getting clusters")
	assert.Equal(clustersIn, clustersOut, "ListClusters returned different clusters")

	return clustersOut
}

// ListClustersFailErrMocked test mocked function
func ListClustersFailErrMocked(t *testing.T, clustersIn []*types.Cluster) []*types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clustersIn)
	assert.Nil(err, "Clusters test data corrupted")

	// call service
	cs.On("Get", APIPathKubernetesClusters).Return(dIn, 200, fmt.Errorf("mocked error"))
	clustersOut, err := ds.ListClusters()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clustersOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return clustersOut
}

// ListClustersFailStatusMocked test mocked function
func ListClustersFailStatusMocked(t *testing.T, clustersIn []*types.Cluster) []*types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clustersIn)
	assert.Nil(err, "Clusters test data corrupted")

	// call service
	cs.On("Get", APIPathKubernetesClusters).Return(dIn, 499, nil)
	clustersOut, err := ds.ListClusters()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clustersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clustersOut
}

// ListClustersFailJSONMocked test mocked function
func ListClustersFailJSONMocked(t *testing.T, clustersIn []*types.Cluster) []*types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathKubernetesClusters).Return(dIn, 200, nil)
	clustersOut, err := ds.ListClusters()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clustersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return clustersOut
}

// GetClusterMocked test mocked function
func GetClusterMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 200, nil)
	clusterOut, err := ds.GetCluster(clusterIn.ID)

	assert.Nil(err, "Error getting cluster")
	assert.Equal(*clusterIn, *clusterOut, "GetCluster returned different cluster")

	return clusterOut
}

// GetClusterFailErrMocked test mocked function
func GetClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	clusterOut, err := ds.GetCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return clusterOut
}

// GetClusterFailStatusMocked test mocked function
func GetClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 499, nil)
	clusterOut, err := ds.GetCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clusterOut
}

// GetClusterFailJSONMocked test mocked function
func GetClusterFailJSONMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 200, nil)
	clusterOut, err := ds.GetCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return clusterOut
}

// CreateClusterMocked test mocked function
func CreateClusterMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Post", APIPathKubernetesClusters, mapIn).Return(dOut, 200, nil)
	clusterOut, err := ds.CreateCluster(mapIn)

	assert.Nil(err, "Error creating cluster")
	assert.Equal(clusterIn, clusterOut, "CreateCluster returned different cluster")

	return clusterOut
}

// CreateClusterFailErrMocked test mocked function
func CreateClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Post", APIPathKubernetesClusters, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	clusterOut, err := ds.CreateCluster(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return clusterOut
}

// CreateClusterFailStatusMocked test mocked function
func CreateClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Post", APIPathKubernetesClusters, mapIn).Return(dOut, 499, nil)
	clusterOut, err := ds.CreateCluster(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clusterOut
}

// CreateClusterFailJSONMocked test mocked function
func CreateClusterFailJSONMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathKubernetesClusters, mapIn).Return(dIn, 200, nil)
	clusterOut, err := ds.CreateCluster(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return clusterOut
}

// UpdateClusterMocked test mocked function
func UpdateClusterMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID), mapIn).Return(dOut, 200, nil)
	clusterOut, err := ds.UpdateCluster(clusterIn.ID, mapIn)

	assert.Nil(err, "Error updating cluster")
	assert.Equal(clusterIn, clusterOut, "UpdateCluster returned different cluster")

	return clusterOut
}

// UpdateClusterFailErrMocked test mocked function
func UpdateClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	clusterOut, err := ds.UpdateCluster(clusterIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return clusterOut
}

// UpdateClusterFailStatusMocked test mocked function
func UpdateClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID), mapIn).Return(dOut, 499, nil)
	clusterOut, err := ds.UpdateCluster(clusterIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clusterOut
}

// UpdateClusterFailJSONMocked test mocked function
func UpdateClusterFailJSONMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID), mapIn).Return(dIn, 200, nil)
	clusterOut, err := ds.UpdateCluster(clusterIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return clusterOut
}

// DeleteClusterMocked test mocked function
func DeleteClusterMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 200, nil)
	clusterOut, err := ds.DeleteCluster(clusterIn.ID)

	assert.Nil(err, "Error deleting cluster")
	assert.Equal(clusterIn, clusterOut, "DeleteCluster returned different cluster")

}

// DeleteClusterFailErrMocked test mocked function
func DeleteClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	clusterOut, err := ds.DeleteCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteClusterFailStatusMocked test mocked function
func DeleteClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 499, nil)
	clusterOut, err := ds.DeleteCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteClusterFailJSONMocked test mocked function
func DeleteClusterFailJSONMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesCluster, clusterIn.ID)).Return(dIn, 200, nil)
	clusterOut, err := ds.DeleteCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")
}

// RetryClusterMocked test mocked function
func RetryClusterMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesClusterRetry, clusterIn.ID), mapIn).Return(dOut, 200, nil)
	clusterOut, err := ds.RetryCluster(clusterIn.ID, mapIn)

	assert.Nil(err, "Error retrying cluster")
	assert.Equal(clusterIn, clusterOut, "RetryCluster returned different cluster")

	return clusterOut
}

// RetryClusterFailErrMocked test mocked function
func RetryClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesClusterRetry, clusterIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	clusterOut, err := ds.RetryCluster(clusterIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return clusterOut
}

// RetryClusterFailStatusMocked test mocked function
func RetryClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesClusterRetry, clusterIn.ID), mapIn).Return(dOut, 499, nil)
	clusterOut, err := ds.RetryCluster(clusterIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clusterOut
}

// RetryClusterFailJSONMocked test mocked function
func RetryClusterFailJSONMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathKubernetesClusterRetry, clusterIn.ID), mapIn).Return(dIn, 200, nil)
	clusterOut, err := ds.RetryCluster(clusterIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return clusterOut
}

// DiscardClusterMocked test mocked function
func DiscardClusterMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesClusterDiscard, clusterIn.ID)).Return(dIn, 200, nil)
	err = ds.DiscardCluster(clusterIn.ID)
	assert.Nil(err, "Error discarding cluster")
}

// DiscardClusterFailErrMocked test mocked function
func DiscardClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesClusterDiscard, clusterIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DiscardCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DiscardClusterFailStatusMocked test mocked function
func DiscardClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathKubernetesClusterDiscard, clusterIn.ID)).Return(dIn, 499, nil)
	err = ds.DiscardCluster(clusterIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// GetClusterPlanMocked test mocked function
func GetClusterPlanMocked(t *testing.T, clusterPlanID string, clusterPlanIn *types.ClusterPlan) *types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterPlanIn)
	assert.Nil(err, "ClusterPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterPlan, clusterPlanID)).Return(dIn, 200, nil)
	clusterPlanOut, err := ds.GetClusterPlan(clusterPlanID)

	assert.Nil(err, "Error getting cluster plan")
	assert.Equal(clusterPlanIn, clusterPlanOut, "GetClusterPlan returned different cluster plan")

	return clusterPlanOut
}

// GetClusterPlanFailErrMocked test mocked function
func GetClusterPlanFailErrMocked(
	t *testing.T,
	clusterPlanID string,
	clusterPlanIn *types.ClusterPlan,
) *types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterPlanIn)
	assert.Nil(err, "ClusterPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterPlan, clusterPlanID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	clusterPlanOut, err := ds.GetClusterPlan(clusterPlanID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterPlanOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return clusterPlanOut
}

// GetClusterPlanFailStatusMocked test mocked function
func GetClusterPlanFailStatusMocked(
	t *testing.T,
	clusterPlanID string,
	clusterPlanIn *types.ClusterPlan,
) *types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterPlanIn)
	assert.Nil(err, "ClusterPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterPlan, clusterPlanID)).Return(dIn, 499, nil)
	clusterPlanOut, err := ds.GetClusterPlan(clusterPlanID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clusterPlanOut
}

// GetClusterPlanFailJSONMocked test mocked function
func GetClusterPlanFailJSONMocked(
	t *testing.T,
	clusterPlanID string,
	clusterPlanIn *types.ClusterPlan,
) *types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathKubernetesClusterPlan, clusterPlanID)).Return(dIn, 200, nil)
	clusterPlanOut, err := ds.GetClusterPlan(clusterPlanID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return clusterPlanOut
}
