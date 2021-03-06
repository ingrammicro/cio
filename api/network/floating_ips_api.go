// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathNetworkFloatingIPs = "/network/floating_ips"
const APIPathCloudServerFloatingIPs = "/cloud/servers/%s/floating_ips"
const APIPathNetworkFloatingIP = "/network/floating_ips/%s"
const APIPathNetworkFloatingIPAttachedServer = "/network/floating_ips/%s/attached_server"
const APIPathNetworkFloatingIPDiscard = "/network/floating_ips/%s/discard"

// FloatingIPService manages floating IP operations
type FloatingIPService struct {
	concertoService utils.ConcertoService
}

// NewFloatingIPService returns a Concerto FloatingIP service
func NewFloatingIPService(concertoService utils.ConcertoService) (*FloatingIPService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &FloatingIPService{
		concertoService: concertoService,
	}, nil
}

// ListFloatingIPs returns the list of FloatingIPs as an array of FloatingIP
func (fips *FloatingIPService) ListFloatingIPs(serverID string) (floatingIPs []*types.FloatingIP, err error) {
	log.Debug("ListFloatingIPs")

	path := APIPathNetworkFloatingIPs
	if serverID != "" {
		path = fmt.Sprintf(APIPathCloudServerFloatingIPs, serverID)

	}
	data, status, err := fips.concertoService.Get(path)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIPs); err != nil {
		return nil, err
	}

	return floatingIPs, nil
}

// GetFloatingIP returns a FloatingIP by its ID
func (fips *FloatingIPService) GetFloatingIP(floatingIPID string) (floatingIP *types.FloatingIP, err error) {
	log.Debug("GetFloatingIP")

	data, status, err := fips.concertoService.Get(fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIP); err != nil {
		return nil, err
	}

	return floatingIP, nil
}

// CreateFloatingIP creates a FloatingIP
func (fips *FloatingIPService) CreateFloatingIP(
	floatingIPParams *map[string]interface{},
) (floatingIP *types.FloatingIP, err error) {
	log.Debug("CreateFloatingIP")

	data, status, err := fips.concertoService.Post(APIPathNetworkFloatingIPs, floatingIPParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIP); err != nil {
		return nil, err
	}

	return floatingIP, nil
}

// UpdateFloatingIP updates a FloatingIP by its ID
func (fips *FloatingIPService) UpdateFloatingIP(
	floatingIPID string,
	floatingIPParams *map[string]interface{},
) (floatingIP *types.FloatingIP, err error) {
	log.Debug("UpdateFloatingIP")

	data, status, err := fips.concertoService.Put(
		fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPID),
		floatingIPParams,
	)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIP); err != nil {
		return nil, err
	}

	return floatingIP, nil
}

// AttachFloatingIP attaches a FloatingIP by its ID
func (fips *FloatingIPService) AttachFloatingIP(
	floatingIPID string,
	floatingIPParams *map[string]interface{},
) (server *types.Server, err error) {
	log.Debug("AttachFloatingIP")

	data, status, err := fips.concertoService.Post(
		fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPID),
		floatingIPParams,
	)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// DetachFloatingIP detaches a FloatingIP by its ID
func (fips *FloatingIPService) DetachFloatingIP(floatingIPID string) (err error) {
	log.Debug("DetachFloatingIP")

	data, status, err := fips.concertoService.Delete(
		fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPID),
	)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DeleteFloatingIP deletes a FloatingIP by its ID
func (fips *FloatingIPService) DeleteFloatingIP(floatingIPID string) (err error) {
	log.Debug("DeleteFloatingIP")

	data, status, err := fips.concertoService.Delete(fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DiscardFloatingIP discards a FloatingIP by its ID
func (fips *FloatingIPService) DiscardFloatingIP(floatingIPID string) (err error) {
	log.Debug("DiscardFloatingIP")

	data, status, err := fips.concertoService.Delete(fmt.Sprintf(APIPathNetworkFloatingIPDiscard, floatingIPID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
