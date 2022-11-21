// Copyright (c) 2017-2022 Ingram Micro Inc.

package storage

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/cmd/cli/cloud"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fServerId := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerId,
		Usage: "Identifier of a server to return only the volumes that are attached with that server"}

	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Volume Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the volume"}

	fSize := cmd.FlagContext{Type: cmd.String, Name: cmd.Size, Required: true, Usage: "Size for the volume, in GB"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which the volume is"}

	fStoragePlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.StoragePlanId, Required: true,
		Usage: "Identifier of the storage plan on which the volume is based"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with volume"}

	fServerIdAttach := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerId, Required: true,
		Usage: "Identifier of the server to attach the volume"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "volume", Hidden: true,
		Usage: "Resource Type"}

	volumesCmd := cmd.NewCommand(storageCmd, &cmd.CommandContext{
		Use:   "volumes",
		Short: "Provides information on storage volumes"},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all existing volumes",
		RunMethod:    VolumeList,
		FlagContexts: []cmd.FlagContext{fServerId, fLabelsFilter}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the volume identified by the given id",
		RunMethod:    VolumeShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new volume",
		RunMethod:    VolumeCreate,
		FlagContexts: []cmd.FlagContext{fName, fSize, fCloudAccountId, fStoragePlanId, fLabels}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing volume identified by the given id",
		RunMethod:    VolumeUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "attach",
		Short:        "Attaches the volume to server",
		RunMethod:    VolumeAttach,
		FlagContexts: []cmd.FlagContext{fId, fServerIdAttach}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "detach",
		Short:        "Detaches a volume from server",
		RunMethod:    VolumeDetach,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a volume",
		RunMethod:    VolumeDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "discard",
		Short:        "Discards a volume but does not delete it from the cloud provider",
		RunMethod:    VolumeDiscard,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(volumesCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// VolumeList subcommand function
func VolumeList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	volumes, err := svc.ListStorageVolumes(ctx, viper.GetString(cmd.ServerId))
	if err != nil {
		formatter.PrintError("Couldn't receive volumes data", err)
		return err
	}
	if err = cloud.FormatVolumesResponse(ctx, volumes, formatter); err != nil {
		return err
	}
	return nil
}

// VolumeShow subcommand function
func VolumeShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	volume, err := svc.GetStorageVolume(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive volume data", err)
		return err
	}
	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	volume.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VolumeCreate subcommand function
func VolumeCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	volumeIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"size":             viper.GetInt(cmd.Size),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
		"storage_plan_id":  viper.GetString(cmd.StoragePlanId),
	}

	ctx := cmd.GetContext()
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		volumeIn["label_ids"], err = labels.LabelResolution(
			ctx,
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	volume, err := svc.CreateStorageVolume(ctx, &volumeIn)
	if err != nil {
		formatter.PrintError("Couldn't create volume", err)
		return err
	}

	volume.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VolumeUpdate subcommand function
func VolumeUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	volumeIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	ctx := cmd.GetContext()
	volume, err := svc.UpdateStorageVolume(ctx, viper.GetString(cmd.Id), &volumeIn)
	if err != nil {
		formatter.PrintError("Couldn't update volume", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	volume.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VolumeAttach subcommand function
func VolumeAttach() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	volumeIn := map[string]interface{}{
		"attached_server_id": viper.GetString(cmd.ServerId),
	}

	ctx := cmd.GetContext()
	server, err := svc.AttachStorageVolume(ctx, viper.GetString(cmd.Id), &volumeIn)
	if err != nil {
		formatter.PrintError("Couldn't attach volume", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VolumeDetach subcommand function
func VolumeDetach() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DetachStorageVolume(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't detach volume", err)
		return err
	}
	return nil
}

// VolumeDelete subcommand function
func VolumeDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteStorageVolume(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete volume", err)
		return err
	}
	return nil
}

// VolumeDiscard subcommand function
func VolumeDiscard() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DiscardStorageVolume(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't discard volume", err)
		return err
	}
	return nil
}
