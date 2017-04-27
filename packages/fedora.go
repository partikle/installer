package packages

import (
	"io"

	"github.com/partikle/installer/pkg/exec"
)

var fedoraPackages = []string{
	"xen-devel.i686",
	"xen-devel.x86_64",
	"curl",
	"git",
	"@development-tools",
	"zlib-devel",
	"which",
}

func installFedora(out io.Writer) error {
	if err := exec.RunCommand(out, "dnf", "update", "-y"); err != nil {
		return err
	}
	args := []string{
		"dnf",
		"install",
		"-y",
	}
	args = append(args, fedoraPackages...)
	if err := exec.RunCommand(out, args...); err != nil {
		return err
	}
	return nil
}
