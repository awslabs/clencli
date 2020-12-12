package tester

import (
	"bytes"

	"github.com/spf13/cobra"
)

func emptyRun(*cobra.Command, []string) {}

// ExecuteCommand execute `root` Cobra command and return its outputs
func ExecuteCommand(root *cobra.Command, args ...string) (err error) {
	_, _, err = executeCommandC(root, args...)
	return err
}

// ExecuteCommandOutput execute `root` Cobra command and return its outputs
func ExecuteCommandOutput(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

// InitRootCmd initializes Cobra `root` command and add the `childCmd` to it
func InitRootCmd(rootCmd *cobra.Command, childCmd *cobra.Command) *cobra.Command {
	rootCmd.AddCommand(childCmd)
	return rootCmd
}

// InitRootAndChildCmd initializes Cobra `root` command and add the `childCmd` to it
func InitRootAndChildCmd(rootCmd *cobra.Command, childCmd *cobra.Command) (*cobra.Command, *cobra.Command) {
	rootCmd.AddCommand(childCmd)
	return rootCmd, childCmd
}
