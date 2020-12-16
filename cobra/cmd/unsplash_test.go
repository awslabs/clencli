package cmd

import (
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/tester"
	"github.com/stretchr/testify/assert"
)

// profiles:
// - name: unit-testing
//   description: Unit Testing
//   enabled: false
//   createdAt: 2020-12-16 20:29:25.990242206 +0800 AWST m=+0.007112078
//   updatedAt: 2020-12-16 20:31:51.50260984 +0800 AWST m=+20.969133917
//   credentials:
//   - name: clencli-unit
//     description: Unsplash Credentials
//     enabled: true
//     createdAt: 2020-12-16 20:29:25.990261473 +0800 AWST m=+0.007131351
//     updatedAt: 2020-12-16 20:32:05.910609998 +0800 AWST m=+35.377134125
//     provider: unsplash
//     accessKey: 
//     secretkey: 

func TestUnsplashEmpty(t *testing.T) {
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash")
	// assert.Contains(t, out, "Usage")
	assert.Contains(t, err.Error(), "Unsplash credential not found")
}

func TestUnsplashWithUntiTestingProfile(t *testing.T) {
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash", "--profile", "unit-testing")
	// assert.Contains(t, out, "Usage")
	assert.Contains(t, err.Error(), "Unsplash credential not found")
}

func TestUnsplashWithQuery(t *testing.T) {
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash", "--query", "horse")
	assert.Contains(t, err.Error(), "invalid argument")
}
