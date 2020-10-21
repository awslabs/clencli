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

import controller "github.com/awslabs/clencli/cobra/controller"

var renderCmd = controller.RenderCmd()

func init() {
	rootCmd.AddCommand(renderCmd)
	renderCmd.Flags().StringP("name", "n", "readme", "Template name to be rendered")
	renderCmd.Flags().StringP("unsplash-profile", "u", "default", "Unsplash profile ")
}