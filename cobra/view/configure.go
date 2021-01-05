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

package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/awslabs/clencli/cobra/model"
	"github.com/spf13/cobra"
)

// CREDENTIALS

// CreateCredentials create the credentials
func CreateCredentials(cmd *cobra.Command, name string) model.Credentials {
	cmd.Println("> Credentials")
	var credentials model.Credentials
	cProfile := createCredentialProfile(cmd, name)
	credentials.Profiles = append(credentials.Profiles, cProfile)
	return credentials
}

func createCredentialProfile(cmd *cobra.Command, name string) model.CredentialProfile {
	cmd.Println(">> Profile: " + name)
	var cProfile model.CredentialProfile
	cProfile.Name = name
	cProfile.Description = "managed by clencli"
	cProfile.Enabled = true // enabling profile by default
	cProfile.CreatedAt = time.Now().String()
	cProfile.UpdatedAt = time.Now().String()

	var credential model.Credential
	credential.Enabled = true
	credential = createOrUpdateCredential(cmd, credential)
	cProfile.Credentials = append(cProfile.Credentials, credential)

	for {
		answer := GetUserInputAsBool(cmd, "Would you like to setup another credential?", false)
		if answer {
			var newCred model.Credential
			newCred.Enabled = true
			newCred = createOrUpdateCredential(cmd, newCred)
			cProfile.Credentials = append(cProfile.Credentials, newCred)
		} else {
			break
		}
	}

	return cProfile
}

func createOrUpdateCredential(cmd *cobra.Command, credential model.Credential) model.Credential {
	cmd.Println(">>> Credential")
	credential.Name = GetUserInputAsString(cmd, ">>>> Name", credential.Name)
	credential.Description = GetUserInputAsString(cmd, ">>>> Description", credential.Description)
	credential.Enabled = GetUserInputAsBool(cmd, ">>>> Enabled", credential.Enabled)
	credential.CreatedAt = time.Now().String()
	credential.UpdatedAt = time.Now().String()
	credential.Provider = GetUserInputAsString(cmd, ">>>> Provider", credential.Provider)
	credential.AccessKey = getSensitiveUserInputAsString(cmd, ">>>> Access Key", credential.AccessKey)
	credential.SecretKey = getSensitiveUserInputAsString(cmd, ">>>> Secret Key", credential.SecretKey)
	credential.SessionToken = getSensitiveUserInputAsString(cmd, ">>>> Session Token", credential.SessionToken)

	return credential
}

// CONFIGURATIONS

// CreateConfigurations create the configuration file with the given profile name
func CreateConfigurations(cmd *cobra.Command, name string) model.Configurations {
	cmd.Println("> Configurations")
	var configurations model.Configurations
	cProfile := createConfigurationProfile(cmd, name)
	configurations.Profiles = append(configurations.Profiles, cProfile)
	return configurations
}

// createConfigurationProfile create the given profile name into the configurations file, return the profile created
func createConfigurationProfile(cmd *cobra.Command, name string) model.ConfigurationProfile {
	cmd.Println(">> Profile: " + name)
	var cProfile model.ConfigurationProfile
	cProfile.Name = name
	cProfile.Description = "managed by clencli"
	cProfile.Enabled = true // enabling profile by default
	cProfile.CreatedAt = time.Now().String()
	cProfile.UpdatedAt = time.Now().String()

	var configuration model.Configuration
	configuration.Enabled = true
	configuration = createOrUpdateConfiguration(cmd, configuration)
	cProfile.Configurations = append(cProfile.Configurations, configuration)

	for {
		answer := GetUserInputAsBool(cmd, "Would you like to setup another configuration?", false)
		if answer {
			var newConf model.Configuration
			newConf.Enabled = true
			newConf = createOrUpdateConfiguration(cmd, newConf)
			cProfile.Configurations = append(cProfile.Configurations, newConf)
		} else {
			break
		}
	}

	return cProfile
}

