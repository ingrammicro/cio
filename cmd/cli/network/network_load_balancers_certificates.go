// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
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

	cmd.NewCommand(certificatesCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all certificates of a load balancer",
		RunMethod:    CertificateList,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId}},
	)
	cmd.NewCommand(certificatesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the certificate identified by the given id",
		RunMethod:    CertificateShow,
		FlagContexts: []cmd.FlagContext{fId, fLoadBalancerId}},
	)
	cmd.NewCommand(certificatesCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new certificate in a load balancer",
		RunMethod:    CertificateCreate,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId, fName, fPublicKey, fChain, fPrivateKey}},
	)
	cmd.NewCommand(certificatesCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing certificate identified by the given id",
		RunMethod:    CertificateUpdate,
		FlagContexts: []cmd.FlagContext{fId, fLoadBalancerId, fName}},
	)
	cmd.NewCommand(certificatesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a certificate",
		RunMethod:    CertificateDelete,
		FlagContexts: []cmd.FlagContext{fId, fLoadBalancerId}},
	)
}

// CertificateList subcommand function
func CertificateList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	certificates, err := svc.ListCertificates(cmd.GetContext(), viper.GetString(cmd.LoadBalancerId))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer certificates data", err)
		return err
	}
	if err = formatter.PrintList(certificates); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CertificateShow subcommand function
func CertificateShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	certificate, err := svc.GetCertificate(
		cmd.GetContext(),
		viper.GetString(cmd.LoadBalancerId),
		viper.GetString(cmd.Id),
	)
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer certificate data", err)
		return err
	}
	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CertificateCreate subcommand function
func CertificateCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	certificateIn := map[string]interface{}{
		"name":        viper.GetString(cmd.Name),
		"public_key":  viper.GetString(cmd.PublicKey),
		"chain":       viper.GetString(cmd.Chain),
		"private_key": viper.GetString(cmd.PrivateKey),
	}

	certificate, err := svc.CreateCertificate(cmd.GetContext(), viper.GetString(cmd.LoadBalancerId), &certificateIn)
	if err != nil {
		formatter.PrintError("Couldn't create load balancer certificate", err)
		return err
	}

	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CertificateUpdate subcommand function
func CertificateUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	certificateIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	certificate, err := svc.UpdateCertificate(cmd.GetContext(),
		viper.GetString(cmd.LoadBalancerId),
		viper.GetString(cmd.Id),
		&certificateIn,
	)
	if err != nil {
		formatter.PrintError("Couldn't update load balancer certificate", err)
		return err
	}

	if err = formatter.PrintItem(*certificate); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CertificateDelete subcommand function
func CertificateDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteCertificate(cmd.GetContext(), viper.GetString(cmd.LoadBalancerId), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete load balancer certificate", err)
		return err
	}
	return nil
}
