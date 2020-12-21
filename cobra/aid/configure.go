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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/awslabs/clencli/helper"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ConfigurationDirectoryExist returns `true` if the configuration directory exist, `false` otherwise
func ConfigurationDirectoryExist() bool {
	return helper.DirOrFileExists(GetAppInfo().ConfigurationsDir)
}

// ConfigurationsFileExist returns `true` if the configuration file exist, `false` otherwise
func ConfigurationsFileExist() bool {
	return helper.DirOrFileExists(GetAppInfo().ConfigurationsPath)
}

// CreateConfigDir creates the configuration directory, returns `true` if the configuration directory exist, `false` otherwise
func CreateConfigDir() bool {
	dir := GetAppInfo().ConfigurationsDir
	fmt.Printf("CLENCLI configuration directory created at: \n%s\n", dir)
	return helper.CreateDir(dir)
}

// CredentialsFileExist returns `true` if the credentials file exist, `false` otherwise
func CredentialsFileExist() bool {
	return helper.DirOrFileExists(GetAppInfo().CredentialsPath)
}

// ReadConfig returns the viper instance of the given configuration `name`
func ReadConfig(name string) (*viper.Viper, error) {
	v := viper.New()
	app := GetAppInfo()

	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AddConfigPath(app.ConfigurationsDir)

	err := v.ReadInConfig()
	if err != nil {
		return v, fmt.Errorf("error when trying to read local configurations \n%s", err)
	}
	return v, err

}

// WriteInterfaceToFile write the given interface into a file
func WriteInterfaceToFile(in interface{}, path string) error {
	b, err := yaml.Marshal(&in)
	if err != nil {
		_, ok := err.(*json.UnsupportedTypeError)
		if ok {
			return fmt.Errorf("tried to marshal an invalid Type")
		}
	}

	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to update: %s \n%v", path, err)
	}

	return err
}
