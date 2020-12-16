package cmd

import (
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/tester"
	"github.com/stretchr/testify/assert"
)

func TestUnsplashEmpty(t *testing.T) {
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash")
	// assert.Contains(t, out, "Usage")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestUnsplashWithQuery(t *testing.T) {
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash", "--query", "horse")
	assert.Contains(t, err.Error(), "invalid argument")
}
