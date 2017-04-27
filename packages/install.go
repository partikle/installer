package packages

import (
	"fmt"
	"github.com/ilackarms/pkg/errors"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	ubuntu = "Ubuntu"
	fedora = "Fedora"
)

func Install(out io.Writer) error {
	versionData, err := getVersion()
	if err != nil {
		return errors.New("failed to determine versionData", err)
	}
	if strings.Contains(versionData, ubuntu) {
		return installUbuntu(out)
	}
	if strings.Contains(versionData, fedora) {
		return installFedora(out)
	}
	return errors.New("unknown distro type: "+string(versionData), nil)
}

func getVersion() (string, error) {
	versionData, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return "", errors.New("failed to get /proc/version", err)
	}
	return string(versionData), nil
}

func command(out io.Writer, args ...string) error {
	fmt.Fprintf(out, "running command: %v\n", args)
	cmd := exec.Command(args[0])
	cmd.Args = args
	cmd.Stdout = out
	cmd.Stderr = out
	if err := cmd.Run(); err != nil {
		return errors.New("failed running command "+strings.Join(args, " "), err)
	}
	return nil
}
