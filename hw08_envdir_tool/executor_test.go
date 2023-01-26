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
}
