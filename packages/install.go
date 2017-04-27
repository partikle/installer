package packages

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/ilackarms/pkg/errors"
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
