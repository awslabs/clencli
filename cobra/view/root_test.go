package cmd

import (
	"testing"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/stretchr/testify/assert"
)

func TestRootWithNoArgAndNoFlags(t *testing.T) {
	rootCmd := RootCmd()
	output, err := cau.ExecuteCommand(rootCmd)

	assert.NotEqual(t, rootCmd.Use, "")
	assert.NotEqual(t, rootCmd.Short, "")
	assert.NotEqual(t, rootCmd.Long, "")
	assert.NotEqual(t, output, "")
	assert.Equal(t, err, nil)

}
