package controller

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/awslabs/clencli/cobra/model"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	if profile == "" {
		profile = "default"
	}

	fmt.Println("Started to configure profile:" + profile)

	// ensure config directory exists, or it will be created
	configDirExist := doesConfigDirExist()

	if !configDirExist {
		created, err := createConfigDir()
		if err != nil {
			return fmt.Errorf("Unexpected error during config directory creation \n%v", err)
		}

		if created {
			err := createProfile(profile)
			if err != nil {
				return fmt.Errorf("Unexpected error when trying to setup profile: %s\n%s", profile, err)
			}
		}

	} else {
		err := updateProfile(profile)
		if err != nil {
			return fmt.Errorf("Unexpected error when trying to update profile: %s\n%s", profile, err)
		}

		// check existing config
		// if not existing, create new ones

	}

	fmt.Println("Finished to configure profile:" + profile)

	return nil
}

func doesConfigDirExist() bool {
	return cau.DirOrFileExists(getConfigDirPath())
}

func getConfigDirPath() string {
	home := getHomeDir()
	path := home + "/.clencli"

	return path
}

func getCredentialsPath() string {
	return getConfigDirPath() + "/credentials.yaml"
}

func getConfigPath() string {
	return getConfigDirPath() + "/config.yaml"
}

func getHomeDir() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Unable to detect the home directory\n%v", err)
		os.Exit(1)
	}
	return home
}

func createConfigDir() (bool, error) {
	created := false
	path := getConfigDirPath()

	if cau.CreateDir(path) {
		fmt.Println("CLENCLI configuration directory created at " + path)
		created = true
	}

	return created, nil
}

func createProfile(profile string) error {

	// credentials
	err := createCredentials(profile)
	if err != nil {
		return fmt.Errorf("Unexpected error during credentials setup")
	}

	// configuratons
	err = createConfig(profile)
	if err != nil {
		return fmt.Errorf("Unexpected error during config setup")
	}

	return err
}

func createCredentials(profile string) error {
	fmt.Println("Started to setup credentials for profile: " + profile)
	var credentials model.Credentials
	// load current credentials

	// configure unsplash
	answer, err := getUserInput("Would you like to setup Unsplash Credentials? (Y/y)")
	if err != nil {
		return fmt.Errorf("Unable to get answer if user wants to setup Unsplash Credentials")
	}

	if answer == "Y" || answer == "y" {
		// if empty, create new credential
		var cp model.CredentialProfile
		cp.Name = profile
		cp.Enabled = true // enabling profile by default
		cp.Credential, err = createCredential("unsplash")
		if err != nil {
			return fmt.Errorf("Unable to configure Unsplash credentials")
		}
		credentials.Profiles = append(credentials.Profiles, cp)
		err = saveCredentials(credentials)
		if err != nil {
			return fmt.Errorf("Unable to save credentials during setup \n%v", err)
		}
	} else {
		fmt.Println("Skipping Unplash configuration ...")
	}

	fmt.Println("Finished to setup credentials for profile: " + profile)
	return err
}

func createCredential(provider string) (model.Credential, error) {
	fmt.Println("=== Starting to configure credential for provider: " + provider)

	var err error
	var credential model.Credential
	credential.Provider = provider
	credential.AccessKey, err = getUserInput("Unsplash API Access Key")
	if err != nil {
		return credential, fmt.Errorf("Unable to read Unsplash API Access Key")
	}
	credential.SecretKey, err = getUserInput("Unsplash API Secret Key")
	if err != nil {
		return credential, fmt.Errorf("Unable to read Unsplash API Secret Key")
	}

	fmt.Println("=== Finished configuring credential for provider: " + provider)
	return credential, err
}

func maskString(s string, showLastChars int) string {
	maskSize := len(s) - showLastChars
	if maskSize <= 0 {
		return s
	}

	return strings.Repeat("*", maskSize) + s[maskSize:]
}

