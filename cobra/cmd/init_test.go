package cmd

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/tester"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	tester.SetupAll()
	// pwd, nwd := tester.SetupAll()
	code := m.Run()
	// comment the line below if you want to keep the test results
	// tester.TeardownAll(pwd, nwd)
	os.Exit(code)
}

func command(args ...string) error {
	cmd := controller.InitCmd()
	cmd.SetArgs(args)
	err := cmd.Execute()
	return err
}

func TestInitEmpty(t *testing.T) {
	err := command("")
	// assert.Contains(t, out, "Usage")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestInitWithInvalidArg(t *testing.T) {
	err := command("foo")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestInitProjectWithNoName(t *testing.T) {
	err := command("project")
	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
}

func TestInitProjectWithEmptyName(t *testing.T) {
	err := command("project", "--name")
	assert.Contains(t, err.Error(), "flag needs an argument: --name")
}

func TestInitProjectWithName(t *testing.T) {
	pwd, nwd := tester.Setup(t)
	pPath := pwd + "/" + nwd + "/" + "foo"
	err := command("project", "--name", "foo")

	assert.Nil(t, err)
	assert.DirExists(t, pPath)
	assert.FileExists(t, pPath+"/.gitignore")
	assert.DirExists(t, pPath+"/clencli")

	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
	assert.FileExists(t, pPath+"/clencli/readme.yaml")
}

func TestInitProjectWithNameAndEmptyType(t *testing.T) {
	err := command("project", "--name", "foo", "--type")
	assert.Contains(t, err.Error(), "flag needs an argument: --type")
}

func TestInitProjectWithNameAndWrongType(t *testing.T) {
	err := command("project", "--name", "foo", "--type", "nil")
	assert.Contains(t, err.Error(), "Unknown project type provided")
}

func TestInitProjectWithNameAndBasicType(t *testing.T) {
	pwd, nwd := tester.Setup(t)
	pPath := pwd + "/" + nwd + "/" + "foo"
	err := command("project", "--name", "foo", "--type", "basic")

	assert.Nil(t, err)
	assert.DirExists(t, pPath)
	assert.FileExists(t, pPath+"/.gitignore")
	assert.DirExists(t, pPath+"/clencli")

	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
	assert.FileExists(t, pPath+"/clencli/readme.yaml")
}

func TestInitProjectWithNameAndCloudType(t *testing.T) {
	pwd, nwd := tester.Setup(t)
	pPath := pwd + "/" + nwd + "/" + "foo"
	err := command("project", "--name", "foo", "--type", "cloud")

	assert.Nil(t, err)
	assert.DirExists(t, pPath)
	assert.FileExists(t, pPath+"/.gitignore")
	assert.DirExists(t, pPath+"/clencli")

	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
	assert.FileExists(t, pPath+"/clencli/readme.yaml")

	assert.FileExists(t, pPath+"/clencli/hld.tmpl")
	assert.FileExists(t, pPath+"/clencli/hld.yaml")
}

func TestInitProjectWithNameAndCloudFormationType(t *testing.T) {
	pwd, nwd := tester.Setup(t)
	pPath := pwd + "/" + nwd + "/" + "foo"
	err := command("project", "--name", "foo", "--type", "cloudformation")

	assert.Nil(t, err)
	assert.DirExists(t, pPath)
	assert.FileExists(t, pPath+"/.gitignore")
	assert.DirExists(t, pPath+"/clencli")

	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
	assert.FileExists(t, pPath+"/clencli/readme.yaml")

	assert.FileExists(t, pPath+"/clencli/hld.tmpl")
	assert.FileExists(t, pPath+"/clencli/hld.yaml")

	assert.DirExists(t, pPath+"/environments/dev")
	assert.DirExists(t, pPath+"/environments/prod")

	assert.FileExists(t, pPath+"/skeleton.yaml")
	assert.FileExists(t, pPath+"/skeleton.json")
}

func TestInitProjectWithNameAndTerraformType(t *testing.T) {
	pwd, nwd := tester.Setup(t)
	pPath := pwd + "/" + nwd + "/" + "foo"
	err := command("project", "--name", "foo", "--type", "terraform")

	assert.Nil(t, err)
	assert.DirExists(t, pPath)
	assert.FileExists(t, pPath+"/.gitignore")
	assert.DirExists(t, pPath+"/clencli")

	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
	assert.FileExists(t, pPath+"/clencli/readme.yaml")

	assert.FileExists(t, pPath+"/clencli/hld.tmpl")
	assert.FileExists(t, pPath+"/clencli/hld.yaml")

	assert.FileExists(t, pPath+"/main.tf")
	assert.FileExists(t, pPath+"/variables.tf")
	assert.FileExists(t, pPath+"/outputs.tf")

	assert.DirExists(t, pPath+"/environments")
	assert.FileExists(t, pPath+"/environments/dev.tf")
	assert.FileExists(t, pPath+"/environments/prod.tf")

	assert.FileExists(t, pPath+"/Makefile")
	assert.FileExists(t, pPath+"/LICENSE")
}
