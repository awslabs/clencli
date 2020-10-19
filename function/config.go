package function

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// getConfig returns a viper instance for the config YAML file
func getConfig(name string) (*viper.Viper, error) {
	c := viper.New()
	c.AddConfigPath("clencli")
	c.SetConfigName(name)
	c.SetConfigType("yaml")
	c.SetConfigPermissions(os.ModePerm)

	err := c.ReadInConfig() // Find and read the c file
	if err != nil {         // Handle errors reading the c file
		return c, fmt.Errorf("Unable to read c "+name+" via Viper"+"\n%v", err)
	}

	return c, nil
}

func getHLD() (*viper.Viper, error) {
	return getConfig("hld")
}

func getReadMeConfig() (*viper.Viper, error) {
	return getConfig("readme")
}

// GetGlobalConfig returns a GlobalConfig struct from the global config
func GetGlobalConfig() (GlobalConfig, error) {
	v := viper.GetViper()
	c := GlobalConfig{}

	err := v.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("Unable to Unmarshall GlobalConfig struct \n%v", err)
	}

	return c, nil
}

// GetLocalReadMeConfig unmarshall local readme.yaml return as ReadMe struct
func GetLocalReadMeConfig() (ReadMe, error) {
	r := ReadMe{}
	c, err := getReadMeConfig()
	if err != nil {
		return r, fmt.Errorf("Unable to get readme config with Viper \n%v", err)
	}

	err = c.Unmarshal(&r)
	if err != nil {
		return r, fmt.Errorf("Unable to Unmarshall ReadMe struct \n%v", err)
	}

	return r, err
}

// GetGlobalReadMeConfig return the ReadMe section as struct from the global config
func GetGlobalReadMeConfig() (ReadMe, error) {
	c, err := GetGlobalConfig()
	r := c.ReadMe
	if err != nil {
		return r, fmt.Errorf("Unable to get global config \n%v", err)
	}

	return r, err
}

// MarshallAndSaveReadMe receive a readme struct and saves it back as file
func MarshallAndSaveReadMe(readme ReadMe) error {
	// Marshal back into yaml
	r, err := yaml.Marshal(&readme)
	if err != nil {
		return fmt.Errorf("Unable to Marshall ReadMe struct %v", err)
	}

	err = ioutil.WriteFile("clencli/readme.yaml", r, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to update clencli/readme.yaml file %v", err)
	}

	return nil
}

// UpdateReadMe updates local config with global config
func UpdateReadMe() error {
	g, err := GetGlobalReadMeConfig()
	if err != nil {
		return fmt.Errorf("Unable to get global readme config \n%v", err)
	}

	l, err := GetLocalReadMeConfig()
	if err != nil {
		return fmt.Errorf("Unable to get local readme config \n%v", err)
	}

	updated := false

	if g.Logo.Theme != "" {
		l.Logo.Theme = g.Logo.Theme
		updated = true
	}

	if g.Logo.URL != "" {
		l.Logo.URL = g.Logo.URL
		updated = true
	}

	if g.License != "" {
		l.License = g.License
		updated = true
	}

	if g.Copyright != "" {
		l.Copyright = g.Copyright
		updated = true
	}

	if updated {
		err = MarshallAndSaveReadMe(l)
		if err != nil {
			return fmt.Errorf("Unable to update cleancli/readme.yaml  \n%v", err)
		}
	}
	return nil
}

// UpdateReadMeLogoURL fetches random photo based readme.logo.theme from config
func UpdateReadMeLogoURL() error {
	l, err := GetLocalReadMeConfig()
	if err != nil {
		return fmt.Errorf("Unable to get local readme config \n%v", err)
	}

	if l.Logo.Theme != "" {
		ru, err := GetRandomPhotoDefaults(l.Logo.Theme)
		if err != nil {
			return fmt.Errorf("Unexpected error while getting random photo from Unsplash with default values \n%v", err)
		}
		l.Logo.URL = ru.Urls.Regular
		err = MarshallAndSaveReadMe(l)
		if err != nil {
			return fmt.Errorf("Unable to update cleancli/readme.yaml  \n%v", err)
		}
	}

	return nil
}
