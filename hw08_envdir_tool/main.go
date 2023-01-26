package main

import (
	"os"
)

func main() {
	args := os.Args
	env, err := ReadDir(args[1])
	if err != nil {
		panic(err)
	}

	code := RunCmd(args[2:], env)
	if code > 0 {
		panic(code)
	}
}
