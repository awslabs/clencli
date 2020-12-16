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
	"github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

var profile string

// RootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	man := helper.GetManual("root")
	cmd := &cobra.Command{
		Use:   man.Use,
		Short: man.Short,
		Long:  man.Long,
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here will be global for your application.
	cmd.PersistentFlags().StringVar(&profile, "profile", "default", "Use a specific profile from your configurations file")

	// Cobra also supports local flags, which will only run when this action is called directly.
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}
