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

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/dao"
	"github.com/awslabs/clencli/cobra/model"
	"github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

var unsplashPhotoSizes = []string{"all", "thumb", "small", "regular", "full", "raw"}

// UnsplashCmd command to download photos from Unsplash.com
func UnsplashCmd() *cobra.Command {
	man := helper.GetManual("unsplash")
	cmd := &cobra.Command{
		Use:     man.Use,
		Short:   man.Short,
		Long:    man.Long,
		Example: man.Example,
		PreRunE: unsplashPreRun,
		RunE:    unsplashRun,
	}

	cmd.Flags().StringP("collections", "c", "", "Public collection ID(‘s) to filter selection. If multiple, comma-separated")
	cmd.Flags().BoolP("featured", "f", false, "Limit selection to featured photos. Valid values: false, true.")
	cmd.Flags().StringP("filter", "l", "low", "Limit results by content safety. Default: low. Valid values are low and high.")
	cmd.Flags().StringP("orientation", "", "landscape", "Filter by photo orientation. Valid values: landscape, portrait, squarish.")
	cmd.Flags().StringP("query", "q", "mountains", "Limit selection to photos matching a search term.")
	cmd.Flags().StringP("size", "s", "all", "Photos size. Valid values: all, thumb, small, regular, full, raw. Default: all")
	cmd.Flags().StringP("username", "u", "", "Limit selection to a single user.")

	return cmd
}

func unsplashPreRun(cmd *cobra.Command, args []string) error {
	// TODO: validate all fields

	params := aid.GetModelFromFlags(cmd)
	if !helper.ContainsString(unsplashPhotoSizes, params.Size) {
		return fmt.Errorf("unknown photo size provided: %s", params.Size)
	}

	return nil
}

func unsplashRun(cmd *cobra.Command, args []string) error {
	params := aid.GetModelFromFlags(cmd)
	var cred model.Credential = dao.GetCredentialByProvider(profile, "unsplash")
	if (model.Credential{}) == cred {
		return fmt.Errorf("Unsplash credential not found")
	}

	return aid.DownloadPhoto(params, cred, unsplashPhotoSizes)
}
