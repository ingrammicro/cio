// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import (
	"github.com/ingrammicro/cio/api/types"
)

// GetLocationData loads test data
func GetLocationData() []*types.Location {

	return []*types.Location{
		{
			ID:   "fakeID0",
			Name: "fakeName0",
		},
		{
			ID:   "fakeID1",
			Name: "fakeName1",
		},
	}
}
