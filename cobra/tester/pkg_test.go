package tester

import (
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

// ExecuteCommand adds cmd (which should be from controller package) into the Root command, executes the Cobra command given the args and returns error
func executeCommand(t *testing.T, cmd *cobra.Command, args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := createTestDirectory(t)
	os.Chdir(dir)

	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetArgs(args)
	err = rootCmd.Execute()

	os.Chdir(wd)

	return err
}

// CreateTestDirectory create the testing directory and enters it
func createTestDirectory(t *testing.T) string {
	path, err := helper.CreateDirectoryNamedPath(t.Name())
	if err != nil {
		log.Fatal("Unable to create directory")
	}

	return path
}
