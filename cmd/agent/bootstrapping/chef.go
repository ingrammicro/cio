// Copyright (c) 2017-2022 Ingram Micro Inc.

package bootstrapping

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd/agent"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

// Subsidiary routine for commands processing
func applyChefPolicyfiles(
	ctx context.Context,
	blueprintConfig *types.BootstrappingConfiguration,
	svc *api.ServerAPI,
	bsProcess *bootstrappingProcess,
	formatter format.Formatter,
) error {
	logger.DebugFuncInfo()

	// Process tarballs policies
	return processChefPolicyfiles(ctx, blueprintConfig, svc, bsProcess)

}

// saveAttributes stores the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
func saveAttributes(bsProcess *bootstrappingProcess, policyfileName string) error {
	logger.DebugFuncInfo()

	bsProcess.attributes.rawData["policy_group"] = "local"
	bsProcess.attributes.rawData["policy_name"] = policyfileName
	attrs, err := json.Marshal(bsProcess.attributes.rawData)
	if err != nil {
		return err
	}
	if err := os.WriteFile(bsProcess.attributes.FilePath(bsProcess.directoryPath), attrs, 0600); err != nil {
		return err
	}
	return nil
}

func runCommand(fn func(chunk string) error, command string, thresholdLines int) error {
	logger.DebugFuncInfo()

	exitCode, err := agent.RunContinuousCmd(fn, command, -1, thresholdLines)
	if err == nil && exitCode != 0 {
		err = fmt.Errorf("policyfile application exited with %d code", exitCode)
	}
	if err != nil {
		return err
	}
	log.Info("completed: ", exitCode)
	return nil
}

// processChefPolicyfiles applies for each policy the required chef commands, reporting in bunches of N lines
func processChefPolicyfiles(
	ctx context.Context,
	blueprintConfig *types.BootstrappingConfiguration,
	svc *api.ServerAPI,
	bsProcess *bootstrappingProcess,
) error {
	logger.DebugFuncInfo()

	for _, bsPolicyfile := range bsProcess.policyfiles {
		command, renamedPolicyfileDir, policyfileDir, err := preparePolicyfileCommand(bsProcess, bsPolicyfile)
		if err != nil {
			return err
		}
		log.Debug(command)
		bsProcess.cmsVersion = ""
		// Custom method for chunks processing
		fn := getBootstrapLogReporter(ctx, svc, bsProcess, blueprintConfig)
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
	logger.DebugFuncInfo()

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
