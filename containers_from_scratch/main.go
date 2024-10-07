//go:build linux

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// go run main.go run <command> <args>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("unknown command")
	}
}

func run() {
	fmt.Printf("Running Main Run Function: %v as pid: %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS, // <= don't share information on mounted fs from the container back to the host
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running Inside Child: %v as pid: %d\n", os.Args[2:], os.Getpid())
	syscall.Sethostname([]byte("container"))

	// here I'm mounting a ubuntu fs downloaded and
	// extracted from https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-root.tar.xz to a path on the host ("/Users/root_for_container")
	syscall.Chroot("/Users/root_for_container")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	must(cmd.Run())
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")

	err := os.Mkdir(filepath.Join(pids, "trawler"), 0755)
	if err != nil {
		panic(err)
	}

	must(ioutil.WriteFile(filepath.Join(pids, "trawler/pids.max"), []byte("20"), 0700)) // no more than 20 processes

	// If the notify_on_release flag is enabled (1) in a cgroup, then
	// whenever the last task in the cgroup leaves (exits or attaches to
	// some other cgroup) and the last child cgroup of that cgroup
	// is removed, then the kernel runs the command specified by the contents
	// of the "release_agent" file in that hierarchy's root directory,
	// supplying the pathname (relative to the mount point of the cgroup
	// file system) of the abandoned cgroup.  This enables automatic
	// removal of abandoned cgroups.
	must(ioutil.WriteFile(filepath.Join(pids, "trawler/notify_on_release"), []byte("1"), 0700))
	// cgroup.procs: list of thread group IDs in the cgroup.
	must(ioutil.WriteFile(filepath.Join(pids, "trawler/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
