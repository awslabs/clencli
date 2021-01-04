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
	// profile, _ := cmd.Flags().GetString("profile")

	if len(args) == 0 {
		if !aid.ConfigurationsDirectoryExist() {
			if aid.CreateConfigurationsDirectory() {
				// TODO: ask if user wants to setup credentials
				dao.CreateCredentials(profile)
				// TODO: ask if user wants to setup configurations
				dao.CreateConfigurations(profile)
			}
		} else if aid.ConfigurationsDirectoryExist() &&
			(!aid.CredentialsFileExist() || !aid.ConfigurationsFileExist()) {

			if !aid.CredentialsFileExist() {
				dao.CreateCredentials(profile)
			}

			if !aid.ConfigurationsFileExist() {
				dao.CreateConfigurations(profile)
			}
		} else {
			if aid.ConfigurationsDirectoryExist() && aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {
				crp, err := dao.CredentialsProfileExist(profile)
				if err != nil {
					return err
				}

				cop, err := dao.ConfigurationsProfileExist(profile)
				if err != nil {
					return err
				}

				if crp && cop {
					dao.UpdateCredentials(profile)
					dao.UpdateConfigurations(profile)
				}

				if !crp {
					dao.AddCredentialProfile(profile)
				}

				if !cop {
					dao.AddConfigurationProfile(profile)
				}
			}
		}
	} else if len(args) > 0 && args[0] == "delete" {
		if !aid.ConfigurationsDirectoryExist() {
			return fmt.Errorf("clencli configuration directory not found")
		}
		if !aid.CredentialsFileExist() {
			return fmt.Errorf("clencli credentials file not found")
		}
		if !aid.ConfigurationsFileExist() {
			return fmt.Errorf("clencli configurations file not found")
		}

		if aid.ConfigurationsDirectoryExist() && aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {
			dao.DeleteCredentialProfile(profile)
			dao.DeleteConfigurationProfile(profile)
		}

	}

	return nil

}
