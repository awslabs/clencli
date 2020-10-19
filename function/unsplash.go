package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// RandomUnsplash is a struct from Unsplash API Randm Request
type RandomUnsplash struct {
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

// GetRandomPhotoDefaults retrieves a single random photo with default values.
func GetRandomPhotoDefaults(query string) (RandomUnsplash, error) {
	// landscape orientation is better for README files
	return GetRandomPhoto(query, "", "", "", "landscape", "low")
}

// GetRandomPhoto retrieves a single random photo, given optional filters.
func GetRandomPhoto(query string, collections string, featured string, username string, orientation string, filter string) (RandomUnsplash, error) {
	var unsplash RandomUnsplash

	clientID := viper.Get("unsplash.access_key")
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s&query=%s", clientID, query)

	if len(collections) > 0 {
		url += fmt.Sprintf("&collections=%s", collections)
	}

	if len(featured) > 0 {
		url += fmt.Sprintf("&featured=%s", featured)
	}

	if len(username) > 0 {
		url += fmt.Sprintf("&username=%s", username)
	}

	if len(orientation) > 0 {
		url += fmt.Sprintf("&orientation=%s", orientation)
	}

	if len(filter) > 0 {
		url += fmt.Sprintf("&filter=%s", filter)
	}

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return unsplash, fmt.Errorf("Unexpected error while performing GET on Unsplash API \n%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return unsplash, fmt.Errorf("Unexpected error while reading Unsplash response \n%v", err)
		}

		json.Unmarshal(bodyBytes, &unsplash)
	}

	return unsplash, err
}

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

	return DownloadFile(url, dirPath, fileName)
}

// GetPhotoURLBySize return the photo URL based on the given size
func GetPhotoURLBySize(size string, u RandomUnsplash) string {
	switch size {
	case "thumb":
		return u.Urls.Thumb
	case "small":
		return u.Urls.Small
	case "regular":
		return u.Urls.Regular
	case "full":
		return u.Urls.Full
	case "raw":
		return u.Urls.Raw
	default:
		return u.Urls.Small
	}

}
