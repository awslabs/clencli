package controller

import (
	"fmt"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/dao"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/spf13/cobra"
)

// ConfigureCmd command to display CLENCLI current version
func ConfigureCmd() *cobra.Command {
	man := cau.GetManual("configure")

	return &cobra.Command{
		Use:   man.Use,
		Short: man.Short,
		Long:  man.Long,
		RunE:  configureRun,
	}
}

func configureRun(cmd *cobra.Command, args []string) error {
	profile, _ := cmd.Flags().GetString("profile")

	if !aid.ConfigDirExist() {
		if aid.CreateConfigDir() {
			fmt.Println("CLENCLI configuration directory created")

			credentials := aid.CreateCredentials(profile)
			dao.SaveCredentials(credentials)

			configurations := aid.CreateConfigurations(profile)
			dao.SaveConfigurations(configurations)

		}
	} else if aid.ConfigDirExist() &&
		(!aid.CredentialsFileExist() || !aid.ConfigurationsFileExist()) {

		if !aid.CredentialsFileExist() {
			credentials := aid.CreateCredentials(profile)
			dao.SaveCredentials(credentials)
		}

		if !aid.ConfigurationsFileExist() {
			configurations := aid.CreateConfigurations(profile)
			dao.SaveConfigurations(configurations)
		}
	} else {
		if aid.ConfigDirExist() && aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {
			credentials := dao.GetCredentials()
			credentials = aid.UpdateCredentials(credentials)
			dao.SaveCredentials(credentials)

			configurations := dao.GetConfigurations()
			configurations = aid.UpdateConfigurations(configurations)
			dao.SaveConfigurations(configurations)
		}
	}

	return nil
}
