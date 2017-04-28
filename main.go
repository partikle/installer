package main

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"github.com/partikle/installer/packages"
	"github.com/partikle/installer/rump"
	"io"
	"os"
	"path/filepath"
)

func main() {
	workdir := flag.String("w", "", "directory to build rump in")
	rumpdir := flag.String("o", "", "output directory for rump binaries")
	platform := flag.String("p", "", "platform to build for. can be hw, xen, or both")
	flag.Parse()
	log := logrus.New()
	for name, value := range map[string]string{
		"workdir":  *workdir,
		"rumpdir":  *rumpdir,
		"platform": *platform,
	} {
		if value == "" {
			log.Fatalf("%s must be provided using its given flag. see installer -h for help", name)
		}
	}
	if *platform != "hw" && *platform != "xen" && *platform != "both" {
		log.Fatal("platform must be one of the following, \"hw\", \"xen\", or\"both\"")
	}
	wd, err := filepath.Abs(*workdir)
	if err != nil {
		log.Fatalf("failed to find absolute path to %s", *workdir)
	}
	rd, err := filepath.Abs(*rumpdir)
	if err != nil {
		log.Fatalf("failed to find absolute path to %s", *rumpdir)
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
	if err := packages.Install(w); err != nil {
		log.Fatal(err)
	}
	if err := rump.BuildRump(w, wd, rd, *platform); err != nil {
		log.Fatal(err)
	}
	log.Info("done!")
}
