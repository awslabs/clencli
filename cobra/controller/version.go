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
	"fmt"
	"runtime"

	"github.com/awslabs/clencli/box"
	"github.com/awslabs/clencli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func versionRun(cmd *cobra.Command, args []string) error {
	// Get the version defined in the VERSION file
	version, found := box.Get("/VERSION")
	if !found {
		logrus.Errorf("unable to find VERSION file under box/resources")
		return fmt.Errorf("version not available")
	}

	goOS := runtime.GOOS
	goVersion := runtime.Version()
	goArch := runtime.GOARCH

	cmd.Printf("CLENCLI v%s %s %s %s\n", version, goVersion, goOS, goArch)

	return nil
}

// VersionCmd command to display CLENCLI current version
func VersionCmd() *cobra.Command {
	man := helper.GetManual("version")

	return &cobra.Command{
		Use:   man.Use,
		Short: man.Short,
		Long:  man.Long,
		RunE:  versionRun,
	}
}
