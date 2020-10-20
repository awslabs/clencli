package function

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/spf13/cobra"
)

// EmptyRun does ...
func emptyRun(*cobra.Command, []string) {}

// ExecuteCommand does ...
func ExecuteCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

// ExecuteCommandC does ...
func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

// InitRootAndChildCmd does ...
func InitRootAndChildCmd(rootCmd *cobra.Command, childCmd *cobra.Command) (*cobra.Command, *cobra.Command) {
	rootCmd.AddCommand(childCmd)
	return rootCmd, childCmd
}

// Setup does ...
func Setup(t *testing.T) (pwd string, nwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get the current directory path")
	}

	nwd, err = CreateDirectoryNamedPath(t.Name())
	os.Chdir(nwd)

	return cwd, nwd
}

// SetupAll does ...
func SetupAll() (pwd string, nwd string) {
	format := "2006-01-02-15-04-05.000000000"
	dt := time.Now().Format(format)

	cwd, err := os.Getwd()
	path := cwd + "/../test/" + dt

	if err != nil {
		log.Fatal("Unable to get the current directory path")
	}

	nwd, err = CreateDirectoryNamedPath(path)
	if err == nil {
		os.Chdir(nwd)
	}

	return cwd, nwd
}

// Teardown does ...
func Teardown(pwd string, nwd string) {
	os.Chdir(pwd)
	err := os.RemoveAll(nwd)
	if err != nil {
		log.Fatal(err)
	}
}

// TeardownAll does ...
func TeardownAll(pwd string, nwd string) {
	err := os.RemoveAll(nwd)
	if err != nil {
		log.Fatal(err)
	}
}
