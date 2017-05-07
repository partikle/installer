package rump_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRump(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rump Suite")
}
