package utils

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/codegangsta/cli"

	log "github.com/Sirupsen/logrus"
)

// TODO remove after migration
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

func Untar(ctx context.Context, source, target string) error {

	if err := os.MkdirAll(target, 0600); err != nil {
		return err
	}

	tarExecutable := "tar"
	if runtime.GOOS == "windows" {
		tarExecutable = "C:\\opscode\\chef\\bin\\tar.exe"
	}
	cmd := exec.CommandContext(ctx, tarExecutable, "-xzf", source, "-C", target)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ScrapeErrorMessage(message string, regExpression string) string {

	re, err := regexp.Compile(regExpression)
	scrapped := re.FindStringSubmatch(message)

	if scrapped == nil || err != nil || len(scrapped) < 2 {
		// couldn't scrape, return generic error
		message = "Error executing operation"
	} else {
		// return scrapped response
		message = scrapped[1]
	}

	return message
}

func CheckReturnCode(res int, mesg []byte) {
	if res >= 300 {

		message := string(mesg[:])
		log.Debugf("IMCO API response: %s", message)

		f := func(c rune) bool {
			return c == ',' || c == ':' || c == '{' || c == '}' || c == '"' || c == ']' || c == '['
		}

		// check if response is a web page.
		if strings.Contains(message, "<html>") {
			scrapResponse := "<title>(.*?)</title>"
			message = ScrapeErrorMessage(message, scrapResponse)
		} else if strings.Contains(message, "{\"errors\":{") {
			scrapResponse := "{\"errors\":(.*?)}"

			message = ScrapeErrorMessage(message, scrapResponse)
			result := strings.Split(message, ",")
			if result != nil && len(result) >= 1 {
				message = result[0]
			}
			// Separate into fields with func.
			fields := strings.FieldsFunc(message, f)
			message = strings.Join(fields[:], " ")

		} else if strings.Contains(message, "{\"error\":") {
			scrapResponse := "{\"error\":\"(.*?)\"}"
			message = ScrapeErrorMessage(message, scrapResponse)
		}
		// if it's not a web page or json-formatted message, return the raw message
		log.Fatal(fmt.Sprintf("HTTP request failed: (%d) [%s]", res, message))
	}
}

// CheckStandardStatus return error if status is not OK
func CheckStandardStatus(status int, mesg []byte) error {

	if status < 300 {
		return nil
	}

	message := string(mesg[:])

	f := func(c rune) bool {
		return c == ',' || c == ':' || c == '{' || c == '}' || c == '"' || c == ']' || c == '['
	}

	if strings.Contains(message, "{\"errors\":{") {
		scrapResponse := "{\"errors\":(.*?)}"

		message = ScrapeErrorMessage(message, scrapResponse)
		result := strings.Split(message, ",")
		if result != nil && len(result) >= 1 {
			message = result[0]
		}
		// Separate into fields with func.
		fields := strings.FieldsFunc(message, f)
		message = strings.Join(fields[:], " ")

	} else if strings.Contains(message, "{\"error\":") {
		scrapResponse := "{\"error\":\"(.*?)\"}"
		message = ScrapeErrorMessage(message, scrapResponse)
	}

	// if it's not a web page or json-formatted message, return the raw message
	return fmt.Errorf("HTTP request failed: (%d) [%s]", status, message)
}

// FileExists checks file existence
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// CheckRequiredFlags checks for required flags, and show usage if requirements not met
func CheckRequiredFlags(c *cli.Context, flags []string) {
	missing := ""
	for _, flag := range flags {
		if !c.IsSet(flag) {
			missing = fmt.Sprintf("%s\t--%s\n", missing, flag)
		}
	}

	if missing != "" {
		fmt.Printf("Incorrect usage. Please use parameters:\n%s\n", missing)
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(2)
	}
}

// RandomString generates a random string from lowercase letters and numbers
func RandomString(strlen int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

// RemoveDuplicates returns the slice removing duplicates if exist
func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// Contains evaluates whether s contains x.
func Contains(s []string, x string) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}

// Subset returns true if the first slice is completely contained in the second slice.
// There must be at least the same number of duplicate values in second as there are in first.
func Subset(s1, s2 []string) bool {
	if len(s1) > len(s2) {
		return false
	}
	for _, e := range s1 {
		if !Contains(s2, e) {
			return false
		}
	}
	return true
}

func RemoveFileInfo(fileInfo os.FileInfo, fileInfoName string) error {
	if fileInfo.IsDir() {
		d, err := os.Open(fileInfoName)
		if err != nil {
			return err
		}
		defer d.Close()
		names, err := d.Readdirnames(-1)
		if err != nil {
			return err
		}
		for _, name := range names {
			err = os.RemoveAll(filepath.Join(fileInfoName, name))
			if err != nil {
				return err
			}
		}
	}

	if err := os.Remove(fileInfoName); err != nil {
		return err
	}
	return nil
}
