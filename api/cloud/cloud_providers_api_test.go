package cloud

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCloudProviderServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudProviderService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetCloudProviderList(t *testing.T) {
	cloudProvidersIn := testdata.GetCloudProviderData()
	GetCloudProviderListMocked(t, cloudProvidersIn)
	GetCloudProviderListFailErrMocked(t, cloudProvidersIn)
	GetCloudProviderListFailStatusMocked(t, cloudProvidersIn)
	GetCloudProviderListFailJSONMocked(t, cloudProvidersIn)
}

func TestGetServerStoragePlanList(t *testing.T) {
	storagePlansIn := testdata.GetStoragePlanData()
	GetServerStoragePlanListMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	GetServerStoragePlanListFailErrMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	GetServerStoragePlanListFailStatusMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	GetServerStoragePlanListFailJSONMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
}

func TestListLoadBalancerPlans(t *testing.T) {
	loadBalancerPlansIn := testdata.GetLoadBalancerPlanData()
	ListLoadBalancerPlansMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
	ListLoadBalancerPlansFailErrMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
	ListLoadBalancerPlansFailStatusMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
	ListLoadBalancerPlansFailJSONMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
}
