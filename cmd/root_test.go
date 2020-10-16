package cmd

import (
	"testing"

	fun "github.com/awslabs/clencli/function"
	"github.com/stretchr/testify/assert"
)

func TestRootWithNoArgAndNoFlags(t *testing.T) {
	rootCmd := RootCmd()
	output, err := fun.ExecuteCommand(rootCmd)

	assert.NotEqual(t, rootCmd.Use, "")
	assert.NotEqual(t, rootCmd.Short, "")
	assert.NotEqual(t, rootCmd.Long, "")
	assert.NotEqual(t, output, "")
	assert.Equal(t, err, nil)

}
