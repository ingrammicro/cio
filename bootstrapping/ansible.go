package bootstrapping

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ingrammicro/cio/api/blueprint"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	inventoryFile = "inventory.yml"
	variableFile  = "variables.yml"
	ansibleScript = "apply.sh"
)

// Subsidiary routine for commands processing
func applyAnsiblePolicyfiles(
	ctx context.Context,
	bsProcess *bootstrappingProcess,
	bootstrappingSvc *blueprint.BootstrappingService,
	formatter format.Formatter,
) error {
	err := prepareAnsibleInventory(ctx, bsProcess)
	if err != nil {
		formatter.PrintError("couldn't prepare inventory:", err)
		return err
	}
	err = prepareAnsibleVariables(ctx, bsProcess)
	if err != nil {
		formatter.PrintError("couldn't prepare variables:", err)
		return err
	}
	err = processAnsiblePolicyfiles(bootstrappingSvc, bsProcess)
	if err != nil {
		formatter.PrintError("couldn't process policyfiles:", err)
		return err
	}
	return nil
}

func prepareAnsibleInventory(ctx context.Context, bsProcess *bootstrappingProcess) error {
	log.Debug("prepareAnsibleInventory")
	file, err := os.Create(inventoryFilePath(bsProcess.directoryPath))
	if err != nil {
		return fmt.Errorf("opening inventory file to write: %w", err)
	}
	defer file.Close()
	inventory := map[string]interface{}{
		"all": map[string]interface{}{
			"hosts": map[string]interface{}{
				"localhost": map[string]interface{}{
					"ansible_connection": "local",
				},
			},
		},
	}
	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	err = encoder.Encode(inventory)
	if err != nil {
		return fmt.Errorf("encoding inventory: %w", err)
	}
	return nil
}

func prepareAnsibleVariables(ctx context.Context, bsProcess *bootstrappingProcess) error {
	log.Debug("prepareAnsibleVariables")
	file, err := os.Create(variableFilePath(bsProcess.directoryPath))
	if err != nil {
		return fmt.Errorf("opening variable file to write: %w", err)
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	err = encoder.Encode(bsProcess.attributes.rawData)
	if err != nil {
		return fmt.Errorf("encoding variables: %w", err)
	}
	return nil
}

func inventoryFilePath(dir string) string {
	return filepath.Join(dir, inventoryFile)
}

func variableFilePath(dir string) string {
	return filepath.Join(dir, variableFile)
}

// processAnsiblePolicyfiles applies for each policy the required ansible-galaxy and ansible-playbook commands, reporting in bunches of N lines
func processAnsiblePolicyfiles(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) error {
	log.Debug("processAnsiblePolicyfiles")
	for _, bsPolicyfile := range bsProcess.policyfiles {
		policyfileDir := bsPolicyfile.Path(bsProcess.directoryPath)
		command := fmt.Sprintf(
			"cd %s && sh %s %s %s",
			policyfileDir,
			ansibleScript, inventoryFilePath(bsProcess.directoryPath), variableFilePath(bsProcess.directoryPath))
		log.Debug(command)
		bsProcess.cmsVersion = ""
		// Custom method for chunks processing
		fn := getBootstrapLogReporter(bootstrappingSvc, bsProcess)
		if err := runCommand(fn, command, bsProcess.thresholdLines); err != nil {
			return err
		}
		bsProcess.appliedPolicyfileRevisionIDs[bsPolicyfile.ID] = bsPolicyfile.RevisionID
	}
	return nil
}