func createOrUpdateConfiguration(cmd *cobra.Command, conf model.Configuration) model.Configuration {
	cmd.Println(">>> Configuration")
	conf.Name = GetUserInputAsString(cmd, ">>>> Name", conf.Name)
	conf.Description = GetUserInputAsString(cmd, ">>>> Description", conf.Description)
	conf.Enabled = GetUserInputAsBool(cmd, ">>>> Enabled", conf.Enabled)
	conf.CreatedAt = time.Now().String()
	conf.UpdatedAt = time.Now().String()
	conf.Unsplash = createOrUpdateUnsplashConfiguration(cmd, conf.Unsplash)

	// configuration.CreatedAt = time.Now().String()
	// configuration = view.AskAboutConfiguration(configuration)

	// profile.Configurations = append(profile.Configurations, configuration)

	// for {
	// 	answer := view.GetUserInputAsBool("Would you like to setup another configuration?", false)
	// 	if answer {
	// 		var newConf model.Configuration
	// 		newConf = view.AskAboutConfiguration(newConf)
	// 		profile.Configurations = append(profile.Configurations, newConf)
	// 	} else {
	// 		break
	// 	}
	// }

	return conf
}

// AskAboutConfiguration ask user about configuration
func createOrUpdateUnsplashConfiguration(cmd *cobra.Command, unsplash model.Unsplash) model.Unsplash {
	// configuration can have many types: Unsplash, AWS, etc
	answer := GetUserInputAsBool(cmd, "Would you like to setup Unsplash configuration?", false)
	if answer {
		cmd.Println(">>> Unsplash")
		unsplash.Name = GetUserInputAsString(cmd, ">>>>> Name", unsplash.Name)
		unsplash.Description = GetUserInputAsString(cmd, ">>>>> Description", unsplash.Description)
		unsplash.Enabled = GetUserInputAsBool(cmd, ">>>>> Enabled", unsplash.Enabled)
		// unsplash.CreatedAt = time.Now().String()
		unsplash.UpdatedAt = time.Now().String()
		unsplash = createOrUpdateUnsplashRandomPhotoConfiguration(cmd, unsplash)

	} else {
		cmd.Println("Skipping Unplash configuration ...")
	}
	return unsplash
}

func createOrUpdateUnsplashRandomPhotoConfiguration(cmd *cobra.Command, unsplash model.Unsplash) model.Unsplash {
	// unsplash configuration may have multiple nested configuration, such as random photo, etc...
	answer := GetUserInputAsBool(cmd, "Would you like to setup Unsplash Random Photo Parameters?", false)

	if answer {
		randomPhoto := unsplash.RandomPhoto
		cmd.Println(">>>>> Random Photo")
		randomPhoto.Name = GetUserInputAsString(cmd, ">>>>>> Name", randomPhoto.Name)
		randomPhoto.Description = GetUserInputAsString(cmd, ">>>>>> Description", randomPhoto.Description)
		randomPhoto.Enabled = GetUserInputAsBool(cmd, ">>>>>> Enabled", randomPhoto.Enabled)
		randomPhoto.UpdatedAt = time.Now().String()
		unsplash.RandomPhoto = randomPhoto

		params := randomPhoto.Parameters
		cmd.Println(">>>>>> Parameters")
		params.Collections = GetUserInputAsString(cmd, ">>>>>>> Collections (Public collection ID to filter selection. If multiple, comma-separated) ", params.Collections)
		params.Featured = GetUserInputAsBool(cmd, ">>>>>>> Featured (Limit selection to featured photos. Valid values: false and true)", params.Featured)
		params.Filter = GetUserInputAsString(cmd, ">>>>>>> Filter (Limit results by content safety. Valid values are low and high)", params.Filter)
		params.Orientation = GetUserInputAsString(cmd, ">>>>>>> Orientation (Filter by photo orientation. Valid values: landscape, portrait, squarish)", params.Orientation)
		params.Query = GetUserInputAsString(cmd, ">>>>>>> Query (Limit selection to photos matching a search term)", params.Query)
		params.Size = GetUserInputAsString(cmd, ">>>>>>> Size (Photos size. Valid values: all, thumb, small, regular, full, raw)", params.Size)
		params.Username = GetUserInputAsString(cmd, ">>>>>>> Username (Limit selection to a single user)", params.Username)
		unsplash.RandomPhoto.Parameters = params

	} else {
		cmd.Println("Skipping Unplash Random Photo configuration ...")
	}

	return unsplash
}

