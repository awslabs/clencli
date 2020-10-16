package cmd

import (
	"testing"

	fun "github.com/awslabs/clencli/function"
	"github.com/stretchr/testify/assert"
)

func TestRenderWithNoArgAndNoFlags(t *testing.T) {
	rootCmd, _ := fun.InitRootAndChildCmd(RootCmd(), RenderCmd())
	output, err := fun.ExecuteCommand(rootCmd, "render")

	assert.Contains(t, output, "one the following arguments are required")
	assert.Contains(t, err.Error(), "one the following arguments are required")
}

func TestRenderWithInvalidArgAndNoFlags(t *testing.T) {
	rootCmd, _ := fun.InitRootAndChildCmd(RootCmd(), RenderCmd())
	output, err := fun.ExecuteCommand(rootCmd, "render", "null")

	assert.Contains(t, output, "invalid argument")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestRenderWithValidArgAndNoFlags(t *testing.T) {
	rootCmd, _ := fun.InitRootAndChildCmd(RootCmd(), RenderCmd())
	output, err := fun.ExecuteCommand(rootCmd, "render", "template")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestRenderWithNameOnly(t *testing.T) {
	rootCmd, renderCmd := fun.InitRootAndChildCmd(RootCmd(), RenderCmd())
	renderCmd.Flags().StringP("name", "n", "readme", "Template name to be rendered")
	output, err := fun.ExecuteCommand(rootCmd, "render", "template")

	assert.Contains(t, output, "Missing database at clencli/readme.yaml")
	assert.Contains(t, err.Error(), "Missing database at clencli/readme.yaml")
}

func TestRenderWithInitBasicProject(t *testing.T) {

	// init a basic project
	pwd, nwd := fun.Setup(t)

	rootCmd, initCmd := fun.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := fun.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	rootCmd, renderCmd := fun.InitRootAndChildCmd(RootCmd(), RenderCmd())
	renderCmd.Flags().StringP("name", "n", "readme", "Template name to be rendered")
	output, err = fun.ExecuteCommand(rootCmd, "render", "template")

	assert.Contains(t, output, "Missing database at clencli/readme.yaml")
	assert.Contains(t, err.Error(), "Missing database at clencli/readme.yaml")

	fun.Teardown(pwd, nwd)
}
