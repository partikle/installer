package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/partikle/installer/packages"
)

func main() {
	log := logrus.New()
	log.Info("Installing Partikle!")
	if err := packages.Install(log.Writer()); err != nil {
		log.Fatal(err)
	}
	log.Info("done!")
}
