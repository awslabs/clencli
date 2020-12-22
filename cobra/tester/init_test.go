package tester

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func TestInitWithNoArgs(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init")
	assert.Contains(t, err.Error(), "this command requires one argument")
}

func TestInitWithEmptyArgs(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestInitWithInvalidArg(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "foo")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestInitProjectWithNoName(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project")
	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
}

func TestInitProjectWithEmptyName(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name")
	assert.Contains(t, err.Error(), "flag needs an argument: --name")
}

func assertBasicProject(t *testing.T, err error) (string, string) {
	sep := string(os.PathSeparator)
	dir := t.Name() + sep + "foo"

	assert.Nil(t, err)
	assert.DirExists(t, dir)
	assert.FileExists(t, dir+sep+".gitignore")
	assert.DirExists(t, dir+sep+"clencli")

	assert.FileExists(t, dir+sep+"clencli"+sep+"readme.tmpl")
	assert.FileExists(t, dir+sep+"clencli"+sep+"readme.yaml")

	return dir, sep
}

func assertCloudProject(t *testing.T, err error) (string, string) {
	dir, sep := assertBasicProject(t, err)
	assert.FileExists(t, dir+sep+"clencli"+sep+"hld.tmpl")
	assert.FileExists(t, dir+sep+"clencli"+sep+"hld.yaml")
	return dir, sep
}

func TestInitProjectWithName(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo")
	assertBasicProject(t, err)

}

func TestInitProjectWithNameAndEmptyType(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type")
	assert.Contains(t, err.Error(), "flag needs an argument: --type")
}

func TestInitProjectWithNameAndWrongType(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "nil")
	assert.Contains(t, err.Error(), "unknown project type provided")
}

func TestInitProjectWithNameAndBasicType(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "basic")
	assertBasicProject(t, err)
}

func TestInitProjectWithNameAndCloudType(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "cloud")
	assertCloudProject(t, err)
}

func TestInitProjectWithNameAndCloudFormationType(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "cloudformation")
	dir, sep := assertCloudProject(t, err)

	assert.DirExists(t, dir+sep+"environments"+sep+"dev")
	assert.DirExists(t, dir+sep+"environments"+sep+"prod")

	assert.FileExists(t, dir+sep+"skeleton.yaml")
	assert.FileExists(t, dir+sep+"skeleton.json")
}

func TestInitProjectWithNameAndTerraformType(t *testing.T) {
	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "terraform")
	dir, sep := assertCloudProject(t, err)

	assert.FileExists(t, dir+sep+"main.tf")
	assert.FileExists(t, dir+sep+"variables.tf")
	assert.FileExists(t, dir+sep+"outputs.tf")

	assert.DirExists(t, dir+sep+"environments")
	assert.FileExists(t, dir+sep+"environments"+sep+"dev.tf")
	assert.FileExists(t, dir+sep+"environments"+sep+"prod.tf")

	assert.FileExists(t, dir+sep+"Makefile")
	assert.FileExists(t, dir+sep+"LICENSE")
}
