package wizard

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWizardServerPlanServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewWizardServerPlanService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListWizardServerPlans(t *testing.T) {
	AppID := "fakeAppID"
	LocID := "fakeLocID"
	ProviderID := "fakeProviderID"
	serverPlansIn := testdata.GetServerPlanData()
	ListWizardServerPlansMocked(t, serverPlansIn, AppID, LocID, ProviderID)
	ListWizardServerPlansFailErrMocked(t, serverPlansIn, AppID, LocID, ProviderID)
	ListWizardServerPlansFailStatusMocked(t, serverPlansIn, AppID, LocID, ProviderID)
	ListWizardServerPlansFailJSONMocked(t, serverPlansIn, AppID, LocID, ProviderID)
}
