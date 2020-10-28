package cmd

import (
	"testing"

	con "github.com/awslabs/clencli/cobra/controller"
	helper "github.com/awslabs/clencli/helper"
	"github.com/stretchr/testify/assert"
)

func TestConfigureWithNoArgAndNoFlags(t *testing.T) {
	rootCmd, _ := helper.InitRootAndChildCmd(con.RootCmd(), con.ConfigureCmd())
	output, err := helper.ExecuteCommand(rootCmd, "configure")

	assert.Contains(t, output, "one the following arguments are required")
	assert.Contains(t, err.Error(), "one the following arguments are required")
}
