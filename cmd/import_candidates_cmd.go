package cmd

import (
	"github.com/ingrammicro/cio/api/clientbrownfield"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
	"strings"
)

// WireUpImportCandidate prepares common resources to send request to Concerto API
func WireUpImportCandidate(c *cli.Context) (ds *clientbrownfield.ImportCandidateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = clientbrownfield.NewImportCandidateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Import Candidate service", err)
	}

	return ds, f
}

// ImportCandidateServerList subcommand function
func ImportCandidateServerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serversImportCandidates, err := cloudAccountSvc.ListServers(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list servers import candidates", err)
	}

	if err = formatter.PrintList(serversImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVPCList subcommand function
func ImportCandidateVPCList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	vpcsImportCandidates, err := cloudAccountSvc.ListVPCs(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list vpcs import candidates", err)
	}

	if err = formatter.PrintList(vpcsImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateFloatingIPList subcommand function
func ImportCandidateFloatingIPList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	floatingIPsImportCandidates, err := cloudAccountSvc.ListFloatingIPs(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list floating ips import candidates", err)
	}

	if err = formatter.PrintList(floatingIPsImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVolumeList subcommand function
func ImportCandidateVolumeList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	volumesImportCandidates, err := cloudAccountSvc.ListVolumes(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list volume import candidates", err)
	}

	if err = formatter.PrintList(volumesImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateServerShow subcommand function
func ImportCandidateServerShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := importCandidateSvc.GetServer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive server data", err)
	}

	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVPCShow subcommand function
func ImportCandidateVPCShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	vpc, err := importCandidateSvc.GetVPC(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive vpc data", err)
	}

	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateFloatingIPShow subcommand function
func ImportCandidateFloatingIPShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	floatingIP, err := importCandidateSvc.GetFloatingIP(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive floating IP data", err)
	}

	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVolumeShow subcommand function
func ImportCandidateVolumeShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	volume, err := importCandidateSvc.GetVolume(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive volume data", err)
	}

	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateServerImport subcommand function
func ImportCandidateServerImport(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverIn := map[string]interface{}{}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		serverIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}
	if c.IsSet("ssh-profile-id") {
		serverIn["ssh_profile_id"] = c.String("ssh-profile-id")
	}
	if c.IsSet("ssh-profile-ids") {
		serverIn["ssh_profile_ids"] = strings.Split(c.String("ssh-profile-ids"), ",")
	}

	server, err := importCandidateSvc.ImportServer(c.String("id"), &serverIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import server", err)
	}

	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVPCImport subcommand function
func ImportCandidateVPCImport(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	vpcIn := map[string]interface{}{}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		vpcIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	vpc, err := importCandidateSvc.ImportVPC(c.String("id"), &vpcIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import vpc", err)
	}

	vpc.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateFloatingIPImport subcommand function
func ImportCandidateFloatingIPImport(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	floatingIPIn := map[string]interface{}{}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		floatingIPIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	floatingIP, err := importCandidateSvc.ImportFloatingIP(c.String("id"), &floatingIPIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import floating IP", err)
	}

	floatingIP.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVolumeImport subcommand function
func ImportCandidateVolumeImport(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	volumeIn := map[string]interface{}{}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		volumeIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	volume, err := importCandidateSvc.ImportVolume(c.String("id"), &volumeIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import volume", err)
	}

	volume.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
