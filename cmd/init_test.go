package cmd

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	fun "github.com/awslabs/clencli/function"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func emptyRun(*cobra.Command, []string) {}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
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

func initRootAndChildCmd() (*cobra.Command, *cobra.Command) {
	rootCmd := &cobra.Command{Use: "root", Args: cobra.NoArgs, Run: emptyRun}
	childCmd := InitCmd()
	rootCmd.AddCommand(childCmd)

	return rootCmd, childCmd
}

func TestMain(m *testing.M) {
	format := "2006-01-02-15-04-05.000000000"
	dt := time.Now().Format(format)
	path := "../test/" + dt

	err := fun.CreateDirectoryNamedPath(path)
	if err == nil {
		// dir created
		os.Chdir(path)
		code := m.Run()
		os.Exit(code)
	} else {
		log.Fatal("Something went wrong", err)
	}
}

func TestInitWithNoArgAndNoFlags(t *testing.T) {
	rootCmd, _ := initRootAndChildCmd()
	output, err := executeCommand(rootCmd, "init")

	assert.Contains(t, output, "Please provide an argument")
	assert.Contains(t, err.Error(), "Please provide an argument")
}

func TestInitWithInvalidArgAndNoFlags(t *testing.T) {
	rootCmd, _ := initRootAndChildCmd()
	output, err := executeCommand(rootCmd, "init", "null")

	assert.Contains(t, output, "invalid argument")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestInitWithValidArgAndNoFlags(t *testing.T) {

	rootCmd, _ := initRootAndChildCmd()
	output, err := executeCommand(rootCmd, "init", "project")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestInitWithEmptyName(t *testing.T) {

	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	output, err := executeCommand(rootCmd, "init", "project", "--name")

	assert.Contains(t, output, "flag needs an argument")
	assert.Contains(t, err.Error(), "flag needs an argument")
}

func TestInitWithNameOnly(t *testing.T) {
	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--name", "generated-project")

	if err != nil {
		log.Fatal("Something went wrong during project initialization: ", err, output)
	}
}