package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("copy full", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata/input.txt", tmp.Name(), 0, 0)

		expected, _ := os.ReadFile("testdata/out_offset0_limit0.txt")
		actual, _ := os.ReadFile(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("copy 0:10", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata/input.txt", tmp.Name(), 0, 10)

		expected, _ := os.ReadFile("testdata/out_offset0_limit10.txt")
		actual, _ := os.ReadFile(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("copy 0:1000", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata/input.txt", tmp.Name(), 0, 1000)

		expected, _ := os.ReadFile("testdata/out_offset0_limit1000.txt")
		actual, _ := os.ReadFile(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("copy 0:10000", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata/input.txt", tmp.Name(), 0, 10000)

		expected, _ := os.ReadFile("testdata/out_offset0_limit10000.txt")
		actual, _ := os.ReadFile(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("copy 100:1000", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata/input.txt", tmp.Name(), 100, 1000)

		expected, _ := os.ReadFile("testdata/out_offset100_limit1000.txt")
		actual, _ := os.ReadFile(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("copy 6000:1000", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata/input.txt", tmp.Name(), 6000, 1000)

		expected, _ := os.ReadFile("testdata/out_offset6000_limit1000.txt")
		actual, _ := os.ReadFile(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("copy 10000:1000", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata/input.txt", tmp.Name(), 10000, 1000)
		require.Error(t, err)
	})

	t.Run("copy no length", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("/dev/urandom", tmp.Name(), 0, 0)
		require.Error(t, err)
	})

	t.Run("copy dir", func(t *testing.T) {
		tmp, _ := os.CreateTemp("", "")
		err := Copy("testdata", tmp.Name(), 0, 0)
		require.Error(t, err)
	})

	t.Run("copy self", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/input.txt", 0, 0)
		require.Error(t, err)
	})
}