func createConfig(profile string) error {
	fmt.Println("Started to setup configuration for profile: " + profile)

	var config model.Config
	// load current configuration

	// configure unsplash
	answer, err := getUserInput("Would you like to setup Unsplash Random Photo Parameters? (Y/y)")
	if err != nil {
		return fmt.Errorf("Unable to get answer if user wants to setup Unsplash Random Photo Parameters")
	}

	if answer == "Y" || answer == "y" {
		var cp model.ConfigProfile
		cp.Name = profile
		cp.Enabled = true // enabling profile by default
		cp.Unsplash, err = createUnsplash()
		if err != nil {
			return fmt.Errorf("Unable to configure Unsplash")
		}
		config.Profiles = append(config.Profiles, cp)
		err = saveConfig(config)
		if err != nil {
			return fmt.Errorf("Unable to save config during setup \n%v", err)
		}
	} else {
		fmt.Println("Skipping Unplash configuration ...")
	}

	fmt.Println("Finished to setup configuration for profile: " + profile)
	return err
}

func createUnsplash() (model.Unsplash, error) {
	fmt.Println("=== Unsplash Random Photo Parameters")

	var unsplash model.Unsplash
	unsplash, err := getUserInputAboutUnsplash()
	if err != nil {
		return unsplash, fmt.Errorf("Unable to update Unsplash config from user input \n%v", err)
	}

	fmt.Println("=== Unsplash Random Photo Parameters configured successfully!")
	return unsplash, err
}

func getUserInput(text string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text + ": ")
	input, err := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		return input, fmt.Errorf("Unable to read user input \n%v", err)
	}
	return input, err
}

func getUserInputAboutUnsplash() (model.Unsplash, error) {

	var unsplash model.Unsplash
	unsplash.Enabled = true
	var err error

	unsplash.RandomPhoto.Parameters.Collections, err = getUserInput("Public collection ID(‘s) to filter selection. If multiple, comma-separated.\nCollections ")
	if err != nil {
		return unsplash, fmt.Errorf("Unable to get user input Unsplash.Parameters.Collections")
	}

	answer, err := getUserInput("Limit selection to featured photos. Valid values: false and true. Default: false\nFeatured")
	if answer == "true" {
		unsplash.RandomPhoto.Parameters.Featured = true

	} else if answer == "false" || answer == "" {
		unsplash.RandomPhoto.Parameters.Featured = false
	} else {
		return unsplash, errors.New("Invalid value. Only false or true are valid values")
	}

	if err != nil {

	}

	unsplash.RandomPhoto.Parameters.Filter, err = getUserInput("Limit results by content safety. Default: low. Valid values are low and high.\nFilter")
	if err != nil {
		return unsplash, fmt.Errorf("Unable to get user input Unsplash.Parameters.Filter")
	}

	unsplash.RandomPhoto.Parameters.Orientation, err = getUserInput("Filter by photo orientation. Valid values: landscape, portrait, squarish.\nOrientation")
	if err != nil {
		return unsplash, fmt.Errorf("Unable to get user input Unsplash.Parameters.Orientation")
	}

	unsplash.RandomPhoto.Parameters.Query, err = getUserInput("Limit selection to photos matching a search term.\nQuery")
	if err != nil {
		return unsplash, fmt.Errorf("Unable to get user input Unsplash.Parameters.Query")
	}

	// unsplash.RandomPhoto.Parameters.Size, err = getUserInput("Photos size. Valid values: all, thumb, small, regular, full, raw. Default: all.\nSize")
	// if err != nil {
	// 	return unsplash, fmt.Errorf("Unable to get user input Unsplash.Parameters.Size")
	// }

	// unsplash.RandomPhoto.Parameters.Username, err = getUserInput("Limit selection to a single user.\nUsername")
	// if err != nil {
	// 	return unsplash, fmt.Errorf("Unable to get user input Unsplash.Parameters.Username")
	// }

	return unsplash, err
}

func saveCredentials(credentials model.Credentials) error {
	path := getConfigDirPath()
	path += "/credentials.yaml"
	return writeInterfaceToFile(credentials, path)
}

func saveConfig(config model.Config) error {
	path := getConfigDirPath()
	path += "/config.yaml"
	return writeInterfaceToFile(config, path)
}

