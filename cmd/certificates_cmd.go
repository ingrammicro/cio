// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/network"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpCertificate prepares common resources to send request to Concerto API
func WireUpCertificate(c *cli.Context) (ds *network.CertificateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewCertificateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Certificate service", err)
	}

	return ds, f
}

// CertificateList subcommand function
func CertificateList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCertificate(c)

	checkRequiredFlags(c, []string{"load-balancer-id"}, formatter)
	certificates, err := svc.ListCertificates(c.String("load-balancer-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer certificates data", err)
	}
	if err = formatter.PrintList(certificates); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CertificateShow subcommand function
func CertificateShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCertificate(c)

	checkRequiredFlags(c, []string{"id", "load-balancer-id"}, formatter)
	certificate, err := svc.GetCertificate(c.String("load-balancer-id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer certificate data", err)
	}
	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CertificateCreate subcommand function
func CertificateCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCertificate(c)

	checkRequiredFlags(c, []string{"load-balancer-id", "name", "public-key", "chain", "private-key"}, formatter)
	certificateIn := map[string]interface{}{
		"name":        c.String("name"),
		"public_key":  c.String("public-key"),
		"chain":       c.String("chain"),
		"private_key": c.String("private-key"),
	}

	certificate, err := svc.CreateCertificate(c.String("load-balancer-id"), &certificateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create load balancer certificate", err)
	}

	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CertificateUpdate subcommand function
func CertificateUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCertificate(c)

	checkRequiredFlags(c, []string{"id", "load-balancer-id", "name"}, formatter)
	certificateIn := map[string]interface{}{
		"name": c.String("name"),
	}

	certificate, err := svc.UpdateCertificate(c.String("load-balancer-id"), c.String("id"), &certificateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update load balancer certificate", err)
	}

	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CertificateDelete subcommand function
func CertificateDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCertificate(c)

	checkRequiredFlags(c, []string{"id", "load-balancer-id"}, formatter)
	err := svc.DeleteCertificate(c.String("load-balancer-id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete load balancer certificate", err)
	}
	return nil
}
