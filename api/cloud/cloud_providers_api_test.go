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

func TestListCloudProviders(t *testing.T) {
	cloudProvidersIn := testdata.GetCloudProviderData()
	ListCloudProvidersMocked(t, cloudProvidersIn)
	ListCloudProvidersFailErrMocked(t, cloudProvidersIn)
	ListCloudProvidersFailStatusMocked(t, cloudProvidersIn)
	ListCloudProvidersFailJSONMocked(t, cloudProvidersIn)
}

func TestListServerStoragePlans(t *testing.T) {
	storagePlansIn := testdata.GetStoragePlanData()
	ListServerStoragePlansMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	ListServerStoragePlansFailErrMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	ListServerStoragePlansFailStatusMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	ListServerStoragePlansFailJSONMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
}

func TestListLoadBalancerPlans(t *testing.T) {
	loadBalancerPlansIn := testdata.GetLoadBalancerPlanData()
	ListLoadBalancerPlansMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
	ListLoadBalancerPlansFailErrMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
	ListLoadBalancerPlansFailStatusMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
	ListLoadBalancerPlansFailJSONMocked(t, loadBalancerPlansIn, loadBalancerPlansIn[0].CloudProviderID)
}
