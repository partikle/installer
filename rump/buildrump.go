package rump

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/ilackarms/pkg/errors"
	"github.com/partikle/installer/pkg/exec"
)

func PrepareRumpRepo(out io.Writer, workdir string) error {
	if err := exec.Command(out,
		"git",
		"clone",
		"https://github.com/rumpkernel/rumprun",
	).Dir(workdir).Run(); err != nil {
		return err
	}
	if err := exec.Command(out,
		"git",
		"checkout",
		"16a7c6eb44523c60ea714a0ec2c7ea6ab3c8fb02",
	).Dir(filepath.Join(workdir, "rumprun")).Run(); err != nil {
		return err
	}
	if err := exec.Command(out,
		"git",
		"submodule",
		"update",
		"--init",
	).Dir(filepath.Join(workdir, "rumprun")).Run(); err != nil {
		return err
	}
	return nil
}

func BuildRump(out io.Writer, workdir, outdir, platform string) error {
	fmt.Fprintf(out, "Building Rumprun for %s, using workdir %s. "+
		"Binaries will be saved to %s\n", platform, workdir, outdir)
	if platform == "both" {
		if err := runBuildrumpSH(out, workdir, outdir, "hw"); err != nil {
			return errors.New("failed installing for hw", err)
		}
		if err := runBuildrumpSH(out, workdir, outdir, "xen"); err != nil {
			return errors.New("failed installing for hw", err)
		}
	} else {
		if err := runBuildrumpSH(out, workdir, outdir, platform); err != nil {
			return errors.New("failed installing for "+platform, err)
		}
	}
	return nil
}

func runBuildrumpSH(out io.Writer, workdir, outdir, platform string) error {
	buildRR := filepath.Join(workdir, "rumprun", "build-rr.sh")
	if err := exec.Command(out,
		buildRR,
		"-d",
		outdir,
		"-o",
		"./obj",
		platform,
		"build",
		"--",
		"",
	).Dir(filepath.Join(workdir, "rumprun")).Run(); err != nil {
		return err
	}
	if err := exec.Command(out,
		buildRR,
		"-d",
		outdir,
		"-o",
		"./obj",
		platform,
		"install",
	).Dir(filepath.Join(workdir, "rumprun")).Run(); err != nil {
		return err
	}
	return nil
}
