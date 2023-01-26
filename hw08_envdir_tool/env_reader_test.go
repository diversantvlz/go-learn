package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("read dir", func(t *testing.T) {
		envs, err := ReadDir("testdata/env")

		require.NoError(t, err)
		require.Equal(t, "bar", envs["BAR"].Value)
		require.Equal(t, "", envs["EMPTY"].Value)
		require.Equal(t, "   foo\nwith new line", envs["FOO"].Value)
		require.Equal(t, "\"hello\"", envs["HELLO"].Value)
		require.Equal(t, "", envs["UNSET"].Value)
	})
}
