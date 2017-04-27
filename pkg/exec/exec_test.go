package exec_test

import (
	. "github.com/partikle/installer/pkg/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Exec", func() {
	Describe("Command", func() {
		It("returns a command object which can be run", func() {
			cmd := Command(os.Stdout, "ls", os.Getenv("HOME"))
			err := cmd.Run()
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