func writeInterfaceToFile(in interface{}, path string) error {
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

func updateProfile(profile string) error {
	var credentials model.Credentials
	// var config model.Config
	var err error

	mustCreateCredentials := shouldCreateCredentials(profile)
	if mustCreateCredentials {
		err = createCredentials(profile)
		if err != nil {
			return fmt.Errorf("Unexpected error while creating credentials\n%v", err)
		}
	} else {
		// update credentials
		credentials, err = readProfileCredentials(profile)
		if err != nil {
			return err
		}

		err = updateCredentials(credentials)
		if err != nil {
			return fmt.Errorf("Unable to update the profile's credentials")
		}
	}

	mustCreateConfig := shouldCreateConfig(profile)
	if mustCreateConfig {
		err = createConfig(profile)
		if err != nil {
			return fmt.Errorf("Unexpected error while creating config\n%v", err)
		}
	}

	// update configs

	return nil
}

func shouldCreateCredentials(profile string) bool {
	// check if files are empty
	size, _ := cau.FileSize(getCredentialsPath())
	if size <= 0 {
		return true
	}

	// check if profile is empty
	isEmpty, _ := isProfileCredentialsEmpty(profile)

	if isEmpty {
		return true
	}

	return false
}

func shouldCreateConfig(profile string) bool {
	// check if files are empty
	size, _ := cau.FileSize(getConfigPath())
	if size <= 0 {
		return true
	}

	// check if profile is empty
	isEmpty, _ := isProfileConfigEmpty(profile)
	if isEmpty {
		return true
	}

	return false
}

func readAllCredentials() (model.Credentials, error) {
	c := model.Credentials{}
	v, err := getViperInstance("credentials") // file within the configuration directory
	if err != nil {
		return c, fmt.Errorf("Unable to load existing credentials\n%v", err)
	}

	err = v.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("Unable to unmarshall config \n%v", err)
	}

	return c, err
}

func readAllConfig() (model.Config, error) {
	c := model.Config{}
	v, err := getViperInstance("credentials") // file within the configuration directory
	if err != nil {
		return c, fmt.Errorf("Unable to load existing credentials\n%v", err)
	}

	err = v.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("Unable to unmarshall config \n%v", err)
	}

	return c, err
}

func readProfileCredentials(profile string) (model.Credentials, error) {
	var credentials model.Credentials
	allCredentials, err := readAllCredentials()
	if err != nil {
		return credentials, fmt.Errorf("Unable to read credentials\n%v", err)
	}
	for _, p := range allCredentials.Profiles {
		if p.Name == profile {
			credentials.Profiles = append(credentials.Profiles, p)
		}
	}

	return credentials, err
}

func readProfileConfig(profile string) (model.Config, error) {
	var config model.Config
	allConfigs, err := readAllConfig()
	if err != nil {
		return config, fmt.Errorf("Unable to read config\n%v", err)
	}
	for _, p := range allConfigs.Profiles {
		if p.Name == profile {
			config.Profiles = append(config.Profiles, p)
		}
	}

	return config, err
}

func isProfileConfigEmpty(profile string) (bool, error) {
	// check credentials
	var answer bool = true
	config, err := readProfileConfig(profile)
	if err != nil {
		return answer, fmt.Errorf("Unable to read profiles from config\n%v", err)
	}

	for _, p := range config.Profiles {
		if p.Name == profile {
			answer = false
		}
	}

	return answer, err
}

func isProfileCredentialsEmpty(profile string) (bool, error) {
	credentials, err := readProfileCredentials(profile)
	if err != nil {
		return false, fmt.Errorf("Unable to read profiles from credentials\n%v", err)
	}

	for _, p := range credentials.Profiles {
		if p.Name == profile {
			return false, err
		}
	}

	return true, err
}

func readConfig() (model.Config, error) {
	v := viper.GetViper()
	c := model.Config{}

	err := v.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("Unable to unmarshall config \n%v", err)
	}

	return c, err
}

func getViperInstance(name string) (*viper.Viper, error) {
	lc := viper.New()

	lc.SetConfigName(name)
	lc.SetConfigType("yaml")

	configDirPath := getConfigDirPath()
	lc.AddConfigPath(configDirPath)

	err := lc.ReadInConfig()
	if err != nil {
		return lc, fmt.Errorf("Error when trying to read local config \n%s", err)
	}
	return lc, err
}

