package install

import (
	"io"
	"path/filepath"

	"github.com/partikle/installer/packages"
	"github.com/partikle/installer/rump"
)

// Main API method to install Rumprun dependencies and binaries
// w is a writer for logging & output of exec comands
// workdir is where we will clone rumprun repo to
// outdir is where all output binaries are placed
// platform is hw|xen|both
func Run(w io.Writer, workdir, outdir, platform string) error {
	if err := packages.Install(w); err != nil {
		return err
	}
	if err := rump.PrepareRumpRepo(w, workdir); err != nil {
		return err
	}
	if err := rump.BuildRump(w, workdir, outdir, platform); err != nil {
		return err
	}
	if err := rump.ApplyPatches(filepath.Join(outdir, "rumprun"), platform); err != nil {
		return err
	}
	if err := rump.BuildRump(w, workdir, outdir, platform); err != nil {
		return err
	}
	return nil
}
