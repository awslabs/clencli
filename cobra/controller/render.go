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
	"errors"
	"fmt"
	"strings"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/helper"
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

	if err := aid.BuildTemplate(name); err != nil {
		logrus.Errorf("Unexpected error: %v", err)
		return fmt.Errorf("error: unable to render template "+name+"\n%v", err)
	}

	cmd.Println("Template " + name + ".tmpl rendered as " + strings.ToUpper(name) + ".md.")

	logrus.Traceln("end: command render run")
	return nil
}
