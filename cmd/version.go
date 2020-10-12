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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show clencli version",
	Long:  `Returns the clencli tree's version string. It is either the commit hash and date at the time of the build or, when possible, a release tag like "clencli1.0".`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the version defined in the VERSION file
		version, err := ioutil.ReadFile("VERSION")
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR][cmd/version]: %v\n", err)
		} else {
			fmt.Printf("CLENCLI %s\n", version)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
