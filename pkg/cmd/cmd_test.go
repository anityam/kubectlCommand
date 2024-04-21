package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKubectlCommandHeadersHooks(t *testing.T) {
	rootCmd := NewKubingCommand()
	require.NoError(t, rootCmd.Execute(), "expected no error")
}
