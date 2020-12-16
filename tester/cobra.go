package tester

import (
	"github.com/awslabs/clencli/cobra/controller"
	"github.com/spf13/cobra"
)

// ExecuteCommand adds cmd (which should be from controller package) into the Root command, executes given the args and returns error
func ExecuteCommand(cmd *cobra.Command, args ...string) error {
	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return err
}
