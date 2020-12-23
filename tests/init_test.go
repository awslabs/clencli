package tests

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func TestInitCmd(t *testing.T) {
	tests := map[string]struct {
		args []string
		out  string
		err  string
	}{
		// argument
		"empty":     {args: []string{"init"}, out: "", err: "this command requires one argument"},
		"empty arg": {args: []string{"init", ""}, out: "", err: "invalid argument"},
		"wrong arg": {args: []string{"init", "foo"}, out: "", err: "invalid argument"},

		// flags
		"wrong flag": {args: []string{"init", "project", "--foo"}, out: "", err: "unknown flag: --foo"},

		// # projects
		"no project name": {args: []string{"init", "project"}, out: "", err: "--project-name must be defined"},

		// ## --project-name
		"emtpy project name": {args: []string{"init", "project", "--project-name"}, out: "", err: "flag needs an argument"},

		// ## --project-type
		"empty project type":   {args: []string{"init", "project", "--project-type"}, out: "", err: "flag needs an argument"},
		"invalid project type": {args: []string{"init", "project", "--project-type", "nil"}, out: "", err: "--project-name must be defined"},

		// ## --project-name && --project-type
		"with project name and empty project type":   {args: []string{"init", "project", "--project-name", "foo", "--project-type"}, out: "", err: "flag needs an argument"},
		"with project name and invalid project type": {args: []string{"init", "project", "--project-name", "foo", "--project-type", "bar"}, out: "", err: "unknow project type"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := executeCommand(t, controller.InitCmd(), tc.args)
			assert.Contains(t, out, tc.out)
			assert.Contains(t, err.Error(), tc.err)
		})
	}
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
	args := []string{"init", "project", "--project-name", "foo"}
	out, err := executeCommand(t, controller.InitCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "initialized with success")
	assertBasicProject(t, err)
}

// func TestInitProjectWithNameAndBasicType(t *testing.T) {
// 	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "basic")
// 	assertBasicProject(t, err)
// }

// func TestInitProjectWithNameAndCloudType(t *testing.T) {
// 	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "cloud")
// 	assertCloudProject(t, err)
// }

// func TestInitProjectWithNameAndCloudFormationType(t *testing.T) {
// 	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "cloudformation")
// 	dir, sep := assertCloudProject(t, err)

// 	assert.DirExists(t, dir+sep+"environments"+sep+"dev")
// 	assert.DirExists(t, dir+sep+"environments"+sep+"prod")

// 	assert.FileExists(t, dir+sep+"skeleton.yaml")
// 	assert.FileExists(t, dir+sep+"skeleton.json")
// }

// func TestInitProjectWithNameAndTerraformType(t *testing.T) {
// 	err := executeCommand(t, controller.InitCmd(), "init", "project", "--name", "foo", "--type", "terraform")
// 	dir, sep := assertCloudProject(t, err)

// 	assert.FileExists(t, dir+sep+"main.tf")
// 	assert.FileExists(t, dir+sep+"variables.tf")
// 	assert.FileExists(t, dir+sep+"outputs.tf")

// 	assert.DirExists(t, dir+sep+"environments")
// 	assert.FileExists(t, dir+sep+"environments"+sep+"dev.tf")
// 	assert.FileExists(t, dir+sep+"environments"+sep+"prod.tf")

// 	assert.FileExists(t, dir+sep+"Makefile")
// 	assert.FileExists(t, dir+sep+"LICENSE")
// }
