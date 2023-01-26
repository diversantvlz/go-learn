package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, val := range env {
		var err error
		if val.NeedRemove {
			err = os.Unsetenv(name)
		} else {
			err = os.Setenv(name, val.Value)
		}

		if err != nil {
			panic(err)
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return exiterr.ExitCode()
		}

		panic(err)
	}

	return 0
}
