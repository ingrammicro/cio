package wizard

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWizServerPlanServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewWizServerPlanService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetWizServerPlanList(t *testing.T) {
	AppID := "fakeAppID"
	LocID := "fakeLocID"
	ProviderID := "fakeProviderID"
	serverPlansIn := testdata.GetServerPlanData()
	GetWizServerPlanListMocked(t, serverPlansIn, AppID, LocID, ProviderID)
	GetWizServerPlanListFailErrMocked(t, serverPlansIn, AppID, LocID, ProviderID)
	GetWizServerPlanListFailStatusMocked(t, serverPlansIn, AppID, LocID, ProviderID)
	GetWizServerPlanListFailJSONMocked(t, serverPlansIn, AppID, LocID, ProviderID)
}
