package network

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ListTargetGroupsMocked test mocked function
func ListTargetGroupsMocked(t *testing.T, loadBalancerID string, targetGroupsIn []*types.TargetGroup) []*types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupsIn)
	assert.Nil(err, "TargetGroups test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID)).Return(dIn, 200, nil)
	targetGroupsOut, err := ds.ListTargetGroups(loadBalancerID)

	assert.Nil(err, "Error getting target groups")
	assert.Equal(targetGroupsIn, targetGroupsOut, "ListTargetGroups returned different target groups")

	return targetGroupsOut
}

// ListTargetGroupsFailErrMocked test mocked function
func ListTargetGroupsFailErrMocked(t *testing.T, loadBalancerID string, targetGroupsIn []*types.TargetGroup) []*types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupsIn)
	assert.Nil(err, "TargetGroups test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	targetGroupsOut, err := ds.ListTargetGroups(loadBalancerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetGroupsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return targetGroupsOut
}

// ListTargetGroupsFailStatusMocked test mocked function
func ListTargetGroupsFailStatusMocked(t *testing.T, loadBalancerID string, targetGroupsIn []*types.TargetGroup) []*types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupsIn)
	assert.Nil(err, "TargetGroups test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID)).Return(dIn, 499, nil)
	targetGroupsOut, err := ds.ListTargetGroups(loadBalancerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetGroupsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return targetGroupsOut
}

// ListTargetGroupsFailJSONMocked test mocked function
func ListTargetGroupsFailJSONMocked(t *testing.T, loadBalancerID string, targetGroupsIn []*types.TargetGroup) []*types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID)).Return(dIn, 200, nil)
	targetGroupsOut, err := ds.ListTargetGroups(loadBalancerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetGroupsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return targetGroupsOut
}

// GetTargetGroupMocked test mocked function
func GetTargetGroupMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 200, nil)
	targetGroupOut, err := ds.GetTargetGroup(targetGroupIn.ID)

	assert.Nil(err, "Error getting target group")
	assert.Equal(*targetGroupIn, *targetGroupOut, "GetTargetGroup returned different target group")

	return targetGroupOut
}

// GetTargetGroupFailErrMocked test mocked function
func GetTargetGroupFailErrMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	targetGroupOut, err := ds.GetTargetGroup(targetGroupIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return targetGroupOut
}

// GetTargetGroupFailStatusMocked test mocked function
func GetTargetGroupFailStatusMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 499, nil)
	targetGroupOut, err := ds.GetTargetGroup(targetGroupIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return targetGroupOut
}

// GetTargetGroupFailJSONMocked test mocked function
func GetTargetGroupFailJSONMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 200, nil)
	targetGroupOut, err := ds.GetTargetGroup(targetGroupIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return targetGroupOut
}

// CreateTargetGroupMocked test mocked function
func CreateTargetGroupMocked(t *testing.T, loadBalancerID string, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID), mapIn).Return(dOut, 200, nil)
	targetGroupOut, err := ds.CreateTargetGroup(loadBalancerID, mapIn)

	assert.Nil(err, "Error creating target group")
	assert.Equal(targetGroupIn, targetGroupOut, "CreateTargetGroup returned different target group")

	return targetGroupOut
}

// CreateTargetGroupFailErrMocked test mocked function
func CreateTargetGroupFailErrMocked(t *testing.T, loadBalancerID string, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	targetGroupOut, err := ds.CreateTargetGroup(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return targetGroupOut
}

// CreateTargetGroupFailStatusMocked test mocked function
func CreateTargetGroupFailStatusMocked(t *testing.T, loadBalancerID string, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID), mapIn).Return(dOut, 499, nil)
	targetGroupOut, err := ds.CreateTargetGroup(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return targetGroupOut
}

// CreateTargetGroupFailJSONMocked test mocked function
func CreateTargetGroupFailJSONMocked(t *testing.T, loadBalancerID string, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID), mapIn).Return(dIn, 200, nil)
	targetGroupOut, err := ds.CreateTargetGroup(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return targetGroupOut
}

