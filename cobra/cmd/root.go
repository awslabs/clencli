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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/helper"

	"github.com/spf13/viper"
)

var rootCmd = controller.RootCmd()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("verbosity", "v", "error", "Valid log level:panic,fatal,error,warn,info,debug,trace).")
	rootCmd.PersistentFlags().Bool("log", true, "Enable or disable logs (can be found at ./clencli/log.json). Log outputs will be redirected default output if disabled.")
	rootCmd.PersistentFlags().String("log-file-path", helper.BuildPath("clencli/log.json"), "Log file path. Requires log=true, ignored otherwise.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	app := aid.GetAppInfo()
	viper.AddConfigPath(app.ConfigurationsDir) // global directory
	viper.SetConfigName(app.ConfigurationsName)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}

	verbosity, err := rootCmd.Flags().GetString("verbosity")
	if err != nil {
		fmt.Printf("unable to read flag verbosity\n%v", err)
	}

	log, err := rootCmd.Flags().GetBool("log")
	if err != nil {
		fmt.Printf("unable to read flag err\n%v", err)
	}

	logFilePath, err := rootCmd.Flags().GetString("log-file-path")
	if err != nil {
		fmt.Printf("unable to read flag log-file-path\n%v", err)
	}

	if log && logFilePath != "" {
		if err := aid.SetupLoggingLevel(verbosity); err == nil {
			fmt.Printf("logging level: %s\n", verbosity)
		}

		if err := aid.SetupLoggingOutput(logFilePath); err == nil {
			fmt.Printf("logging path: %s\n", logFilePath)
		}
	}

}
