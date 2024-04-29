package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Sometimes we just want to completely replace the current Go process with another

	// Search for "ls" binary command's absolute path
	command, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// command and arguments in slice form
	args := []string{"ls", "-a", "-l", "-h"}

	// load enviromnent, required by syscall.Exec
	env := os.Environ()

	// execute command
	// if the next call is succesful the execution of this
	// process ends here and is replaced by ""
	execErr := syscall.Exec(command, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
