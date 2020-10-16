/*
Copyright © 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/awslabs/clencli/function"
	fun "github.com/awslabs/clencli/function"
	gomplateV3 "github.com/hairyhenderson/gomplate/v3"
	"github.com/spf13/cobra"
)

var renderValidArgs = []string{"template"}

func renderPreRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("one the following arguments are required: %s", renderValidArgs)
	}

	return nil
}

func renderRun(cmd *cobra.Command, args []string) error {
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

	err = initGomplate(name)
	if err == nil {
		fmt.Println("Template " + name + ".tmpl rendered as " + strings.ToUpper(name) + ".md.")
	} else {
		log.Fatalf("Unexpected error: %v", err)
	}

	return nil
}

// RenderCmd command to render templates
func RenderCmd() *cobra.Command {
	man := fun.GetManual("render")
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

// renderCmd represents the render command
var renderCmd = RenderCmd()

func init() {
	rootCmd.AddCommand(renderCmd)
	renderCmd.Flags().StringP("name", "n", "readme", "Template name to be rendered")
}

func initGomplate(name string) error {
	function.UpdateReadMe()

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
