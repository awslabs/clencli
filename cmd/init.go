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

	fun "github.com/awslabs/clencli/function"
	"github.com/spf13/cobra"
)

var validArgs = []string{"project"}

func preRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide an argument")
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetString("name")
	t, _ := cmd.Flags().GetString("type")
	s, _ := cmd.Flags().GetString("structure")
	o, _ := cmd.Flags().GetBool("only-customized-structure")

	switch t {
	case "basic":
		fun.Init(n)
		if !o {
			fun.InitBasic()
		}
		fun.InitCustomProjectLayout(t, "default")
		fun.InitCustomProjectLayout(t, s)
	case "cloudformation":
		fun.Init(n)
		if !o {
			fun.InitBasic()
			fun.InitHLD(n)
			fun.InitCloudFormation()
		}
		fun.InitCustomProjectLayout("basic", "default")
		fun.InitCustomProjectLayout(t, s)
	case "terraform":
		fun.Init(n)
		if !o {
			fun.InitBasic()
			fun.InitHLD(n)
			fun.InitTerraform()
		}
		fun.InitCustomProjectLayout("basic", "default")
		fun.InitCustomProjectLayout(t, s)
	default:
		log.Fatal("Unknown project type")
	}

	// Update clencli/*.yaml based on clencli's config
	fun.UpdateReadMe()
}

// InitCmd command to initialize projects
func InitCmd() *cobra.Command {
	man := fun.GetManual("init")
	return &cobra.Command{
		// Use:       "init ",
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		ValidArgs: validArgs,
		Args:      cobra.OnlyValidArgs,
		PreRun:    preRun,
		Run:       run,
	}
}

// initCmd represents the init command
var initCmd = InitCmd()

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("name", "n", "", "The project name.")
	initCmd.Flags().StringP("type", "t", "basic", "The project type.")
	initCmd.Flags().StringP("structure", "s", "default", "The project structure name defined on main config.")
	initCmd.Flags().BoolP("only-customized-structure", "o", false, "Only customized structure to be used when initializing the project")

	initCmd.MarkFlagRequired("name")

}
