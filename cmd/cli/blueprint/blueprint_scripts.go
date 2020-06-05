// Copyright (c) 2017-2022 Ingram Micro Inc.

package blueprint

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"
	"strings"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/spf13/viper"
)

func init() {
	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Script Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the script"}
	fNameReq := fName
	fNameReq.Required = true

	fDescription := cmd.FlagContext{Type: cmd.String, Name: cmd.Description,
		Usage: "Description of the script's purpose"}
	fDescriptionReq := fDescription
	fDescriptionReq.Required = true

	fCode := cmd.FlagContext{Type: cmd.String, Name: cmd.Code, Usage: "The script's code"}
	fCodeReq := fCode
	fCodeReq.Required = true

	fParameters := cmd.FlagContext{Type: cmd.String, Name: cmd.Parameters,
		Usage: "A comma-separated list of names of the script parameters"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with script"}

	fAttachmentName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true,
		Usage: "Name of the attachment"}

	fAttachmentPath := cmd.FlagContext{Type: cmd.String, Name: cmd.Filepath, Required: true,
		Usage: "Path to attachment file"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true,
		Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "script", Hidden: true,
		Usage: "Resource Type"}

	scriptsCmd := cmd.NewCommand(blueprintCmd, &cmd.CommandContext{
		Use:   "scripts",
		Short: "Allow the user to manage the scripts they want to run on the servers",
	},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all available scripts",
		RunMethod:    ScriptsList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about a specific script",
		RunMethod:    ScriptShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new script to be used in the templates",
		RunMethod:    ScriptCreate,
		FlagContexts: []cmd.FlagContext{fNameReq, fDescriptionReq, fCodeReq, fParameters, fLabels}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing script",
		RunMethod:    ScriptUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName, fDescription, fCode, fParameters}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a script",
		RunMethod:    ScriptDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "add-attachment",
		Short:        "Adds an attachment to a script",
		RunMethod:    ScriptAttachmentAdd,
		FlagContexts: []cmd.FlagContext{fId, fAttachmentName, fAttachmentPath}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "list-attachments",
		Short:        "List the attachments a script has",
		RunMethod:    ScriptAttachmentList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(scriptsCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// ScriptsList subcommand function
func ScriptsList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	scripts, err := svc.ListScripts(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}

	labelables := make([]types.Labelable, len(scripts))
	for i := 0; i < len(scripts); i++ {
		labelables[i] = types.Labelable(scripts[i])
	}

	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	filteredLabelables := labels.LabelFiltering(labelables, labelIDsByName)
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	scripts = make([]*types.Script, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		s, ok := labelable.(*types.Script)
		if !ok {
			formatter.PrintFatal(cmd.LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.Script, got a %T", labelable))
		}
		scripts[i] = s
	}

	if err = formatter.PrintList(scripts); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// ScriptShow subcommand function
func ScriptShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	script, err := svc.GetScript(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	script.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// ScriptCreate subcommand function
func ScriptCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	scriptIn := map[string]interface{}{
		cmd.Name:        viper.GetString(cmd.Name),
		cmd.Description: viper.GetString(cmd.Description),
		cmd.Code:        viper.GetString(cmd.Code),
	}
	if viper.IsSet(cmd.Parameters) {
		scriptIn[cmd.Parameters] = strings.Split(viper.GetString(cmd.Parameters), ",")
	}

	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()

	if viper.IsSet(cmd.Labels) {
		scriptIn["label_ids"] = labels.LabelResolution(viper.GetString(cmd.Labels), &labelNamesByID, &labelIDsByName)
	}

	script, err := svc.CreateScript(cmd.GetContext(), &scriptIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create script", err)
	}

	script.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// ScriptUpdate subcommand function
func ScriptUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	scriptIn := map[string]interface{}{}
	if viper.IsSet(cmd.Name) {
		scriptIn[cmd.Name] = viper.GetString(cmd.Name)
	}
	if viper.IsSet(cmd.Description) {
		scriptIn[cmd.Description] = viper.GetString(cmd.Description)
	}
	if viper.IsSet(cmd.Code) {
		scriptIn[cmd.Code] = viper.GetString(cmd.Code)
	}
	if viper.IsSet(cmd.Parameters) {
		scriptIn[cmd.Parameters] = strings.Split(viper.GetString(cmd.Parameters), ",")
	}

	script, err := svc.UpdateScript(cmd.GetContext(), viper.GetString(cmd.Id), &scriptIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update script", err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	script.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// ScriptDelete subcommand function
func ScriptDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteScript(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't delete script", err)
	}
	return nil
}

// ScriptAttachmentAdd subcommand function
func ScriptAttachmentAdd() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	sourceFilePath := viper.GetString(cmd.Filepath)
	if !utils.FileExists(sourceFilePath) {
		formatter.PrintFatal("Invalid file path", fmt.Errorf("no such file or directory: %s", sourceFilePath))
	}

	attachmentIn := map[string]interface{}{
		cmd.Name: viper.GetString(cmd.Name),
	}

	// adds new attachment
	attachment, err := svc.AddScriptAttachment(cmd.GetContext(), viper.GetString(cmd.Id), &attachmentIn)
	if err != nil {
		formatter.PrintFatal("Couldn't add attachment to script", err)
	}

	// uploads new attachment file
	err = svc.UploadFile(cmd.GetContext(), sourceFilePath, attachment.UploadURL)
	if err != nil {
		cleanAttachment(attachment.ID)
		formatter.PrintFatal("Couldn't upload attachment data", err)
	}

	// marks the attachment as "uploaded"
	attachmentID := attachment.ID
	attachment, err = svc.UploadedScriptAttachment(cmd.GetContext(), attachment.ID, &attachmentIn)
	if err != nil {
		cleanAttachment(attachmentID)
		formatter.PrintFatal("Couldn't set attachment as uploaded", err)
	}

	if err = formatter.PrintItem(*attachment); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// cleanAttachment deletes Attachment. Ideally for cleaning at uploading error cases
func cleanAttachment(attachmentID string) {
	svc, _, formatter := cli.WireUpAPIClient()
	if err := svc.DeleteAttachment(cmd.GetContext(), attachmentID); err != nil {
		formatter.PrintError("Couldn't clean failed attachment", err)
	}
}

// ScriptAttachmentList subcommand function
func ScriptAttachmentList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	attachments, err := svc.ListScriptAttachments(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't list attachments script", err)
	}

	if err = formatter.PrintList(attachments); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
