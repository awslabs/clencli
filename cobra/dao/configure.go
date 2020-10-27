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

package dao

import (
	"fmt"
	"time"

	"github.com/awslabs/clencli/cobra/view"

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
	configuration = view.AskAboutConfiguration(configuration)

	profile.Configurations = append(profile.Configurations, configuration)

	for {
		answer := view.GetUserInputAsBool("Would you like to setup another configuration?", false)
		if answer {
			var newConf model.Configuration
			newConf = view.AskAboutConfiguration(newConf)
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
	credential = view.AskAboutCredential(credential)

	profile.Credentials = append(profile.Credentials, credential)

	for {
		answer := view.GetUserInputAsBool("Would you like to setup another credential?", false)
		if answer {
			var newCred model.Credential
			newCred = view.AskAboutCredential(newCred)
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
	answer := view.GetUserInputAsBool("Do you want to delete the profile '"+name+"' from "+aid.GetAppInfo().ConfigurationsPath+"?", false)
	if answer {
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
}

// DeleteCredentialProfile delete a profile preserving the credentials order
func DeleteCredentialProfile(name string) {
	answer := view.GetUserInputAsBool("Do you want to delete the profile '"+name+"' from "+aid.GetAppInfo().CredentialsPath+"?", false)
	if answer {
		allCredentials := GetCredentials()
		var newCredentials model.Credentials
		for _, profile := range allCredentials.Profiles {
			if profile.Name != name {
				newCredentials.Profiles = append(newCredentials.Profiles, profile)
			}
		}

		SaveCredentials(newCredentials)
	}
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
			profile = view.AskAboutConfigurationProfile(profile)

			for j, conf := range profile.Configurations {
				profile.Configurations[j] = view.AskAboutConfiguration(conf)
			}

			for {
				answer := view.GetUserInputAsBool("Would you like to setup another configuration?", false)
				if answer {
					var newConf model.Configuration
					newConf = view.AskAboutConfiguration(newConf)
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
			profile = view.AskAboutCredentialProfile(profile)

			for j, cred := range profile.Credentials {
				profile.Credentials[j] = view.AskAboutCredential(cred)
			}

			for {
				answer := view.GetUserInputAsBool("Would you like to setup another credential?", false)
				if answer {
					var newCred model.Credential
					newCred = view.AskAboutCredential(newCred)
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
