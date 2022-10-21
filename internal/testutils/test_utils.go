// Copyright (c) 2017-2022 Ingram Micro Inc.

package testutils

import (
	"encoding/json"
	"encoding/xml"
	"github.com/ingrammicro/cio/configuration"
	"net/http"
	"net/http/httptest"
	"time"
)

// TODO COMMON
func InitConfig() *configuration.Config {
	config := new(configuration.Config)
	config.XMLName = xml.Name{
		Space: "",
		Local: "",
	}
	config.APIEndpoint = "https://clients.test.imco.io/v3"
	config.LogFile = "/var/log/concerto-client.log"
	config.LogLevel = "info"
	config.Certificate = configuration.Cert{
		Cert: "../../../configuration/testdata/ssl/cert.crt",
		Key:  "../../../configuration/testdata/ssl/private/cert.key",
		Ca:   "../../../configuration/testdata/ssl/ca_cert.pem",
	}
	config.BootstrapConfig = configuration.BootstrapConfig{
		IntervalSeconds:      600,
		SplaySeconds:         300,
		ApplyAfterIterations: 4,
		RunOnce:              false,
	}
	config.ConfLocation = TEST
	config.ConfFile = "../../../configuration/testdata/client.xml"
	config.ConfFileLastLoadedAt = time.Now()
	config.IsHost = false
	config.ConcertoURL = TEST
	config.BrownfieldToken = ""
	config.CommandPollingToken = ""
	config.ServerID = TEST
	config.CurrentUserName = TEST
	config.CurrentUserIsAdmin = false
	return config
}

const TEST = "test"

func NewServer(returnStatus int, returnData any) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if r.URL.Path != "/storage/volumes" {
		//	t.Errorf("Expected to request '/storage/volumes', got: %s", r.URL.Path)
		//}
		//if r.Header.Get("Accept") != "application/json" {
		//	t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		//}
		w.WriteHeader(returnStatus)
		d, _ := json.Marshal(returnData)
		w.Write(d)
	}))
	return server
}

func AddHandleFunc(mux *http.ServeMux, pattern string, status int, v any) {
	mux.HandleFunc(pattern,
		func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(status)
			d, _ := json.Marshal(v)
			res.Write(d)
		})
}

type Value struct {
	Status int
	V      any
}

func AddHandleFuncMultiple(mux *http.ServeMux, pattern string, variants map[string]Value) {
	mux.HandleFunc(pattern, func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			value := variants["GET"]
			res.WriteHeader(value.Status)
			d, _ := json.Marshal(value.V)
			res.Write(d)
		}
		if req.Method == "POST" {
			value := variants["POST"]
			res.WriteHeader(value.Status)
			d, _ := json.Marshal(value.V)
			res.Write(d)
		}
	})
}
