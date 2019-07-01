package blueprint

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetBootstrappingConfigurationMocked test mocked function
func GetBootstrappingConfigurationMocked(t *testing.T, bcConfIn *types.BootstrappingConfiguration) *types.BootstrappingConfiguration {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dIn, err := json.Marshal(bcConfIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	// call service
	cs.On("Get", "/blueprint/configuration").Return(dIn, 200, nil)
	bcConfOut, status, err := ds.GetBootstrappingConfiguration()
	assert.Nil(err, "Error getting bootstrapping configuration")
	assert.Equal(status, 200, "GetBootstrappingConfiguration returned invalid response")
	assert.Equal(*bcConfIn, *bcConfOut, "GetBootstrappingConfiguration returned different services")
	return bcConfOut
}

// GetBootstrappingConfigurationFailErrMocked test mocked function
func GetBootstrappingConfigurationFailErrMocked(t *testing.T, bcConfIn *types.BootstrappingConfiguration) *types.BootstrappingConfiguration {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dIn, err := json.Marshal(bcConfIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	// call service
	cs.On("Get", "/blueprint/configuration").Return(dIn, 404, fmt.Errorf("mocked error"))
	bcConfOut, _, err := ds.GetBootstrappingConfiguration()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(bcConfOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return bcConfOut
}

// GetBootstrappingConfigurationFailStatusMocked test mocked function
func GetBootstrappingConfigurationFailStatusMocked(t *testing.T, bcConfIn *types.BootstrappingConfiguration) *types.BootstrappingConfiguration {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dIn, err := json.Marshal(bcConfIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	// call service
	cs.On("Get", "/blueprint/configuration").Return(dIn, 499, nil)
	bcConfOut, status, err := ds.GetBootstrappingConfiguration()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(bcConfOut, "Expecting nil output")
	assert.Equal(499, status, "Expecting http code 499")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return bcConfOut
}

// GetBootstrappingConfigurationFailJSONMocked test mocked function
func GetBootstrappingConfigurationFailJSONMocked(t *testing.T, bcConfIn *types.BootstrappingConfiguration) *types.BootstrappingConfiguration {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/blueprint/configuration").Return(dIn, 200, nil)
	bcConfOut, _, err := ds.GetBootstrappingConfiguration()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(bcConfOut, "Expecting nil output")

	return bcConfOut
}

// ReportBootstrappingAppliedConfigurationMocked test mocked function
func ReportBootstrappingAppliedConfigurationMocked(t *testing.T, commandIn *types.BootstrappingConfiguration) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dOut, err := json.Marshal(commandIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/blueprint/applied_configuration"), &payload).Return(dOut, 200, nil)
	err = ds.ReportBootstrappingAppliedConfiguration(&payload)
	assert.Nil(err, "Error getting bootstrapping command")
}

// ReportBootstrappingAppliedConfigurationFailErrMocked test mocked function
func ReportBootstrappingAppliedConfigurationFailErrMocked(t *testing.T, commandIn *types.BootstrappingConfiguration) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/blueprint/applied_configuration"), &payload).Return(dIn, 499, fmt.Errorf("mocked error"))
	err = ds.ReportBootstrappingAppliedConfiguration(&payload)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// ReportBootstrappingAppliedConfigurationFailStatusMocked test mocked function
func ReportBootstrappingAppliedConfigurationFailStatusMocked(t *testing.T, commandIn *types.BootstrappingConfiguration) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/blueprint/applied_configuration"), &payload).Return(dIn, 499, fmt.Errorf("error 499 Mocked error"))
	err = ds.ReportBootstrappingAppliedConfiguration(&payload)
	assert.NotNil(err, "We are expecting a status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// ReportBootstrappingAppliedConfigurationFailJSONMocked test mocked function
func ReportBootstrappingAppliedConfigurationFailJSONMocked(t *testing.T, commandIn *types.BootstrappingConfiguration) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// wrong json
	dIn := []byte{0}

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/blueprint/applied_configuration"), &payload).Return(dIn, 499, nil)
	err = ds.ReportBootstrappingAppliedConfiguration(&payload)
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// ReportBootstrappingLogMocked test mocked function
func ReportBootstrappingLogMocked(t *testing.T, commandIn *types.BootstrappingContinuousReport) *types.BootstrappingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dOut, err := json.Marshal(commandIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/blueprint/bootstrap_logs"), &payload).Return(dOut, 200, nil)
	commandOut, status, err := ds.ReportBootstrappingLog(&payload)

	assert.Nil(err, "Error posting report command")
	assert.Equal(status, 200, "ReportBootstrappingLog returned invalid response")
	assert.Equal(commandOut.Stdout, "Bootstrap log created", "ReportBootstrapLog returned unexpected message")

	return commandOut
}

// ReportBootstrappingLogFailErrMocked test mocked function
func ReportBootstrappingLogFailErrMocked(t *testing.T, commandIn *types.BootstrappingContinuousReport) *types.BootstrappingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/blueprint/bootstrap_logs"), &payload).Return(dIn, 499, fmt.Errorf("mocked error"))
	commandOut, _, err := ds.ReportBootstrappingLog(&payload)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return commandOut
}

// ReportBootstrappingLogFailStatusMocked test mocked function
func ReportBootstrappingLogFailStatusMocked(t *testing.T, commandIn *types.BootstrappingContinuousReport) *types.BootstrappingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "Bootstrapping test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/blueprint/bootstrap_logs"), &payload).Return(dIn, 499, fmt.Errorf("error 499 Mocked error"))
	commandOut, status, err := ds.ReportBootstrappingLog(&payload)

	assert.Equal(status, 499, "ReportBootstrappingLog returned an unexpected status code")
	assert.NotNil(err, "We are expecting a status code error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return commandOut
}

// ReportBootstrappingLogFailJSONMocked test mocked function
func ReportBootstrappingLogFailJSONMocked(t *testing.T, commandIn *types.BootstrappingContinuousReport) *types.BootstrappingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/blueprint/bootstrap_logs"), &payload).Return(dIn, 200, nil)
	commandOut, _, err := ds.ReportBootstrappingLog(&payload)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return commandOut
}

// DownloadPolicyfileMocked test mocked function
func DownloadPolicyfileMocked(t *testing.T, dataIn map[string]string) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	urlSource := dataIn["fakeURLToFile"]
	pathFile := dataIn["fakeFileDownloadFile"]

	// call service
	cs.On("GetFile", urlSource, pathFile).Return(pathFile, 200, nil)
	realFileName, status, err := ds.DownloadPolicyfile(urlSource, pathFile)
	assert.Nil(err, "Error downloading bootstrapping policy file")
	assert.Equal(status, 200, "DownloadPolicyfile returned invalid response")
	assert.Equal(realFileName, pathFile, "Invalid downloaded file path")
}

// DownloadPolicyfileFailErrMocked test mocked function
func DownloadPolicyfileFailErrMocked(t *testing.T, dataIn map[string]string) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBootstrappingService(cs)
	assert.Nil(err, "Couldn't load bootstrapping service")
	assert.NotNil(ds, "Bootstrapping service not instanced")

	urlSource := dataIn["fakeURLToFile"]
	pathFile := dataIn["fakeFileDownloadFile"]

	// call service
	cs.On("GetFile", urlSource, pathFile).Return("", 499, fmt.Errorf("mocked error"))
	_, status, err := ds.DownloadPolicyfile(urlSource, pathFile)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(status, 499, "DownloadPolicyfile returned an unexpected status code")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}
