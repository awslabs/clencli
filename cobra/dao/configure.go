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

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/model"
	"github.com/awslabs/clencli/cobra/view"

	log "github.com/sirupsen/logrus"
)

// AddConfigurationProfile add the given profile name into the configurations file
func AddConfigurationProfile(name string) {
	configurations := GetConfigurations()
	configurations.Profiles = append(configurations.Profiles, CreateConfigurationProfile(name))
	saveConfigurations(configurations)
}

// AddCredentialProfile add the given profile name into the credentials file
func AddCredentialProfile(name string) {
	credentials := GetCredentials()
	credentials.Profiles = append(credentials.Profiles, CreateCredentialProfile(name))
	saveCredentials(credentials)
}

// ConfigurationsProfileExist return `true` if the configuration file exist, `false` if otherwise
func ConfigurationsProfileExist(name string) bool {
	configurations := GetConfigurations()
	for _, profile := range configurations.Profiles {
		if profile.Name == name {
			return true
		}
	}
	return false

}

// CreateConfigurationProfile create the given profile name into the configurations file, return the profile created
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

// CreateConfigurations create the configuration file with the given profile name
func CreateConfigurations(name string) {
	fmt.Println("> Configurations")
	var configurations model.Configurations
	var profile model.ConfigurationProfile
	profile = CreateConfigurationProfile(name)
	configurations.Profiles = append(configurations.Profiles, profile)
	saveConfigurations(configurations)
}

// CreateCredentialProfile create the given profile name into the credentials file, return the profile created
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

// CreateCredentials create the credentials file with the given profile name
func CreateCredentials(name string) {
	fmt.Println("> Credentials")
	var credentials model.Credentials
	profile := CreateCredentialProfile(name)
	credentials.Profiles = append(credentials.Profiles, profile)
	saveCredentials(credentials)
}

// CredentialsProfileExist returns `true` if the profile name given exist in the credentials file
func CredentialsProfileExist(name string) bool {
	credentials := GetCredentials()
	for _, profile := range credentials.Profiles {
		if profile.Name == name {
			return true
		}
	}
	return false

}

// DeleteConfigurationProfile delete a profile from the configurations file
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
		saveConfigurations(newConfigurations)
	}
}

// DeleteCredentialProfile delete a profile from the credentials file
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

		saveCredentials(newCredentials)
	}
}

// GetConfigurations read the current configurations file and return its model
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

// GetCredentials read the current credentials file and return its model
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

// GetCredentialProfile returns credentials of a profile
func GetCredentialProfile(name string) model.CredentialProfile {
	credentials := GetCredentials()
	for _, profile := range credentials.Profiles {
		if profile.Name == name {
			return profile
		}
	}
	return (model.CredentialProfile{})
}

// GetCredentialByProvider return credentials based on the given provider, if non-existent, return an empty credential
func GetCredentialByProvider(profile string, provider string) model.Credential {
	var cp model.CredentialProfile = GetCredentialProfile(profile)
	for _, c := range cp.Credentials {
		if c.Provider == provider {
			return c
		}
	}

	return (model.Credential{})
}

func saveConfigurations(configurations model.Configurations) error {
	return aid.WriteInterfaceToFile(configurations, aid.GetAppInfo().ConfigurationsPath)
}

func saveCredentials(credentials model.Credentials) error {
	return aid.WriteInterfaceToFile(credentials, aid.GetAppInfo().CredentialsPath)
}

// UpdateConfigurations updates the given profile name in the configurations file
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

	saveConfigurations(configurations)
}

// UpdateCredentials updates the given profile name in the credentials file
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

	saveCredentials(credentials)
}
