package function

import (
	"fmt"
	"log"

	"github.com/awslabs/clencli/box"
	yaml "gopkg.in/yaml.v2"
)

// Help mapping the help
type Help struct {
	Usage string `yaml:"usage"`
}

// GetHelp retrieve help struct
func GetHelp(help string) Help {
	var h Help
	helpBlob, status := box.Get("/help/" + help + ".yaml")
	if status {
		err := yaml.Unmarshal(helpBlob, &h)
		if err != nil {
			fmt.Println("Not able to decode YAML file, error:", err)
		}
	} else {
		log.Fatal("Not able to read help from box")
	}

	return h
}
