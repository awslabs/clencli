package tester

import (
	"log"
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

// ExecuteCommand adds cmd (which should be from controller package) into the Root command, executes given the args and returns error
func executeCommand(t *testing.T, cmd *cobra.Command, args ...string) error {
	// reset current directory to the temporary master directory
	// os.Chdir(tDir)

	// td := createTestDirectory(t)
	// os.Chdir(td)

	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()

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
