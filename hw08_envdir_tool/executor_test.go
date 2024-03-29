package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("exec", func(t *testing.T) {
		envs, _ := ReadDir("testdata/env")
		code := RunCmd([]string{
			"echo",
			"123",
		}, envs)

		require.Equal(t, 0, code)
	})

	t.Run("invalid exec", func(t *testing.T) {
		envs, _ := ReadDir("testdata/env")
		code := RunCmd([]string{
			"/bin/bash",
			"echo 123",
		}, envs)

		require.Equal(t, 127, code)
	})

	t.Run("empty exec", func(t *testing.T) {
		envs, _ := ReadDir("testdata/env")
		code := RunCmd([]string{}, envs)

		require.Equal(t, 1, code)
	})

	t.Run("only command exec", func(t *testing.T) {
		envs, _ := ReadDir("testdata/env")
		code := RunCmd([]string{"echo"}, envs)

		require.Equal(t, 0, code)
	})
}
