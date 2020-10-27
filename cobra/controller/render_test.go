/*
Copyright © 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"testing"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/stretchr/testify/assert"
)

func TestRenderWithNoArgAndNoFlags(t *testing.T) {
	rootCmd, _ := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
	output, err := cau.ExecuteCommand(rootCmd, "render")

	assert.Contains(t, output, "one the following arguments are required")
	assert.Contains(t, err.Error(), "one the following arguments are required")
}

func TestRenderWithInvalidArgAndNoFlags(t *testing.T) {
	rootCmd, _ := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
	output, err := cau.ExecuteCommand(rootCmd, "render", "null")

	assert.Contains(t, output, "invalid argument")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestRenderWithValidArgAndNoFlags(t *testing.T) {
	rootCmd, _ := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
	output, err := cau.ExecuteCommand(rootCmd, "render", "template")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestRenderWithNameOnly(t *testing.T) {
	rootCmd, renderCmd := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
	renderCmd.Flags().StringP("name", "n", "readme", "Template name to be rendered")
	output, err := cau.ExecuteCommand(rootCmd, "render", "template")

	assert.Contains(t, output, "Missing database at clencli/readme.yaml")
	assert.Contains(t, err.Error(), "Missing database at clencli/readme.yaml")
}

func TestRenderWithInitBasicProject(t *testing.T) {

	// init a basic project
	pwd, nwd := cau.Setup(t)

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "generated-project", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	rootCmd, renderCmd := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
	renderCmd.Flags().StringP("name", "n", "readme", "Template name to be rendered")
	output, err = cau.ExecuteCommand(rootCmd, "render", "template")

	// Ensure project was initialized correctly
	assert.Equal(t, output, "")
	assert.Equal(t, err, nil)

	if !cau.FileExists("clencli/readme.tmpl") {
		t.Error("clencli/readme.tmpl not found, project initialization failed")
	}

	if !cau.FileExists("clencli/readme.yaml") {
		t.Error("clencli/readme.yaml not found, project initialization failed")
	}

	if !cau.FileExists("README.md") {
		t.Error("README.md not found, rendering failed")
	}

	cau.Teardown(pwd, nwd)
}

func TestRenderWithInitTerraformProject(t *testing.T) {

	// init a basic project
	pwd, nwd := cau.Setup(t)

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "generated-project", "The project name.")
	initCmd.Flags().StringP("type", "t", "terraform", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	rootCmd, renderCmd := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
	renderCmd.Flags().StringP("name", "n", "hld", "Template name to be rendered")
	output, err = cau.ExecuteCommand(rootCmd, "render", "template")

	// Ensure project was initialized correctly
	assert.Equal(t, output, "")
	assert.Equal(t, err, nil)

	if !cau.FileExists("clencli/hld.tmpl") {
		t.Error("clencli/hld.tmpl not found, project initialization failed")
	}

	if !cau.FileExists("clencli/hld.yaml") {
		t.Error("clencli/hld.yaml not found, project initialization failed")
	}

	if !cau.FileExists("HLD.md") {
		t.Error("HLD.md not found, rendering failed")
	}

	cau.Teardown(pwd, nwd)
}

func TestRenderWithInitCloudFormationProject(t *testing.T) {

	// init a basic project
	pwd, nwd := cau.Setup(t)

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "generated-project", "The project name.")
	initCmd.Flags().StringP("type", "t", "cloudformation", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	rootCmd, renderCmd := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
	renderCmd.Flags().StringP("name", "n", "hld", "Template name to be rendered")
	output, err = cau.ExecuteCommand(rootCmd, "render", "template")

	// Ensure project was initialized correctly
	assert.Equal(t, output, "")
	assert.Equal(t, err, nil)

	if !cau.FileExists("clencli/hld.tmpl") {
		t.Error("clencli/hld.tmpl not found, project initialization failed")
	}

	if !cau.FileExists("clencli/hld.yaml") {
		t.Error("clencli/hld.yaml not found, project initialization failed")
	}

	if !cau.FileExists("HLD.md") {
		t.Error("HLD.md not found, rendering failed")
	}

	cau.Teardown(pwd, nwd)
}

// TODO: find a way to use secret keys on Github Actions

// func TestRenderWithUpdatedTheme(t *testing.T) {

// 	pwd, nwd := cau.Setup(t)

// 	// init a basic project
// 	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
// 	initCmd.Flags().StringP("name", "n", "generated-project", "The project name.")
// 	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
// 	output, err := cau.ExecuteCommand(rootCmd, "init", "project")

// 	if err != nil {
// 		t.Errorf("Project wasn't able to initialize: %v", output)
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	readme, err := cau.GetLocalReadMeConfig()
// 	if err != nil {
// 		t.Errorf("Unexpected error \n%v", err)
// 	}

// 	// if theme is set, URL must change
// 	readme.Logo.Theme = "dogs"
// 	err = cau.MarshallAndSaveReadMe(readme)
// 	if err != nil {
// 		t.Errorf("Unexpected error \n%v", err)
// 	}

// 	rootCmd, renderCmd := cau.InitRootAndChildCmd(RootCmd(), RenderCmd())
// 	renderCmd.Flags().StringP("name", "n", "readme", "Template name to be rendered")
// 	output, err = cau.ExecuteCommand(rootCmd, "render", "template")

// 	// Ensure project was initialized correctly
// 	assert.Equal(t, output, "")
// 	assert.Equal(t, err, nil)

// 	readme, err = cau.GetLocalReadMeConfig()
// 	if err != nil {
// 		t.Errorf("Unexpected error \n%v", err)
// 	}

// 	assert.NotEqual(t, readme.Logo.URL, "")

// 	cau.Teardown(pwd, nwd)
// }
