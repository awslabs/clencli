package tests

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/helper"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRenderCmd(t *testing.T) {
	tests := map[string]struct {
		args []string
		out  string
		err  string
	}{
		// argument
		"no arg":                      {args: []string{"render"}, out: "", err: "this command requires one argument"},
		"empty arg":                   {args: []string{"render", ""}, out: "", err: "invalid argument"},
		"wrong arg":                   {args: []string{"render", "foo"}, out: "", err: "invalid argument"},
		"unknown flag":                {args: []string{"render", "--foo"}, out: "", err: "unknown flag: --foo"},
		"missing database":            {args: []string{"render", "template"}, out: "", err: "missing database at clencli/readme.yaml"},
		"flag needs an argument name": {args: []string{"render", "template", "--name"}, out: "", err: "flag needs an argument: --name"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := executeCommand(t, controller.RenderCmd(), tc.args)
			assert.Contains(t, out, tc.out)
			assert.Contains(t, err.Error(), tc.err)
		})
	}
}

// /* README */

func initProject(t *testing.T, pType string) {
	args := []string{"init", "project", "--project-name", "foo", "--project-type", pType}
	wd, out, err := executeCommandGetWorkingDirectory(t, controller.InitCmd(), args)

	assert.Nil(t, err)
	assert.Contains(t, out, "was successfully initialized")

	if err := os.Chdir(helper.BuildPath(wd + "/" + t.Name() + "/" + "foo")); err != nil {
		logrus.Fatal("unable to change current working directory")
	}

}

func TestRenderDefault(t *testing.T) {
	initProject(t, "basic")

	args := []string{"render", "template"}
	out, err := executeCommandOnly(t, controller.RenderCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "Template readme.tmpl rendered as README.md")
}

func TestRenderReadme(t *testing.T) {
	initProject(t, "basic")

	args := []string{"render", "template", "--name", "readme"}
	out, err := executeCommandOnly(t, controller.RenderCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "Template readme.tmpl rendered as README.md")
}

func TestRenderHLD(t *testing.T) {
	initProject(t, "cloud")

	args := []string{"render", "template", "--name", "hld"}
	out, err := executeCommandOnly(t, controller.RenderCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "Template hld.tmpl rendered as HLD.md")
}

// func assertBasicProject(t *testing.T, err error) (string, string) {
// 	sep := string(os.PathSeparator)
// 	dir := t.Name() + sep + "foo"

// 	assert.Nil(t, err)
// 	assert.DirExists(t, dir)
// 	assert.FileExists(t, dir+sep+".gitignore")
// 	assert.DirExists(t, dir+sep+"clencli")

// 	assert.FileExists(t, dir+sep+"clencli"+sep+"readme.tmpl")
// 	assert.FileExists(t, dir+sep+"clencli"+sep+"readme.yaml")

// 	return dir, sep
// }

// func assertCloudProject(t *testing.T, err error) (string, string) {
// 	dir, sep := assertBasicProject(t, err)
// 	assert.FileExists(t, dir+sep+"clencli"+sep+"hld.tmpl")
// 	assert.FileExists(t, dir+sep+"clencli"+sep+"hld.yaml")
// 	return dir, sep
// }

// func TestrenderBasicProjectWithNameAndType(t *testing.T) {
// 	args := []string{"render", "project", "--project-name", "foo", "--project-type", "basic"}
// 	out, err := executeCommand(t, controller.renderCmd(), args)
// 	assert.Nil(t, err)
// 	assert.Contains(t, out, "was successfully renderialized as a basic project")
// 	assertBasicProject(t, err)
// }

// /* PROJECT: CLOUD */

// func TestrenderCloudProjectWithName(t *testing.T) {
// 	args := []string{"render", "project", "--project-name", "foo", "--project-type", "cloud"}
// 	out, err := executeCommand(t, controller.renderCmd(), args)
// 	assert.Contains(t, out, "was successfully renderialized as a cloud project")
// 	assertCloudProject(t, err)
// }

// /* PROJECT: CLOUDFORMATION */

// func TestrenderProjectWithNameAndCloudFormationType(t *testing.T) {
// 	args := []string{"render", "project", "--project-name", "foo", "--project-type", "cloudformation"}
// 	out, err := executeCommand(t, controller.renderCmd(), args)
// 	assert.Contains(t, out, "was successfully renderialized as a cloudformation project")

// 	dir, sep := assertCloudProject(t, err)
// 	assert.DirExists(t, dir+sep+"environments"+sep+"dev")
// 	assert.DirExists(t, dir+sep+"environments"+sep+"prod")

// 	assert.FileExists(t, dir+sep+"skeleton.yaml")
// 	assert.FileExists(t, dir+sep+"skeleton.json")
// }

// /* PROJECT: TERRAFORM */

// func TestrenderProjectWithNameAndTerraformType(t *testing.T) {
// 	args := []string{"render", "project", "--project-name", "foo", "--project-type", "terraform"}
// 	out, err := executeCommand(t, controller.renderCmd(), args)
// 	assert.Contains(t, out, "was successfully renderialized as a terraform project")

// 	dir, sep := assertCloudProject(t, err)

// 	assert.FileExists(t, dir+sep+"Makefile")
// 	assert.FileExists(t, dir+sep+"LICENSE")

// 	assert.DirExists(t, dir+sep+"environments")
// 	assert.FileExists(t, dir+sep+"environments"+sep+"dev.tf")
// 	assert.FileExists(t, dir+sep+"environments"+sep+"prod.tf")

// 	assert.FileExists(t, dir+sep+"main.tf")
// 	assert.FileExists(t, dir+sep+"variables.tf")
// 	assert.FileExists(t, dir+sep+"outputs.tf")
// }
