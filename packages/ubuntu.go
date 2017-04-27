package packages

import (
	"io"

	"github.com/partikle/installer/pkg/exec"
)

var ubuntuPackages = []string{
	"libxen-dev",
	"curl",
	"git",
	"build-essential",
}

func installUbuntu(out io.Writer) error {
	if err := exec.RunCommand(out, "apt", "update"); err != nil {
		return err
	}
	args := []string{
		"apt",
		"install",
		"-y",
	}
	args = append(args, ubuntuPackages...)
	if err := exec.RunCommand(out, args...); err != nil {
		return err
	}
	return nil
}
