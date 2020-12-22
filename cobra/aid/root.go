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

package aid

import (
	"fmt"
	"os"

	"github.com/awslabs/clencli/cobra/model"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

// GetAppInfo return information about CLENCLI settings
func GetAppInfo() model.App {
	var err error
	var app model.App
	app.Name = "clencli"
	app.HomeDir = getHomeDir()
	app.ConfigurationsDir = app.HomeDir + "/" + "." + app.Name
	app.ConfigurationsName = "configurations"
	app.ConfigurationsType = "yaml"
	app.ConfigurationsPath = app.ConfigurationsDir + "/" + app.ConfigurationsName + "." + app.ConfigurationsType
	app.ConfigurationsPermissions = os.ModePerm
	app.CredentialsName = "credentials"
	app.CredentialsType = "yaml"
	app.CredentialsPath = app.ConfigurationsDir + "/" + app.CredentialsName + "." + app.CredentialsType
	app.CredentialsPermissions = os.ModePerm
	app.LogsDir = app.ConfigurationsDir
	app.LogsName = "logs"
	app.LogsType = "json"
	app.LogsPath = app.LogsDir + "/" + app.LogsName + "." + app.LogsType
	app.LogsPermissions = os.ModePerm
	app.WorkingDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("Unable to detect the current directory\n%v", err)
		os.Exit(1)
	}

	return app
}

func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Unable to detect the home directory\n%v", err)
		os.Exit(1)
	}
	return home
}

// SetupLogging set up logging for the application
func SetupLogging(level string) error {
	app := GetAppInfo()
	file, err := os.OpenFile(app.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Unexpected error while opening log file\n%v", err)
		os.Exit(1)
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("Unexpected error while parsing log level\n%v", err)
		os.Exit(1)
	}

	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(file)

	return nil

}
