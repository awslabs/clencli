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

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/dao"

	"github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

var configureValidArgs = []string{"delete"}

// ConfigureCmd command to display CLENCLI current version
func ConfigureCmd() *cobra.Command {
	man := helper.GetManual("configure")
	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: configureValidArgs,
		Args:      cobra.OnlyValidArgs,
		RunE:      configureRun,
	}

	cmd.Flags().StringP("profile", "p", "default", "Profile name")

	return cmd
}

func configureRun(cmd *cobra.Command, args []string) error {
	profile, _ := cmd.Flags().GetString("profile")

	if len(args) == 0 {
		if !aid.ConfigurationDirectoryExist() {
			if aid.CreateConfigDir() {
				dao.CreateCredentials(profile)
				dao.CreateConfigurations(profile)
			}
		} else if aid.ConfigurationDirectoryExist() &&
			(!aid.CredentialsFileExist() || !aid.ConfigurationsFileExist()) {

			if !aid.CredentialsFileExist() {
				dao.CreateCredentials(profile)
			}

			if !aid.ConfigurationsFileExist() {
				dao.CreateConfigurations(profile)
			}
		} else {
			if aid.ConfigurationDirectoryExist() && aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {
				if dao.CredentialsProfileExist(profile) && dao.ConfigurationsProfileExist(profile) {
					dao.UpdateCredentials(profile)
					dao.UpdateConfigurations(profile)
				}

				if !dao.CredentialsProfileExist(profile) {
					dao.AddCredentialProfile(profile)
				}

				if !dao.ConfigurationsProfileExist(profile) {
					dao.AddConfigurationProfile(profile)
				}
			}
		}
	} else if len(args) > 0 && args[0] == "delete" {
		if !aid.ConfigurationDirectoryExist() {
			return fmt.Errorf("CLENCLI configuration directory not found")
		}
		if !aid.CredentialsFileExist() {
			return fmt.Errorf("CLENCLI credentials file not found")
		}
		if !aid.ConfigurationsFileExist() {
			return fmt.Errorf("CLENCLI configurations file not found")
		}

		if aid.ConfigurationDirectoryExist() && aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {
			dao.DeleteCredentialProfile(profile)
			dao.DeleteConfigurationProfile(profile)
		}

	}

	return nil

}
