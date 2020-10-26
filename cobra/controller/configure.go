package controller

import (
	"fmt"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/dao"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/spf13/cobra"
)

var configureValidArgs = []string{"delete"}

// ConfigureCmd command to display CLENCLI current version
func ConfigureCmd() *cobra.Command {
	man := cau.GetManual("configure")

	return &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		ValidArgs: configureValidArgs,
		Args:      cobra.OnlyValidArgs,
		RunE:      configureRun,
	}
}

func configureRun(cmd *cobra.Command, args []string) error {
	profile, _ := cmd.Flags().GetString("profile")
	arg := args[0]

	if arg != "delete" {
		// todo: cases to cover, configure --profile name

		if !aid.ConfigDirExist() {
			if aid.CreateConfigDir() {
				fmt.Println("CLENCLI configuration directory created")

				credentials := dao.CreateCredentials(profile)
				dao.SaveCredentials(credentials)

				configurations := dao.CreateConfigurations(profile)
				dao.SaveConfigurations(configurations)

			}
		} else if aid.ConfigDirExist() &&
			(!aid.CredentialsFileExist() || !aid.ConfigurationsFileExist()) {

			if !aid.CredentialsFileExist() {
				credentials := dao.CreateCredentials(profile)
				dao.SaveCredentials(credentials)
			}

			if !aid.ConfigurationsFileExist() {
				configurations := dao.CreateConfigurations(profile)
				dao.SaveConfigurations(configurations)
			}
		} else {
			if aid.ConfigDirExist() && aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {
				credentials := dao.GetCredentials()
				configurations := dao.GetConfigurations()

				if aid.CredentialsProfileExist(profile, credentials) && aid.ConfigurationsProfileExist(profile, configurations) {
					credentials = dao.UpdateCredentials(profile, credentials)
					dao.SaveCredentials(credentials)

					configurations = dao.UpdateConfigurations(profile, configurations)
					dao.SaveConfigurations(configurations)
				}

				if !aid.CredentialsProfileExist(profile, credentials) {
					credentials.Profiles = append(credentials.Profiles, dao.CreateCredentialProfile(profile))
					dao.SaveCredentials(credentials)
				}

				if !aid.ConfigurationsProfileExist(profile, configurations) {
					configurations.Profiles = append(configurations.Profiles, dao.CreateConfigurationProfile(profile))
					dao.SaveConfigurations(configurations)
				}
			}
		}
	} else {
		if !aid.ConfigDirExist() {
			return fmt.Errorf("CLENCLI configuration directory not found")
		}
		if !aid.CredentialsFileExist() {
			return fmt.Errorf("CLENCLI credentials file not found")
		}
		if !aid.ConfigurationsFileExist() {
			return fmt.Errorf("CLENCLI configurations file not found")
		}

		if aid.ConfigDirExist() && aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {
			dao.DeleteCredentialProfile(profile)
			dao.DeleteConfigurationProfile(profile)
		}

	}

	return nil

}
