package dao

import (
	"fmt"
	"time"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/model"

	log "github.com/sirupsen/logrus"
)

// AddConfigurationProfile TODO
func AddConfigurationProfile(name string) {
	configurations := GetConfigurations()
	configurations.Profiles = append(configurations.Profiles, CreateConfigurationProfile(name))
	SaveConfigurations(configurations)
}

// AddCredentialProfile TODO
func AddCredentialProfile(name string) {
	credentials := GetCredentials()
	credentials.Profiles = append(credentials.Profiles, CreateCredentialProfile(name))
	SaveCredentials(credentials)
}

// ConfigurationsProfileExist TODO
func ConfigurationsProfileExist(name string) bool {
	configurations := GetConfigurations()
	for _, profile := range configurations.Profiles {
		if profile.Name == name {
			return true
		}
	}
	return false

}

// CreateConfigurationProfile TODO
func CreateConfigurationProfile(name string) model.ConfigurationProfile {
	fmt.Println(">> Profile")
	var profile model.ConfigurationProfile
	profile.Name = name
	profile.CreatedAt = time.Now().String()
	profile.Enabled = true // enabling profile by default

	var configuration model.Configuration
	configuration.CreatedAt = time.Now().String()
	configuration.Enabled = true // enabling configuration by default
	configuration = aid.AskAboutConfiguration(configuration)

	profile.Configurations = append(profile.Configurations, configuration)

	for {
		answer := aid.GetUserInputAsBool("Would you like to setup another configuration?", false)
		if answer {
			var newConf model.Configuration
			newConf = aid.AskAboutConfiguration(newConf)
			profile.Configurations = append(profile.Configurations, newConf)
		} else {
			break
		}
	}

	return profile
}

// CreateConfigurations does TODO
func CreateConfigurations(name string) {
	fmt.Println("> Configurations")
	var configurations model.Configurations
	var profile model.ConfigurationProfile
	profile = CreateConfigurationProfile(name)
	configurations.Profiles = append(configurations.Profiles, profile)
	SaveConfigurations(configurations)
}

// CreateCredentialProfile TODO
func CreateCredentialProfile(name string) model.CredentialProfile {
	fmt.Println(">> Profile")
	var profile model.CredentialProfile
	profile.Name = name
	profile.CreatedAt = time.Now().String()
	profile.Enabled = true // enabling profile by default

	var credential model.Credential
	credential.CreatedAt = time.Now().String()
	credential.Enabled = true
	credential = aid.AskAboutCredential(credential)

	profile.Credentials = append(profile.Credentials, credential)

	for {
		answer := aid.GetUserInputAsBool("Would you like to setup another credential?", false)
		if answer {
			var newCred model.Credential
			newCred = aid.AskAboutCredential(newCred)
			profile.Credentials = append(profile.Credentials, newCred)
		} else {
			break
		}
	}

	return profile
}

// CreateCredentials TODO
func CreateCredentials(name string) {
	fmt.Println("> Credentials")
	var credentials model.Credentials
	profile := CreateCredentialProfile(name)
	credentials.Profiles = append(credentials.Profiles, profile)
	SaveCredentials(credentials)
}

// CredentialsProfileExist TODO
func CredentialsProfileExist(name string) bool {
	credentials := GetCredentials()
	for _, profile := range credentials.Profiles {
		if profile.Name == name {
			return true
		}
	}
	return false

}

// DeleteConfigurationProfile delete a profile preserving it order
func DeleteConfigurationProfile(name string) {
	allConfigurations := GetConfigurations()
	var newConfigurations model.Configurations
	for _, profile := range allConfigurations.Profiles {
		// only append profile that doesn't match
		if profile.Name != name {
			newConfigurations.Profiles = append(newConfigurations.Profiles, profile)
		}
	}
	SaveConfigurations(newConfigurations)
}

// DeleteCredentialProfile delete a profile preserving the credentials order
func DeleteCredentialProfile(name string) {
	allCredentials := GetCredentials()
	var newCredentials model.Credentials
	for _, profile := range allCredentials.Profiles {
		if profile.Name != name {
			newCredentials.Profiles = append(newCredentials.Profiles, profile)
		}
	}

	SaveCredentials(newCredentials)
}

// GetConfigurations does TODO
func GetConfigurations() model.Configurations {
	var confs model.Configurations
	v, err := aid.ReadConfig(aid.GetAppInfo().ConfigurationsName)
	if err != nil {
		log.Fatalf("Unable to read configurations\n%v", err)
	}

	err = v.Unmarshal(&confs)
	if err != nil {
		log.Fatalf("Unable to unmarshall configurations \n%v", err)
	}

	return confs
}

// GetCredentials does TODO
func GetCredentials() model.Credentials {
	var creds model.Credentials
	v, err := aid.ReadConfig(aid.GetAppInfo().CredentialsName)
	if err != nil {
		log.Fatalf("Unable to read credentials\n%v", err)
	}

	err = v.Unmarshal(&creds)
	if err != nil {
		log.Fatalf("Unable to unmarshall credentials \n%v", err)
	}

	return creds
}

// SaveConfigurations TODO
func SaveConfigurations(configurations model.Configurations) error {
	return aid.WriteInterfaceToFile(configurations, aid.GetAppInfo().ConfigurationsPath)
}

// SaveCredentials TODO
func SaveCredentials(credentials model.Credentials) error {
	return aid.WriteInterfaceToFile(credentials, aid.GetAppInfo().CredentialsPath)
}

// UpdateConfigurations does TODO
func UpdateConfigurations(name string) {
	fmt.Println("> Configurations")
	configurations := GetConfigurations()
	for i, profile := range configurations.Profiles {
		if profile.Name == name {
			profile = aid.AskAboutConfigurationProfile(profile)

			for j, conf := range profile.Configurations {
				profile.Configurations[j] = aid.AskAboutConfiguration(conf)
			}

			for {
				answer := aid.GetUserInputAsBool("Would you like to setup another configuration?", false)
				if answer {
					var newConf model.Configuration
					newConf = aid.AskAboutConfiguration(newConf)
					profile.Configurations = append(profile.Configurations, newConf)
				} else {
					break
				}
			}

			configurations.Profiles[i] = profile
		}

	}

	SaveConfigurations(configurations)
}

// UpdateCredentials does TODO
func UpdateCredentials(name string) {
	fmt.Println("> Credentials")
	credentials := GetCredentials()
	for i, profile := range credentials.Profiles {

		if profile.Name == name {
			profile = aid.AskAboutCredentialProfile(profile)

			for j, cred := range profile.Credentials {
				profile.Credentials[j] = aid.AskAboutCredential(cred)
			}

			for {
				answer := aid.GetUserInputAsBool("Would you like to setup another credential?", false)
				if answer {
					var newCred model.Credential
					newCred = aid.AskAboutCredential(newCred)
					profile.Credentials = append(profile.Credentials, newCred)
				} else {
					break
				}
			}

			credentials.Profiles[i] = profile
		}
	}

	SaveCredentials(credentials)
}
