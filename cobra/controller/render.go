package controller

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	cau "github.com/awslabs/clencli/cauldron"
	function "github.com/awslabs/clencli/cauldron"
	gomplateV3 "github.com/hairyhenderson/gomplate/v3"
	"github.com/spf13/cobra"
)

var renderValidArgs = []string{"template"}

// RenderCmd command to render templates
func RenderCmd() *cobra.Command {
	man := cau.GetManual("render")
	return &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		ValidArgs: renderValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   renderPreRun,
		RunE:      renderRun,
	}
}

func renderPreRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("one the following arguments are required: %s", renderValidArgs)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return errors.New("required flag name not set")
	}

	if !function.FileExists("clencli/" + name + ".yaml") {
		return errors.New("Missing database at clencli/" + name + ".yaml")
	}

	if !function.FileExists("clencli/" + name + ".tmpl") {
		return errors.New("Missing template at clencli/" + name + ".tmpl")
	}

	return nil
}

func renderRun(cmd *cobra.Command, args []string) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return fmt.Errorf("Unable to render template "+name+"\n%v", err)
	}

	// err = cau.UpdateReadMe()
	// if err != nil {
	// 	return fmt.Errorf("Unable to update local config with global config values \n%v", err)
	// }

	// err = cau.UpdateReadMeLogoURL()
	// if err != nil {
	// 	return fmt.Errorf("Unable to update local config with new URL from Unsplash \n%v", err)
	// }

	err = initGomplate(name)
	if err == nil {
		fmt.Println("Template " + name + ".tmpl rendered as " + strings.ToUpper(name) + ".md.")
	} else {
		log.Fatalf("Unexpected error: %v", err)
	}

	return nil
}

func initGomplate(name string) error {
	var inputFiles = []string{}
	var outputFiles = []string{}

	if function.FileExists("clencli/" + name + ".tmpl") {
		inputFiles = append(inputFiles, "clencli/"+name+".tmpl")
		outputFiles = append(outputFiles, strings.ToUpper(name)+".md")
	}

	var config gomplateV3.Config
	config.InputFiles = inputFiles
	config.OutputFiles = outputFiles

	dataSources := []string{}
	if function.FileExists("clencli/" + name + ".yaml") {
		dataSources = append(dataSources, "db=./clencli/"+name+".yaml")
	}

	config.DataSources = dataSources

	err := gomplateV3.RunTemplates(&config)
	if err != nil {
		log.Fatalf("Gomplate.RunTemplates() failed with %s\n", err)
	}

	return err
}

func writeInputs() error {
	variables, err := os.Open("variables.tf")
	if err != nil {
		log.Fatal(err)
	}
	defer variables.Close()

	// create INPUTS.md
	inputs, err := os.OpenFile("INPUTS.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer inputs.Close()

	if _, err := inputs.WriteString("| Name | Description | Type | Default | Required |\n|------|-------------|:----:|:-----:|:-----:|\n"); err != nil {
		log.Println(err)
	}

	var varName, varType, varDescription, varDefault string
	varRequired := "no"

	// startBlock := false
	scanner := bufio.NewScanner(variables)
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if len(line) > 0 {
			if strings.Contains(line, "variable") && strings.Contains(line, "{") {
				out, found := function.GetStringBetweenDoubleQuotes(line)
				if found {
					varName = out
				}

			}

			if strings.Contains(line, "type") && strings.Contains(line, "=") {
				slc := function.GetStringTrimmed(line, "=")
				if slc[0] == "type" {
					varType = slc[1]
					if strings.Contains(varType, "({") {
						slc = function.GetStringTrimmed(varType, "({")
						varType = slc[0]
					}
				}
			}

			if strings.Contains(line, "description") && strings.Contains(line, "=") {
				slc := function.GetStringTrimmed(line, "=")
				if slc[0] == "description" {
					out, found := function.GetStringBetweenDoubleQuotes(slc[1])
					if found {
						varDescription = out
					}
				}
			}

			if strings.Contains(line, "default") && strings.Contains(line, "=") {
				slc := function.GetStringTrimmed(line, "=")
				if slc[0] == "default" {
					varDefault = slc[1]
					if strings.Contains(varDefault, "{") {
						varDefault = "<map>"
					}
				}
			}

			// end of the variable declaration
			if strings.Contains(line, "}") && len(line) == 1 {
				if len(varName) > 0 && len(varType) > 0 && len(varDescription) > 0 {

					var result string
					if len(varDefault) == 0 {
						varRequired = "yes"
						result = fmt.Sprintf("| %s | %s | %s | %s | %s |\n", varName, varDescription, varType, varDefault, varRequired)
					} else {
						result = fmt.Sprintf("| %s | %s | %s | `%s` | %s |\n", varName, varDescription, varType, varDefault, varRequired)
					}

					if _, err := inputs.WriteString(result); err != nil {
						log.Println(err)
					}
					varName, varType, varDescription, varDefault, varRequired = "", "", "", "", "no"
				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return err
}

func writeOutputs() error {
	outputs, err := os.Open("outputs.tf")
	if err != nil {
		log.Fatal(err)
	}
	defer outputs.Close()

	// create INPUTS.md
	outs, err := os.OpenFile("OUTPUTS.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer outs.Close()

	if _, err := outs.WriteString("| Name | Description |\n|------|-------------|\n"); err != nil {
		log.Println(err)
	}

	var outName, outDescription string

	scanner := bufio.NewScanner(outputs)
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if len(line) > 0 {
			if strings.Contains(line, "output") && strings.Contains(line, "{") {
				out, found := function.GetStringBetweenDoubleQuotes(line)
				if found {
					outName = out
				}
			}

			if strings.Contains(line, "description") && strings.Contains(line, "=") {
				slc := function.GetStringTrimmed(line, "=")
				if slc[0] == "description" {
					out, found := function.GetStringBetweenDoubleQuotes(slc[1])
					if found {
						outDescription = out
					}
				}
			}

			// end of the output declaration
			if strings.Contains(line, "}") && len(line) == 1 {
				if len(outName) > 0 && len(outDescription) > 0 {

					result := fmt.Sprintf("| %s | %s | |\n", outName, outDescription)

					if _, err := outs.WriteString(result); err != nil {
						log.Println(err)
					}
					outName, outDescription = "", ""
				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return err
}
