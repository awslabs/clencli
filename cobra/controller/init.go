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

// InitCmd command to initialize projects
func InitCmd() *cobra.Command {
	man := helper.GetManual("init")
	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Example:   man.Example,
		Long:      man.Long,
		ValidArgs: initValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   initPreRun,
		RunE:      initRun,
	}

	cmd.Flags().StringP("name", "n", "", "The project name.")
	cmd.Flags().StringP("type", "t", "basic", "The project type.")
	cmd.MarkFlagRequired("name")

	return cmd
}

func initPreRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("this command requires one argument")
	}

	err := validateProjectType(cmd, args)
	return err
}

func validateProjectType(cmd *cobra.Command, args []string) error {
	// ensure the project types
	if args[0] == "project" {
		pType, err := cmd.Flags().GetString("type")
		// flag accessed but not defined

		if err != nil || pType == "" {
			return errors.New("Project type must be defined")
		}

		if pType != "basic" && pType != "cloudformation" && pType != "terraform" {
			return fmt.Errorf("Unknown project type provided: %s", pType)
		}
	} else {
		return fmt.Errorf("Unknown argument provided: %s", args[0])
	}

	return nil
}

func initRun(cmd *cobra.Command, args []string) error {
	pName, err := cmd.Flags().GetString("name")
	pType, err := cmd.Flags().GetString("type")

	if err != nil {
		return fmt.Errorf("Something went wrong \n%s", err)
	}

	// structure, _ = cmd.Flags().GetString("structure")
	// onlyCustomizedStructure, _ = cmd.Flags().GetBool("only-customized-structure")

	switch pType {
	case "basic":
		aid.CreateBasicProject(pName)
	default:
		return errors.New("Unknow project type")
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
