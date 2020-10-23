package aid

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/awslabs/clencli/cobra/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// -- EXPORTED

// ConfigDirExist TODO
func ConfigDirExist() bool {
	return cau.DirOrFileExists(GetAppInfo().ConfigurationsDir)
}

// ConfigurationsFileExist does TODO
func ConfigurationsFileExist() bool {
	return cau.DirOrFileExists(GetAppInfo().ConfigurationsPath)
}

// CredentialsFileExist does TODO
func CredentialsFileExist() bool {
	return cau.DirOrFileExists(GetAppInfo().CredentialsPath)
}

// CreateConfigDir TODO
func CreateConfigDir() bool {
	return cau.CreateDir(GetAppInfo().ConfigurationsDir)
}

// WriteInterfaceToFile does TODO
func WriteInterfaceToFile(in interface{}, path string) error {
	b, err := yaml.Marshal(&in)
	if err != nil {
		_, ok := err.(*json.UnsupportedTypeError)
		if ok {
			return fmt.Errorf("Tried to marshal an invalid Type")
		}
	}

	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to update: %s \n%v", path, err)
	}

	return err
}

// GetCredentials does TODO
func GetCredentials() (model.Credentials, error) {
	var creds model.Credentials
	v, err := readConfig(GetAppInfo().CredentialsName)
	if err != nil {
		return creds, fmt.Errorf("Unable to read credentials\n%v", err)
	}

	err = v.Unmarshal(&creds)
	if err != nil {
		return creds, fmt.Errorf("Unable to unmarshall credentials \n%v", err)
	}

	return creds, err
}

// GetConfigurations does TODO
func GetConfigurations() (model.Configurations, error) {
	var confs model.Configurations
	v, err := readConfig(GetAppInfo().ConfigurationsName)
	if err != nil {
		return confs, fmt.Errorf("Unable to read credentials\n%v", err)
	}

	err = v.Unmarshal(&confs)
	if err != nil {
		return confs, fmt.Errorf("Unable to unmarshall configurations \n%v", err)
	}

	return confs, err
}

// CreateCredentials TODO
func CreateCredentials(name string) model.Credentials {
	fmt.Println("> Credentials")
	var credentials model.Credentials
	profile := createCredentialProfile(name)
	credentials.Profiles = append(credentials.Profiles, profile)

	return credentials
}

// createCredentialProfile does TODO
func createCredentialProfile(name string) model.CredentialProfile {
	fmt.Println(">> Profile")
	var profile model.CredentialProfile
	profile.Name = name
	profile.CreatedAt = time.Now().String()
	profile.Enabled = true // enabling profile by default

	var credential model.Credential
	credential.CreatedAt = time.Now().String()
	credential.Enabled = true
	credential = askAboutCredential(credential)

	profile.Credentials = append(profile.Credentials, credential)

	for {
		answer := getUserInputAsBool("Would you like to setup another credential?", false)
		if answer {
			var newCred model.Credential
			newCred = askAboutCredential(newCred)
			profile.Credentials = append(profile.Credentials, newCred)
		} else {
			break
		}
	}

	return profile
}

// UpdateCredentials does TODO
func UpdateCredentials(creds model.Credentials) model.Credentials {
	fmt.Println("> Credentials")
	for i, profile := range creds.Profiles {
		profile = askAboutCredentialProfile(profile)

		for j, cred := range profile.Credentials {
			profile.Credentials[j] = askAboutCredential(cred)
		}

		for {
			answer := getUserInputAsBool("Would you like to setup another credential?", false)
			if answer {
				var newCred model.Credential
				newCred = askAboutCredential(newCred)
				profile.Credentials = append(profile.Credentials, newCred)
			} else {
				break
			}
		}

		creds.Profiles[i] = profile
	}

	return creds
}

// CreateConfigurations does TODO
func CreateConfigurations(name string) model.Configurations {
	fmt.Println("> Configurations")
	var configurations model.Configurations
	var profile model.ConfigurationProfile
	profile = createConfigurationProfile(name)
	configurations.Profiles = append(configurations.Profiles, profile)

	return configurations
}

