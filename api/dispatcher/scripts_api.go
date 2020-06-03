package dispatcher

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// DispatcherService manages dispatcher operations
type DispatcherService struct {
	concertoService utils.ConcertoService
}

// NewDispatcherService returns a dispatcher service
func NewDispatcherService(concertoService utils.ConcertoService) (*DispatcherService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &DispatcherService{
		concertoService: concertoService,
	}, nil
}

// GetDispatcherScriptCharacterizationsByType returns script characterizations list for a given phase
func (ds *DispatcherService) GetDispatcherScriptCharacterizationsByType(phase string) (scriptCharacterizations []*types.ScriptCharacterization, err error) {
	log.Debug("GetDispatcherScriptCharacterizationsByType")

	data, status, err := ds.concertoService.Get(fmt.Sprintf("/blueprint/script_characterizations?type=%s", phase))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &scriptCharacterizations); err != nil {
		return nil, err
	}

	return scriptCharacterizations, nil
}

// GetDispatcherScriptCharacterizationByUUID returns script characterizations list for a given UUID
func (ds *DispatcherService) GetDispatcherScriptCharacterizationByUUID(scriptCharacterizationUUID string) (*types.ScriptCharacterization, error) {
	log.Debug("GetDispatcherScriptCharacterizationByUUID")

	data, status, err := ds.concertoService.Get(fmt.Sprintf("/blueprint/script_characterizations/%s", scriptCharacterizationUUID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}
	var scriptCharacterization types.ScriptCharacterization
	if err = json.Unmarshal(data, &scriptCharacterization); err != nil {
		return nil, err
	}

	return &scriptCharacterization, nil
}

// ReportScriptConclusions reports a result
func (ds *DispatcherService) ReportScriptConclusions(scriptConclusions *map[string]interface{}) (command *types.ScriptConclusion, status int, err error) {
	log.Debug("ReportScriptConclusions")

	data, status, err := ds.concertoService.Post("/blueprint/script_conclusions", scriptConclusions)
	if err != nil {
		return nil, status, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}

// DownloadAttachment gets a file from given url saving file into given file path
func (ds *DispatcherService) DownloadAttachment(url string, filePath string) (realFileName string, status int, err error) {
	log.Debug("DownloadAttachment")

	realFileName, status, err = ds.concertoService.GetFile(url, filePath, true)
	if err != nil {
		return realFileName, status, err
	}

	return realFileName, status, nil
}
