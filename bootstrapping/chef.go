package bootstrapping

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ingrammicro/cio/api/blueprint"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
)

// Subsidiary routine for commands processing
func applyChefPolicyfiles(
	ctx context.Context,
	bsProcess *bootstrappingProcess,
	bootstrappingSvc *blueprint.BootstrappingService,
	formatter format.Formatter,
) error {
	// Process tarballs policies
	err := processPolicyfiles(bootstrappingSvc, bsProcess)
	// Finishing time
	bsProcess.finishedAt = time.Now().UTC()

	// Inform the platform of applied changes via a `PUT /blueprint/applied_configuration` request with a JSON payload
	// similar to
	log.Debug("reporting applied policy files")
	reportErr := reportAppliedConfiguration(bootstrappingSvc, bsProcess)
	if reportErr != nil {
		formatter.PrintError("couldn't report applied status for policy files", err)
		return err
	}
	return err
}

// saveAttributes stores the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
func saveAttributes(bsProcess *bootstrappingProcess, policyfileName string) error {
	log.Debug("saveAttributes")
	bsProcess.attributes.rawData["policy_group"] = "local"
	bsProcess.attributes.rawData["policy_name"] = policyfileName
	attrs, err := json.Marshal(bsProcess.attributes.rawData)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(bsProcess.attributes.FilePath(bsProcess.directoryPath), attrs, 0600); err != nil {
		return err
	}
	return nil
}

func runCommand(fn func(chunk string) error, command string, thresholdLines int) error {
	log.Debug("runCommand")
	exitCode, err := utils.RunContinuousCmd(fn, command, -1, thresholdLines)
	if err == nil && exitCode != 0 {
		err = fmt.Errorf("policyfile application exited with %d code", exitCode)
	}
	if err != nil {
		return err
	}
	log.Info("completed: ", exitCode)
	return nil
}

// processPolicyfiles applies for each policy the required chef commands, reporting in bunches of N lines
func processPolicyfiles(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) error {
	log.Debug("processPolicyfiles")
	for _, bsPolicyfile := range bsProcess.policyfiles {
		command, renamedPolicyfileDir, policyfileDir, err := preparePolicyfileCommand(bsProcess, bsPolicyfile)
		if err != nil {
			return err
		}
		log.Debug(command)
		bsProcess.cmsVersion = ""
		// Custom method for chunks processing
		fn := getBootstrapLogReporter(bootstrappingSvc, bsProcess)
		if err = runCommand(fn, command, bsProcess.thresholdLines); err != nil {
			return err
		}

		bsProcess.appliedPolicyfileRevisionIDs[bsPolicyfile.ID] = bsPolicyfile.RevisionID

		if renamedPolicyfileDir != "" {
			err = os.Rename(policyfileDir, renamedPolicyfileDir)
			if err != nil {
				return fmt.Errorf("could not rename %s as %s back: %v", policyfileDir, renamedPolicyfileDir, err)
			}
		}
	}
	return nil
}

func preparePolicyfileCommand(bsProcess *bootstrappingProcess, bsPolicyfile policyfile) (
	string, string, string, error,
) {
	log.Debug("preparePolicyfileCommand")
	// Store the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
	err := saveAttributes(bsProcess, bsPolicyfile.ID)
	if err != nil {
		return "", "", "", fmt.Errorf("couldn't save attributes for policy file %q: %w", bsPolicyfile.ID, err)
	}
	command := fmt.Sprintf("chef-client -z -j %s", bsProcess.attributes.FilePath(bsProcess.directoryPath))
	policyfileDir := bsPolicyfile.Path(bsProcess.directoryPath)
	var renamedPolicyfileDir string
	if runtime.GOOS == "windows" {
		renamedPolicyfileDir = policyfileDir
		policyfileDir = filepath.Join(bsProcess.directoryPath, "active")
		err := os.Rename(renamedPolicyfileDir, policyfileDir)
		if err != nil {
			return "", "", "", fmt.Errorf("could not rename %s as %s: %v", renamedPolicyfileDir, policyfileDir, err)
		}
		command = fmt.Sprintf(
			"SET \"PATH=%%PATH%%;C:\\ruby\\bin;C:\\cinc-project\\cinc\\bin;C:\\cinc-project\\cinc\\embedded\\bin;"+
				"C:\\opscode\\chef\\bin;C:\\opscode\\chef\\embedded\\bin\"\n%s",
			command,
		)
	}
	command = fmt.Sprintf("cd %s\n%s", policyfileDir, command)
	return command, renamedPolicyfileDir, policyfileDir, nil
}
