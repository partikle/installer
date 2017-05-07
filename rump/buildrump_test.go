package rump_test

import (
	. "github.com/partikle/installer/rump"

	"bytes"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"

	"github.com/partikle/installer/pkg/exec"
	"path/filepath"
)

var _ = Describe("Buildrump", func() {
	var rumpbase, outdir string
	BeforeEach(func() {
		base, err := ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())
		rumpbase = base
		out, err := ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())
		outdir = out
	})
	AfterEach(func() {
		os.RemoveAll(rumpbase)
		os.RemoveAll(outdir)
	})
	Describe("PrepareRumpRepo", func() {
		It("downloads and initializes git repo for rumprun", func() {
			out := &bytes.Buffer{}
			err := PrepareRumpRepo(out, rumpbase)
			if err != nil {
				err = fmt.Errorf("preparing rump repo failed with message %s: %s", out.String(), err.Error())
			}
			Expect(err).NotTo(HaveOccurred())
			out.Reset()
			err = exec.Command(out, "git", "show").Dir(rumpbase).Run()
			Expect(err).NotTo(HaveOccurred())
			Expect(out.String()).To(ContainSubstring("16a7c6eb44523c60ea714a0ec2c7ea6ab3c8fb02"))
			_, err = os.Stat(filepath.Join(rumpbase, "src-netbsd", "Makefile"))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("BuildRump", func() {
		It("installs rump to the workdir", func() {
			out := &bytes.Buffer{}
			err := PrepareRumpRepo(out, rumpbase)
			if err != nil {
				err = fmt.Errorf("preparing rump repo failed with message %s: %s", out.String(), err.Error())
			}
			Expect(err).NotTo(HaveOccurred())
			out.Reset()
			err = BuildRump(out, outdir, rumpbase, "hw")
			if err != nil {
				err = fmt.Errorf("building rump failed with message %s: %s", out.String(), err.Error())
			}
			Expect(err).NotTo(HaveOccurred())
			out.Reset()
			err = exec.Command(out, "ls", outdir).Run()
			Expect(err).NotTo(HaveOccurred())
			lsResult := out.String()
			for _, binaryName := range []string{
				"rumprun",
				"rumprun-bake",
				"x86_64-rumprun-netbsd-ar",
				"x86_64-rumprun-netbsd-as",
				"x86_64-rumprun-netbsd-cookfs",
				"x86_64-rumprun-netbsd-cpp",
				"x86_64-rumprun-netbsd-gcc",
				"x86_64-rumprun-netbsd-ld",
				"x86_64-rumprun-netbsd-nm",
				"x86_64-rumprun-netbsd-objcopy",
				"x86_64-rumprun-netbsd-objdump",
				"x86_64-rumprun-netbsd-ranlib",
				"x86_64-rumprun-netbsd-readelf",
				"x86_64-rumprun-netbsd-size",
				"x86_64-rumprun-netbsd-strings",
				"x86_64-rumprun-netbsd-strip",
			} {
				Expect(lsResult).To(ContainSubstring(binaryName))
			}
		})
	})
})
