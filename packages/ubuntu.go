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
		return errors.New("fedora not yet supported", nil)
	}
	return errors.New("unknown distro type: "+string(versionData), nil)
}

func getVersion() (string, error) {
	versionData, err := ioutil.ReadFile("/proc/version")
	if err != nil {
		return "", errors.New("failed to get /proc/version", err)
	}
	return string(versionData), nil
}

var deps = []string{
	"libxen-dev",
	"curl",
	"git",
	"build-essential",
}

func installUbuntu(out io.Writer) error {
	if err := command(out, "apt", "update"); err != nil {
		return err
	}
	args := []string{
		"apt",
		"install",
		"-y",
	}
	args = append(args, deps...)
	if err := command(out, args...); err != nil {
		return err
	}
	return nil
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
