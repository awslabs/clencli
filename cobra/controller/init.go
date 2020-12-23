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
	"errors"
	"fmt"
	"log"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

var initValidArgs = []string{"project"}
var initValidProjectTypes = []string{"basic", "cloud", "cloudformation", "terraform"}

// InitCmd command to initialize projects
func InitCmd() *cobra.Command {
	man := helper.GetManual("init")
	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: initValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   initPreRun,
		RunE:      initRun,
	}

	cmd.Flags().String("project-name", "", "The project name.")
	cmd.Flags().String("project-type", "basic", "The project type.")
	cmd.MarkFlagRequired("name")

	return cmd
}

func initPreRun(cmd *cobra.Command, args []string) error {
	if err := validateInitArgs(args); err != nil {
		return err
	}

	if err := validateProjectName(cmd, args); err != nil {
		return err
	}

	if err := validateProjectType(cmd, args); err != nil {
		return err
	}

	return nil
}

func validateInitArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("error: this command requires one argument")
	}

	if len(args) > 1 {
		return fmt.Errorf("error: this command accepts only one argument at a time")
	}

	if !helper.ContainsString(initValidArgs, args[0]) {
		return fmt.Errorf("error: unknown argument provided: %s", args[0])
	}

	return nil
}

func validateProjectName(cmd *cobra.Command, args []string) error {
	if args[0] == "project" {
		pName, err := cmd.Flags().GetString("project-name")
		if err != nil {
			return err
		}

		if pName == "" {
			return errors.New("project name must be defined")
		}
	} else {
		return fmt.Errorf("error: unknown argument provided: %s", args[0])
	}

	return nil
}

func validateProjectType(cmd *cobra.Command, args []string) error {
	// ensure the project types
	if args[0] == "project" {
		pType, err := cmd.Flags().GetString("project-type")

		if err != nil || pType == "" {
			return errors.New("project type must be defined")
		}

		if !helper.ContainsString(initValidProjectTypes, pType) {
			return fmt.Errorf("error: unknown project type provided: %s", pType)
		}
	} else {
		return fmt.Errorf("error: unknown argument provided: %s", args[0])
	}

	return nil
}

func initRun(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	if args[0] == "project" {
		pName, err := cmd.Flags().GetString("project-name")
		if err != nil {
			return fmt.Errorf("error: something went wrong \n%s", err)
		}

		pType, err := cmd.Flags().GetString("project-type")
		if err != nil {
			return fmt.Errorf("error: something went wrong \n%s", err)
		}

		switch pType {
		case "basic":
			err = aid.CreateBasicProject(pName)
		case "cloud":
			err = aid.CreateCloudProject(pName)
		case "cloudformation":
			err = aid.CreateCloudFormationProject(pName)
		case "terraform":
			err = aid.CreateTerraformProject(pName)
		default:
			return errors.New("unknow project type")
		}

		if err != nil {
			return fmt.Errorf("error: unable to initialize project sucessfully \n%s", err)
		}
	}

	return nil
}

func initGetFlags(cmd *cobra.Command) (name string, typee string, structure string, onlyCustomizedStructure bool) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatal("required flag name not set")
	}

	return name, typee, structure, onlyCustomizedStructure
}