// UpdateTargetGroupMocked test mocked function
func UpdateTargetGroupMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID), mapIn).Return(dOut, 200, nil)
	targetGroupOut, err := ds.UpdateTargetGroup(targetGroupIn.ID, mapIn)

	assert.Nil(err, "Error updating target group")
	assert.Equal(targetGroupIn, targetGroupOut, "UpdateTargetGroup returned different target group")

	return targetGroupOut
}

// UpdateTargetGroupFailErrMocked test mocked function
func UpdateTargetGroupFailErrMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	targetGroupOut, err := ds.UpdateTargetGroup(targetGroupIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return targetGroupOut
}

// UpdateTargetGroupFailStatusMocked test mocked function
func UpdateTargetGroupFailStatusMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID), mapIn).Return(dOut, 499, nil)
	targetGroupOut, err := ds.UpdateTargetGroup(targetGroupIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return targetGroupOut
}

// UpdateTargetGroupFailJSONMocked test mocked function
func UpdateTargetGroupFailJSONMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID), mapIn).Return(dIn, 200, nil)
	targetGroupOut, err := ds.UpdateTargetGroup(targetGroupIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return targetGroupOut
}

// DeleteTargetGroupMocked test mocked function
func DeleteTargetGroupMocked(t *testing.T, targetGroupIn *types.TargetGroup) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 200, nil)
	targetGroupOut, err := ds.DeleteTargetGroup(targetGroupIn.ID)

	assert.Nil(err, "Error deleting target group")
	assert.Equal(targetGroupIn, targetGroupOut, "DeleteTargetGroup returned different target group")

}

// DeleteTargetGroupFailErrMocked test mocked function
func DeleteTargetGroupFailErrMocked(t *testing.T, targetGroupIn *types.TargetGroup) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	targetGroupOut, err := ds.DeleteTargetGroup(targetGroupIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteTargetGroupFailStatusMocked test mocked function
func DeleteTargetGroupFailStatusMocked(t *testing.T, targetGroupIn *types.TargetGroup) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 499, nil)
	targetGroupOut, err := ds.DeleteTargetGroup(targetGroupIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteTargetGroupFailJSONMocked test mocked function
func DeleteTargetGroupFailJSONMocked(t *testing.T, targetGroupIn *types.TargetGroup) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf("/network/target_groups/%s", targetGroupIn.ID)).Return(dIn, 200, nil)
	targetGroupOut, err := ds.DeleteTargetGroup(targetGroupIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")
}

// RetryTargetGroupMocked test mocked function
func RetryTargetGroupMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s/retry", targetGroupIn.ID), mapIn).Return(dOut, 200, nil)
	targetGroupOut, err := ds.RetryTargetGroup(targetGroupIn.ID, mapIn)

	assert.Nil(err, "Error retrying target group")
	assert.Equal(targetGroupIn, targetGroupOut, "RetryTargetGroup returned different target group")

	return targetGroupOut
}

// RetryTargetGroupFailErrMocked test mocked function
func RetryTargetGroupFailErrMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s/retry", targetGroupIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	targetGroupOut, err := ds.RetryTargetGroup(targetGroupIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return targetGroupOut
}

// RetryTargetGroupFailStatusMocked test mocked function
func RetryTargetGroupFailStatusMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// to json
	dOut, err := json.Marshal(targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s/retry", targetGroupIn.ID), mapIn).Return(dOut, 499, nil)
	targetGroupOut, err := ds.RetryTargetGroup(targetGroupIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return targetGroupOut
}

// RetryTargetGroupFailJSONMocked test mocked function
func RetryTargetGroupFailJSONMocked(t *testing.T, targetGroupIn *types.TargetGroup) *types.TargetGroup {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetGroupIn)
	assert.Nil(err, "TargetGroup test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/target_groups/%s/retry", targetGroupIn.ID), mapIn).Return(dIn, 200, nil)
	targetGroupOut, err := ds.RetryTargetGroup(targetGroupIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetGroupOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return targetGroupOut
}

// ListTargetsMocked test mocked function
func ListTargetsMocked(t *testing.T, targetGroupID string, targetsIn []*types.Target) []*types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetsIn)
	assert.Nil(err, "Targets test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID)).Return(dIn, 200, nil)
	targetsOut, err := ds.ListTargets(targetGroupID)

	assert.Nil(err, "Error getting targets")
	assert.Equal(targetsIn, targetsOut, "ListTargets returned different targets")

	return targetsOut
}

