package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("read dir", func(t *testing.T) {
		envs, err := ReadDir("testdata/env")

		tests := map[string]string{
			"BAR":   "bar",
			"EMPTY": "",
			"FOO":   "   foo\nwith new line",
			"HELLO": "\"hello\"",
			"UNSET": "",
		}

		require.NoError(t, err)
		for key, value := range tests {
			require.Equal(t, value, envs[key].Value, "key "+key+" not equal")
		}
	})
}