func createConfigurationProfile(name string) model.ConfigurationProfile {
	fmt.Println(">> Profile")
	var profile model.ConfigurationProfile
	profile.Name = name
	profile.CreatedAt = time.Now().String()
	profile.Enabled = true // enabling profile by default

	var configuration model.Configuration
	configuration.CreatedAt = time.Now().String()
	configuration.Enabled = true // enabling configuration by default
	configuration = askAboutConfiguration(configuration)

	profile.Configurations = append(profile.Configurations, configuration)

	for {
		answer := getUserInputAsBool("Would you like to setup another configuration?", false)
		if answer {
			var newConf model.Configuration
			newConf = askAboutConfiguration(newConf)
			profile.Configurations = append(profile.Configurations, newConf)
		} else {
			break
		}
	}

	return profile
}

// UpdateConfigurations does TODO
func UpdateConfigurations(confs model.Configurations) model.Configurations {
	fmt.Println("> Configurations")
	for i, profile := range confs.Profiles {
		profile = askAboutConfigurationProfile(profile)

		for j, conf := range profile.Configurations {
			profile.Configurations[j] = askAboutConfiguration(conf)
		}

		for {
			answer := getUserInputAsBool("Would you like to setup another configuration?", false)
			if answer {
				var newConf model.Configuration
				newConf = askAboutConfiguration(newConf)
				profile.Configurations = append(profile.Configurations, newConf)
			} else {
				break
			}
		}

		confs.Profiles[i] = profile
	}

	return confs
}

// --- UNEXPORTED

func askAboutConfigurationProfile(profile model.ConfigurationProfile) model.ConfigurationProfile {
	fmt.Println(">> Profile")
	profile.Name = getUserInputAsString(">>> Name", profile.Name)
	profile.Enabled = getUserInputAsBool(">>> Enabled", profile.Enabled)
	profile.Description = getUserInputAsString(">>> Description", profile.Description)
	profile.UpdatedAt = time.Now().String()
	return profile
}

func askAboutCredentialProfile(profile model.CredentialProfile) model.CredentialProfile {
	fmt.Println(">> Profile")
	profile.Name = getUserInputAsString(">>> Name", profile.Name)
	profile.Description = getUserInputAsString(">>> Description", profile.Description)
	profile.Enabled = getUserInputAsBool(">>> Enabled", profile.Enabled)
	profile.UpdatedAt = time.Now().String()
	return profile
}

func readConfig(name string) (*viper.Viper, error) {
	v := viper.New()
	app := GetAppInfo()

	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AddConfigPath(app.ConfigurationsDir)

	err := v.ReadInConfig()
	if err != nil {
		return v, fmt.Errorf("Error when trying to read local configurations \n%s", err)
	}
	return v, err

}

func askAboutCredential(credential model.Credential) model.Credential {
	fmt.Println(">>> Credential")
	credential.Name = getUserInputAsString(">>>> Name", credential.Name)
	credential.Enabled = getUserInputAsBool(">>>> Enabled", credential.Enabled)
	credential.Description = getUserInputAsString(">>>> Description", credential.Description)
	credential.UpdatedAt = time.Now().String()
	credential.Provider = getUserInputAsString(">>>> Provider", credential.Provider)
	credential.AccessKey = getSensitiveUserInputAsString(">>>> Access Key", credential.AccessKey)
	credential.SecretKey = getSensitiveUserInputAsString(">>>> Secret Key", credential.SecretKey)
	return credential
}

func getUserInput(text string, info string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if info == "" {
		fmt.Print(text + ": ")
	} else {
		fmt.Print(text + " [" + info + "]: ")
	}

	input, err := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		return input, fmt.Errorf("Unable to read user input \n%v", err)
	}
	return input, err
}

func getSensitiveUserInput(text string, info string) (string, error) {
	return getUserInput(text+" ["+maskString(info, 3)+"]", "")
}

