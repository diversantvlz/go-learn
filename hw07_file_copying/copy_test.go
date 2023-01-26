package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	outputFile := "testdata/output.txt"

	t.Run("copy full", func(t *testing.T) {
		err := Copy("testdata/input.txt", outputFile, 0, 0)

		expected, _ := os.ReadFile("testdata/out_offset0_limit0.txt")
		actual, _ := os.ReadFile(outputFile)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
		_ = os.Remove(outputFile)
	})

	t.Run("copy 0:10", func(t *testing.T) {
		err := Copy("testdata/input.txt", outputFile, 0, 10)

		expected, _ := os.ReadFile("testdata/out_offset0_limit10.txt")
		actual, _ := os.ReadFile(outputFile)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
		_ = os.Remove(outputFile)
	})

	t.Run("copy 0:1000", func(t *testing.T) {
		err := Copy("testdata/input.txt", outputFile, 0, 1000)

		expected, _ := os.ReadFile("testdata/out_offset0_limit1000.txt")
		actual, _ := os.ReadFile(outputFile)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
		_ = os.Remove(outputFile)
	})

	t.Run("copy 0:10000", func(t *testing.T) {
		err := Copy("testdata/input.txt", outputFile, 0, 10000)

		expected, _ := os.ReadFile("testdata/out_offset0_limit10000.txt")
		actual, _ := os.ReadFile(outputFile)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
		_ = os.Remove(outputFile)
	})

	t.Run("copy 100:1000", func(t *testing.T) {
		err := Copy("testdata/input.txt", outputFile, 100, 1000)

		expected, _ := os.ReadFile("testdata/out_offset100_limit1000.txt")
		actual, _ := os.ReadFile(outputFile)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
		_ = os.Remove(outputFile)
	})

	t.Run("copy 6000:1000", func(t *testing.T) {
		err := Copy("testdata/input.txt", outputFile, 6000, 1000)

		expected, _ := os.ReadFile("testdata/out_offset6000_limit1000.txt")
		actual, _ := os.ReadFile(outputFile)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
		_ = os.Remove(outputFile)
	})

	t.Run("copy 10000:1000", func(t *testing.T) {
		err := Copy("testdata/input.txt", outputFile, 10000, 1000)
		require.Error(t, err)
	})

	t.Run("copy no length", func(t *testing.T) {
		err := Copy("/dev/urandom", outputFile, 0, 0)
		require.Error(t, err)
	})

	t.Run("copy dir", func(t *testing.T) {
		err := Copy("testdata", outputFile, 0, 0)
		require.Error(t, err)
	})
}
