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
	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var profile string

//The verbose flag value
var verbosity string

// RootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	man := helper.GetManual("root")
	cmd := &cobra.Command{
		Use:               man.Use,
		Short:             man.Short,
		Long:              man.Long,
		PersistentPreRunE: rootPersistentPreRun,
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here will be global for your application.
	cmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "Use a specific profile from your credentials and configurations file")
	cmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", logrus.InfoLevel.String(), "Valid log level:panic,fatal,error,warn,info,debug,trace)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

func rootPersistentPreRun(cmd *cobra.Command, args []string) error {
	if err := aid.SetupLogging(verbosity); err != nil {
		return err
	}

	return nil
}
