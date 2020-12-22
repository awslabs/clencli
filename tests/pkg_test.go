package tests

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/* SETUP */

func beforeSetup() {
	format := "2006-01-02-15-04-05.000000000"
	dt := time.Now().Format(format)

	dir := helper.CreateTempDir(os.TempDir(), "clencli-"+dt)

	// enter the new directory
	err := os.Chdir(dir)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	beforeSetup()
	os.Exit(m.Run())
}

/* COBRA */

// ExecuteCommand adds cmd (which should be from controller package) into the Root command, executes the Cobra command given the args and returns error
// func executeCommand(t *testing.T, cmd *cobra.Command, args []string]) error {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}

// 	dir := createTestDirectory(t)
// 	os.Chdir(dir)

// 	rootCmd := controller.RootCmd()
// 	rootCmd.AddCommand(cmd)
// 	rootCmd.SetArgs(args)
// 	err = rootCmd.Execute()

// 	os.Chdir(wd)

// 	return err
// }

// CreateTestDirectory create the testing directory and enters it
func createTestDirectory(t *testing.T) string {
	created := helper.MkDirsIfNotExist(t.Name())
	if !created {
		log.Fatal("Unable to create directory")
	}

	return t.Name()
}

func emptyRun(*cobra.Command, []string) {}

/* NEW COBRA */

// createAndEnterTestDirectory create the test directory, enter it and
// return the current working directory, useful to return to the previous directory
func createAndEnterTestDirectory(t *testing.T) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get current working directory")
	}

	dir := createTestDirectory(t)
	os.Chdir(dir)
	return wd
}

// ExecuteCommand execute `root` Cobra command and return its outputs
func executeCommand(t *testing.T, cmd *cobra.Command, args []string) (output string, err error) {
	wd := createAndEnterTestDirectory(t)
	cmd, stdout, stderr, err := executeCommandC(cmd, args)
	logrus.Debugf("Test %s has the following stdout:\n%s", t.Name(), stdout)
	logrus.Debugf("Test %s has the following stderr:\n%s", t.Name(), stderr)
	os.Chdir(wd)

	return output, err
}

func executeCommandR(t *testing.T, cmd *cobra.Command, args []string) (output string, err error) {
	wd := createAndEnterTestDirectory(t)

	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	_, output, err = executeCommandRC(rootCmd, args)
	logrus.Debugf("Test %s has the following stdout:\n%s", t.Name(), output)
	os.Chdir(wd)

	return output, err
}

func executeCommandC(cmd *cobra.Command, args []string) (*cobra.Command, string, string, error) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs(args)

	cmd, err := rootCmd.ExecuteC()

	return cmd, stdout.String(), stderr.String(), err
}

func executeCommandRC(cmd *cobra.Command, args []string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	c, err = cmd.ExecuteC()

	return c, buf.String(), err
}

// InitRootAndChildCmd initializes Cobra `root` command and add the `childCmd` to it
func InitRootAndChildCmd(rootCmd *cobra.Command, childCmd *cobra.Command) (*cobra.Command, *cobra.Command) {
	rootCmd.AddCommand(childCmd)
	return rootCmd, childCmd
}