func updateCredentials(credentials model.Credentials) error {
	var err error
	for i, profile := range credentials.Profiles {
		credentials.Profiles[i].Credential, err = updateCredential(profile.Credential)
		if err != nil {
			return fmt.Errorf("Unable to update credential\n%v", err)
		}
	}

	err = saveCredentials(credentials)
	if err != nil {
		return fmt.Errorf("Unable to save credentials\n%v", err)
	}

	return err
}

func updateCredential(credential model.Credential) (model.Credential, error) {
	var err error
	if credential.AccessKey != "" {
		accessKey, err := getUserInput("Unsplash API Access Key [" + maskString(credential.AccessKey, 3) + "]")
		if err != nil {
			return credential, fmt.Errorf("Unable to get user input about access key\n%v", err)
		}

		if accessKey != "" {
			credential.AccessKey = accessKey
		}

	} else {
		createCredential(credential.Provider)
	}

	if credential.SecretKey != "" {
		secretKey, err := getUserInput("Unsplash API Secret Key [" + maskString(credential.SecretKey, 3) + "]")
		if err != nil {
			return credential, fmt.Errorf("Unable to get user input about access key\n%v", err)
		}

		if secretKey != "" {
			credential.SecretKey = secretKey
		}

	} else {
		createCredential(credential.Provider)
	}

	return credential, err
}

// // mergeConfig merges a new configuration with an existing config.
// func mergeConfig(local *viper.Viper, source ReadMe) error {
// 	s, _ := json.Marshal(source)
// 	return local.MergeConfig(bytes.NewBuffer(s))
// }

// // getLocalReadMeConfig unmarshall local readme.yaml return as ReadMe struct
// func getLocalReadMeConfig() (ReadMe, error) {
// 	r := ReadMe{}
// 	c, err := getLocalConfigV3("readme")
// 	if err != nil {
// 		return r, fmt.Errorf("Unable to get readme config with Viper \n%v", err)
// 	}

// 	err = c.Unmarshal(&r)
// 	if err != nil {
// 		return r, fmt.Errorf("Unable to Unmarshall ReadMe struct \n%v", err)
// 	}

// 	return r, err
// }

// func getLocalUnplashConfig() (Unsplash, error) {
// 	un := Unsplash{}
// 	c, err := getLocalConfigV3("unsplash")
// 	if err != nil {
// 		return un, fmt.Errorf("Unable to get unsplash config with Viper \n%v", err)
// 	}

// 	err = c.Unmarshal(&un)
// 	if err != nil {
// 		return un, fmt.Errorf("Unable to unmarshall Unsplash struct \n%v", err)
// 	}

// 	return un, err
// }

// // GetGlobalReadMeConfig return the ReadMe section as struct from the global config
// // func GetGlobalReadMeConfig() (ReadMe, error) {
// // 	c, err := getGlobalConfig()
// // 	r := c.Config.ReadMe
// // 	if err != nil {
// // 		return r, fmt.Errorf("Unable to get global config \n%v", err)
// // 	}

// // 	return r, err
// // }

// // MarshallAndSaveReadMe receive a readme struct and saves it back as file
// func MarshallAndSaveReadMe(readme ReadMe) error {
// 	// Marshal back into yaml
// 	r, err := yaml.Marshal(&readme)
// 	if err != nil {
// 		return fmt.Errorf("Unable to Marshall ReadMe struct %v", err)
// 	}

// 	err = ioutil.WriteFile("clencli/readme.yaml", r, os.ModePerm)
// 	if err != nil {
// 		return fmt.Errorf("Unable to update clencli/readme.yaml file %v", err)
// 	}

// 	return nil
// }

// // saveConfig does TODO

// // UpdateReadMe updates local config with global config
// // func UpdateReadMe() error {
// // 	g, err := GetGlobalReadMeConfig()
// // 	if err != nil {
// // 		return fmt.Errorf("Unable to get global readme config \n%v", err)
// // 	}

