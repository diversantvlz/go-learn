package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Envdir is required")
		os.Exit(1)
	}

	env, err := ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(RunCmd(args[2:], env))
}
