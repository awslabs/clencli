package helper

import (
	"fmt"

	"github.com/awslabs/clencli/box"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Manual mapping the fields used by a Cobra command
type Manual struct {
	Use     string `yaml:"use"`
	Example string `yaml:"example"`
	Short   string `yaml:"short"`
	Long    string `yaml:"long"`
}

// GetManual retrieve information about the given command
func GetManual(command string) Manual {
	var man Manual
	manualBlob, status := box.Get("/manual/" + command + ".yaml")
	if status {
		err := yaml.Unmarshal(manualBlob, &man)
		if err != nil {
			fmt.Println("Not able to decode YAML file, error:", err)
		}
	} else {
		logrus.Fatal("Not able to read manual from box")
	}

	return man
}
