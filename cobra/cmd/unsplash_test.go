package cmd

import (
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func executeUnsplash(args ...string) error {
	rootCmd := controller.RootCmd()
	cmd := controller.UnsplashCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return err
}

func TestUnsplashEmpty(t *testing.T) {
	err := executeUnsplash("unsplash")
	// assert.Contains(t, out, "Usage")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestUnsplashWithQuery(t *testing.T) {
	err := executeUnsplash("--query", "horse")
	assert.Contains(t, err.Error(), "invalid argument")
}
