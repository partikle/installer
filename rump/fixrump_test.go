package rump_test

import (
	. "github.com/partikle/installer/rump"

	"bytes"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
)

var _ = FDescribe("Fixrump", func() {
	var rumpbase string
	BeforeEach(func() {
		base, err := ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())
		rumpbase = base
	})
	AfterEach(func() {
		os.RemoveAll(rumpbase)
	})
	Describe("ApplyPatches", func() {
		It("Installs the necessary patches for the specified platform", func() {
			out := &bytes.Buffer{}
			err := PrepareRumpRepo(out, rumpbase)
			if err != nil {
				err = fmt.Errorf("preparing rump repo failed with message %s: %s", out.String(), err.Error())
			}
			Expect(err).NotTo(HaveOccurred())
			out.Reset()

			platform := "both"
			err = ApplyPatches(rumpbase, platform)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
