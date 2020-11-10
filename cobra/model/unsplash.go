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

// Unsplash does ...
type Unsplash struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Enabled     bool   `yaml:"enabled"`
	CreatedAt   string `yaml:"createdAt"`
	UpdatedAt   string `yaml:"updatedAt"`
	RandomPhoto struct {
		Name        string                        `yaml:"name,omitempty"`
		Description string                        `yaml:"description,omitempty"`
		Enabled     bool                          `yaml:"enabled"`
		CreatedAt   string                        `yaml:"createdAt"`
		UpdatedAt   string                        `yaml:"updatedAt"`
		Parameters  UnsplashRandomPhotoParameters `yaml:"parameters"`
		Response    UnsplashRandomPhotoResponse   `yaml:"response,omitempty"`
	} `yaml:"random_photo"`
}

// UnsplashRandomPhotoParameters struct from Unsplash command
type UnsplashRandomPhotoParameters struct {
	Collections string `yaml:"collections"`
	Featured    bool   `yaml:"featured"`
	Filter      string `yaml:"filter"`
	Orientation string `yaml:"orientation"`
	Query       string `yaml:"query"`
	Size        string `yaml:"size"`
	Username    string `yaml:"username"`
}

// UnsplashRandomPhotoResponse is a struct from Unsplash API Randm Request
type UnsplashRandomPhotoResponse struct {
	ID             string      `json:"id"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
	PromotedAt     interface{} `json:"promoted_at"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Color          string      `json:"color"`
	Description    string      `json:"description"`
	AltDescription string      `json:"alt_description"`
	Urls           struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	Categories             []interface{} `json:"categories"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"liked_by_user"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	Sponsorship            interface{}   `json:"sponsorship"`
	User                   struct {
		ID              string `json:"id"`
		UpdatedAt       string `json:"updated_at"`
		Username        string `json:"username"`
		Name            string `json:"name"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		TwitterUsername string `json:"twitter_username"`
		PortfolioURL    string `json:"portfolio_url"`
		Bio             string `json:"bio"`
		Location        string `json:"location"`
		Links           struct {
			Self      string `json:"self"`
			HTML      string `json:"html"`
			Photos    string `json:"photos"`
			Likes     string `json:"likes"`
			Portfolio string `json:"portfolio"`
			Following string `json:"following"`
			Followers string `json:"followers"`
		} `json:"links"`
		ProfileImage struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"profile_image"`
		InstagramUsername string `json:"instagram_username"`
		TotalCollections  int    `json:"total_collections"`
		TotalLikes        int    `json:"total_likes"`
		TotalPhotos       int    `json:"total_photos"`
		AcceptedTos       bool   `json:"accepted_tos"`
	} `json:"user"`
	Exif struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposure_time"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focal_length"`
		Iso          int    `json:"iso"`
	} `json:"exif"`
	Location struct {
		Title    string      `json:"title"`
		Name     string      `json:"name"`
		City     interface{} `json:"city"`
		Country  string      `json:"country"`
		Position struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
	} `json:"location"`
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
}
