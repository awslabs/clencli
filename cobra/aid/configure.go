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

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ConfigDirExist TODO
func ConfigDirExist() bool {
	return cau.DirOrFileExists(GetAppInfo().ConfigurationsDir)
}

// ConfigurationsFileExist does TODO
func ConfigurationsFileExist() bool {
	return cau.DirOrFileExists(GetAppInfo().ConfigurationsPath)
}

// CreateConfigDir TODO
func CreateConfigDir() bool {
	return cau.CreateDir(GetAppInfo().ConfigurationsDir)
}

// CredentialsFileExist does TODO
func CredentialsFileExist() bool {
	return cau.DirOrFileExists(GetAppInfo().CredentialsPath)
}

// ReadConfig TODO
func ReadConfig(name string) (*viper.Viper, error) {
	v := viper.New()
	app := GetAppInfo()

	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AddConfigPath(app.ConfigurationsDir)

	err := v.ReadInConfig()
	if err != nil {
		return v, fmt.Errorf("Error when trying to read local configurations \n%s", err)
	}
	return v, err

}

// WriteInterfaceToFile does TODO
func WriteInterfaceToFile(in interface{}, path string) error {
	b, err := yaml.Marshal(&in)
	if err != nil {
		_, ok := err.(*json.UnsupportedTypeError)
		if ok {
			return fmt.Errorf("Tried to marshal an invalid Type")
		}
	}

	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to update: %s \n%v", path, err)
	}

	return err
}
