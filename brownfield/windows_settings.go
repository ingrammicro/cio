// Copyright (c) 2017-2021 Ingram Micro Inc.

//go:build windows
// +build windows

package brownfield

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
)

type Settings struct {
	SSHPublicKeys []string `json:"ssh_public_keys"`
}

func applyConcertoSettings(cs *utils.HTTPConcertoservice, f format.Formatter, username, password string) {
	settings, err := obtainSettings(cs)
	if err != nil {
		f.PrintFatal("Cannot obtain settings", err)
	}
	err = sendUsernamePassword(cs, username, password)
	if err != nil {
		f.PrintFatal("Cannot send server credentials", err)
	}
	dir, err := ioutil.TempDir("", "brownfield-configure")
	if err != nil {
		f.PrintFatal("Cannot create temp dir", err)
	}
	defer os.RemoveAll(dir) // clean up

	scriptFilePath := fmt.Sprintf("%s\\configure.bat", dir)

	// Writes content to file
	if err := ioutil.WriteFile(scriptFilePath, []byte(createScriptTemplate(settings)), 0600); err != nil {
		log.Fatalf("Error creating temp file: %v", err)
	}

	output, err := exec.Command("cmd", "/C", scriptFilePath).CombinedOutput()
	if err != nil {
		f.PrintFatal("Error happened running setup script", fmt.Errorf("%s: %v", output, err))
	}
	fmt.Printf("Setup script ran successfully\n")
}

func obtainSettings(cs *utils.HTTPConcertoservice) (settings *Settings, err error) {
	body, status, err := cs.Get("/brownfield/settings")
	if err != nil {
		return
	}
	if status == 403 {
		err = fmt.Errorf("server responded with 403 code: authentication was not successful")
		return
	}
	if status >= 300 {
		err = fmt.Errorf("server responded with %d code: %s", status, string(body))
		return
	}
	settings = &Settings{}
	err = json.Unmarshal(body, settings)
	if err != nil {
		err = fmt.Errorf("cannot parse as JSON server response %v: %v", string(body), err)
		return
	}
	return
}

func sendUsernamePassword(cs *utils.HTTPConcertoservice, username, password string) error {
	payload := &map[string]interface{}{
		"settings": map[string]interface{}{
			"username":    username,
			"user_passwd": password,
		},
	}
	body, status, err := cs.Put("/brownfield/settings", payload)
	if err != nil {
		return err
	}
	if status == 403 {
		return fmt.Errorf("server responded with 403 code: authentication was not successful")
	}
	if status >= 300 {
		return fmt.Errorf("server responded with %d code: %s", status, string(body))
	}
	return nil
}

func createScriptTemplate(settings *Settings) string {
	return strings.Join([]string{
		`powershell -command "Set-Service -Name sshd -StartupType Automatic"`,
		`powershell -command "Start-Service sshd"`,
		`powershell -command "Set-Content -path C:\ProgramData\ssh\administrators_authorized_keys '` + settings.SSHPublicKeys[0] + `'"`,
		`powershell -command "icacls C:\ProgramData\ssh\administrators_authorized_keys /inheritance:d"`,
		`powershell -command "icacls C:\ProgramData\ssh\administrators_authorized_keys /remove 'NT AUTHORITY\Authenticated Users'"`,
		`powershell -command "Restart-Service sshd"`,
		`powershell -command "Set-ItemProperty -Path 'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System\' -Name 'shutdownwithoutlogon' -Value 0"`,
		`powershell -command "Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp\' -Name 'UserAuthentication' -Value 0"`,
		`powershell -command "Set-ItemProperty -Path 'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System\' -Name 'dontdisplaylastusername' -Value 1"`,
	}, " && ")
}
