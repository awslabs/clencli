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

package model

// Init struct to initalize things: projects, etc
type Init struct {
	Types []struct {
		Type    string `yaml:"type"`
		Name    string `yaml:"name"`
		Enabled bool   `yaml:"enabled"`
		Files   []struct {
			File struct {
				Path  string `yaml:"path"`
				Src   string `yaml:"src"`
				Dest  string `yaml:"dest"`
				State string `yaml:"state"`
			} `yaml:"file,omitempty"`
		} `yaml:"files"`
	} `yaml:"types"`
}
