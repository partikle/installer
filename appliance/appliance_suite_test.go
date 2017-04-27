package appliance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAppliance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Appliance Suite")
}
