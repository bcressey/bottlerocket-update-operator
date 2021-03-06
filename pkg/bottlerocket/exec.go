package bottlerocket

import "syscall"

const RootFS = "/.bottlerocket/rootfs"

// PlatformBin is where platform interfacing executables are located.
const PlatformBin = "/usr/bin"

// ProcessAttrs may be used to exec a process in the PlatformBin.
func ProcessAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Chroot: RootFS,
	}
}
