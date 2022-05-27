// Copyright (c) 2017-2022 Ingram Micro Inc.

package types

import (
	"fmt"
	"github.com/ingrammicro/cio/configuration"
	"strings"
)

type TemporaryArchive struct {
	ID           string `json:"id"            header:"ID"`
	UploadURL    string `json:"upload_url"    header:"UPLOAD_URL"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type TemporaryArchiveImport struct {
	ID           string `json:"id"            header:"ID"`
	Type         string `json:"type"          header:"TYPE"`
	LabelName    string `json:"label_name"    header:"LABEL_NAME"`
	LabelID      string `json:"label_id"      header:"LABEL_ID"`
	State        string `json:"state"         header:"STATE"`
	Outputs      string `json:"outputs"       header:"OUTPUTS"       show:"nolist"`
	UserID       string `json:"user_id"       header:"USER_ID"`
	ArchiveID    string `json:"archive_id"    header:"ARCHIVE_ID"`
	ErrorMessage string `json:"error_message" header:"ERROR_MESSAGE" show:"nolist"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type TemporaryArchiveExport struct {
	ID     string `json:"id"      header:"ID"`
	TaskID string `json:"task_id" header:"TASK_ID"`
}

// TemporaryArchiveExportTask tosca_task
type TemporaryArchiveExportTask struct {
	ID           string `json:"id"            header:"ID"`
	Type         string `json:"type"          header:"TYPE"`
	State        string `json:"state"         header:"STATE"`
	FileAccess   string `json:"file_access"   header:"FILE_ACCESS"`
	UserID       string `json:"user_id"       header:"USER_ID"`
	ArchiveID    string `json:"archive_id"    header:"ARCHIVE_ID"`
	ErrorMessage string `json:"error_message" header:"ERROR_MESSAGE" show:"nolist"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

// DownloadURL prepares adequate download url
func (taet *TemporaryArchiveExportTask) DownloadURL(configAPIEndpoint string) string {
	path := strings.Replace(
		configAPIEndpoint,
		strings.Join([]string{"/", configuration.VERSION_API_USER_MODE}, ""), "", -1)
	return fmt.Sprintf("%s%s", path, taet.FileAccess)
}
