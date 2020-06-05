// Copyright (c) 2017-2022 Ingram Micro Inc.

// +build linux darwin

package brownfield

import (
	"fmt"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/types"

	"github.com/ingrammicro/cio/api"

	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"github.com/ingrammicro/cio/utils/format"
)

func applySettings(svc *api.ServerAPI, f format.Formatter, _, _ string) {
	settings, err := obtainSettings(svc)
	if err != nil {
		f.PrintFatal("Cannot obtain settings", err)
	}

	var tmpFileName string
	tmpFile, err := ioutil.TempFile("", "concerto-setup")
	if err != nil {
		f.PrintFatal("Cannot not open temp file to write setup script", err)
	}
	defer func() {
		if tmpFileName == "" {
			tmpFile.Close()
		}
		os.Remove(tmpFileName)
	}()
	err = scriptTemplate.Execute(tmpFile, settings)
	if err != nil {
		f.PrintFatal("Cannot not instantiate setup script", err)
	}
	err = tmpFile.Close()
	tmpFileName = tmpFile.Name()
	if err != nil {
		f.PrintFatal("Cannot instantiate setup script", err)
	}
	_, err = exec.Command("bash", tmpFileName).Output()
	if err != nil {
		f.PrintFatal("Error happened running setup script", err)
	}
	fmt.Printf("Setup script ran successfully\n")
}

func obtainSettings(svc *api.ServerAPI) (settings *types.Settings, err error) {
	settings, status, err := svc.GetBrownfieldSettings(cmd.GetContext())
	if err != nil {
		return
	}
	if status == 403 {
		err = fmt.Errorf("server responded with 403 code: authentication was not successful")
		return
	}
	if status >= 300 {
		err = fmt.Errorf("server responded with %d code: %s", status, settings)
		return
	}
	return
}

var scriptTemplate = template.Must(template.New("configFile").Parse(`#! /bin/bash

## SSH settings ##
mkdir -p $HOME/.ssh
{{range .SSHPublicKeys}}
echo {{.}} >> $HOME/.ssh/authorized_keys
{{end}}


sed -i -e "s/^#PubkeyAuthentication[ \t]*yes/PubkeyAuthentication yes/g" -e "s/^PubkeyAuthentication[ \t]*no/PubkeyAuthentication yes/g" /etc/ssh/sshd_config
sed -i 's/root:x:0:0:root:\\/root:\\/sbin\\/nologin/root:x:0:0:root:\\/root:\\/bin\\/bash/' /etc/passwd
sed -i -e 's/^AllowUsers /#AllowUsers /' -e 's/^PermitRootLogin /#PermitRootLogin /' /etc/ssh/sshd_config
/etc/init.d/ssh* restart
`))
