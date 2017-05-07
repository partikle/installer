package rump

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ilackarms/pkg/errors"
	"github.com/partikle/installer/pkg/fileutils"
	"github.com/partikle/installer/rump/bindata"
)

func ApplyPatches(rumproot, platform string) error {
	switch platform {
	case "hw":
		if err := applyHWPatches(rumproot); err != nil {
			return errors.New("applying hw patches", err)
		}
	case "xen":
		if err := applyXenPatches(rumproot); err != nil {
			return errors.New("applying xen patches", err)
		}
	case "both":
		if err := applyHWPatches(rumproot); err != nil {
			return errors.New("applying hw patches", err)
		}
		if err := applyXenPatches(rumproot); err != nil {
			return errors.New("applying xen patches", err)
		}
	default:
		return errors.New("unknown platform "+platform, nil)
	}
	if err := applyGenericPatches(rumproot); err != nil {
		return errors.New("applying generic rumprun patches", err)
	}
	return nil
}

// hw patches
func applyHWPatches(rumproot string) error {
	// ppb patch
	if err := applyPPBpatch(rumproot); err != nil {
		return errors.New("applying pci-pci bridge patch", err)
	}
	// add ppb pride for vmware network cards
	if err := addPPBtoMakefile(rumproot); err != nil {
		return errors.New("adding ppb.c to makefile", err)
	}
	// add scsi driver for vmware hard drives
	if err := addSCSIdriver(rumproot); err != nil {
		return errors.New("adding scsi driver for vmmware", err)
	}
	// patch rumprun-bake.conf
	if err := patchRumpbakeConf(rumproot); err != nil {
		return errors.New("patching rumprun-bake.conf", err)
	}
	// patch kernel linker script
	if err := patchKernelLinkerscript(rumproot); err != nil {
		return errors.New("patching kernel linker script", err)
	}
	return nil
}

func applyPPBpatch(rumproot string) error {
	hwPCIConf := filepath.Join(rumproot, "src-netbsd/sys/rump/dev/lib/libpci/PCI.ioconf")
	if err := fileutils.Append(hwPCIConf, []byte(`pci*    at ppb? bus ?`)); err != nil {
		return errors.New("appending file "+hwPCIConf, err)
	}
	if err := fileutils.Append(hwPCIConf, []byte(`ppb*    at pci? dev ? function ?`)); err != nil {
		return errors.New("appending file "+hwPCIConf, err)
	}
	return nil
}

func addPPBtoMakefile(rumproot string) error {
	makefile := filepath.Join(rumproot, "src-netbsd/src-netbsd/sys/rump/dev/lib/libpci/Makefile")
	data, err := ioutil.ReadFile(makefile)
	if err != nil {
		return errors.New("failed to read makefile at "+makefile, err)
	}
	rxp := regexp.MustCompile("SRCS+=\tpci.c")
	data = rxp.ReplaceAll(data, []byte("SRCS+=	ppb.c pci.c"))
	fileinfo, err := os.Stat(makefile)
	if err != nil {
		return errors.New("failed to get file info for "+makefile, err)
	}
	if err := ioutil.WriteFile(makefile, data, fileinfo.Mode()); err != nil {
		return errors.New("writing modified makefile", err)
	}
	return nil
}

func addSCSIdriver(rumproot string) error {
	//equivalent to touch
	os.Create(filepath.Join(rumproot, "src-netbsd/sys/dev/ic/bio.h"))
	if err := extractFiles(filepath.Join(rumproot, "src-netbsd/sys/rump/dev/lib"),
		"libpci_scsi/Makefile",
		"libpci_scsi/PCI_SCSI.ioconf",
	); err != nil {
		return errors.New("extracting libpci_sci files", err)
	}
	if err := extractFiles(filepath.Join(rumproot, "src-netbsd/sys/rump/dev/lib/libscsipi"),
		"scsipi_component.c",
	); err != nil {
		return errors.New("extracting scsipi_component.c", err)
	}
	makefile := filepath.Join(rumproot, "src-netbsd/sys/rump/dev/Makefile.rumpdevcomp")
	data, err := ioutil.ReadFile(makefile)
	if err != nil {
		return errors.New("failed to read makefile at "+makefile, err)
	}
	rxp := regexp.MustCompile("RUMPPCIDEVS+=\tmiiphy")
	data = rxp.ReplaceAll(data, []byte("RUMPPCIDEVS+=  pci_scsi miiphy"))
	fileinfo, err := os.Stat(makefile)
	if err != nil {
		return errors.New("failed to get file info for "+makefile, err)
	}
	if err := ioutil.WriteFile(makefile, data, fileinfo.Mode()); err != nil {
		return errors.New("writing modified makefile", err)
	}
	return nil
}

func patchRumpbakeConf(rumproot string) error {
	if err := extractFiles(filepath.Join(rumproot, "app-tools"),
		"rumprun-bake.conf",
	); err != nil {
		return errors.New("extracting rumprun-bake.conf", err)
	}
	return nil
}

func patchKernelLinkerscript(rumproot string) error {
	if err := extractFiles(filepath.Join(rumproot, "platform/hw/arch/amd64"),
		"rump/kern.ldscript",
	); err != nil {
		return errors.New("extracting file", err)
	}
	return nil
}

// xen patches
func applyXenPatches(rumproot string) error {
	if err := extractFiles(filepath.Join(rumproot, "platform/xen/xen/arch/x86"),
		"rump/minios-x86_64.lds",
	); err != nil {
		return errors.New("extracting minios-x86_64 linker script", err)
	}
	return nil
}

// both providers
func applyGenericPatches(rumproot string) error {
	if err := extractFiles(filepath.Join(rumproot, "lib/librumprun_base"),
		"rump/rumprun.c",
	); err != nil {
		return errors.New("extracting rumprun.c", err)
	}
	if err := extractFiles(filepath.Join(rumproot, "app-tools"),
		"rump/rumprun-bake.in.c",
	); err != nil {
		return errors.New("extracting rumprun-bake.in", err)
	}
	if err := extractFiles(filepath.Join(rumproot, "buildrump.sh/brlib/libnetconfig/dhcp_configure.c"),
		"buildrump.sh/brlib/libnetconfig/dhcp_configure.c",
	); err != nil {
		return errors.New("extracting dhcp_configure.c", err)
	}
	return nil
}

// helper function
func extractFiles(destDir string, files ...string) error {
	for _, file := range files {
		file = filepath.Join("patches", file)
		extractedPath := filepath.Join(destDir, filepath.Base(file))
		os.MkdirAll(destDir, 0755)
		data, err := bindata.Asset(file)
		if err != nil {
			return errors.New("failed to open bindata asset "+file, err)
		}
		info, err := bindata.AssetInfo(file)
		if err != nil {
			return errors.New("failed to get file info for bindata asset "+file, err)
		}
		if err := ioutil.WriteFile(extractedPath, data, info.Mode()); err != nil {
			return errors.New("writing contents of bindata asset "+file+" to "+extractedPath, err)
		}
	}
	return nil
}
