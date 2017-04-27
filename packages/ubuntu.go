package packages

import (
	"io"
)

var ubuntuPackages = []string{
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
	args = append(args, ubuntuPackages...)
	if err := command(out, args...); err != nil {
		return err
	}
	return nil
}
