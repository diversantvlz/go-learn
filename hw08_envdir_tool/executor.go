package main

import (
	"bytes"
	"fmt"
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

	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(out.String())
	return 0
}