// // 	l, err := getLocalReadMeConfig()
// // 	if err != nil {
// // 		return fmt.Errorf("Unable to get local readme config \n%v", err)
// // 	}

// // 	updated := false

// // 	if g.Logo.Provider == "unsplash" {

// // 		// Unsplash.Query != "" {
// // 		// l.Logo.Unsplash.Query = g.Logo.Unsplash.Query
// // 		// updated = true
// // 	}

// // 	if g.Logo.URL != "" {
// // 		l.Logo.URL = g.Logo.URL
// // 		updated = true
// // 	}

// // 	if g.License != "" {
// // 		l.License = g.License
// // 		updated = true
// // 	}

// // 	if g.Copyright != "" {
// // 		l.Copyright = g.Copyright
// // 		updated = true
// // 	}

// // 	if updated {
// // 		err = MarshallAndSaveReadMe(l)
// // 		if err != nil {
// // 			return fmt.Errorf("Unable to update cleancli/readme.yaml  \n%v", err)
// // 		}
// // 	}
// // 	return nil
// // }

// // UpdateReadMeLogoURL fetches random photo based readme.logo.theme from config
// // func UpdateReadMeLogoURL() error {

// // 	global, err := getGlobalConfig()
// // 	if err != nil {
// // 		return fmt.Errorf("Unable to get global config \n%v", err)
// // 	}

// // 	for _, cred := range global.Credentials {
// // 		// verify if Unsplash credentials are set
// // 		if cred.Provider == "unsplash" {
// // 			if cred.AccessKey != "" && cred.SecretKey != "" {
// // 				fmt.Printf("Unsplash credentials found, using profile name: %s\n", cred.Name)

// // 				// check if global confign has unsplash as provider
// // 				if global.Config.ReadMe.Logo.Provider == "unsplash" {
// // 					// check if unsplash is enabled globally
// // 					if global.Config.Unsplash.Enabled {
// // 						// check if random photo parameters are set
// // 						if (UnsplashRandomPhotoParameters{} != global.Config.Unsplash.UnsplashRandomPhotoParameters) {
// // 							randomPhotoResponse, err := getUnsplashRandomPhoto(global.Config.Unsplash.UnsplashRandomPhotoParameters)
// // 							if err != nil {
// // 								return fmt.Errorf("Unable to get unsplash random photo using global config \n%v", err)
// // 							}
// // 							randomPhotoResponse.Urls.Regular = ""
// // 						}
// // 					}
// // 					// check if random photo parameters exists in global config

// // 				} else {
// // 					// do something

// // 					// lReadMeConfig, err := getLocalReadMeConfig()
// // 					// if err != nil {
// // 					// 	return fmt.Errorf("Unable to load local readme config \n%v", err)
// // 					// }
// // 					// lReadMeConfig.Logo.URL =

// // 				}
// // 			}
// // 		}
// // 	}

// // 	return nil
// // }

// // func updateLocalUnsplashUnsplashRandomPhotoResponse(response UnsplashRandomPhotoResponse) error {
// // 	luc, err := getLocalUnplashConfig()
// // 	if err != nil {
// // 		return fmt.Errorf("Unable to get local Unsplash config \n%v", err)
// // 	}

// // 	luc.UnsplashRandomPhotoResponse = response
// // 	err = saveConfig(luc, "clencli/unsplash.yaml")
// // 	if err != nil {
// // 		return fmt.Errorf("Unable to save local Unsplash config \n%v", err)
// // 	}

// // 	return err
// // }

// func updateLocalReadmeLogoURL(params UnsplashRandomPhotoParameters, response UnsplashRandomPhotoResponse) error {
// 	// save url to readme
// 	lr, err := getLocalReadMeConfig()
// 	if err != nil {
// 		return fmt.Errorf("Unable to read local config: readme.yaml \n%v", err)
// 	}

// 	lr.Logo.URL = GetPhotoURLBySize(params, response)
// 	err = saveConfig(lr, "clencli/readme.yaml")
// 	if err != nil {
// 		return fmt.Errorf("Unable to save local config: readme.yaml \n%v", err)
// 	}

// 	return err
// }
