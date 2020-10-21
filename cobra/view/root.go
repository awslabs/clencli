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

package view

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cau "github.com/awslabs/clencli/cauldron"
	controller "github.com/awslabs/clencli/cobra/controller"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgFile string
var profile string
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clencli.yaml)")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "default", "Use a specific profile from your config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		configDirPath := home + "/.clencli"
		viper.AddConfigPath(configDirPath) // global directory
		viper.AddConfigPath("clencli")     // local directory
		viper.SetConfigName("config")

		if cau.DirOrFileExists(configDirPath) {
			// If the file doesn't exist, create it or append to the file
			file, err := os.OpenFile(configDirPath+"/logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			log.SetFormatter(&log.JSONFormatter{})
			log.SetOutput(file)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}
