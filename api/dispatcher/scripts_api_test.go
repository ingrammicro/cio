package dispatcher

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewDispatcherServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewDispatcherService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetDispatcherScriptCharacterizationsByType(t *testing.T) {
	dIn := testdata.GetScriptCharacterizationsData()
	GetDispatcherScriptCharacterizationsByTypeMocked(t, "boot", dIn)
	GetDispatcherScriptCharacterizationsByTypeFailErrMocked(t, "boot", dIn)
	GetDispatcherScriptCharacterizationsByTypeFailStatusMocked(t, "boot", dIn)
	GetDispatcherScriptCharacterizationsByTypeFailJSONMocked(t, "boot", dIn)
}

func TestGetDispatcherScriptCharacterizationByUUID(t *testing.T) {
	dIn := testdata.GetScriptCharacterizationsData()
	GetDispatcherScriptCharacterizationByUUIDMocked(t, "fakeUUID1", dIn[0])
	GetDispatcherScriptCharacterizationByUUIDFailErrMocked(t, "fakeUUID1", dIn[0])
	GetDispatcherScriptCharacterizationByUUIDFailStatusMocked(t, "fakeUUID1", dIn[0])
	GetDispatcherScriptCharacterizationByUUIDFailJSONMocked(t, "fakeUUID1", dIn[0])
}

func TestReportScriptConclusions(t *testing.T) {
	dIn := testdata.GetScriptConclusionData()
	ReportScriptConclusionsMocked(t, dIn)
	ReportScriptConclusionsFailErrMocked(t, dIn)
	ReportScriptConclusionsFailStatusMocked(t, dIn)
	ReportScriptConclusionsFailJSONMocked(t, dIn)
}

func TestDownloadAttachment(t *testing.T) {
	dataIn := testdata.GetDownloadAttachmentData()
	DownloadAttachmentMocked(t, dataIn)
	DownloadAttachmentFailErrMocked(t, dataIn)
}
