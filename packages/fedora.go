package packages

import (
	"io"
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
	if err := command(out, "dnf", "update"); err != nil {
		return err
	}
	args := []string{
		"dnf",
		"install",
		"-y",
	}
	args = append(args, fedoraPackages...)
	if err := command(out, args...); err != nil {
		return err
	}
	return nil
}
