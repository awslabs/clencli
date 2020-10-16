package cmd

import (
	"testing"

	fun "github.com/awslabs/clencli/function"
	"github.com/stretchr/testify/assert"
)

func TestRenderWithNoArgAndNoFlags(t *testing.T) {
	rootCmd, _ := fun.InitRootAndChildCmd(RenderCmd())
	output, err := fun.ExecuteCommand(rootCmd, "render")

	assert.Contains(t, output, "Please provide an argument")
	assert.Contains(t, err.Error(), "Please provide an argument")
}
