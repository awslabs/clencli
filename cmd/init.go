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
	"errors"
	"fmt"
	"log"

	fun "github.com/awslabs/clencli/function"
	"github.com/spf13/cobra"
)

var validArgs = []string{"project"}

func preRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Please provide an argument")
	}

	// https://github.com/spf13/cobra/issues/655
	_, err := cmd.Flags().GetString("name")
	// flag accessed but not defined
	if err != nil {
		return errors.New("required flag name not set")
	}

	// ensure the project types
	if args[0] == "project" {
		t, err := cmd.Flags().GetString("type")
		// flag accessed but not defined
		if err != nil {
			return errors.New("When initializing a project")
		}
		if t == "" {
			return errors.New("Project type cannot be empty")
		}

		if t != "basic" && t != "cloudformation" && t != "terraform" {
			return fmt.Errorf("Unknown project type provided: %s", t)
		}

	}

	return nil
}

func getFlags(cmd *cobra.Command) (name string, typee string, structure string, onlyCustomizedStructure bool) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatal("required flag name not set")
	}

	typee, _ = cmd.Flags().GetString("type")
	structure, _ = cmd.Flags().GetString("structure")
	onlyCustomizedStructure, _ = cmd.Flags().GetBool("only-customized-structure")
	return name, typee, structure, onlyCustomizedStructure
}

func initBasicProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
	fun.Init(name)
	if !onlyCustomizedStructure {
		fun.InitBasic()
	}
	fun.InitCustomProjectLayout(typee, "default")
	fun.InitCustomProjectLayout(typee, structure)
	fun.UpdateReadMe()
}

func initCloudFormationProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
	fun.Init(name)
	if !onlyCustomizedStructure {
		fun.InitBasic()
		fun.InitHLD(name)
		fun.InitCloudFormation()
	}
	fun.InitCustomProjectLayout(typee, "default")
	fun.InitCustomProjectLayout(typee, structure)
	fun.UpdateReadMe()
}

func initTerraformProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
	fun.Init(name)
	if !onlyCustomizedStructure {
		fun.InitBasic()
		fun.InitHLD(name)
		fun.InitTerraform()
	}
	fun.InitCustomProjectLayout(typee, "default")
	fun.InitCustomProjectLayout(typee, structure)
	fun.UpdateReadMe()
}

func run(cmd *cobra.Command, args []string) error {
	name, typee, structure, onlyCustomizedStructure := getFlags(cmd)

	if args[0] == "project" {
		switch typee {
		case "basic":
			initBasicProject(name, typee, structure, onlyCustomizedStructure)
		case "cloudformation":
			initCloudFormationProject(name, typee, structure, onlyCustomizedStructure)
		case "terraform":
			initTerraformProject(name, typee, structure, onlyCustomizedStructure)

		default:
			return errors.New("Unknow project type")
		}
	} else {
		return errors.New("invalid argument")
	}

	return nil
}

// InitCmd command to initialize projects
func InitCmd() *cobra.Command {
	man := fun.GetManual("init")
	return &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		ValidArgs: validArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   preRun,
		RunE:      run,
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
