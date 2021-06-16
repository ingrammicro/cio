// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"

	"github.com/ingrammicro/cio/api/network"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpDomain prepares common resources to send request to Concerto API
func WireUpDomain(c *cli.Context) (ds *network.DomainService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewDomainService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Domain service", err)
	}

	return ds, f
}

// DomainList subcommand function
func DomainList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	domains, err := svc.ListDomains()
	if err != nil {
		formatter.PrintFatal("Couldn't receive dns domains data", err)
	}

	labelables := make([]types.Labelable, len(domains))
	for i := 0; i < len(domains); i++ {
		labelables[i] = types.Labelable(domains[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	domains = make([]*types.Domain, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		s, ok := labelable.(*types.Domain)
		if !ok {
			formatter.PrintFatal(
				LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.Domain, got a %T", labelable),
			)
		}
		domains[i] = s
	}

	if err = formatter.PrintList(domains); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainShow subcommand function
func DomainShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	domain, err := svc.GetDomain(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive dns domain data", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainCreate subcommand function
func DomainCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"name", "cloud-account-id"}, formatter)
	domainIn := map[string]interface{}{
		"name":             c.String("name"),
		"cloud_account_id": c.String("cloud-account-id"),
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		domainIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	domain, err := svc.CreateDomain(&domainIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create dns domain", err)
	}

	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainDelete subcommand function
func DomainDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	domain, err := svc.DeleteDomain(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete dns domain", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainRetry subcommand function
func DomainRetry(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	domain, err := svc.RetryDomain(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't retry dns domain", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	domain.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainListRecords subcommand function
func DomainListRecords(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"domain-id"}, formatter)
	records, err := svc.ListRecords(c.String("domain-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive dns records data", err)
	}
	if err = formatter.PrintList(records); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainShowRecord subcommand function
func DomainShowRecord(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	record, err := svc.GetRecord(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive dns record data", err)
	}
	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainCreateRecord subcommand function
func DomainCreateRecord(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"domain-id", "name", "type"}, formatter)
	recordType := c.String("type")
	recordIn := map[string]interface{}{
		"name": c.String("name"),
		"type": recordType,
	}

	setParamString(c, "content", "content", recordIn)
	setParamInt(c, "ttl", "ttl", recordIn)

	// If provided, only include in adequate context
	switch recordType {
	case "a":
		if c.IsSet("content") &&
			(c.IsSet("server-id") || c.IsSet("floating-ip-id")) {
			return fmt.Errorf(
				"invalid parameters detected. Please provide only one: 'content', 'server-id' or 'floating-ip-id'",
			)
		}
		// one and only one of the fields must be provided.
		if c.IsSet("server-id") && c.IsSet("floating-ip-id") {
			return fmt.Errorf("invalid parameters detected. Please provide only one: 'server-id' or 'floating-ip-id'")
		}
		setParamString(c, "instance_id", "server-id", recordIn)
		setParamString(c, "floating_ip_id", "floating-ip-id", recordIn)
	case "cname":
		if c.IsSet("content") && c.IsSet("load-balancer-id") {
			return fmt.Errorf("invalid parameters detected. Please provide only one: 'content' or 'load-balancer-id'")
		}
		setParamString(c, "load_balancer_id", "load-balancer-id", recordIn)
	case "mx":
		setParamInt(c, "priority", "priority", recordIn)
	case "srv":
		setParamInt(c, "priority", "priority", recordIn)
		setParamInt(c, "weight", "weight", recordIn)
		setParamInt(c, "port", "port", recordIn)
		//case "aaaa":
		//case "txt":
	default:
	}

	record, err := svc.CreateRecord(c.String("domain-id"), &recordIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create dns record", err)
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainUpdateRecord subcommand function
func DomainUpdateRecord(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	recordIn := map[string]interface{}{
		"name": c.String("name"),
	}

	if c.IsSet("content") {
		recordIn["content"] = c.String("content")
	}
	if c.IsSet("ttl") {
		recordIn["ttl"] = c.Int("ttl")
	}
	if c.IsSet("priority") {
		recordIn["priority"] = c.Int("priority")
	}
	if c.IsSet("weight") {
		recordIn["weight"] = c.Int("weight")
	}
	if c.IsSet("port") {
		recordIn["port"] = c.Int("port")
	}

	record, err := svc.UpdateRecord(c.String("id"), &recordIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update dns record", err)
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainDeleteRecord subcommand function
func DomainDeleteRecord(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	record, err := svc.DeleteRecord(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete dns record", err)
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// DomainRetryRecord subcommand function
func DomainRetryRecord(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	record, err := svc.RetryRecord(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't retry dns record", err)
	}

	if err = formatter.PrintItem(*record); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
