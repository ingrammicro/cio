package network

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// TargetGroupService manages target group operations
type TargetGroupService struct {
	concertoService utils.ConcertoService
}

// NewTargetGroupService returns a Concerto target group service
func NewTargetGroupService(concertoService utils.ConcertoService) (*TargetGroupService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &TargetGroupService{
		concertoService: concertoService,
	}, nil
}

// ListTargetGroups returns the list of target groups in a load balancer by its ID, as an array of TargetGroup
func (tgs *TargetGroupService) ListTargetGroups(loadBalancerID string) (targetGroups []*types.TargetGroup, err error) {
	log.Debug("ListTargetGroups")

	data, status, err := tgs.concertoService.Get(fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &targetGroups); err != nil {
		return nil, err
	}

	return targetGroups, nil
}

// GetTargetGroup returns a target group by its ID
func (tgs *TargetGroupService) GetTargetGroup(targetGroupID string) (targetGroup *types.TargetGroup, err error) {
	log.Debug("GetTargetGroup")

	data, status, err := tgs.concertoService.Get(fmt.Sprintf("/network/target_groups/%s", targetGroupID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &targetGroup); err != nil {
		return nil, err
	}

	return targetGroup, nil
}

// CreateTargetGroup creates a target group in a load balancer by its ID
func (tgs *TargetGroupService) CreateTargetGroup(loadBalancerID string, targetGroupParams *map[string]interface{}) (targetGroup *types.TargetGroup, err error) {
	log.Debug("CreateTargetGroup")

	data, status, err := tgs.concertoService.Post(fmt.Sprintf("/network/load_balancers/%s/target_groups", loadBalancerID), targetGroupParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &targetGroup); err != nil {
		return nil, err
	}

	return targetGroup, nil
}

// UpdateTargetGroup updates a target group by its ID
func (tgs *TargetGroupService) UpdateTargetGroup(targetGroupID string, targetGroupParams *map[string]interface{}) (targetGroup *types.TargetGroup, err error) {
	log.Debug("UpdateTargetGroup")

	data, status, err := tgs.concertoService.Put(fmt.Sprintf("/network/target_groups/%s", targetGroupID), targetGroupParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &targetGroup); err != nil {
		return nil, err
	}

	return targetGroup, nil
}

// DeleteTargetGroup deletes a target group by its ID
func (tgs *TargetGroupService) DeleteTargetGroup(targetGroupID string) (targetGroup *types.TargetGroup, err error) {
	log.Debug("DeleteTargetGroup")

	data, status, err := tgs.concertoService.Delete(fmt.Sprintf("/network/target_groups/%s", targetGroupID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &targetGroup); err != nil {
		return nil, err
	}

	return targetGroup, nil
}

// RetryTargetGroup retries a target group by its ID
func (tgs *TargetGroupService) RetryTargetGroup(targetGroupID string, targetGroupParams *map[string]interface{}) (targetGroup *types.TargetGroup, err error) {
	log.Debug("RetryTargetGroup")

	data, status, err := tgs.concertoService.Put(fmt.Sprintf("/network/target_groups/%s/retry", targetGroupID), targetGroupParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &targetGroup); err != nil {
		return nil, err
	}

	return targetGroup, nil
}

// ListTargets returns the list of targets in a target group by its ID, as an array of Target
func (tgs *TargetGroupService) ListTargets(targetGroupID string) (targets []*types.Target, err error) {
	log.Debug("ListTargets")

	data, status, err := tgs.concertoService.Get(fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &targets); err != nil {
		return nil, err
	}

	return targets, nil
}

// CreateTarget creates a target in a target group by its ID
func (tgs *TargetGroupService) CreateTarget(targetGroupID string, targetParams *map[string]interface{}) (target *types.Target, err error) {
	log.Debug("CreateTarget")

	data, status, err := tgs.concertoService.Post(fmt.Sprintf("/network/target_groups/%s/targets", targetGroupID), targetParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	return target, nil
}

// DeleteTarget deletes a target in a target group by given IDs and resource type
func (tgs *TargetGroupService) DeleteTarget(targetGroupID string, targetResourceType string, targetResourceID string) (err error) {
	log.Debug("DeleteTarget")

	data, status, err := tgs.concertoService.Delete(fmt.Sprintf("/network/target_groups/%s/targets/%s/%s", targetGroupID, targetResourceType, targetResourceID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
