// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/spf13/viper"
)

func init() {
	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Domain Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the DNS domain"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which the domain shall be registered"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with domain"}

	fDomainId := cmd.FlagContext{Type: cmd.String, Name: cmd.DomainId, Required: true,
		Usage: "Identifier of the DNS domain"}

	fIdRecord := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Identifier of the DNS record"}

	fType := cmd.FlagContext{Type: cmd.String, Name: cmd.Type, Required: true,
		Usage: "Type of the  DNS record, among 'a', 'aaaa', 'cname', 'mx', 'srv', 'txt'"}

	fNameRecord := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the DNS record"}
	fNameRecordRec := fNameRecord
	fNameRecordRec.Required = true

	fContent := cmd.FlagContext{Type: cmd.String, Name: cmd.Content, Usage: "Content of the DNS record"}

	fTtl := cmd.FlagContext{Type: cmd.Int, Name: cmd.Ttl,
		Usage: "TTL of the DNS record. Defaults to 3600 if not provided"}

	fServerId := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerId,
		Usage: "Identifier of the Server that is wanted to be attached to the record. " +
			"Only valid for records of type 'a'"}

	fFloatingIpId := cmd.FlagContext{Type: cmd.String, Name: cmd.FloatingIpId,
		Usage: "Identifier of the floating IP that is wanted to be attached to the record. " +
			"Only valid for records of type 'a'"}

	fLoadBalancerId := cmd.FlagContext{Type: cmd.String, Name: cmd.LoadBalancerId,
		Usage: "Identifier of the load balancer that is wanted to be attached to the record. " +
			"Only valid for records of type 'cname'"}

	fPriority := cmd.FlagContext{Type: cmd.String, Name: cmd.Priority,
		Usage: "Priority of the record. Only valid for 'mx' and 'srv' types. Defaults to 0. " +
			"Only valid for records of types 'mx' and 'srv'"}

	fWeight := cmd.FlagContext{Type: cmd.String, Name: cmd.Weight,
		Usage: "Weight of the record. Only valid for 'srv' type. Defaults to 0. Only valid for records of type 'srv'"}

	fPort := cmd.FlagContext{Type: cmd.Int, Name: cmd.Port,
		Usage: "Port of the record. Only valid for 'srv' type. Defaults to 0. Only valid for records of type 'srv'"}

	domainsCmd := cmd.NewCommand(networkCmd, &cmd.CommandContext{
		Use:   "dns-domains",
		Short: "Provides information about DNS domains and records"},
	)
	cmd.NewCommand(domainsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all DNS domains",
		RunMethod:    DomainList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(domainsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the DNS domain identified by the given id",
		RunMethod:    DomainShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(domainsCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new DNS domain",
		RunMethod:    DomainCreate,
		FlagContexts: []cmd.FlagContext{fName, fCloudAccountId, fLabels}},
	)
	cmd.NewCommand(domainsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a DNS domain",
		RunMethod:    DomainDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(domainsCmd, &cmd.CommandContext{
		Use:          "retry",
		Short:        "Retries the application of DNS domain",
		RunMethod:    DomainRetry,
		FlagContexts: []cmd.FlagContext{fId}},
	)

	recordsCmd := cmd.NewCommand(domainsCmd, &cmd.CommandContext{
		Use:   "records",
		Short: "Provides information about DNS records"},
	)
	cmd.NewCommand(recordsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all DNS records of a domain",
		RunMethod:    DomainListRecords,
		FlagContexts: []cmd.FlagContext{fDomainId}},
	)
	cmd.NewCommand(recordsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the DNS record identified by the given id",
		RunMethod:    DomainShowRecord,
		FlagContexts: []cmd.FlagContext{fIdRecord}},
	)
	cmd.NewCommand(recordsCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates a DNS record in a domain",
		RunMethod: DomainCreateRecord,
		FlagContexts: []cmd.FlagContext{fDomainId, fNameRecordRec, fType, fContent, fTtl, fServerId, fFloatingIpId,
			fLoadBalancerId, fPriority, fWeight, fPort}},
	)
	cmd.NewCommand(recordsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates a DNS record in a domain",
		RunMethod:    DomainUpdateRecord,
		FlagContexts: []cmd.FlagContext{fIdRecord, fNameRecord, fContent, fTtl, fPriority, fWeight, fPort}},
	)
	cmd.NewCommand(recordsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a DNS record in a domain",
		RunMethod:    DomainDeleteRecord,
		FlagContexts: []cmd.FlagContext{fIdRecord}},
	)
	cmd.NewCommand(recordsCmd, &cmd.CommandContext{
		Use:          "retry",
		Short:        "Retries the application of DNS record in a domain",
		RunMethod:    DomainRetryRecord,
		FlagContexts: []cmd.FlagContext{fIdRecord}},
	)
}

// DomainList subcommand function
func DomainList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	domains, err := svc.ListDomains(ctx)
	if err != nil {
		formatter.PrintError("Couldn't receive dns domains data", err)
		return err
	}

	labelables := make([]types.Labelable, len(domains))
	for i := 0; i < len(domains); i++ {
		labelables[i] = types.Labelable(domains[i])
	}
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	filteredLabelables, err := labels.LabelFiltering(labelables, labelIDsByName)
	if err != nil {
		return err
	}
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	domains = make([]*types.Domain, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		s, ok := labelable.(*types.Domain)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.Domain, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		domains[i] = s
	}

	if err = formatter.PrintList(domains); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainShow subcommand function
func DomainShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	domain, err := svc.GetDomain(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive dns domain data", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainCreate subcommand function
func DomainCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	domainIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
	}

	ctx := cmd.GetContext()
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	if viper.IsSet(cmd.Labels) {
		domainIn["label_ids"], err = labels.LabelResolution(
			ctx,
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	domain, err := svc.CreateDomain(ctx, &domainIn)
	if err != nil {
		formatter.PrintError("Couldn't create dns domain", err)
		return err
	}

	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainDelete subcommand function
func DomainDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	domain, err := svc.DeleteDomain(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete dns domain", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainRetry subcommand function
func DomainRetry() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	domain, err := svc.RetryDomain(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't retry dns domain", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainListRecords subcommand function
func DomainListRecords() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	records, err := svc.ListRecords(cmd.GetContext(), viper.GetString(cmd.DomainId))
	if err != nil {
		formatter.PrintError("Couldn't receive dns records data", err)
		return err
	}
	if err = formatter.PrintList(records); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainShowRecord subcommand function
func DomainShowRecord() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	record, err := svc.GetRecord(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive dns record data", err)
		return err
	}
	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

func setRecordA(recordIn map[string]interface{}) error {
	if viper.IsSet(cmd.Content) &&
		(viper.IsSet(cmd.ServerId) || viper.IsSet(cmd.FloatingIpId)) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'content', 'server-id' or 'floating-ip-id'",
		)
	}
	// one and only one of the fields must be provided.
	if viper.IsSet(cmd.ServerId) && viper.IsSet(cmd.FloatingIpId) {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'server-id' or 'floating-ip-id'")
	}
	cmd.SetParamString("instance_id", cmd.ServerId, recordIn)
	cmd.SetParamString("floating_ip_id", cmd.FloatingIpId, recordIn)
	return nil
}
func setRecordCName(recordIn map[string]interface{}) error {
	if viper.IsSet(cmd.Content) && viper.IsSet(cmd.LoadBalancerId) {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'content' or 'load-balancer-id'")
	}
	cmd.SetParamString("load_balancer_id", cmd.LoadBalancerId, recordIn)
	return nil
}

func setRecordMx(recordIn map[string]interface{}) {
	cmd.SetParamInt("priority", cmd.Priority, recordIn)
}

func setRecordSrv(recordIn map[string]interface{}) {
	cmd.SetParamInt("priority", cmd.Priority, recordIn)
	cmd.SetParamInt("weight", cmd.Weight, recordIn)
	cmd.SetParamInt("port", cmd.Port, recordIn)
}

func setRecordInByType(recordType string, recordIn map[string]interface{}) error {
	// If provided, only include in adequate context
	switch recordType {
	case "a":
		if err := setRecordA(recordIn); err != nil {
			return err
		}
	case "cname":
		if err := setRecordCName(recordIn); err != nil {
			return err
		}
	case "mx":
		setRecordMx(recordIn)
	case "srv":
		setRecordSrv(recordIn)
		//case "aaaa":
		//case "txt":
	default:
	}
	return nil
}

// DomainCreateRecord subcommand function
func DomainCreateRecord() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	recordType := viper.GetString(cmd.Type)
	recordIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
		"type": recordType,
	}
	cmd.SetParamString("content", cmd.Content, recordIn)
	cmd.SetParamString("ttl", cmd.Ttl, recordIn)

	// If provided, only include in adequate context
	if err := setRecordInByType(recordType, recordIn); err != nil {
		formatter.PrintError("Couldn't set dns record type", err)
		return err
	}

	record, err := svc.CreateRecord(cmd.GetContext(), viper.GetString(cmd.DomainId), &recordIn)
	if err != nil {
		formatter.PrintError("Couldn't create dns record", err)
		return err
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainUpdateRecord subcommand function
func DomainUpdateRecord() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	recordIn := map[string]interface{}{}
	cmd.SetParamString("name", cmd.Name, recordIn)
	cmd.SetParamString("content", cmd.Content, recordIn)
	cmd.SetParamInt("ttl", cmd.Ttl, recordIn)
	// Params only supported by adequate record type!?
	cmd.SetParamInt("priority", cmd.Priority, recordIn)
	cmd.SetParamInt("weight", cmd.Weight, recordIn)
	cmd.SetParamInt("port", cmd.Port, recordIn)

	record, err := svc.UpdateRecord(cmd.GetContext(), viper.GetString(cmd.Id), &recordIn)
	if err != nil {
		formatter.PrintError("Couldn't update dns record", err)
		return err
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainDeleteRecord subcommand function
func DomainDeleteRecord() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	record, err := svc.DeleteRecord(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete dns record", err)
		return err
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// DomainRetryRecord subcommand function
func DomainRetryRecord() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	record, err := svc.RetryRecord(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't retry dns record", err)
		return err
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
