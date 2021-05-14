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
	RandomPhoto UnsplashRandomPhoto `yaml:"randomPhoto"`
}

// UnsplashRandomPhoto ...
type UnsplashRandomPhoto struct {
	Parameters UnsplashRandomPhotoParameters `yaml:"parameters"`
	Response   UnsplashRandomPhotoResponse   `yaml:"response,omitempty"`
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
	CreatedAt      string      `json:"createdAt"`
	UpdatedAt      string      `json:"updatedAt"`
	PromotedAt     interface{} `json:"promotedAt"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Color          string      `json:"color"`
	Description    string      `json:"description"`
	AltDescription string      `json:"altDescription"`
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
		DownloadLocation string `json:"downloadLocation"`
	} `json:"links"`
	Categories             []interface{} `json:"categories"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"likedByUser"`
	CurrentUserCollections []interface{} `json:"currentUserCollections"`
	Sponsorship            interface{}   `json:"sponsorship"`
	User                   struct {
		ID              string `json:"id"`
		UpdatedAt       string `json:"updatedAt"`
		Username        string `json:"username"`
		Name            string `json:"name"`
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		TwitterUsername string `json:"twitterUsername"`
		PortfolioURL    string `json:"portfolioUrl"`
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
		} `json:"profileImage"`
		InstagramUsername string `json:"instagramUsername"`
		TotalCollections  int    `json:"totalCollections"`
		TotalLikes        int    `json:"totalLikes"`
		TotalPhotos       int    `json:"totalPhotos"`
		AcceptedTos       bool   `json:"acceptedTos"`
	} `json:"user"`
	Exif struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposureTime"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focalLength"`
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

// UnsplashGetPhotoResponse TODO ...
type UnsplashGetPhotoResponse struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	BlurHash       string `json:"blur_hash"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
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
		Title    string `json:"title"`
		Name     string `json:"name"`
		City     string `json:"city"`
		Country  string `json:"country"`
		Position struct {
			Latitude  interface{} `json:"latitude"`
			Longitude interface{} `json:"longitude"`
		} `json:"position"`
	} `json:"location"`
	Meta struct {
		Index bool `json:"index"`
	} `json:"meta"`
	Tags []struct {
		Type   string `json:"type"`
		Title  string `json:"title"`
		Source struct {
			Ancestry struct {
				Type struct {
					Slug       string `json:"slug"`
					PrettySlug string `json:"pretty_slug"`
				} `json:"type"`
				Category struct {
					Slug       string `json:"slug"`
					PrettySlug string `json:"pretty_slug"`
				} `json:"category"`
				Subcategory struct {
					Slug       string `json:"slug"`
					PrettySlug string `json:"pretty_slug"`
				} `json:"subcategory"`
			} `json:"ancestry"`
			Title           string `json:"title"`
			Subtitle        string `json:"subtitle"`
			Description     string `json:"description"`
			MetaTitle       string `json:"meta_title"`
			MetaDescription string `json:"meta_description"`
			CoverPhoto      struct {
				ID             string      `json:"id"`
				CreatedAt      string      `json:"created_at"`
				UpdatedAt      string      `json:"updated_at"`
				PromotedAt     string      `json:"promoted_at"`
				Width          int         `json:"width"`
				Height         int         `json:"height"`
				Color          string      `json:"color"`
				BlurHash       string      `json:"blur_hash"`
				Description    interface{} `json:"description"`
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
			} `json:"cover_photo"`
		} `json:"source,omitempty"`
	} `json:"tags"`
	RelatedCollections struct {
		Total   int    `json:"total"`
		Type    string `json:"type"`
		Results []struct {
			ID              string      `json:"id"`
			Title           string      `json:"title"`
			Description     interface{} `json:"description"`
			PublishedAt     string      `json:"published_at"`
			LastCollectedAt string      `json:"last_collected_at"`
			UpdatedAt       string      `json:"updated_at"`
			Curated         bool        `json:"curated"`
			Featured        bool        `json:"featured"`
			TotalPhotos     int         `json:"total_photos"`
			Private         bool        `json:"private"`
			ShareKey        string      `json:"share_key"`
			Tags            []struct {
				Type   string `json:"type"`
				Title  string `json:"title"`
				Source struct {
					Ancestry struct {
						Type struct {
							Slug       string `json:"slug"`
							PrettySlug string `json:"pretty_slug"`
						} `json:"type"`
						Category struct {
							Slug       string `json:"slug"`
							PrettySlug string `json:"pretty_slug"`
						} `json:"category"`
						Subcategory struct {
							Slug       string `json:"slug"`
							PrettySlug string `json:"pretty_slug"`
						} `json:"subcategory"`
					} `json:"ancestry"`
					Title           string `json:"title"`
					Subtitle        string `json:"subtitle"`
					Description     string `json:"description"`
					MetaTitle       string `json:"meta_title"`
					MetaDescription string `json:"meta_description"`
					CoverPhoto      struct {
						ID             string      `json:"id"`
						CreatedAt      string      `json:"created_at"`
						UpdatedAt      string      `json:"updated_at"`
						PromotedAt     string      `json:"promoted_at"`
						Width          int         `json:"width"`
						Height         int         `json:"height"`
						Color          string      `json:"color"`
						BlurHash       string      `json:"blur_hash"`
						Description    interface{} `json:"description"`
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
							ID              string      `json:"id"`
							UpdatedAt       string      `json:"updated_at"`
							Username        string      `json:"username"`
							Name            string      `json:"name"`
							FirstName       string      `json:"first_name"`
							LastName        string      `json:"last_name"`
							TwitterUsername interface{} `json:"twitter_username"`
							PortfolioURL    interface{} `json:"portfolio_url"`
							Bio             string      `json:"bio"`
							Location        interface{} `json:"location"`
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
					} `json:"cover_photo"`
				} `json:"source,omitempty"`
			} `json:"tags"`
			Links struct {
				Self    string `json:"self"`
				HTML    string `json:"html"`
				Photos  string `json:"photos"`
				Related string `json:"related"`
			} `json:"links"`
			User struct {
				ID              string      `json:"id"`
				UpdatedAt       string      `json:"updated_at"`
				Username        string      `json:"username"`
				Name            string      `json:"name"`
				FirstName       string      `json:"first_name"`
				LastName        string      `json:"last_name"`
				TwitterUsername interface{} `json:"twitter_username"`
				PortfolioURL    interface{} `json:"portfolio_url"`
				Bio             interface{} `json:"bio"`
				Location        interface{} `json:"location"`
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
				InstagramUsername interface{} `json:"instagram_username"`
				TotalCollections  int         `json:"total_collections"`
				TotalLikes        int         `json:"total_likes"`
				TotalPhotos       int         `json:"total_photos"`
				AcceptedTos       bool        `json:"accepted_tos"`
			} `json:"user"`
			CoverPhoto struct {
				ID             string      `json:"id"`
				CreatedAt      string      `json:"created_at"`
				UpdatedAt      string      `json:"updated_at"`
				PromotedAt     interface{} `json:"promoted_at"`
				Width          int         `json:"width"`
				Height         int         `json:"height"`
				Color          string      `json:"color"`
				BlurHash       string      `json:"blur_hash"`
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
					ID              string      `json:"id"`
					UpdatedAt       string      `json:"updated_at"`
					Username        string      `json:"username"`
					Name            string      `json:"name"`
					FirstName       string      `json:"first_name"`
					LastName        string      `json:"last_name"`
					TwitterUsername interface{} `json:"twitter_username"`
					PortfolioURL    string      `json:"portfolio_url"`
					Bio             string      `json:"bio"`
					Location        string      `json:"location"`
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
					InstagramUsername interface{} `json:"instagram_username"`
					TotalCollections  int         `json:"total_collections"`
					TotalLikes        int         `json:"total_likes"`
					TotalPhotos       int         `json:"total_photos"`
					AcceptedTos       bool        `json:"accepted_tos"`
				} `json:"user"`
			} `json:"cover_photo"`
			PreviewPhotos []struct {
				ID        string `json:"id"`
				CreatedAt string `json:"created_at"`
				UpdatedAt string `json:"updated_at"`
				BlurHash  string `json:"blur_hash"`
				Urls      struct {
					Raw     string `json:"raw"`
					Full    string `json:"full"`
					Regular string `json:"regular"`
					Small   string `json:"small"`
					Thumb   string `json:"thumb"`
				} `json:"urls"`
			} `json:"preview_photos"`
		} `json:"results"`
	} `json:"related_collections"`
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
}
