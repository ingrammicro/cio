// Copyright (c) 2017-2022 Ingram Micro Inc.

package testdata

import (
	"github.com/ingrammicro/cio/types"
)

// GetGenericImageData loads test data
func GetGenericImageData() []*types.GenericImage {

	return []*types.GenericImage{
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
