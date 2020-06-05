// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fLoadBalancerId := cmd.FlagContext{Type: cmd.String, Name: cmd.LoadBalancerId, Required: true,
		Usage: "Identifier of the load balancer"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Certificate Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the certificate"}

	fPublicKey := cmd.FlagContext{Type: cmd.String, Name: cmd.PublicKey, Required: true,
		Usage: "The public key of the certificate"}

	fChain := cmd.FlagContext{Type: cmd.String, Name: cmd.Chain, Required: true, Usage: "Chain of the certificate"}

	fPrivateKey := cmd.FlagContext{Type: cmd.String, Name: cmd.PrivateKey, Required: true,
		Usage: "The private key of the certificate"}

	cmd.NewCommand(CertificatesCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all certificates of a load balancer",
		RunMethod:    CertificateList,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId}},
	)
	cmd.NewCommand(CertificatesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the certificate identified by the given id",
		RunMethod:    CertificateShow,
		FlagContexts: []cmd.FlagContext{fId, fLoadBalancerId}},
	)
	cmd.NewCommand(CertificatesCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new certificate in a load balancer",
		RunMethod:    CertificateCreate,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId, fName, fPublicKey, fChain, fPrivateKey}},
	)
	cmd.NewCommand(CertificatesCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing certificate identified by the given id",
		RunMethod:    CertificateUpdate,
		FlagContexts: []cmd.FlagContext{fId, fLoadBalancerId, fName}},
	)
	cmd.NewCommand(CertificatesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a certificate",
		RunMethod:    CertificateDelete,
		FlagContexts: []cmd.FlagContext{fId, fLoadBalancerId}},
	)
}

// CertificateList subcommand function
func CertificateList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	certificates, err := svc.ListCertificates(viper.GetString(cmd.LoadBalancerId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer certificates data", err)
	}
	if err = formatter.PrintList(certificates); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CertificateShow subcommand function
func CertificateShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	certificate, err := svc.GetCertificate(viper.GetString(cmd.LoadBalancerId), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer certificate data", err)
	}
	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CertificateCreate subcommand function
func CertificateCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	certificateIn := map[string]interface{}{
		"name":        viper.GetString(cmd.Name),
		"public_key":  viper.GetString(cmd.PublicKey),
		"chain":       viper.GetString(cmd.Chain),
		"private_key": viper.GetString(cmd.PrivateKey),
	}

	certificate, err := svc.CreateCertificate(viper.GetString(cmd.LoadBalancerId), &certificateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create load balancer certificate", err)
	}

	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CertificateUpdate subcommand function
func CertificateUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	certificateIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	certificate, err := svc.UpdateCertificate(
		viper.GetString(cmd.LoadBalancerId),
		viper.GetString(cmd.Id),
		&certificateIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't update load balancer certificate", err)
	}

	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CertificateDelete subcommand function
func CertificateDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	err := svc.DeleteCertificate(viper.GetString(cmd.LoadBalancerId), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't delete load balancer certificate", err)
	}
	return nil
}
