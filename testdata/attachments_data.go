// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import (
	"github.com/ingrammicro/cio/api/types"
)

// GetAttachmentData loads test data
func GetAttachmentData() []*types.Attachment {
	return []*types.Attachment{
		{
			ID:          "fakeID0",
			Name:        "fakeName0",
			UploadURL:   "fakeUploadURL0",
			DownloadURL: "fakeDownloadURL0",
			Uploaded:    true,
		},
		{
			ID:          "fakeID1",
			Name:        "fakeName1",
			UploadURL:   "fakeUploadURL1",
			DownloadURL: "fakeDownloadURL1",
			Uploaded:    false,
		},
	}
}
