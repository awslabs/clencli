package function

import (
	"fmt"
	"log"

	"github.com/awslabs/clencli/box"
	yaml "gopkg.in/yaml.v2"
)

// Manual mapping the fields used by a Cobra command
type Manual struct {
	Use   string `yaml:"use"`
	Short string `yaml:"short"`
	Long  string `yaml:"long"`
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
		log.Fatal("Not able to read manual from box")
	}

	return man
}
