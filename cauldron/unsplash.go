package function

// getUnsplashRandomPhotoDefaults retrieves a single random photo with default values.
// func getUnsplashRandomPhotoDefaults(query string) (RandomPhotoResponse, error) {
// 	// landscape orientation is better for README files
// 	return getUnsplashRandomPhoto(query, "", "", "", "landscape", "low")
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

// getUnsplashRandomPhoto retrieves a single random photo, given optional filters.
// func getUnsplashRandomPhoto(params RandomPhotoParameters) (RandomPhotoResponse, error) {
// 	var response RandomPhotoResponse

// 	if (RandomPhotoParameters{} == params) {
// 		return response, fmt.Errorf("Unable to download Unsplash photo if all fields from query as empty")
// 	}

// 	gc, err := getGlobalConfig()
// 	if err != nil {
// 		return response, fmt.Errorf("Unexpected error while getting global configuration \n%v", err)
// 	}

// 	// check if Credentials has unsplash credential
// 	// return response, fmt.Errorf("No Unsplash credentials found in the global configuration \n%v", err)

// 	// check if config has Unsplash credentials
// 	for _, cred := range gc.Credentials {
// 		if cred.Provider == "unsplash" {
// 			if cred.AccessKey != "" && cred.SecretKey != "" {

// 				url := buildURL(params, cred.Credential)

// 				var client http.Client
// 				resp, err := client.Get(url)
// 				if err != nil {
// 					return response, fmt.Errorf("Unexpected error while performing GET on Unsplash API \n%v", err)
// 				}
// 				defer resp.Body.Close()

// 				if resp.StatusCode == http.StatusOK {
// 					bodyBytes, err := ioutil.ReadAll(resp.Body)
// 					if err != nil {
// 						return response, fmt.Errorf("Unexpected error while reading Unsplash response \n%v", err)
// 					}

// 					json.Unmarshal(bodyBytes, &response)
// 				}
// 			}
// 		}
// 	}

// 	return response, err
// }

// func buildURL(params RandomPhotoParameters, cred Credential) string {
// 	clientID := cred.AccessKey
// 	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s", clientID)

// 	if len(params.Collections) > 0 {
// 		url += fmt.Sprintf("&collections=%s", params.Collections)
// 	}

// 	if len(params.Query) > 0 {
// 		url += fmt.Sprintf("&query=%s", params.Query)
// 	}

// 	if len(params.Featured) > 0 {
// 		url += fmt.Sprintf("&featured=%s", params.Featured)
// 	}

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

// // DownloadPhoto downloads a photo and saves into downloads/unsplash/ folder
// // It creates the downloads/ folder if it doesn't exists
// func DownloadPhoto(url string, size string, query string) error {

// 	// Get the photo identifier
// 	start := strings.Index(url, "photo")
// 	end := strings.Index(url, "?")

// 	// Create a Rune from the URL
// 	runes := []rune(url)

// 	// Generate the directory path
// 	dirPath := "downloads/unsplash/" + query

// 	// Generate the filename
// 	fileName := string(runes[start:end])
// 	fileName += "-" + size + ".jpg"

// 	return DownloadFile(url, dirPath, fileName)
// }

// // GetPhotoURLBySize return the photo URL based on the given size
// func GetPhotoURLBySize(p RandomPhotoParameters, r RandomPhotoResponse) string {
// 	switch p.Size {
// 	case "thumb":
// 		return r.Urls.Thumb
// 	case "small":
// 		return r.Urls.Small
// 	case "regular":
// 		return r.Urls.Regular
// 	case "full":
// 		return r.Urls.Full
// 	case "raw":
// 		return r.Urls.Raw
// 	default:
// 		return r.Urls.Small
// 	}

// }
