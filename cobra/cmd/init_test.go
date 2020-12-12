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

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/tester"
	"github.com/spf13/cobra"
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

func command() *cobra.Command {
	return tester.InitRootCmd(controller.RootCmd(), GetInitCmd())
}

func TestInitWithNoArgAndNoFlags(t *testing.T) {
	err := tester.ExecuteCommand(command(), "init")
	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
}

func TestInitWithInvalidArgAndNoFlags(t *testing.T) {
	err := tester.ExecuteCommand(command(), "init", "null")
	assert.Contains(t, err.Error(), "invalid argument")
}

func TestInitWithValidArgAndNoFlags(t *testing.T) {
	err := tester.ExecuteCommand(command(), "init", "project")
	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
}

func TestInitWithEmptyName(t *testing.T) {
	err := tester.ExecuteCommand(command(), "init", "project", "--name")
	assert.Contains(t, err.Error(), "flag needs an argument")
}

func TestInitWithValidTypeOnly(t *testing.T) {
	err := tester.ExecuteCommand(command(), "init", "project", "--type", "basic")
	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
}

func TestInitWithInvalidTypeOnly(t *testing.T) {
	err := tester.ExecuteCommand(command(), "init", "project", "--type", "null")
	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
}

func TestInitWithNameAndInvalidType(t *testing.T) {
	err := tester.ExecuteCommand(command(), "init", "project", "--name", "generated-project", "--type", "null")
	assert.Contains(t, err.Error(), "Unknown project type provided")
}

func TestInitWithNameOnly(t *testing.T) {
	pwd, nwd := tester.Setup(t)
	fmt.Println(pwd)
	fmt.Println(nwd)
	err := tester.ExecuteCommand(command(), "init", "project", "--name", "generated-project")
	assert.Nil(t, err)
	// exec.Command("pwd")
	cmd := exec.Command("ls -ltha")
	cmd.Run()

	assert.DirExists(t, "TestInitWithNameOnly/")
	tester.Teardown(pwd, nwd)
}

// func TestInitWithNameAndType(t *testing.T) {
// 	pwd, nwd := tester.Setup(t)

// 	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
// 	err := tester.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "basic")

// 	if err != nil {
// 		t.Errorf("Project wasn't able to initialize: %v", output)
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	if !helper.DirOrFileExists("clencli") {
// 		t.Errorf("CLENCLI directory missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/readme.tmpl") {
// 		t.Errorf("CLENCLI readme.tmpl is missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/readme.yaml") {
// 		t.Errorf("CLENCLI readme.yaml is missing")
// 	}

// 	tester.Teardown(pwd, nwd)
// }

// func TestInitWithNameAndCloudFormationType(t *testing.T) {
// 	pwd, nwd := tester.Setup(t)

// 	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
// 	err := tester.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "cloudformation")

// 	if err != nil {
// 		t.Errorf("Project wasn't able to initialize: %v", output)
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	if !helper.DirOrFileExists("clencli") {
// 		t.Errorf("CLENCLI directory missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/readme.tmpl") {
// 		t.Errorf("CLENCLI readme.tmpl is missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/readme.yaml") {
// 		t.Errorf("CLENCLI readme.yaml is missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/hld.tmpl") {
// 		t.Errorf("CLENCLI hld.tmpl is missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/hld.yaml") {
// 		t.Errorf("CLENCLI hld.yaml is missing")
// 	}

// 	tester.Teardown(pwd, nwd)
// }

// func TestInitWithNameAndTerraformType(t *testing.T) {
// 	pwd, nwd := tester.Setup(t)

// 	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
// 	err := tester.ExecuteCommand(rootCmd, "init", "project", "--name", "generated-project", "--type", "terraform")

// 	if err != nil {
// 		t.Errorf("Project wasn't able to initialize: %v", output)
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	if !helper.DirOrFileExists("clencli") {
// 		t.Errorf("CLENCLI directory missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/readme.tmpl") {
// 		t.Errorf("CLENCLI readme.tmpl is missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/readme.yaml") {
// 		t.Errorf("CLENCLI readme.yaml is missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/hld.tmpl") {
// 		t.Errorf("CLENCLI hld.tmpl is missing")
// 	}
// 	if !helper.DirOrFileExists("clencli/hld.yaml") {
// 		t.Errorf("CLENCLI hld.yaml is missing")
// 	}

// 	tester.Teardown(pwd, nwd)
// }
