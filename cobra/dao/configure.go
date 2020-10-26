package dao

import (
	"fmt"
	"time"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/model"

	log "github.com/sirupsen/logrus"
)

// Data Access Object

// Optional<T> get(long id);
// List<T> getAll();
// void save(T t);
// void update(T t, String[] params);
// void delete(T t);

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
func CreateConfigurations(name string) model.Configurations {
	fmt.Println("> Configurations")
	var configurations model.Configurations
	var profile model.ConfigurationProfile
	profile = CreateConfigurationProfile(name)
	configurations.Profiles = append(configurations.Profiles, profile)

	return configurations
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
func CreateCredentials(name string) model.Credentials {
	fmt.Println("> Credentials")
	var credentials model.Credentials
	profile := CreateCredentialProfile(name)
	credentials.Profiles = append(credentials.Profiles, profile)

	return credentials
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
func UpdateConfigurations(name string, confs model.Configurations) model.Configurations {
	fmt.Println("> Configurations")
	for i, profile := range confs.Profiles {
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

			confs.Profiles[i] = profile
		}

	}

	return confs
}

// UpdateCredentials does TODO
func UpdateCredentials(name string, creds model.Credentials) model.Credentials {
	fmt.Println("> Credentials")
	for i, profile := range creds.Profiles {

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

			creds.Profiles[i] = profile
		}
	}

	return creds
}
