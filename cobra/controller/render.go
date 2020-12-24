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

package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/awslabs/clencli/helper"
	gomplateV3 "github.com/hairyhenderson/gomplate/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var renderValidArgs = []string{"template"}

// RenderCmd command to render templates
func RenderCmd() *cobra.Command {
	man := helper.GetManual("render")
	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		ValidArgs: renderValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   renderPreRun,
		RunE:      renderRun,
	}

	cmd.Flags().StringP("name", "n", "readme", "Database file name of the template to be rendered (it must be under clencli/ directory.")

	return cmd
}

func renderPreRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command render pre-run")

	if err := helper.ValidateCmdArgs(cmd, args, "render"); err != nil {
		return err
	}

	if err := helper.ValidateCmdArgAndFlag(cmd, args, "render", "template", "name"); err != nil {
		return err
	}

	name, _ := cmd.Flags().GetString("name")

	if !helper.FileExists("clencli/" + name + ".yaml") {
		return errors.New("missing database at clencli/" + name + ".yaml")
	}

	if !helper.FileExists("clencli/" + name + ".tmpl") {
		return errors.New("missing template at clencli/" + name + ".tmpl")
	}

	logrus.Traceln("end: command render pre-run")
	return nil
}

func renderRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command render run")

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Errorf("error: unable to render template "+name+"\n%v", err)
		return fmt.Errorf("error: unable to render template "+name+"\n%v", err)
	}

	// TODO: update readme and logo url based on global configuration

	if err := initGomplate(name); err != nil {
		logrus.Errorf("Unexpected error: %v", err)
		return fmt.Errorf("error: unable to render template "+name+"\n%v", err)
	}

	cmd.Println("Template " + name + ".tmpl rendered as " + strings.ToUpper(name) + ".md.")

	logrus.Traceln("end: command render run")
	return nil
}

func initGomplate(name string) error {
	var inputFiles = []string{}
	var outputFiles = []string{}

	if helper.FileExists("clencli/" + name + ".tmpl") {
		inputFiles = append(inputFiles, "clencli/"+name+".tmpl")
		outputFiles = append(outputFiles, strings.ToUpper(name)+".md")
	}

	var config gomplateV3.Config
	config.InputFiles = inputFiles
	config.OutputFiles = outputFiles

	dataSources := []string{}
	if helper.FileExists("clencli/" + name + ".yaml") {
		dataSources = append(dataSources, "db=./clencli/"+name+".yaml")
	}

	config.DataSources = dataSources

	err := gomplateV3.RunTemplates(&config)
	if err != nil {
		logrus.Fatalf("Gomplate.RunTemplates() failed with %s\n", err)
	}

	return err
}

func writeInputs() error {
	variables, err := os.Open("variables.tf")
	if err != nil {
		logrus.Fatal(err)
	}
	defer variables.Close()

	// create INPUTS.md
	inputs, err := os.OpenFile("INPUTS.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Println(err)
	}
	defer inputs.Close()

	if _, err := inputs.WriteString("| Name | Description | Type | Default | Required |\n|------|-------------|:----:|:-----:|:-----:|\n"); err != nil {
		logrus.Println(err)
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
				out, found := helper.GetStringBetweenDoubleQuotes(line)
				if found {
					varName = out
				}

			}

			if strings.Contains(line, "type") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
				if slc[0] == "type" {
					varType = slc[1]
					if strings.Contains(varType, "({") {
						slc = helper.GetStringTrimmed(varType, "({")
						varType = slc[0]
					}
				}
			}

			if strings.Contains(line, "description") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
				if slc[0] == "description" {
					out, found := helper.GetStringBetweenDoubleQuotes(slc[1])
					if found {
						varDescription = out
					}
				}
			}

			if strings.Contains(line, "default") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
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
						logrus.Println(err)
					}
					varName, varType, varDescription, varDefault, varRequired = "", "", "", "", "no"
				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		logrus.Fatal(err)

	}
	return err
}

func writeOutputs() error {
	outputs, err := os.Open("outputs.tf")
	if err != nil {
		logrus.Fatal(err)
	}
	defer outputs.Close()

	// create INPUTS.md
	outs, err := os.OpenFile("OUTPUTS.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Println(err)
	}
	defer outs.Close()

	if _, err := outs.WriteString("| Name | Description |\n|------|-------------|\n"); err != nil {
		logrus.Println(err)
	}

	var outName, outDescription string

	scanner := bufio.NewScanner(outputs)
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if len(line) > 0 {
			if strings.Contains(line, "output") && strings.Contains(line, "{") {
				out, found := helper.GetStringBetweenDoubleQuotes(line)
				if found {
					outName = out
				}
			}

			if strings.Contains(line, "description") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
				if slc[0] == "description" {
					out, found := helper.GetStringBetweenDoubleQuotes(slc[1])
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
						logrus.Println(err)
					}
					outName, outDescription = "", ""
				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		logrus.Fatal(err)

	}
	return err
}
