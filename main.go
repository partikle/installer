package main

import (
	"flag"
	"io"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/partikle/installer/install"
)

func main() {
	workdirFlag := flag.String("w", "", "directory to build rump in")
	rumpdirFlag := flag.String("o", "", "output directory for rump binaries")
	platform := flag.String("p", "", "platform to build for. can be hw, xen, or both")
	flag.Parse()
	log := logrus.New()
	for name, value := range map[string]string{
		"workdir":  *workdirFlag,
		"outdir":   *rumpdirFlag,
		"platform": *platform,
	} {
		if value == "" {
			log.Fatalf("%s must be provided using its given flag. see installer -h for help", name)
		}
	}
	if *platform != "hw" && *platform != "xen" && *platform != "both" {
		log.Fatal("platform must be one of the following, \"hw\", \"xen\", or\"both\"")
	}
	wd, err := filepath.Abs(*workdirFlag)
	if err != nil {
		log.Fatalf("failed to find absolute path to %s", *workdirFlag)
	}
	outdir, err := filepath.Abs(*rumpdirFlag)
	if err != nil {
		log.Fatalf("failed to find absolute path to %s", *rumpdirFlag)
	}

	log.Info("Installing Partikle!")
	var w io.Writer
	if os.Getenv("LOGGING") == "0" {
		w, err = os.Open("/dev/null")
		if err != nil {
			log.Fatalf("error opening /dev/null: %v", err)
		}
	} else {
		w = log.Writer()
	}
	if err := install.Run(w, wd, outdir, *platform); err != nil {
		log.Fatal(err)
	}
	log.Info("done!")
}
