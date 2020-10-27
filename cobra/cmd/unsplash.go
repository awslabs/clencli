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

import controller "github.com/awslabs/clencli/cobra/controller"

// unsplashCmd represents the unsplash command
var unsplashCmd = controller.UnsplashCmd()

func init() {
	rootCmd.AddCommand(unsplashCmd)

	unsplashCmd.Flags().StringP("collections", "c", "", "Public collection ID(‘s) to filter selection. If multiple, comma-separated")
	unsplashCmd.Flags().BoolP("featured", "f", false, "Limit selection to featured photos. Valid values: false, true.")
	unsplashCmd.Flags().StringP("filter", "l", "low", "Limit results by content safety. Default: low. Valid values are low and high.")
	unsplashCmd.Flags().StringP("orientation", "", "", "Filter by photo orientation. Valid values: landscape, portrait, squarish.")
	unsplashCmd.Flags().StringP("query", "q", "mountains", "Limit selection to photos matching a search term. (Deafult: mountains)")
	unsplashCmd.Flags().StringP("size", "s", "all", "Photos size. Valid values: all, thumb, small, regular, full, raw. Default: all")
	unsplashCmd.Flags().StringP("username", "u", "", "Limit selection to a single user.")

}