// AskAboutConfigurationProfile ask user about configuration profile
// func AskAboutConfigurationProfile(profile model.ConfigurationProfile) model.ConfigurationProfile {
// 	cmd.Println(">> Profile")
// 	profile.Name = GetUserInputAsString(">>> Name", profile.Name)
// 	profile.Enabled = GetUserInputAsBool(">>> Enabled", profile.Enabled)
// 	profile.Description = GetUserInputAsString(">>> Description", profile.Description)
// 	profile.UpdatedAt = time.Now().String()
// 	return profile
// }

// AskAboutCredential ask user about credential
// func AskAboutCredential(credential model.Credential) model.Credential {
// 	cmd.Println(">>> Credential")
// 	credential.Name = GetUserInputAsString(">>>> Name", credential.Name)
// 	credential.Enabled = GetUserInputAsBool(">>>> Enabled", credential.Enabled)
// 	credential.Description = GetUserInputAsString(">>>> Description", credential.Description)
// 	credential.UpdatedAt = time.Now().String()
// 	credential.Provider = GetUserInputAsString(">>>> Provider", credential.Provider)
// 	credential.AccessKey = getSensitiveUserInputAsString(">>>> Access Key", credential.AccessKey)
// 	credential.SecretKey = getSensitiveUserInputAsString(">>>> Secret Key", credential.SecretKey)
// 	return credential
// }

// AskAboutCredentialProfile ask user about credential profile
// func AskAboutCredentialProfile(profile model.CredentialProfile) model.CredentialProfile {
// 	cmd.Println(">> Profile")
// 	profile.Name = GetUserInputAsString(">>> Name", profile.Name)
// 	profile.Description = GetUserInputAsString(">>> Description", profile.Description)
// 	profile.Enabled = GetUserInputAsBool(">>> Enabled", profile.Enabled)
// 	profile.UpdatedAt = time.Now().String()
// 	return profile
// }

// func askAboutUnsplashRandomPhotoParameters(params model.UnsplashRandomPhotoParameters) model.UnsplashRandomPhotoParameters {

// 	return params
// }

func getSensitiveUserInput(cmd *cobra.Command, text string, info string) (string, error) {
	return getUserInput(cmd, text+" ["+maskString(info, 3)+"]", "")
}

func getSensitiveUserInputAsString(cmd *cobra.Command, text string, info string) string {
	answer, err := getSensitiveUserInput(cmd, text, info)
	if err != nil {
		log.Fatalf("unable to get user input about profile's name\n%v", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}

func getUserInput(cmd *cobra.Command, text string, info string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if info == "" {
		cmd.Print(text + ": ")
	} else {
		cmd.Print(text + " [" + info + "]: ")
	}

	input, err := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		return input, fmt.Errorf("unable to read user input\n%v", err)
	}

	return input, err
}

// GetUserInputAsBool prints `text` on console and return answer as `boolean`
func GetUserInputAsBool(cmd *cobra.Command, text string, info bool) bool {
	answer, err := getUserInput(cmd, text, strconv.FormatBool(info))
	if err != nil {
		log.Fatalf("unable to get user input as boolean\n%s", err)
	}

	if answer != "" && answer == "true" {
		return true
	}

	return false
}

// GetUserInputAsString prints `text` on console and return answer as `string`
func GetUserInputAsString(cmd *cobra.Command, text string, info string) string {
	answer, err := getUserInput(cmd, text, info)
	if err != nil {
		log.Fatalf("unable to get user input about profile's name\n%v", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}

func maskString(s string, showLastChars int) string {
	maskSize := len(s) - showLastChars
	if maskSize <= 0 {
		return s
	}

	return strings.Repeat("*", maskSize) + s[maskSize:]
}
