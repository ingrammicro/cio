package certificates

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns certificates commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all certificates of a load balancer",
			Action: cmd.CertificateList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the certificate identified by the given id",
			Action: cmd.CertificateShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Certificate Id",
				},
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new certificate in a load balancer",
			Action: cmd.CertificateCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer of the certificate",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the certificate",
				},
				cli.StringFlag{
					Name:  "public-key",
					Usage: "The public key of the certificate",
				},
				cli.StringFlag{
					Name:  "chain",
					Usage: "Chain of the certificate",
				},
				cli.StringFlag{
					Name:  "private-key",
					Usage: "The private key of the certificate",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing certificate identified by the given id",
			Action: cmd.CertificateUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Certificate Id",
				},
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer of the certificate",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the certificate",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a certificate",
			Action: cmd.CertificateDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Certificate Id",
				},
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer of the certificate",
				},
			},
		},
	}
}
