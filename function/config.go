package function

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// getConfig returns a viper instance for the config YAML file
func getConfig(name string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath("clencli")
	config.SetConfigName(name)
	config.SetConfigType("yaml")
	config.SetConfigPermissions(os.ModePerm)

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %v", err)
	}

	return config
}

func getHLD() *viper.Viper {
	return getConfig("hld")
}

func getReadMeConfig() *viper.Viper {
	return getConfig("readme")
}

// ReadMe struct for readme.yaml config
type ReadMe struct {
	Logo struct {
		Label string `yaml:"label"`
		Theme string `yaml:"theme"`
		URL   string `yaml:"url"`
	} `yaml:"logo,omitempty"`
	Shields struct {
		Badges []struct {
			Description string `yaml:"description"`
			Image string `yaml:"image"`
			URL         string `yaml:"url"`
		} `yaml:"badges"`
	} `yaml:"shields,omitempty"`
	App struct {
		Name     string `yaml:"name"`
		Function string `yaml:"function"`
		ID       string `yaml:"id"`
	} `yaml:"app,omitempty"`
	Screenshots []struct {
		Caption string `yaml:"caption"`
		Label   string `yaml:"label"`
		URL     string `yaml:"url"`
	} `yaml:"screenshots,omitempty"`
	Usage         string `yaml:"usage"`
	Prerequisites []struct {
		Description string `yaml:"description"`
		Name        string `yaml:"name"`
		URL         string `yaml:"url"`
	} `yaml:"prerequisites,omitempty"`
	Installing   string   `yaml:"installing,omitempty"`
	Testing      string   `yaml:"testing,omitempty"`
	Deployment   string   `yaml:"deployment,omitempty"`
	Include      []string `yaml:"include,omitempty"`
	Contributors []struct {
		Name  string `yaml:"name"`
		Role  string `yaml:"role"`
		Email string `yaml:"email"`
	} `yaml:"contributors,omitempty"`
	Acknowledgments []struct {
		Name string `yaml:"name"`
		Role string `yaml:"role"`
	} `yaml:"acknowledgments,omitempty"`
	References []struct {
		Description string `yaml:"description"`
		Name        string `yaml:"name"`
		URL         string `yaml:"url"`
	} `yaml:"references,omitempty"`
	License   string `yaml:"license,omitempty"`
	Copyright string `yaml:"copyright,omitempty"`
}

// UpdateReadMe updates local configuration file with global configuration values
func UpdateReadMe() {
	readmeConfig := getReadMeConfig()

	localReadMe := ReadMe{}
	err := readmeConfig.Unmarshal(&localReadMe)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Get values from global config
	url := viper.GetString("readme.logo.url")
	theme := viper.GetString("readme.logo.theme")
	license := viper.GetString("readme.license")
	copyright := viper.GetString("readme.copyright")

	// Give logo.url precedence over logo.theme
	if len(url) > 0 {
		localReadMe.Logo.URL = url

		// if readme.logo.url is set in config, should take precendece over theme
		localReadMe.Logo.Theme = ""
	} else if len(theme) > 0 {
		localReadMe.Logo.Theme = theme

		// replaces the current logo.url by Unsplash's random photo URL based on theme's value
		unsplash := GetRandomPhotoDefaults(theme)
		localReadMe.Logo.URL = unsplash.Urls.Regular
	}

	if len(license) > 0 {
		localReadMe.License = license
	}

	if len(copyright) > 0 {
		localReadMe.Copyright = copyright
	}

	// Marshal back into yaml
	d, err := yaml.Marshal(&localReadMe)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("clencli/readme.yaml", d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}
