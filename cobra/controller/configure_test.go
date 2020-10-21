package controller

import (
	"testing"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/stretchr/testify/assert"
)

func TestConfigureCreateConfigDir(t *testing.T) {
	pwd, nwd := cau.Setup(t)
	createConfigDir()
	cau.Teardown(pwd, nwd)
}

func TestConfigureWithNoArgAndNoFlags(t *testing.T) {
	rootCmd, _ := cau.InitRootAndChildCmd(RootCmd(), ConfigureCmd())
	output, err := cau.ExecuteCommand(rootCmd, "configure")

	assert.Contains(t, output, "one the following arguments are required")
	assert.Contains(t, err.Error(), "one the following arguments are required")
}
