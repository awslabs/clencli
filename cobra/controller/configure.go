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
	"fmt"
	"os"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/dao"
	"github.com/awslabs/clencli/cobra/view"
	"github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

var configureValidArgs = []string{"delete"}

// ConfigureCmd command to display clencli current version
func ConfigureCmd() *cobra.Command {
	man, err := helper.GetManual("configure")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: configureValidArgs,
		Args:      cobra.OnlyValidArgs,
		RunE:      configureRun,
	}

	return cmd
}

func configureRun(cmd *cobra.Command, args []string) error {
	if !aid.ConfigurationsDirectoryExist() {
		if created, dir := aid.CreateConfigurationsDirectory(); created {
			cmd.Printf("clencli configuration directory created at %s\n", dir)

			answer := view.GetUserInputAsBool(cmd, "Would you like to setup credentials?", false)
			if answer {
				credentials := view.CreateCredentials(cmd, profile)
				dao.SaveCredentials(credentials)
			}

			answer = view.GetUserInputAsBool(cmd, "Would you like to setup configurations?", false)
			if answer {
				configurations := view.CreateConfigurations(cmd, profile)
				dao.SaveConfigurations(configurations)
			}
		}
	}

	// clencli configure credential --profile abc
	// clencli configure configuration --profile abc

	// |- config dir doesnt exist
	// |--  create config dir
	// |--  create credentials file and add default cred profile
	// |--  create configurations file and add default config profile

	// |- config dir exist
	// |--  add cred profile if profile doesnt exist
	// |----  update cred profile if profile exist
	// |--  add config profile if profile doesnt exist
	// |----  update config profile if profile exist
	// |--  delete cred and config profile if profile

	// config dir exist
	return nil
}
