// Copyright (c) 2017-2022 Ingram Micro Inc.

package configuration

// Mode contextual working mode: Client (User) / Server (Agent)
type (
	Mode = int
)

const (
	Server Mode = iota + 1
	Client
)

type (
	Context = int
)

const (
	Brownfield Context = iota + 1
	Polling
)

// Internal fields
const (
	CloudOrchestratorPlatformName = "IMCO"
	ConcertoEnvVarPrefixName      = "concerto"
	NAME                          = "CIO"
	AUTHOR                        = "Ingram Micro"
	EMAIL                         = "https://github.com/ingrammicro/cio"
)

// config / build
const (
	WindowsServerConfigFile = "c:\\cio\\client.xml"
	NixServerConfigFile     = "/etc/cio/client.xml"
	DefaultEndpoint         = "https://clients.concerto.io/"

	WindowsServerLogFilePath = "c:\\cio\\log\\concerto-client.log"
	WindowsServerCaCertPath  = "c:\\cio\\client_ssl\\ca_cert.pem"
	WindowsServerCertPath    = "c:\\cio\\client_ssl\\cert.pem"
	WindowsServerKeyPath     = "c:\\cio\\client_ssl\\private\\key.pem"
	NixServerLogFilePath     = "/var/log/concerto-client.log"
	NixServerCaCertPath      = "/etc/cio/client_ssl/ca_cert.pem"
	NixServerCertPath        = "/etc/cio/client_ssl/cert.pem"
	NixServerKeyPath         = "/etc/cio/client_ssl/private/key.pem"
)

// global flags
const (
	CaCert           = "ca-cert"
	ClientCert       = "client-cert"
	ClientKey        = "client-key"
	ConcertoConfig   = "concerto-config"
	ConcertoEndpoint = "concerto-endpoint"
	ConcertoUrl      = "concerto-url"

	ConcertoBrownfieldToken     = "concerto-brownfield-token"
	ConcertoCommandPollingToken = "concerto-command-polling-token"
	ConcertoServerId            = "concerto-server-id"

	Formatter = "formatter"
	Debug     = "debug"
)
