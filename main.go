package main

import (
	"fmt"
	"os"
	"os/exec"
	"mlog"
)

func main() {
	spawnProcess("echo","-e", "Hello World!")
}


func spawnProcess(arg string, args ...string) {
	if len(args) == 0 {
		return
	}
	switch arg {
	case "run":
		run(args[0], args[1:])
	case "proc":
		proc(args[0], args[1:])
	default:
		mlog.Errorf("Unknown command. Exiting..")
	}
}

func run(arg string, args ...string) {
	fmt.Println("[pequod init]")
	cmd := exec.Command("/proc/self/exe", append([]string{"proc", arg}, args[0:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		// new user time space -> clones the hostname from the host name | new process ID | 
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	cmd.Run()
}


func proc(arg string, args ...string) {
//	[TODO:] make a random string generator
	syscall.Sethostname([]byte("container"))
	// Change ps1 style
	cmd := exec.Command(arg)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()
}
