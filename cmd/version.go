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

// Package cmd contains Cobra commands
package cmd

import (
	"fmt"
	"log"

	"github.com/awslabs/clencli/box"
	"github.com/spf13/cobra"
)

// VersionCmd command to display CLENCLI current version
func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show clencli version",
		Long:  `Returns the clencli tree's version string. It is either the commit hash and date at the time of the build or, when possible, a release tag like "clencli1.0".`,
		Run: func(cmd *cobra.Command, args []string) {
			// Get the version defined in the VERSION file
			version, status := box.Get("/VERSION")
			if status {
				fmt.Printf("CLENCLI v%s", version)
			} else {
				log.Fatal("Version not available")
			}
		},
	}
}

// versionCmd represents the version command
var versionCmd = VersionCmd()

func init() {
	rootCmd.AddCommand(versionCmd)
}