// ListTargetsFailErrMocked test mocked function
func ListTargetsFailErrMocked(t *testing.T, targetGroupID string, targetsIn []*types.Target) []*types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetsIn)
	assert.Nil(err, "Targets test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	targetsOut, err := ds.ListTargets(targetGroupID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return targetsOut
}

// ListTargetsFailStatusMocked test mocked function
func ListTargetsFailStatusMocked(t *testing.T, targetGroupID string, targetsIn []*types.Target) []*types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetsIn)
	assert.Nil(err, "Targets test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID)).Return(dIn, 499, nil)
	targetsOut, err := ds.ListTargets(targetGroupID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return targetsOut
}

// ListTargetsFailJSONMocked test mocked function
func ListTargetsFailJSONMocked(t *testing.T, targetGroupID string, targetsIn []*types.Target) []*types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID)).Return(dIn, 200, nil)
	targetsOut, err := ds.ListTargets(targetGroupID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return targetsOut
}

// CreateTargetMocked test mocked function
func CreateTargetMocked(t *testing.T, targetGroupID string, targetIn *types.Target) *types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetIn)
	assert.Nil(err, "Target test data corrupted")

	// to json
	dOut, err := json.Marshal(targetIn)
	assert.Nil(err, "Target test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID), mapIn).Return(dOut, 200, nil)
	targetOut, err := ds.CreateTarget(targetGroupID, mapIn)

	assert.Nil(err, "Error creating target")
	assert.Equal(targetIn, targetOut, "CreateTarget returned different target")

	return targetOut
}

// CreateTargetFailErrMocked test mocked function
func CreateTargetFailErrMocked(t *testing.T, targetGroupID string, targetIn *types.Target) *types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetIn)
	assert.Nil(err, "Target test data corrupted")

	// to json
	dOut, err := json.Marshal(targetIn)
	assert.Nil(err, "Target test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	targetOut, err := ds.CreateTarget(targetGroupID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(targetOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return targetOut
}

// CreateTargetFailStatusMocked test mocked function
func CreateTargetFailStatusMocked(t *testing.T, targetGroupID string, targetIn *types.Target) *types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetIn)
	assert.Nil(err, "Target test data corrupted")

	// to json
	dOut, err := json.Marshal(targetIn)
	assert.Nil(err, "Target test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID), mapIn).Return(dOut, 499, nil)
	targetOut, err := ds.CreateTarget(targetGroupID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(targetOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return targetOut
}

// CreateTargetFailJSONMocked test mocked function
func CreateTargetFailJSONMocked(t *testing.T, targetGroupID string, targetIn *types.Target) *types.Target {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*targetIn)
	assert.Nil(err, "Target test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID), mapIn).Return(dIn, 200, nil)
	targetOut, err := ds.CreateTarget(targetGroupID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(targetOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return targetOut
}

// DeleteTargetMocked test mocked function
func DeleteTargetMocked(t *testing.T, targetGroupID string, targetIn *types.Target) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetIn)
	assert.Nil(err, "Target test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/target_groups/%s/targets/%s/%s", targetGroupID, targetIn.ResourceType, targetIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteTarget(targetGroupID, targetIn.ResourceType, targetIn.ID)

	assert.Nil(err, "Error deleting target")
}

// DeleteTargetFailErrMocked test mocked function
func DeleteTargetFailErrMocked(t *testing.T, targetGroupID string, targetIn *types.Target) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetIn)
	assert.Nil(err, "Target test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/target_groups/%s/targets/%s/%s", targetGroupID, targetIn.ResourceType, targetIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteTarget(targetGroupID, targetIn.ResourceType, targetIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteTargetFailStatusMocked test mocked function
func DeleteTargetFailStatusMocked(t *testing.T, targetGroupID string, targetIn *types.Target) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTargetGroupService(cs)
	assert.Nil(err, "Couldn't load targetGroup service")
	assert.NotNil(ds, "TargetGroup service not instanced")

	// to json
	dIn, err := json.Marshal(targetIn)
	assert.Nil(err, "Target test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/target_groups/%s/targets/%s/%s", targetGroupID, targetIn.ResourceType, targetIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteTarget(targetGroupID, targetIn.ResourceType, targetIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
