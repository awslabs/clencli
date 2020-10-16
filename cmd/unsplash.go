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
	"github.com/awslabs/clencli/function"
	"github.com/spf13/cobra"
)

// UnsplashCmd command to download photos from Unsplash.com
func UnsplashCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "unsplash",
		Short: "Downloads pictures from Unsplash.com",
		Long:  `Retrieve a single random photo, given optional filters.`,
		Run: func(cmd *cobra.Command, args []string) {

			query, _ := cmd.Flags().GetString("query")
			collections, _ := cmd.Flags().GetString("collections")
			featured, _ := cmd.Flags().GetString("featured")
			username, _ := cmd.Flags().GetString("username")
			orientation, _ := cmd.Flags().GetString("orientation")
			filter, _ := cmd.Flags().GetString("filter")
			size, _ := cmd.Flags().GetString("size")

			unsplash := function.GetRandomPhoto(
				query,
				collections,
				featured,
				username,
				orientation,
				filter)
			// size)

			if size == "thumb" || size == "all" {
				function.DownloadPhoto(unsplash.Urls.Thumb, "thumb", query)
			}
			if size == "small" || size == "all" {
				function.DownloadPhoto(unsplash.Urls.Small, "small", query)
			}
			if size == "regular" || size == "all" {
				function.DownloadPhoto(unsplash.Urls.Regular, "regular", query)
			}
			if size == "full" || size == "all" {
				function.DownloadPhoto(unsplash.Urls.Full, "full", query)
			}
			if size == "raw" || size == "all" {
				function.DownloadPhoto(unsplash.Urls.Raw, "raw", query)
			}
		},
	}
}

// unsplashCmd represents the unsplash command
var unsplashCmd = UnsplashCmd()

func init() {
	rootCmd.AddCommand(unsplashCmd)

	unsplashCmd.Flags().StringP("collections", "c", "", "Public collection ID(‘s) to filter selection. If multiple, comma-separated")
	unsplashCmd.Flags().StringP("featured", "f", "", "Limit selection to featured photos.")
	unsplashCmd.Flags().StringP("username", "u", "", "Limit selection to a single user.")
	unsplashCmd.Flags().StringP("query", "q", "mountains", "Limit selection to photos matching a search term. (Deafult: mountains")
	unsplashCmd.Flags().StringP("orientation", "", "", "Filter by photo orientation. Valid values: landscape, portrait, squarish.")
	unsplashCmd.Flags().StringP("filter", "l", "low", "Limit results by content safety. Default: low. Valid values are low and high.")
	unsplashCmd.Flags().StringP("size", "s", "all", "Photos size. Valid values: all, thumb, small, regular, full, raw. Default: all")

}
