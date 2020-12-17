package cmd

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/tester"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	tester.SetupAll()

	// comment the line below if you want to keep the test results
	// pwd, nwd := tester.SetupAll()

	code := m.Run()
	// comment the line below if you want to keep the test results
	// tester.TeardownAll(pwd, nwd)
	os.Exit(code)
}

func executeConfigure(args ...string) error {
	cmd := controller.ConfigureCmd()
	cmd.SetArgs(args)
	err := cmd.Execute()
	return err
}

func TestConfigureEmpty(t *testing.T) {
	err := executeConfigure("")
	// assert.Contains(t, out, "Usage")
	assert.Contains(t, err.Error(), "invalid argument")
}

// TODO: test configure command and provide the input via the test block
