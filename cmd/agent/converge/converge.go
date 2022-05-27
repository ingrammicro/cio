// Copyright (c) 2017-2022 Ingram Micro Inc.

package converge

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"path"
	"regexp"
	"runtime"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	cmd.NewCommand(
		cmd.RootCmd,
		&cmd.CommandContext{Use: "converge", Short: "Converges Host to original Blueprint", RunMethod: Converge},
	)
}

func traceOutput(ls *bufio.Reader) {
	garbageOutput, _ := regexp.Compile("[\\[][^\\[|^\\]]*[\\]]\\s[A-Z]*:\\s")
	output, _ := regexp.Compile("Chef Run")
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
		outputLine := garbageOutput.ReplaceAllString(string(line), "")
		if output.MatchString(outputLine) {
			log.Infof("%s", outputLine)
		} else {
			log.Debugf("%s", outputLine)
		}
	}
}

func Converge() error {
	firstBootJsonChef := path.Join("/etc/chef", "first-boot.json")
	if runtime.GOOS == "windows" {
		firstBootJsonChef = path.Join("c:\\chef", "first-boot.json")
	}

	if utils.FileExists(firstBootJsonChef) {
		cmd := exec.Command("chef-client", "-j", firstBootJsonChef)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Errorf("%s", err.Error())
		}
		ls := bufio.NewReader(stdout)
		err = cmd.Start()
		if err != nil {
			log.Errorf("%s", err.Error())
		}

		traceOutput(ls)

		err = cmd.Wait()
		if err != nil {
			log.Errorf("%s", err.Error())
		}
	} else {
		log.Fatalf("Make sure %s chef client configuration exists.", firstBootJsonChef)
	}
	return nil
}
