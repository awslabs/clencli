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

package controller

import (
	"fmt"
	"strings"

	model "github.com/awslabs/clencli/cobra/model"
	helper "github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

// UnsplashCmd command to download photos from Unsplash.com
func UnsplashCmd() *cobra.Command {
	man := helper.GetManual("unsplash")
	return &cobra.Command{
		Use:   man.Use,
		Short: man.Short,
		Long:  man.Long,
		RunE:  unsplashRun,
	}
}

func unsplashRun(cmd *cobra.Command, args []string) error {
	// params := getModelFromFlags(cmd)

	// unsplash, err := helper.GetRandomPhoto(
	// 	query,
	// 	collections,
	// 	featured,
	// 	username,
	// 	orientation,
	// 	filter)
	// // size)
	// if err != nil {
	// 	return fmt.Errorf("Unexpected error while getting random photo from Unsplash \n%v", err)
	// }

	// if size == "thumb" || size == "all" {
	// 	helper.DownloadPhoto(unsplash.Urls.Thumb, "thumb", query)
	// }
	// if size == "small" || size == "all" {
	// 	helper.DownloadPhoto(unsplash.Urls.Small, "small", query)
	// }
	// if size == "regular" || size == "all" {
	// 	helper.DownloadPhoto(unsplash.Urls.Regular, "regular", query)
	// }
	// if size == "full" || size == "all" {
	// 	helper.DownloadPhoto(unsplash.Urls.Full, "full", query)
	// }
	// if size == "raw" || size == "all" {
	// 	helper.DownloadPhoto(unsplash.Urls.Raw, "raw", query)
	// }
	return nil
}

func getModelFromFlags(cmd *cobra.Command) model.UnsplashRandomPhotoParameters {
	var params model.UnsplashRandomPhotoParameters

	params.Query, _ = cmd.Flags().GetString("query")
	params.Collections, _ = cmd.Flags().GetString("collections")
	params.Featured, _ = cmd.Flags().GetBool("featured")
	params.Username, _ = cmd.Flags().GetString("username")
	params.Orientation, _ = cmd.Flags().GetString("orientation")
	params.Filter, _ = cmd.Flags().GetString("filter")
	params.Size, _ = cmd.Flags().GetString("size")

	return params
}

// requestRandomPhotoDefaults retrieves a single random photo with default values.
// func requestRandomPhotoDefaults(query string) (UnsplashRandomPhotoResponse, error) {
// 	// landscape orientation is better for README files
// 	return requestRandomPhoto(query, "", "", "", "landscape", "low")
// }

// func hasCredentials(provider string, name string) (bool, error) {
// 	global, err := getGlobalConfig()
// 	if err != nil {
// 		return false, fmt.Errorf("Unable to get global config \n%v", err)
// 	}

// 	for _, cred := range global.Credentials {
// 		// verify if Unsplash credentials are set
// 		if cred.Provider == "unsplash" {
// 			if cred.AccessKey != "" && cred.SecretKey != "" {
// 				return cred, nil
// 			}
// 		}

// 	}

// }

// requestRandomPhoto retrieves a single random photo, given optional filters.
func requestRandomPhoto(params model.UnsplashRandomPhotoParameters) (model.UnsplashRandomPhotoResponse, error) {
	var response model.UnsplashRandomPhotoResponse

	if (model.UnsplashRandomPhotoParameters{} == params) {
		return response, fmt.Errorf("Unable to download Unsplash photo if all fields from query as empty")
	}

	// gc, err := getGlobalConfig()
	// if err != nil {
	// 	return response, fmt.Errorf("Unexpected error while getting global configuration \n%v", err)
	// }

	// // check if Credentials has unsplash credential
	// // return response, fmt.Errorf("No Unsplash credentials found in the global configuration \n%v", err)

	// // check if config has Unsplash credentials
	// for _, cred := range gc.Credentials {
	// 	if cred.Provider == "unsplash" {
	// 		if cred.AccessKey != "" && cred.SecretKey != "" {

	// 			url := buildURL(params, cred.Credential)

	// 			var client http.Client
	// 			resp, err := client.Get(url)
	// 			if err != nil {
	// 				return response, fmt.Errorf("Unexpected error while performing GET on Unsplash API \n%v", err)
	// 			}
	// 			defer resp.Body.Close()

	// 			if resp.StatusCode == http.StatusOK {
	// 				bodyBytes, err := ioutil.ReadAll(resp.Body)
	// 				if err != nil {
	// 					return response, fmt.Errorf("Unexpected error while reading Unsplash response \n%v", err)
	// 				}

	// 				json.Unmarshal(bodyBytes, &response)
	// 			}
	// 		}
	// 	}
	// }

	return response, nil
}

// func buildURL(params model.UnsplashRandomPhotoParameters, cred model.Credential) string {
// 	clientID := cred.AccessKey
// 	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s", clientID)

// 	if len(params.Collections) > 0 {
// 		url += fmt.Sprintf("&collections=%s", params.Collections)
// 	}

// 	if len(params.Query) > 0 {
// 		url += fmt.Sprintf("&query=%s", params.Query)
// 	}

// 	url += fmt.Sprintf("&featured=%t", params.Featured)

// 	if len(params.Username) > 0 {
// 		url += fmt.Sprintf("&username=%s", params.Username)
// 	}

// 	if len(params.Orientation) > 0 {
// 		url += fmt.Sprintf("&orientation=%s", params.Orientation)
// 	}

// 	if len(params.Filter) > 0 {
// 		url += fmt.Sprintf("&filter=%s", params.Filter)
// 	}

// 	return url
// }

// DownloadPhoto downloads a photo and saves into downloads/unsplash/ folder
// It creates the downloads/ folder if it doesn't exists
func DownloadPhoto(url string, size string, query string) error {

	// Get the photo identifier
	start := strings.Index(url, "photo")
	end := strings.Index(url, "?")

	// Create a Rune from the URL
	runes := []rune(url)

	// Generate the directory path
	dirPath := "downloads/unsplash/" + query

	// Generate the filename
	fileName := string(runes[start:end])
	fileName += "-" + size + ".jpg"

	return helper.DownloadFile(url, dirPath, fileName)
}

// GetPhotoURLBySize return the photo URL based on the given size
func GetPhotoURLBySize(p model.UnsplashRandomPhotoParameters, r model.UnsplashRandomPhotoResponse) string {
	switch p.Size {
	case "thumb":
		return r.Urls.Thumb
	case "small":
		return r.Urls.Small
	case "regular":
		return r.Urls.Regular
	case "full":
		return r.Urls.Full
	case "raw":
		return r.Urls.Raw
	default:
		return r.Urls.Small
	}

}
