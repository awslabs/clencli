package controller

import (
	"fmt"

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
		RunE:  runConfigureCmd,
	}
}

func runConfigureCmd(cmd *cobra.Command, args []string) error {
	return nil
}

// getLocalConfigV3 does TODO
func getLocalConfig(name string) (*viper.Viper, error) {
	lc := viper.New()

	lc.SetConfigName(name)      // name of config file (without extension)
	lc.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	lc.AddConfigPath("clencli") // path to look for the config file in

	err := lc.ReadInConfig() // Find and read the config file
	if err != nil {          // Handle errors reading the config file
		return lc, fmt.Errorf("Error when trying to read local config \n%s", err)
	}
	return lc, err
}

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
// func saveConfig(in interface{}, path string) error {
// 	b, err := yaml.Marshal(&in)
// 	if err != nil {
// 		_, ok := err.(*json.UnsupportedTypeError)
// 		if ok {
// 			return fmt.Errorf("Tried to Marshal Invalid Type")
// 		}
// 	}

// 	err = ioutil.WriteFile(path, b, os.ModePerm)
// 	if err != nil {
// 		return fmt.Errorf("Unable to update: %s \n%v", path, err)
// 	}

// 	return err
// }

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
// // 						if (RandomPhotoParameters{} != global.Config.Unsplash.RandomPhotoParameters) {
// // 							randomPhotoResponse, err := getUnsplashRandomPhoto(global.Config.Unsplash.RandomPhotoParameters)
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

// // func updateLocalUnsplashRandomPhotoResponse(response RandomPhotoResponse) error {
// // 	luc, err := getLocalUnplashConfig()
// // 	if err != nil {
// // 		return fmt.Errorf("Unable to get local Unsplash config \n%v", err)
// // 	}

// // 	luc.RandomPhotoResponse = response
// // 	err = saveConfig(luc, "clencli/unsplash.yaml")
// // 	if err != nil {
// // 		return fmt.Errorf("Unable to save local Unsplash config \n%v", err)
// // 	}

// // 	return err
// // }

// func updateLocalReadmeLogoURL(params RandomPhotoParameters, response RandomPhotoResponse) error {
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
