package main

import (
	"bytes"
	"os"
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

		content = bytes.ReplaceAll(content, []byte{0x00}, []byte("\n"))
		env := EnvValue{Value: string(content[:])}
		environment[entity.Name()] = env
	}

	return environment, nil
}
