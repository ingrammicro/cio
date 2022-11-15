// Copyright (c) 2017-2022 Ingram Micro Inc.

// +build windows

package brownfield

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/types"

	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/utils/format"
)

func applySettings(svc *api.ServerAPI, f format.Formatter, username, password string) {
	settings, err := obtainSettings(svc)
	if err != nil {
		f.PrintFatal("Cannot obtain settings", err)
	}
	err = sendUsernamePassword(svc, username, password)
	if err != nil {
		f.PrintFatal("Cannot send server credentials", err)
	}
	dir, err := os.MkdirTemp("", "brownfield-configure")
	if err != nil {
		f.PrintFatal("Cannot create temp dir", err)
	}
	defer os.RemoveAll(dir) // clean up

	scriptFilePath := fmt.Sprintf("%s\\configure.bat", dir)

	// Writes content to file
	if err := os.WriteFile(scriptFilePath, []byte(scriptTemplate(settings)), 0600); err != nil {
		log.Fatalf("Error creating temp file: %v", err)
	}

	output, err := exec.Command("cmd", "/C", scriptFilePath).CombinedOutput()
	if err != nil {
		f.PrintFatal("Error happened running setup script", fmt.Errorf("%s: %v", output, err))
	}
	fmt.Printf("Setup script ran successfully\n")
}

func obtainSettings(svc *api.ServerAPI) (settings *types.Settings, err error) {
	// We do not need settings data, but make the API call to log progress on API service log
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

func sendUsernamePassword(svc *api.ServerAPI, username, password string) error {
	payload := &map[string]interface{}{
		"settings": map[string]interface{}{
			"username":    username,
			"user_passwd": password,
		},
	}
	settings, status, err := svc.SetBrownfieldSettings(cmd.GetContext(), payload)
	if err != nil {
		return err
	}
	if status == 403 {
		return fmt.Errorf("server responded with 403 code: authentication was not successful")
	}
	if status >= 300 {
		return fmt.Errorf("server responded with %d code: %s", status, settings)
	}
	return nil
}

func scriptTemplate(settings *types.Settings) string {
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
