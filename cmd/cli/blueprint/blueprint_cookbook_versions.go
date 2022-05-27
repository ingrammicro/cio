// Copyright (c) 2017-2022 Ingram Micro Inc.

package blueprint

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Cookbook version Id"}

	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with cookbook version"}

	fFilepath := cmd.FlagContext{Type: cmd.String, Name: cmd.Filepath, Required: true,
		Usage: "path to cookbook version file"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "cookbook_version",
		Hidden: true, Usage: "Resource Type"}

	cookbookVersionsCmd := cmd.NewCommand(blueprintCmd, &cmd.CommandContext{
		Use:     "cookbook-versions",
		Short:   "Provides information on chef cookbook versions",
		Aliases: []string{"cv"}},
	)
	cmd.NewCommand(cookbookVersionsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all available cookbook versions",
		RunMethod:    CookbookVersionList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(cookbookVersionsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about a specific cookbook version",
		RunMethod:    CookbookVersionShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(cookbookVersionsCmd, &cmd.CommandContext{
		Use:          "upload",
		Short:        "Uploads a new cookbook version",
		RunMethod:    CookbookVersionUpload,
		FlagContexts: []cmd.FlagContext{fFilepath, fLabels}},
	)
	cmd.NewCommand(cookbookVersionsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a cookbook version",
		RunMethod:    CookbookVersionDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(cookbookVersionsCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(cookbookVersionsCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// CookbookVersionList subcommand function
func CookbookVersionList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cookbookVersions, err := svc.ListCookbookVersions(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive cookbook versions data", err)
	}

	labelables := make([]types.Labelable, len(cookbookVersions))
	for i := 0; i < len(cookbookVersions); i++ {
		labelables[i] = types.Labelable(cookbookVersions[i])
	}
	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	filteredLabelables := labels.LabelFiltering(labelables, labelIDsByName)
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)
	cookbookVersions = make([]*types.CookbookVersion, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		cb, ok := labelable.(*types.CookbookVersion)
		if !ok {
			formatter.PrintFatal(cmd.LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.CookbookVersion, got a %T", labelable))
		}
		cookbookVersions[i] = cb
	}

	if err = formatter.PrintList(cookbookVersions); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CookbookVersionShow subcommand function
func CookbookVersionShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cookbookVersion, err := svc.GetCookbookVersion(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cookbook version data", err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	cookbookVersion.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cookbookVersion); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CookbookVersionUpload subcommand function
func CookbookVersionUpload() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	sourceFilePath := viper.GetString(cmd.Filepath)

	if !utils.FileExists(sourceFilePath) {
		formatter.PrintFatal("Invalid file path", fmt.Errorf("no such file or directory: %s", sourceFilePath))
	}

	cbIn := map[string]interface{}{}
	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	if viper.IsSet(cmd.Labels) {
		cbIn["label_ids"] = labels.LabelResolution(viper.GetString(cmd.Labels), &labelNamesByID, &labelIDsByName)
	}

	// creates new cookbook_version
	cookbookVersion, err := svc.CreateCookbookVersion(cmd.GetContext(), &cbIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create cookbook version data", err)
	}

	// uploads new cookbook_version file
	err = svc.UploadFile(cmd.GetContext(), sourceFilePath, cookbookVersion.UploadURL)
	if err != nil {
		cleanCookbookVersion(cookbookVersion.ID)
		formatter.PrintFatal("Couldn't upload cookbook version data", err)
	}

	// processes the new cookbook_version
	cookbookVersionID := cookbookVersion.ID
	cookbookVersionParams := new(map[string]interface{})
	cookbookVersion, err = svc.ProcessCookbookVersion(cmd.GetContext(), cookbookVersion.ID, cookbookVersionParams)
	if err != nil {
		cleanCookbookVersion(cookbookVersionID)
		formatter.PrintFatal("Couldn't process cookbook version", err)
	}

	cookbookVersion.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cookbookVersion); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}

	return nil
}

// cleanCookbookVersion deletes CookbookVersion. Ideally for cleaning at uploading error cases
func cleanCookbookVersion(cookbookVersionID string) {
	svc, _, formatter := cli.WireUpAPIClient()
	if err := svc.DeleteCookbookVersion(cmd.GetContext(), cookbookVersionID); err != nil {
		formatter.PrintError("Couldn't clean failed cookbook version", err)
	}
}

// CookbookVersionDelete subcommand function
func CookbookVersionDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteCookbookVersion(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cookbook version", err)
	}
	return nil
}
