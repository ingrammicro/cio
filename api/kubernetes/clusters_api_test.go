package kubernetes

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClusterServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewClusterService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListClusters(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	ListClustersMocked(t, clustersIn)
	ListClustersFailErrMocked(t, clustersIn)
	ListClustersFailStatusMocked(t, clustersIn)
	ListClustersFailJSONMocked(t, clustersIn)
}

func TestGetCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range clustersIn {
		GetClusterMocked(t, clusterIn)
		GetClusterFailErrMocked(t, clusterIn)
		GetClusterFailStatusMocked(t, clusterIn)
		GetClusterFailJSONMocked(t, clusterIn)
	}
}

func TestCreateCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range clustersIn {
		CreateClusterMocked(t, clusterIn)
		CreateClusterFailErrMocked(t, clusterIn)
		CreateClusterFailStatusMocked(t, clusterIn)
		CreateClusterFailJSONMocked(t, clusterIn)
	}
}

func TestUpdateCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range clustersIn {
		UpdateClusterMocked(t, clusterIn)
		UpdateClusterFailErrMocked(t, clusterIn)
		UpdateClusterFailStatusMocked(t, clusterIn)
		UpdateClusterFailJSONMocked(t, clusterIn)
	}
}

func TestDeleteCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range clustersIn {
		DeleteClusterMocked(t, clusterIn)
		DeleteClusterFailErrMocked(t, clusterIn)
		DeleteClusterFailStatusMocked(t, clusterIn)
		DeleteClusterFailJSONMocked(t, clusterIn)
	}
}

func TestRetryCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range clustersIn {
		RetryClusterMocked(t, clusterIn)
		RetryClusterFailErrMocked(t, clusterIn)
		RetryClusterFailStatusMocked(t, clusterIn)
		RetryClusterFailJSONMocked(t, clusterIn)
	}
}

func TestDiscardCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range clustersIn {
		DiscardClusterMocked(t, clusterIn)
		DiscardClusterFailErrMocked(t, clusterIn)
		DiscardClusterFailStatusMocked(t, clusterIn)
	}
}

func TestGetClusterPlan(t *testing.T) {
	clusterPlansIn := testdata.GetClusterPlanData()
	for _, clusterPlanIn := range clusterPlansIn {
		GetClusterPlanMocked(t, clusterPlanIn.ID, clusterPlanIn)
		GetClusterPlanFailErrMocked(t, clusterPlanIn.ID, clusterPlanIn)
		GetClusterPlanFailStatusMocked(t, clusterPlanIn.ID, clusterPlanIn)
		GetClusterPlanFailJSONMocked(t, clusterPlanIn.ID, clusterPlanIn)
	}
}
