package testdata

import "github.com/ingrammicro/cio/api/types"

// GetCloudApplicationTemplateData loads test data
func GetCloudApplicationTemplateData() []*types.CloudApplicationTemplate {
	return []*types.CloudApplicationTemplate{
		{
			ID:        "fakeID0",
			Name:      "fakeName0",
			Version:   "0.0",
			Global:    false,
			UploadURL: "fakeUploadURL0",
			VendorID:  "fakeVendorID0",
			Inputs:    nil,
			IsMock:    false,
		},
		{
			ID:        "fakeID1",
			Name:      "fakeName1",
			Version:   "0.1",
			Global:    true,
			UploadURL: "fakeUploadURL1",
			VendorID:  "fakeVendorID1",
			Inputs:    nil,
			IsMock:    true,
		},
	}
}
