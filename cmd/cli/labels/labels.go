// Copyright (c) 2017-2022 Ingram Micro Inc.

package labels

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"
	"regexp"
	"strings"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/spf13/viper"
)

func init() {
	labelsCmd := cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "labels",
		Short:   "Provides information about labels",
		Aliases: []string{"lbl"}},
	)
	cmd.NewCommand(labelsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists the current labels existing in the platform for the user",
		RunMethod: LabelList},
	)
}

// LabelList subcommand function
func LabelList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	labels, err := svc.ListLabels(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive labels data", err)
		return err
	}

	if err = formatter.PrintList(labels); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LabelFiltering subcommand function receives a collection of references to labelable objects
// Evaluates the matching of assigned labels with the labels requested for filtering.
func LabelFiltering(items []types.Labelable, labelIDsByName map[string]string) ([]types.Labelable, error) {
	logger.DebugFuncInfo()

	labels := viper.GetString(cmd.Labels)
	if labels != "" {
		_, _, formatter := cli.WireUpAPIClient()
		labelNamesIn, err := LabelsUnifyInputNames(labels, formatter)
		if err != nil {
			return nil, err
		}
		var filteringLabelIDs []string
		for _, name := range labelNamesIn {
			id := labelIDsByName[name]
			filteringLabelIDs = append(filteringLabelIDs, id)
		}
		var result []types.Labelable
		for _, item := range items {
			if item.FilterByLabelIDs(filteringLabelIDs) {
				result = append(result, item)
			}
		}
		return result, nil
	}

	return items, nil
}

// LabelAssignNamesForIDs subcommand function receives a collection of references to labelables objects
// Resolves the Labels names associated to each resource from given Labels ids, loading object with respective labels
// names
func LabelAssignNamesForIDs(items []types.Labelable, labelNamesByID map[string]string) {
	logger.DebugFuncInfo()
	for _, labelable := range items {
		labelable.FillInLabelNames(labelNamesByID)
	}
}

// LabelLoadsMapping subcommand function retrieves the current label list in IMCO; then prepares two mapping structures
// (Name <-> ID and ID <-> Name)
func LabelLoadsMapping() (map[string]string, map[string]string, error) {
	logger.DebugFuncInfo()

	svc, _, formatter := cli.WireUpAPIClient()
	labels, err := svc.ListLabels(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive labels data", err)
		return nil, nil, err
	}

	labelIDsByName := make(map[string]string)
	labelNamesByID := make(map[string]string)

	for _, label := range labels {
		labelIDsByName[label.Name] = label.ID
		labelNamesByID[label.ID] = label.Name
	}
	return labelIDsByName, labelNamesByID, nil
}

// LabelsUnifyInputNames subcommand function evaluates the received labels names (comma separated string).
// Validates, remove duplicates and resolves a slice with unique label names.
func LabelsUnifyInputNames(labelsNames string, formatter format.Formatter) ([]string, error) {
	labelNamesIn := utils.RemoveDuplicates(strings.Split(labelsNames, ","))
	for _, c := range labelNamesIn {
		matched := regexp.MustCompile(`^[A-Za-z0-9 .\s_-]+$`).MatchString(c)
		if !matched {
			e := fmt.Errorf(
				"invalid label format: %v (Labels would be indicated with their name, "+
					"which must satisfy to be composed of spaces, underscores, dots, dashes "+
					"and/or lower/upper -case alphanumeric characters-)",
				c,
			)
			formatter.PrintError("Invalid label name ", e)
			return nil, e
		}
	}
	return labelNamesIn, nil
}

// LabelResolution subcommand function retrieves a labels map(Name<->ID) based on label names received to be processed.
// The function evaluates the received labels names (comma separated string); with them, solves the assigned IDs for the
// given labels names.
// If the label name is not available in IMCO yet, it is created.
// If new label is created, mapping structures labelNamesByID/labelIDsByName are updated
func LabelResolution(
	labelsNames string,
	labelNamesByID *map[string]string,
	labelIDsByName *map[string]string,
) ([]string, error) {
	logger.DebugFuncInfo()

	svc, _, formatter := cli.WireUpAPIClient()
	labelNamesIn, err := LabelsUnifyInputNames(labelsNames, formatter)
	if err != nil {
		return nil, err
	}

	// Obtain output mapped labels Name<->ID; currently in IMCO platform as well as if creation is required
	labelsOutMap := make(map[string]string)
	for _, name := range labelNamesIn {
		// check if the label already exists in IMCO, creates it if it does not exist
		if (*labelIDsByName)[name] == "" {
			labelPayload := make(map[string]interface{})
			labelPayload["name"] = name
			newLabel, err := svc.CreateLabel(cmd.GetContext(), &labelPayload)
			if err != nil {
				formatter.PrintError("Couldn't create label", err)
				return nil, err
			}
			labelsOutMap[name] = newLabel.ID
			// updates the mapping!
			(*labelIDsByName)[name] = newLabel.ID
			(*labelNamesByID)[newLabel.ID] = name
		} else {
			labelsOutMap[name] = (*labelIDsByName)[name]
		}
	}
	labelIds := make([]string, 0)
	for _, mp := range labelsOutMap {
		labelIds = append(labelIds, mp)
	}
	return labelIds, nil
}

// LabelAdd subcommand function assigns a single label from a single labelable resource
func LabelAdd() error {
	logger.DebugFuncInfo()

	svc, _, formatter := cli.WireUpAPIClient()

	labelIDsByName, labelNamesByID, err := LabelLoadsMapping()
	if err != nil {
		return err
	}
	labelIds, err := LabelResolution(viper.GetString(cmd.Label), &labelNamesByID, &labelIDsByName)
	if err != nil {
		return err
	}
	if len(labelIds) > 1 {
		e := fmt.Errorf("invalid parameter: %v - %v", viper.GetString(cmd.Label), labelIds)
		formatter.PrintError("Too many label names. Please, Use only one label name", e)
		return e
	}
	labelID := labelIds[0]

	resData := make(map[string]string)
	resData["id"] = viper.GetString(cmd.Id)
	resData["resource_type"] = viper.GetString(cmd.ResourceType)
	resourcesData := make([]interface{}, 0, 1)
	resourcesData = append(resourcesData, resData)

	labelIn := map[string]interface{}{
		"resources": resourcesData,
	}

	labeledResources, err := svc.AddLabel(cmd.GetContext(), labelID, &labelIn)
	if err != nil {
		formatter.PrintError("Couldn't add label data", err)
		return err
	}
	if err = formatter.PrintList(labeledResources); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LabelRemove subcommand function de-assigns a single label from a single labelable resource
func LabelRemove() error {
	logger.DebugFuncInfo()

	svc, _, formatter := cli.WireUpAPIClient()

	labelsMapNameToID, _, err := LabelLoadsMapping()
	if err != nil {
		return err
	}
	labelID := labelsMapNameToID[viper.GetString(cmd.Label)]

	err = svc.RemoveLabel(cmd.GetContext(), labelID, viper.GetString(cmd.ResourceType), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't remove label", err)
		return err
	}
	return nil
}
