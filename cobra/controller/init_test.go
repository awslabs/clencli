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
	"os"
	"testing"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	pwd, nwd := cau.SetupAll()
	// setupAll()
	code := m.Run()
	// comment the line below if you want to keep the test results
	cau.TeardownAll(pwd, nwd)
	os.Exit(code)
}

func TestInitWithNoArgAndNoFlags(t *testing.T) {
	rootCmd, _ := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	output, err := cau.ExecuteCommand(rootCmd, "init")

	assert.Contains(t, output, "one the following arguments are required")
	assert.Contains(t, err.Error(), "one the following arguments are required")
}

func TestInitWithInvalidArgAndNoFlags(t *testing.T) {
	rootCmd, _ := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	output, err := cau.ExecuteCommand(rootCmd, "init", "null")

	assert.Contains(t, output, "invalid argument")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestInitWithValidArgAndNoFlags(t *testing.T) {

	rootCmd, _ := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	output, err := cau.ExecuteCommand(rootCmd, "init", "project")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestInitWithEmptyName(t *testing.T) {

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--name")

	assert.Contains(t, output, "flag needs an argument")
	assert.Contains(t, err.Error(), "flag needs an argument")
}

func TestInitWithNameOnly(t *testing.T) {
	pwd, nwd := cau.Setup(t)

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	cau.Teardown(pwd, nwd)
}

func TestInitWithValidTypeOnly(t *testing.T) {
	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--type", "basic")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestInitWithInvalidTypeOnly(t *testing.T) {
	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--type", "null")

	assert.Contains(t, output, "required flag name not set")
	assert.Contains(t, err.Error(), "required flag name not set")
}

func TestInitWithNameAndInvalidType(t *testing.T) {
	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "null")

	assert.Contains(t, output, "Unknown project type provided")
	assert.Contains(t, err.Error(), "Unknown project type provided")
}

func TestInitWithNameAndType(t *testing.T) {
	pwd, nwd := cau.Setup(t)

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "basic")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	if !cau.DirOrFileExists("clencli") {
		t.Errorf("CLENCLI directory missing")
	}
	if !cau.DirOrFileExists("clencli/readme.tmpl") {
		t.Errorf("CLENCLI readme.tmpl is missing")
	}
	if !cau.DirOrFileExists("clencli/readme.yaml") {
		t.Errorf("CLENCLI readme.yaml is missing")
	}

	cau.Teardown(pwd, nwd)
}

func TestInitWithNameAndCloudFormationType(t *testing.T) {
	pwd, nwd := cau.Setup(t)

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "cloudformation")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	if !cau.DirOrFileExists("clencli") {
		t.Errorf("CLENCLI directory missing")
	}
	if !cau.DirOrFileExists("clencli/readme.tmpl") {
		t.Errorf("CLENCLI readme.tmpl is missing")
	}
	if !cau.DirOrFileExists("clencli/readme.yaml") {
		t.Errorf("CLENCLI readme.yaml is missing")
	}
	if !cau.DirOrFileExists("clencli/hld.tmpl") {
		t.Errorf("CLENCLI hld.tmpl is missing")
	}
	if !cau.DirOrFileExists("clencli/hld.yaml") {
		t.Errorf("CLENCLI hld.yaml is missing")
	}

	cau.Teardown(pwd, nwd)
}

func TestInitWithNameAndTerraformType(t *testing.T) {
	pwd, nwd := cau.Setup(t)

	rootCmd, initCmd := cau.InitRootAndChildCmd(RootCmd(), InitCmd())
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	output, err := cau.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "terraform")

	if err != nil {
		t.Errorf("Project wasn't able to initialize: %v", output)
		t.Errorf("Unexpected error: %v", err)
	}

	if !cau.DirOrFileExists("clencli") {
		t.Errorf("CLENCLI directory missing")
	}
	if !cau.DirOrFileExists("clencli/readme.tmpl") {
		t.Errorf("CLENCLI readme.tmpl is missing")
	}
	if !cau.DirOrFileExists("clencli/readme.yaml") {
		t.Errorf("CLENCLI readme.yaml is missing")
	}
	if !cau.DirOrFileExists("clencli/hld.tmpl") {
		t.Errorf("CLENCLI hld.tmpl is missing")
	}
	if !cau.DirOrFileExists("clencli/hld.yaml") {
		t.Errorf("CLENCLI hld.yaml is missing")
	}

	cau.Teardown(pwd, nwd)
}
