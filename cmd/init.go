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

// Package cmd contains all Cobra commands
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/awslabs/clencli/box"
	"github.com/awslabs/clencli/function"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	// Use:       "init ",
	Use: `init project --name <value> 
	[ --type [basic|cloudformation|terraform] ]`,
	Short:     "Initialize a project",
	Long:      "Initialize a project with code structure",
	ValidArgs: []string{"project"},
	Args:      cobra.OnlyValidArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide an argument")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetString("name")
		t, _ := cmd.Flags().GetString("type")

		switch t {
		case "basic":
			function.Init(n)
		case "cloudformation":
			function.Init(n)
			function.InitHLD(n)
			initCloudFormation(cmd, n)
		case "terraform":
			function.Init(n)
			function.InitHLD(n)
			initTerraform(cmd, n)
		default:
			log.Fatal("Unknown project type")
		}

		// Update clencli/*.yaml based on clencli's config
		function.UpdateReadMe()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")

	initCmd.MarkFlagRequired("name")

}

func initCloudFormation(cmd *cobra.Command, name string) {

	function.CreateDir("cloudformation/environments")
	function.CreateDir("cloudformation/environments/dev")
	function.CreateDir("cloudformation/environments/prod")
	function.CreateDir("cloudformation/templates")

	initCFStack := "cloudformation/templates/stack.yaml"

	blobCFStack, _ := box.Get("/init/type/clouformation/templates/stack.yaml")
	function.WriteFile(initCFStack, blobCFStack)

	initCFNested := "cloudformation/templates/nested.yaml"
	blobCFNested, _ := box.Get("/init/type/clouformation/templates/nested.yaml")
	function.WriteFile(initCFNested, blobCFNested)

}

func initTerraform(cmd *cobra.Command, name string) {

	initMakefile := "Makefile"
	blobMakefile, _ := box.Get("/init/type/terraform/Makefile")
	function.WriteFile(initMakefile, blobMakefile)

	initLicense := "LICENSE"
	blobLicense, _ := box.Get("/init/type/terraform/LICENSE")
	function.WriteFile(initLicense, blobLicense)

	function.CreateDir("environments")

	initDevEnvironment := "environments/dev.tf"
	blobDevEnvironment, _ := box.Get("/init/type/terraform/environments/dev.tf")
	function.WriteFile(initDevEnvironment, blobDevEnvironment)

	initProdEnvironment := "environments/prod.tf"
	blobProdEnvironment, _ := box.Get("/init/type/terraform/environments/prod.tf")
	function.WriteFile(initProdEnvironment, blobProdEnvironment)

	initMainTF := "main.tf"
	blobMainTF, _ := box.Get("/init/type/terraform/main.tf")
	function.WriteFile(initMainTF, blobMainTF)

	initVariablesTF := "variables.tf"
	blobVariablesTF, _ := box.Get("/init/type/terraform/variables.tf")
	function.WriteFile(initVariablesTF, blobVariablesTF)

	initOutputsTF := "outputs.tf"
	blobOutputsTF, _ := box.Get("/init/type/terraform/outputs.tf")
	function.WriteFile(initOutputsTF, blobOutputsTF)
}
