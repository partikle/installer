package appliance_test

import (
	"bytes"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/partikle/installer/pkg/exec"
	"github.com/partikle/installer/pkg/test"
	"github.com/pborman/uuid"
)

var _ = Describe("Docker", func() {
	var (
		fedoraDockerfile = "Dockerfile.fedora"
		ubuntuDockerfile = "Dockerfile.ubuntu"
		fedoraImageName  = uuid.New()
		ubuntuImageName  = uuid.New()
		wd               string
		err              error
	)
	Describe("docker build", func() {
		BeforeEach(func() {
			wd, err = test.GetCurrentDir()
			Expect(err).NotTo(HaveOccurred())
		})
		for dockerfile, imageName := range map[string]string{
			fedoraDockerfile: fedoraImageName,
			ubuntuDockerfile: ubuntuImageName,
		} {
			It("should build "+dockerfile+" with rumprun installed", func() {
				cmd := exec.Command(os.Stdout, "docker", "build", "-t", imageName, "-f", dockerfile, wd)
				defer exec.RunCommand(os.Stdout, "docker", "rmi", "-f", imageName)
				err = cmd.Run()
				Expect(err).NotTo(HaveOccurred())
				out := &bytes.Buffer{}
				cmd = exec.Command(out, "docker", "run", "--rm", imageName, "/bin/bash", "-c", "ls /usr/local/bin | grep rumprun")
				err = cmd.Run()
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
		}
	})
})
