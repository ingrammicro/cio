/*
	Reports allow the user to have information about the historical uptime of their servers.
	** Admins will have visibility for all the servers of the associated tenant.

	The available commands are:
		list	reports related to all the account groups of the tenant (admins only)
		show	details about a particular report associated to any account group of the tenant (admins only)

	Use "admin reports --help" on the commandline interface for more information about the available subcommands.

	Reports list

	The command `reports list` returns information about the reports related to all the account groups of the tenant.
	The authenticated user must be an admin.

	Usage:

		reports list

	Reports show

	The command `reports show` returns details about a particular report associated to any account group of the tenant.
	The authenticated user must be an admin.
	The report is identified by a unique report_id.

	Usage:

		reports show --id <report_id>

*/
package admin

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/webservice"
	// "time"
)

// type Report struct {
// 	Id             string       `json:"id"`
// 	Year           int          `json:"year"`
// 	Month          time.Month   `json:"month"`
// 	Start_time     time.Time    `json:"start_time"`
// 	End_time       time.Time    `json:"end_time"`
// 	Server_seconds float32      `json:"server_seconds"`
// 	Closed         bool         `json:"closed"`
// 	Li             []Lines      `json:"lines"`
// 	Account_group  AccountGroup `json:"account_group"`
// }

// type Lines struct {
// 	Id                string    `json:"_id"`
// 	Commissioned_at   time.Time `json:"commissioned_at"`
// 	Decommissioned_at time.Time `json:"decommissioned_at"`
// 	Instance_id       string    `json:"instance_id"`
// 	Instance_name     string    `json:"instance_name"`
// 	Instance_fqdn     string    `json:"instance_fqdn"`
// 	Consumption       float32   `json:"consumption"`
// }

// type AccountGroup struct {
// 	Id   string `json:"_id"`
// 	Name string `json:"name"`
// }

// func cmdList(c *cli.Context) {
// 	var reports []Report

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/admin/reports")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &reports)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")

// 	for _, report := range reports {
// 		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", report.Id, report.Year, report.Month, report.StartTime, report.EndTime, report.ServerSeconds, report.Closed)
// 	}

// 	w.Flush()
// }

func cmdShow(c *cli.Context) error {
	var vals types.Report

	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get(fmt.Sprintf("/v1/admin/reports/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &vals)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)

	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\tACCOUNT GROUP ID\tACCOUNT GROUP NAME\r")
	fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\t%s\t%s\n", vals.ID, vals.Year, vals.Month, vals.StartTime, vals.EndTime, vals.ServerSeconds, vals.Closed, vals.AccountGroup.ID, vals.AccountGroup.Name)

	fmt.Fprintln(w, "LINES:\r")
	fmt.Fprintln(w, "ID\tCOMMISSIONED AT\tDECOMMISSIONED AT\tINSTANCE ID\tINSTANCE NAME\tINSTANCE FQDN\tCONSUMPTION\r")

	for _, l := range vals.Lines {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%g\n", l.ID, l.CommissionedAt, l.DecommissionedAt, l.InstanceID, l.InstanceName, l.InstanceFQDN, l.Consumption)
	}
	w.Flush()
	return nil
}

// func SubCommands() []cli.Command {
// 	return []cli.Command{
// 		{
// 			Name:   "list",
// 			Usage:  "Returns information about the reports related to all the account groups of the tenant. The authenticated user must be an admin.",
// 			Action: cmdList,
// 		},
// 		{
// 			Name:   "show",
// 			Usage:  "Returns details about a particular report associated to any account group of the tenant. The authenticated user must be an admin.",
// 			Action: cmdShow,
// 			Flags: []cli.Flag{
// 				cli.StringFlag{
// 					Name:  "id",
// 					Usage: "Report Identifier",
// 				},
// 			},
// 		},
// 	}
// }
