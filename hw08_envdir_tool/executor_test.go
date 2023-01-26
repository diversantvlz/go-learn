package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("exec", func(t *testing.T) {
		envs, _ := ReadDir("testdata/env")
		code := RunCmd([]string{
			"/bin/bash",
			"testdata/echo.sh",
			"arg1=1",
			"arg2=2",
		}, envs)

		require.Equal(t, 0, code)
	})
}
