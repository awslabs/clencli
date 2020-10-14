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
		s, _ := cmd.Flags().GetString("structure")
		o, _ := cmd.Flags().GetBool("only-customized-structure")

		switch t {
		case "basic":
			function.Init(n)
			if !o {
				function.InitBasic()
			}
			function.InitCustomProjectLayout(t, "default")
			function.InitCustomProjectLayout(t, s)
		case "cloudformation":
			function.Init(n)
			if !o {
				function.InitBasic()
				function.InitHLD(n)
				function.InitCloudFormation()
			}
			function.InitCustomProjectLayout("basic", "default")
			function.InitCustomProjectLayout(t, s)
		case "terraform":
			function.Init(n)
			if !o {
				function.InitBasic()
				function.InitHLD(n)
				function.InitTerraform()
			}
			function.InitCustomProjectLayout("basic", "default")
			function.InitCustomProjectLayout(t, s)
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
	initCmd.Flags().StringP("structure", "s", "default", "The project structure name defined on main config.")
	initCmd.Flags().BoolP("only-customized-structure", "o", false, "Only customized structure to be used when initializing the project")

	initCmd.MarkFlagRequired("name")

}
