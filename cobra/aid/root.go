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
	"github.com/sirupsen/logrus"
)

// GetAppInfo return information about clencli settings
func GetAppInfo() model.App {
	var err error

	configDir, err := os.UserConfigDir()
	if err != nil {
		logrus.Fatalln("unable to read user configuration directory\n%v", err)
	}

	sep := string(os.PathSeparator)

	var app model.App
	app.Name = "clencli"
	app.ConfigurationsDir = configDir + sep + app.Name
	app.ConfigurationsName = "configurations"
	app.ConfigurationsType = "yaml"
	app.ConfigurationsPath = app.ConfigurationsDir + sep + app.ConfigurationsName + "." + app.ConfigurationsType
	app.ConfigurationsPermissions = os.ModePerm
	app.CredentialsName = "credentials"
	app.CredentialsType = "yaml"
	app.CredentialsPath = app.ConfigurationsDir + sep + app.CredentialsName + "." + app.CredentialsType
	app.CredentialsPermissions = os.ModePerm
	app.LogsDir = app.ConfigurationsDir
	app.LogsName = "logs"
	app.LogsType = "json"
	app.LogsPath = app.LogsDir + sep + app.LogsName + "." + app.LogsType
	app.LogsPermissions = os.ModePerm

	app.WorkingDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("Unable to detect the current directory\n%v", err)
		os.Exit(1)
	}

	return app
}

// SetupLoggingOutput set logrun output file
func SetupLoggingOutput(path string) error {
	if path == "" {
		app := GetAppInfo()
		path = app.LogsPath
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to open log file\n%v", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(file)

	return nil
}

// SetupLoggingLevel set logrus level
func SetupLoggingLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("unable to set log level\n%v", err)
	}

	logrus.SetLevel(lvl)
	return nil
}
