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
	"os"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/model"
)

// GetConfigurations read the current configurations file and return its model
func GetConfigurations() (model.Configurations, error) {
	var confs model.Configurations
	v, err := aid.ReadConfig(aid.GetAppInfo().ConfigurationsName)
	if err != nil {
		return confs, fmt.Errorf("unable to read configurations\n%v", err)
	}

	err = v.Unmarshal(&confs)
	if err != nil {
		return confs, fmt.Errorf("unable to unmarshall configurations\n%v", err)
	}

	return confs, err
}

// GetReadMe TODO...
func GetReadMe() (model.ReadMe, error) {
	var readMe model.ReadMe
	wd, err := os.Getwd()
	if err != nil {
		return readMe, fmt.Errorf("unable to get the current working directory:\n%v", err)
	}

	configPath := wd + "/clencli"
	configName := "readme"
	configType := "yaml"

	v, err := aid.ReadConfigAsViper(configPath, configName, configType)
	if err != nil {
		return readMe, fmt.Errorf("unable to read configurations\n%v", err)
	}

	err = v.Unmarshal(&readMe)
	if err != nil {
		return readMe, fmt.Errorf("unable to unmarshall configurations\n%v", err)
	}

	return readMe, err
}

// GetConfigurationProfile returns credentials of a profile
func GetConfigurationProfile(name string) (model.ConfigurationProfile, error) {
	configurations, err := GetConfigurations()

	if err != nil {
		return (model.ConfigurationProfile{}), err
	}

	for _, profile := range configurations.Profiles {
		if profile.Name == name {
			return profile, err
		}
	}

	return (model.ConfigurationProfile{}), err
}

// GetUnsplashRandomPhotoParameters TODO ...
func GetUnsplashRandomPhotoParameters(name string) model.UnsplashRandomPhotoParameters {
	profile, _ := GetConfigurationProfile(name)

	if len(profile.Configurations) > 0 {
		// TODO: improve this logic
		for _, conf := range profile.Configurations {
			return conf.Unsplash.RandomPhoto.Parameters
		}
	}

	return (model.UnsplashRandomPhotoParameters{})
}

// GetCredentials read the current credentials file and return its model
func GetCredentials() (model.Credentials, error) {
	var creds model.Credentials
	v, err := aid.ReadConfig(aid.GetAppInfo().CredentialsName)
	if err != nil {
		return creds, fmt.Errorf("unable to read credentials\n%v", err)
	}

	err = v.Unmarshal(&creds)
	if err != nil {
		return creds, fmt.Errorf("unable to unmarshall credentials\n%v", err)
	}

	return creds, err
}

// GetCredentialProfile returns credentials of a profile
func GetCredentialProfile(name string) (model.CredentialProfile, error) {
	credentials, err := GetCredentials()

	if err != nil {
		return (model.CredentialProfile{}), err
	}

	for _, profile := range credentials.Profiles {
		if profile.Name == name {
			return profile, err
		}
	}

	return (model.CredentialProfile{}), err
}

// GetCredentialByProvider return credentials based on the given provider, if non-existent, return an empty credential
func GetCredentialByProvider(profile string, provider string) (model.Credential, error) {
	cp, err := GetCredentialProfile(profile)
	if err != nil {
		return (model.Credential{}), err
	}

	for _, c := range cp.Credentials {
		if c.Provider == provider {
			return c, err
		}
	}

	return (model.Credential{}), err
}

// SaveConfigurations saves the given configuration onto the configurations file
func SaveConfigurations(configurations model.Configurations) error {
	return aid.WriteInterfaceToFile(configurations, aid.GetAppInfo().ConfigurationsPath)
}

// SaveCredentials saves the given credential onto the credentials file
func SaveCredentials(credentials model.Credentials) error {
	return aid.WriteInterfaceToFile(credentials, aid.GetAppInfo().CredentialsPath)
}
