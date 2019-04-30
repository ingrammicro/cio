package cloud

import (
	"testing"

	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewServerServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewServerService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetServerList(t *testing.T) {
	serversIn := testdata.GetServerData()
	GetServerListMocked(t, serversIn)
	GetServerListFailErrMocked(t, serversIn)
	GetServerListFailStatusMocked(t, serversIn)
	GetServerListFailJSONMocked(t, serversIn)
}

func TestGetServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		GetServerMocked(t, serverIn)
		GetServerFailErrMocked(t, serverIn)
		GetServerFailStatusMocked(t, serverIn)
		GetServerFailJSONMocked(t, serverIn)
	}
}

func TestCreateServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		CreateServerMocked(t, serverIn)
		CreateServerFailErrMocked(t, serverIn)
		CreateServerFailStatusMocked(t, serverIn)
		CreateServerFailJSONMocked(t, serverIn)
	}
}

func TestUpdateServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		UpdateServerMocked(t, serverIn)
		UpdateServerFailErrMocked(t, serverIn)
		UpdateServerFailStatusMocked(t, serverIn)
		UpdateServerFailJSONMocked(t, serverIn)

	}
}

func TestBootServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		BootServerMocked(t, serverIn)
		BootServerFailErrMocked(t, serverIn)
		BootServerFailStatusMocked(t, serverIn)
		BootServerFailJSONMocked(t, serverIn)
	}
}

func TestRebootServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		RebootServerMocked(t, serverIn)
		RebootServerFailErrMocked(t, serverIn)
		RebootServerFailStatusMocked(t, serverIn)
		RebootServerFailJSONMocked(t, serverIn)
	}
}

func TestShutdownServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		ShutdownServerMocked(t, serverIn)
		ShutdownServerFailErrMocked(t, serverIn)
		ShutdownServerFailStatusMocked(t, serverIn)
		ShutdownServerFailJSONMocked(t, serverIn)
	}
}

func TestOverrideServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		OverrideServerMocked(t, serverIn)
		OverrideServerFailErrMocked(t, serverIn)
		OverrideServerFailStatusMocked(t, serverIn)
		OverrideServerFailJSONMocked(t, serverIn)
	}
}

func TestDeleteServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {
		DeleteServerMocked(t, serverIn)
		DeleteServerFailErrMocked(t, serverIn)
		DeleteServerFailStatusMocked(t, serverIn)
	}
}

//======= Events ==========
func TestGetEventsList(t *testing.T) {
	serversIn := testdata.GetServerData()
	eventsIn := testdata.GetEventData()
	for _, serverIn := range serversIn {
		GetServerEventListMocked(t, eventsIn, serverIn.ID)
		GetServerEventListFailErrMocked(t, eventsIn, serverIn.ID)
		GetServerEventListFailStatusMocked(t, eventsIn, serverIn.ID)
		GetServerEventListFailJSONMocked(t, eventsIn, serverIn.ID)
	}
}

//======= Operational Scripts ==========
func TestGetOperationalScriptList(t *testing.T) {
	serversIn := testdata.GetServerData()
	scriptsIn := testdata.GetScriptCharData()
	for _, serverIn := range serversIn {
		GetOperationalScriptListMocked(t, scriptsIn, serverIn.ID)
		GetOperationalScriptFailErrMocked(t, scriptsIn, serverIn.ID)
		GetOperationalScriptFailStatusMocked(t, scriptsIn, serverIn.ID)
		GetOperationalScriptFailJSONMocked(t, scriptsIn, serverIn.ID)
	}
}

func TestExecuteOperationalScript(t *testing.T) {
	serversIn := testdata.GetServerData()
	scriptsIn := testdata.GetScriptCharData()
	eventDataIn := testdata.GetEventData()
	for _, serverIn := range serversIn {
		for _, scriptIn := range scriptsIn {
			ExecuteOperationalScriptListMocked(t, scriptIn, serverIn.ID, eventDataIn[0])
			ExecuteOperationalScriptFailErrMocked(t, scriptIn, serverIn.ID)
			ExecuteOperationalScriptFailStatusMocked(t, scriptIn, serverIn.ID)
			ExecuteOperationalScriptFailJSONMocked(t, scriptIn, serverIn.ID)
		}
	}
}