func getUserInputAsBool(text string, info bool) bool {
	answer, err := getUserInput(text, strconv.FormatBool(info))
	if err != nil {
		log.Fatalf("Unable to get user input as boolean\n%s", err)
	}

	if answer != "" && answer == "true" {
		return true
	}

	return false
}

func getUserInputAsString(text string, info string) string {
	answer, err := getUserInput(text, info)
	if err != nil {
		log.Fatalf("Unable to get user input about profile's name\n%v", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}

func getSensitiveUserInputAsString(text string, info string) string {
	answer, err := getSensitiveUserInput(text, info)
	if err != nil {
		log.Fatalf("Unable to get user input about profile's name\n%v", err)
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

func askAboutConfiguration(conf model.Configuration) model.Configuration {
	// configuration can have many types: Unsplash, AWS, etc
	fmt.Println(">>> Configuration")
	conf.Name = getUserInputAsString(">>>> Name", conf.Name)
	conf.Description = getUserInputAsString(">>>> Description", conf.Description)
	conf.Enabled = getUserInputAsBool(">>>> Enabled", conf.Enabled)
	conf.UpdatedAt = time.Now().String()

	answer := getUserInputAsBool("Would you like to setup Unsplash configuration?", false)

	if answer {
		conf.Unsplash = askAboutUnsplashConfiguration(conf.Unsplash)
	} else {
		fmt.Println("Skipping Unplash configuration ...")
	}
	return conf
}

func askAboutUnsplashConfiguration(unsplash model.Unsplash) model.Unsplash {
	// unsplash configuration may have multiple nested configuration, such as random photo, etc...
	fmt.Println(">>>> Unsplash")
	unsplash.Name = getUserInputAsString(">>>>> Name", unsplash.Name)
	unsplash.Description = getUserInputAsString(">>>>> Description", unsplash.Description)
	unsplash.Enabled = getUserInputAsBool(">>>>> Enabled", unsplash.Enabled)
	unsplash.UpdatedAt = time.Now().String()

	answer := getUserInputAsBool("Would you like to setup Unsplash Random Photo Parameters?", false)

	if answer {
		fmt.Println(">>>>> Random Photo")
		unsplash.RandomPhoto.Name = getUserInputAsString(">>>>>> Name", unsplash.RandomPhoto.Name)
		unsplash.RandomPhoto.Description = getUserInputAsString(">>>>>> Description", unsplash.RandomPhoto.Description)
		unsplash.RandomPhoto.Enabled = getUserInputAsBool(">>>>>> Enabled", unsplash.RandomPhoto.Enabled)
		unsplash.RandomPhoto.UpdatedAt = time.Now().String()
		unsplash.RandomPhoto.Parameters = askAboutUnsplashRandomPhotoParameters(unsplash.RandomPhoto.Parameters)
	} else {
		fmt.Println("Skipping Unplash Random Photo configuration ...")
	}

	return unsplash
}

func askAboutUnsplashRandomPhotoParameters(params model.UnsplashRandomPhotoParameters) model.UnsplashRandomPhotoParameters {
	fmt.Println(">>>>>> Parameters")
	params.Collections = getUserInputAsString(">>>>>>> Public collection ID(‘s) to filter selection. If multiple, comma-separated.\nCollections ", params.Collections)
	params.Featured = getUserInputAsBool(">>>>>>> Limit selection to featured photos. Valid values: false and true. Default: false\nFeatured", params.Featured)
	params.Filter = getUserInputAsString(">>>>>>> Limit results by content safety. Valid values are low and high.\nFilter", params.Filter)
	params.Orientation = getUserInputAsString(">>>>>>> Filter by photo orientation. Valid values: landscape, portrait, squarish.\nOrientation", params.Orientation)
	params.Query = getUserInputAsString(">>>>>>> Limit selection to photos matching a search term.\nQuery", params.Query)
	params.Size = getUserInputAsString(">>>>>>> Photos size. Valid values: all, thumb, small, regular, full, raw.\nSize", params.Size)
	params.Username = getUserInputAsString(">>>>>>> Limit selection to a single user.\nUsername", params.Username)
	return params
}
