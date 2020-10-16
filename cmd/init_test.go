package cmd

import (
	"bytes"
	"fmt"
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

func setup(t *testing.T) (pwd string, nwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get the current directory path")
	}

	nwd, err = fun.CreateDirectoryNamedPath(t.Name())
	os.Chdir(nwd)

	return cwd, nwd
}

func setupAll() (pwd string, nwd string) {
	format := "2006-01-02-15-04-05.000000000"
	dt := time.Now().Format(format)

	cwd, err := os.Getwd()
	path := cwd + "/../test/" + dt

	if err != nil {
		log.Fatal("Unable to get the current directory path")
	}

	nwd, err = fun.CreateDirectoryNamedPath(path)
	if err == nil {
		os.Chdir(nwd)
	}

	return cwd, nwd
}

func teardown(pwd string, nwd string) {
	fmt.Println("Going back to ", pwd)
	os.Chdir(pwd)
	// err := os.RemoveAll(nwd)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func teardownAll(pwd string, nwd string) {
	fmt.Println("About to delete ", nwd)
	err := os.RemoveAll(nwd)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	// pwd, nwd := setupAll()
	setupAll()
	code := m.Run()
	// comment the line below if you want to keep the test results
	// teardownAll(pwd, nwd)
	os.Exit(code)
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
	pwd, nwd := setup(t)

	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--name", "generated-project")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	teardown(pwd, nwd)
}

func TestInitWithValidTypeOnly(t *testing.T) {
	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--type", "basic")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestInitWithInvalidTypeOnly(t *testing.T) {
	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--type", "null")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestInitWithNameAndInvalidType(t *testing.T) {
	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "null")

	assert.Contains(t, output, "Unknown project type provided")
	assert.Contains(t, err.Error(), "Unknown project type provided")
}

func TestInitWithNameAndType(t *testing.T) {
	pwd, nwd := setup(t)

	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "basic")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	if !fun.DirOrFileExists("clencli") {
		t.Errorf("CLENCLI directory missing")
	}
	if !fun.DirOrFileExists("clencli/readme.tmpl") {
		t.Errorf("CLENCLI readme.tmpl is missing")
	}
	if !fun.DirOrFileExists("clencli/readme.yaml") {
		t.Errorf("CLENCLI readme.yaml is missing")
	}

	teardown(pwd, nwd)
}

func TestInitWithNameAndCloudFormationType(t *testing.T) {
	pwd, nwd := setup(t)

	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "cloudformation")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	if !fun.DirOrFileExists("clencli") {
		t.Errorf("CLENCLI directory missing")
	}
	if !fun.DirOrFileExists("clencli/readme.tmpl") {
		t.Errorf("CLENCLI readme.tmpl is missing")
	}
	if !fun.DirOrFileExists("clencli/readme.yaml") {
		t.Errorf("CLENCLI readme.yaml is missing")
	}
	if !fun.DirOrFileExists("clencli/hld.tmpl") {
		t.Errorf("CLENCLI hld.tmpl is missing")
	}
	if !fun.DirOrFileExists("clencli/hld.yaml") {
		t.Errorf("CLENCLI hld.yaml is missing")
	}

	teardown(pwd, nwd)
}

func TestInitWithNameAndTerraformType(t *testing.T) {
	pwd, nwd := setup(t)

	rootCmd, initCmd := initRootAndChildCmd()
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := executeCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "terraform")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	if !fun.DirOrFileExists("clencli") {
		t.Errorf("CLENCLI directory missing")
	}
	if !fun.DirOrFileExists("clencli/readme.tmpl") {
		t.Errorf("CLENCLI readme.tmpl is missing")
	}
	if !fun.DirOrFileExists("clencli/readme.yaml") {
		t.Errorf("CLENCLI readme.yaml is missing")
	}
	if !fun.DirOrFileExists("clencli/hld.tmpl") {
		t.Errorf("CLENCLI hld.tmpl is missing")
	}
	if !fun.DirOrFileExists("clencli/hld.yaml") {
		t.Errorf("CLENCLI hld.yaml is missing")
	}

	teardown(pwd, nwd)
}
