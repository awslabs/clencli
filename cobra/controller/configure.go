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
	"gopkg.in/yaml.v2"

	cau "github.com/awslabs/clencli/cauldron"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
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
	p, _ := cmd.Flags().GetString("profile")
	if p == "" {
		p = "default"
	}

	log.Info("Started to configure profile:" + p)

	// credentials
	err := setupCredentials(p)
	if err != nil {
		return fmt.Errorf("Unexpected error during credentials setup")
	}

	// configuratons
	err = setupConfig(p)
	if err != nil {
		return fmt.Errorf("Unexpected error during config setup")
	}

	log.Info("Finished to configure profile:" + p)

	return nil
}

func setupCredentials(p string) error {
	log.Info("Started to setup credentials for profile:" + p)
	log.Info("Finished to setup credentials for profile:" + p)
	return nil
}

func setupConfig(p string) error {
	log.Info("Started to setup configuration for profile:" + p)

	fmt.Println("Configuring named profile:" + p)

	var config model.Config
	// load current configuration

	var profile model.Profile
	profile.Name = p

	// else create new profile if config directory doesn't exist
	status, err := createConfigDir()
	if status {
		// enabling profile by default
		profile.Enabled = true

		// .. create two files: credentials.yaml and config.yaml

		// configure unsplash
		answer, err := getUserInput("Would you like to setup Unsplash Random Photo Parameters? (Y/y)")
		// answer := "Y"
		if answer == "Y" || answer == "y" {
			profile.Unsplash, err = configureUnsplash()
			if err != nil {
				return fmt.Errorf("Unable to configure Unsplash")
			}
		}
		if err != nil {
			return fmt.Errorf("Unable to get answer if user wants to setup Unsplash Random Photo Parameters")
		}

	}

	if err != nil {

	}

	// if profile == "" {

	// }

	// if profile is given,
	// .. check if profile already exits
	// .. if exists, load it and
	// .. update credentials with Credentials struct
	// .. update config with Config struct
	// .. else create new profile with credentials and config struct

	config.Profiles = append(config.Profiles, profile)
	err = saveConfig(config)
	if err != nil {
		return fmt.Errorf("Unable to save config during setup \n%v", err)
	}

	log.Info("Finished to setup configuration for profile:" + p)
	return err
}

func configureUnsplash() (model.Unsplash, error) {
	var unsplash model.Unsplash

	fmt.Println("=== Unsplash Random Photo Parameters")

	unsplash, err := updateUnsplashFromUserInput()
	if err != nil {
		return unsplash, fmt.Errorf("Unable to update Unsplash config from user input \n%v", err)
	}

	fmt.Println("=== Unsplash Random Photo Parameters configured successfully!")

	return unsplash, err
}

func getConfigDirPath() (string, error) {
	var path string
	home, err := homedir.Dir()
	if err != nil {
		return path, fmt.Errorf("Unable to get home directory path \n%v", err)
	}
	path = home + "/.clencli"

	return path, err
}

func createConfigDir() (bool, error) {
	status := false
	path, err := getConfigDirPath()
	if err != nil {
		return status, fmt.Errorf("Unable to get global config dir path \n%v", err)
	}

	if !cau.DirOrFileExists(path) {
		fmt.Println("CLENCLI configuration directory not found")

		if cau.CreateDir(path) {
			fmt.Println("CLENCLI configuration directory created at " + path)
			status = true
		}
	}

	return status, nil
}

func getUserInput(text string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text + ":")
	input, err := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		return input, fmt.Errorf("Unable to read user input \n%v", err)
	}
	return input, err
}

func updateUnsplashFromUserInput() (model.Unsplash, error) {

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

func saveConfig(config model.Config) error {
	path, err := getConfigDirPath()
	if err != nil {
		return fmt.Errorf("Unable to get CLENCLI configuration directory path \n%v", path)
	}
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

// func getLocalConfig(name string) (*viper.Viper, error) {
// 	lc := viper.New()

// 	lc.SetConfigName(name)      // name of config file (without extension)
// 	lc.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
// 	lc.AddConfigPath("clencli") // path to look for the config file in

// 	err := lc.ReadInConfig() // Find and read the config file
// 	if err != nil {          // Handle errors reading the config file
// 		return lc, fmt.Errorf("Error when trying to read local config \n%s", err)
// 	}
// 	return lc, err
// }

// // mergeConfig merges a new configuration with an existing config.
// func mergeConfig(local *viper.Viper, source ReadMe) error {
// 	s, _ := json.Marshal(source)
// 	return local.MergeConfig(bytes.NewBuffer(s))
// }

// // getGlobalConfig returns a GlobalConfig struct from the global config
// func getGlobalConfig() (GlobalConfig, error) {
// 	v := viper.GetViper()
// 	c := GlobalConfig{}

// 	err := v.Unmarshal(&c)
// 	if err != nil {
// 		return c, fmt.Errorf("Unable to Unmarshall GlobalConfig struct \n%v", err)
// 	}

// 	return c, nil
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
