package node

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	// "text/tabwriter"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/webservice"
)

type Node struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Fqdn      string `json:"fqdn"`
	PublicIp  string `json:"public_ip"`
	State     string `json:"state"`
	Os        string `json:"os"`
	Plan      string `json:"plan"`
	FleetName string `json:"fleet_name"`
	Master    bool   `json:"is_master"`
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"cluster", "plan"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["fleet_name"] = c.String("cluster")
	v["plan"] = c.String("plan")

	json, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, code := webservice.Post("/v1/kaas/ships", json)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

}

// func cmdStart(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Put(fmt.Sprintf("/v1/kaas/ships/%s/start", c.String("id")), nil)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)

// }

// func cmdStop(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Put(fmt.Sprintf("/v1/kaas/ships/%s/stop", c.String("id")), nil)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)

// }

// func cmdRestart(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Put(fmt.Sprintf("/v1/kaas/ships/%s/restart", c.String("id")), nil)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)

// }

// func cmdDelete(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/kaas/ships/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)

// }

// func cmdList(c *cli.Context) {
// 	var nodes []Node

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/kaas/ships")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &nodes)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "CLUSTER\tMASTER\tID\tNAME\tFQDN\tIP\tSTATE")

// 	for _, node := range nodes {
// 		if node.Master {
// 			fmt.Fprintf(w, "%s\t*\t%s\t%s\t%s\t%s\t%s\n", node.FleetName, node.Id, node.Name, node.Fqdn, node.PublicIp, node.State)
// 		} else {
// 			fmt.Fprintf(w, "%s\t\t%s\t%s\t%s\t%s\t%s\n", node.FleetName, node.Id, node.Name, node.Fqdn, node.PublicIp, node.State)
// 		}

// 	}

// 	w.Flush()
// }

func cmdDockerHijack(c *cli.Context) error {

	var nodes []Node
	var node Node

	discovered := false

	utils.FlagsRequired(c, []string{"node"})

	var firstArgument string
	if c.Args().Present() {
		firstArgument = c.Args().First()
	} else {
		firstArgument = "help"
	}

	nodeName := c.String("node")

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get("/v1/kaas/ships")
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &nodes)
	utils.CheckError(err)

	// Validating if node exist
	for _, element := range nodes {
		if (element.Name == nodeName) || (element.Id == nodeName) {
			discovered = true
			node = element
		}
	}

	if discovered == true {

		dockerLocation, err := exec.LookPath("docker")
		if err != nil {
			log.Warn(fmt.Sprintf("We could not find docker in your enviroment. Please install it."))
			os.Exit(1)
		}

		log.Debug(fmt.Sprintf("Found docker at %s", dockerLocation))
		config, err := utils.GetConcertoConfig()
		utils.CheckError(err)

		nodeParameters := fmt.Sprintf("--host=tcp://%s:2376", node.Fqdn)
		tls := "--tls=true"
		clientCertificate := fmt.Sprintf("--tlscert=%s", config.Certificate.Cert)
		clientKey := fmt.Sprintf("--tlskey=%s", config.Certificate.Key)
		clientCA := fmt.Sprintf("--tlscacert=%s", config.Certificate.Ca)

		arguments := append([]string{nodeParameters, tls, clientCertificate, clientKey, clientCA, firstArgument}, c.Args().Tail()...)

		log.Debug(fmt.Sprintf("Going to execute %s %s", dockerLocation, arguments))

		cmd := exec.Command(dockerLocation, arguments...)

		stdout, err := cmd.StdoutPipe()
		utils.CheckError(err)

		stderr, err := cmd.StderrPipe()
		utils.CheckError(err)

		// Start command
		err = cmd.Start()
		utils.CheckError(err)
		defer cmd.Wait()

		go io.Copy(os.Stderr, stderr)

		ls := bufio.NewReader(stdout)

		for {
			line, isPrefix, err := ls.ReadLine()
			if isPrefix {
				log.Errorf("%s", errors.New("isPrefix: true"))
			}
			if err != nil {
				if err != io.EOF {
					log.Errorf("%s", err.Error())
				}
				break
			}
			fmt.Printf("%s\n", strings.Replace(string(line), "docker", fmt.Sprintf("concerto nodes docker --node %s", nodeName), -1))
		}

		go func() {
			time.Sleep(30 * time.Second)
			log.Fatal(fmt.Sprintf("Timeout out. Check conectivity to %s", nodeParameters))
		}()

		return nil

	} else {
		log.Warn(fmt.Sprintf("Node \"%s\" is not in your account please create it. Thank you.", nodeName))
		os.Exit(1)
	}
	return nil
}
