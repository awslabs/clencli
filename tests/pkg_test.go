package tests

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/helper"
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

func executeCommand(t *testing.T, cmd *cobra.Command, args []string) (stdout string, err error) {
	wd := createAndEnterTestDirectory(t)
	_, stdout, err = executeCommandC(cmd, args)
	os.Chdir(wd)

	return stdout, err
}

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

// createTestDirectory create the testing directory and enters it
func createTestDirectory(t *testing.T) string {
	created := helper.MkDirsIfNotExist(t.Name())
	if !created {
		log.Fatal("Unable to create directory")
	}

	return t.Name()
}

func executeCommandC(cmd *cobra.Command, args []string) (command *cobra.Command, stdout string, err error) {
	buf := new(bytes.Buffer)

	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)

	command, err = rootCmd.ExecuteC()
	stdout = buf.String()
	return command, stdout, err
}

// InitRootAndChildCmd initializes Cobra `root` command and add the `childCmd` to it
func InitRootAndChildCmd(rootCmd *cobra.Command, childCmd *cobra.Command) (*cobra.Command, *cobra.Command) {
	rootCmd.AddCommand(childCmd)
	return rootCmd, childCmd
}
