package tester

import (
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

// func TestInitWithInvalidArg(t *testing.T) {
// 	err := executeCommand(controller.InitCmd(), "init", "foo")
// 	assert.Contains(t, err.Error(), "invalid argument")
// }

// func TestInitProjectWithNoName(t *testing.T) {
// 	err := executeCommand(controller.InitCmd(), "init", "project")
// 	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
// }

// func TestInitProjectWithEmptyName(t *testing.T) {
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name")
// 	assert.Contains(t, err.Error(), "flag needs an argument: --name")
// }

// func TestInitProjectWithName(t *testing.T) {
// 	// pwd, nwd := tester.Setup(t)
// 	pPath := pwd + "/" + nwd + "/" + "foo"
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name", "foo")

// 	assert.Nil(t, err)
// 	assert.DirExists(t, pPath)
// 	assert.FileExists(t, pPath+"/.gitignore")
// 	assert.DirExists(t, pPath+"/clencli")

// 	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/readme.yaml")
// }

// func TestInitProjectWithNameAndEmptyType(t *testing.T) {
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name", "foo", "--type")
// 	assert.Contains(t, err.Error(), "flag needs an argument: --type")
// }

// func TestInitProjectWithNameAndWrongType(t *testing.T) {
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name", "foo", "--type", "nil")
// 	assert.Contains(t, err.Error(), "Unknown project type provided")
// }

// func TestInitProjectWithNameAndBasicType(t *testing.T) {
// 	// pwd, nwd := tester.Setup(t)
// 	pPath := pwd + "/" + nwd + "/" + "foo"
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name", "foo", "--type", "basic")

// 	assert.Nil(t, err)
// 	assert.DirExists(t, pPath)
// 	assert.FileExists(t, pPath+"/.gitignore")
// 	assert.DirExists(t, pPath+"/clencli")

// 	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/readme.yaml")
// }

// func TestInitProjectWithNameAndCloudType(t *testing.T) {
// 	pwd, nwd := tester.Setup(t)
// 	pPath := pwd + "/" + nwd + "/" + "foo"
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name", "foo", "--type", "cloud")

// 	assert.Nil(t, err)
// 	assert.DirExists(t, pPath)
// 	assert.FileExists(t, pPath+"/.gitignore")
// 	assert.DirExists(t, pPath+"/clencli")

// 	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/readme.yaml")

// 	assert.FileExists(t, pPath+"/clencli/hld.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/hld.yaml")
// }

// func TestInitProjectWithNameAndCloudFormationType(t *testing.T) {
// 	pwd, nwd := tester.Setup(t)
// 	pPath := pwd + "/" + nwd + "/" + "foo"
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name", "foo", "--type", "cloudformation")

// 	assert.Nil(t, err)
// 	assert.DirExists(t, pPath)
// 	assert.FileExists(t, pPath+"/.gitignore")
// 	assert.DirExists(t, pPath+"/clencli")

// 	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/readme.yaml")

// 	assert.FileExists(t, pPath+"/clencli/hld.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/hld.yaml")

// 	assert.DirExists(t, pPath+"/environments/dev")
// 	assert.DirExists(t, pPath+"/environments/prod")

// 	assert.FileExists(t, pPath+"/skeleton.yaml")
// 	assert.FileExists(t, pPath+"/skeleton.json")
// }

// func TestInitProjectWithNameAndTerraformType(t *testing.T) {
// 	pwd, nwd := tester.Setup(t)
// 	pPath := pwd + "/" + nwd + "/" + "foo"
// 	err := executeCommand(controller.InitCmd(), "init", "project", "--name", "foo", "--type", "terraform")

// 	assert.Nil(t, err)
// 	assert.DirExists(t, pPath)
// 	assert.FileExists(t, pPath+"/.gitignore")
// 	assert.DirExists(t, pPath+"/clencli")

// 	assert.FileExists(t, pPath+"/clencli/readme.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/readme.yaml")

// 	assert.FileExists(t, pPath+"/clencli/hld.tmpl")
// 	assert.FileExists(t, pPath+"/clencli/hld.yaml")

// 	assert.FileExists(t, pPath+"/main.tf")
// 	assert.FileExists(t, pPath+"/variables.tf")
// 	assert.FileExists(t, pPath+"/outputs.tf")

// 	assert.DirExists(t, pPath+"/environments")
// 	assert.FileExists(t, pPath+"/environments/dev.tf")
// 	assert.FileExists(t, pPath+"/environments/prod.tf")

// 	assert.FileExists(t, pPath+"/Makefile")
// 	assert.FileExists(t, pPath+"/LICENSE")
// }
