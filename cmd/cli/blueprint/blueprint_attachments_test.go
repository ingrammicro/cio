package blueprint

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/ingrammicro/cio/types"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAttachmentShow(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Attachment),
			server:   testutils.NewServer(http.StatusOK, new(types.Attachment)),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = testutils.TEST
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := AttachmentShow()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

//// TODO
//func TestAttachmentDownload(t *testing.T) {
//
//	att := new(types.Attachment)
//	//att.DownloadURL = "testdata/test_utils.txt"
//	tests := map[string]struct {
//		expected any
//		server   *httptest.Server
//	}{
//		"if defined endpoint for API service is resolving properly": {
//			expected: new(types.Attachment),
//			server:   testutils.NewServer(http.StatusOK, att),
//		},
//		//"if defined endpoint for API service is invalid or cannot be reached": {
//		//	expected: "Cannot execute request",
//		//	server:   nil,
//		//},
//	}
//
//	config := testutils.InitConfig()
//	configuration.SetConfig(config)
//	cmd.RootCmd.SetContext(context.Background())
//
//	for title, test := range tests {
//		t.Run(title, func(t *testing.T) {
//			config.APIEndpoint = testutils.TEST
//			if test.server != nil {
//				server := test.server
//				//server = testutils.NewServer(http.StatusOK, att)
//				//att.DownloadURL = server.URL + "/testdata/test.txt"
//				defer server.Close()
//				config.APIEndpoint = server.URL + "/testdata/test.txt"
//				server.URL = server.URL + "/testdata/test.txt"
//
//			}
//			viper.Set(cmd.Filepath, "testdata/test_utils.txt")
//			err := AttachmentDownload()
//			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
//				t.Errorf("Unexpected error: %v\n", err)
//			}
//		})
//	}
//}

func TestAttachmentDelete(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Attachment),
			server:   testutils.NewServer(http.StatusOK, new(types.Attachment)),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = testutils.TEST
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := AttachmentDelete()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}
