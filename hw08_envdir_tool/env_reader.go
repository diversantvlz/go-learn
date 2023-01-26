package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entities, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environment := Environment{}

	err = os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	for _, entity := range entities {
		content, err := os.ReadFile(entity.Name())
		if err != nil {
			return nil, err
		}

		env := EnvValue{}

		if len(content) == 0 {
			env.NeedRemove = true
		} else {
			_, firstLine, err := bufio.ScanLines(content, true)
			if err != nil {
				return nil, err
			}

			firstLine = bytes.ReplaceAll(firstLine, []byte{0x00}, []byte("\n"))
			env.Value = strings.TrimRight(string(firstLine[:]), " \t")
		}
		environment[entity.Name()] = env
	}

	return environment, nil
}
