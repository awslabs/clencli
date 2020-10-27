package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/awslabs/clencli/cobra/model"
)

// AskAboutConfiguration ask user about configuration
func AskAboutConfiguration(conf model.Configuration) model.Configuration {
	// configuration can have many types: Unsplash, AWS, etc
	fmt.Println(">>> Configuration")
	conf.Name = GetUserInputAsString(">>>> Name", conf.Name)
	conf.Description = GetUserInputAsString(">>>> Description", conf.Description)
	conf.Enabled = GetUserInputAsBool(">>>> Enabled", conf.Enabled)
	conf.UpdatedAt = time.Now().String()

	answer := GetUserInputAsBool("Would you like to setup Unsplash configuration?", false)

	if answer {
		conf.Unsplash = askAboutUnsplashConfiguration(conf.Unsplash)
	} else {
		fmt.Println("Skipping Unplash configuration ...")
	}
	return conf
}

// AskAboutConfigurationProfile ask user about configuration profile
func AskAboutConfigurationProfile(profile model.ConfigurationProfile) model.ConfigurationProfile {
	fmt.Println(">> Profile")
	profile.Name = GetUserInputAsString(">>> Name", profile.Name)
	profile.Enabled = GetUserInputAsBool(">>> Enabled", profile.Enabled)
	profile.Description = GetUserInputAsString(">>> Description", profile.Description)
	profile.UpdatedAt = time.Now().String()
	return profile
}

// AskAboutCredential ask user about credential
func AskAboutCredential(credential model.Credential) model.Credential {
	fmt.Println(">>> Credential")
	credential.Name = GetUserInputAsString(">>>> Name", credential.Name)
	credential.Enabled = GetUserInputAsBool(">>>> Enabled", credential.Enabled)
	credential.Description = GetUserInputAsString(">>>> Description", credential.Description)
	credential.UpdatedAt = time.Now().String()
	credential.Provider = GetUserInputAsString(">>>> Provider", credential.Provider)
	credential.AccessKey = getSensitiveUserInputAsString(">>>> Access Key", credential.AccessKey)
	credential.SecretKey = getSensitiveUserInputAsString(">>>> Secret Key", credential.SecretKey)
	return credential
}

// AskAboutCredentialProfile ask user about credential profile
func AskAboutCredentialProfile(profile model.CredentialProfile) model.CredentialProfile {
	fmt.Println(">> Profile")
	profile.Name = GetUserInputAsString(">>> Name", profile.Name)
	profile.Description = GetUserInputAsString(">>> Description", profile.Description)
	profile.Enabled = GetUserInputAsBool(">>> Enabled", profile.Enabled)
	profile.UpdatedAt = time.Now().String()
	return profile
}

func askAboutUnsplashConfiguration(unsplash model.Unsplash) model.Unsplash {
	// unsplash configuration may have multiple nested configuration, such as random photo, etc...
	fmt.Println(">>>> Unsplash")
	unsplash.Name = GetUserInputAsString(">>>>> Name", unsplash.Name)
	unsplash.Description = GetUserInputAsString(">>>>> Description", unsplash.Description)
	unsplash.Enabled = GetUserInputAsBool(">>>>> Enabled", unsplash.Enabled)
	unsplash.UpdatedAt = time.Now().String()

	answer := GetUserInputAsBool("Would you like to setup Unsplash Random Photo Parameters?", false)

	if answer {
		fmt.Println(">>>>> Random Photo")
		unsplash.RandomPhoto.Name = GetUserInputAsString(">>>>>> Name", unsplash.RandomPhoto.Name)
		unsplash.RandomPhoto.Description = GetUserInputAsString(">>>>>> Description", unsplash.RandomPhoto.Description)
		unsplash.RandomPhoto.Enabled = GetUserInputAsBool(">>>>>> Enabled", unsplash.RandomPhoto.Enabled)
		unsplash.RandomPhoto.UpdatedAt = time.Now().String()
		unsplash.RandomPhoto.Parameters = askAboutUnsplashRandomPhotoParameters(unsplash.RandomPhoto.Parameters)
	} else {
		fmt.Println("Skipping Unplash Random Photo configuration ...")
	}

	return unsplash
}

func askAboutUnsplashRandomPhotoParameters(params model.UnsplashRandomPhotoParameters) model.UnsplashRandomPhotoParameters {
	fmt.Println(">>>>>> Parameters")
	params.Collections = GetUserInputAsString(">>>>>>> Public collection ID(‘s) to filter selection. If multiple, comma-separated.\nCollections ", params.Collections)
	params.Featured = GetUserInputAsBool(">>>>>>> Limit selection to featured photos. Valid values: false and true. Default: false\nFeatured", params.Featured)
	params.Filter = GetUserInputAsString(">>>>>>> Limit results by content safety. Valid values are low and high.\nFilter", params.Filter)
	params.Orientation = GetUserInputAsString(">>>>>>> Filter by photo orientation. Valid values: landscape, portrait, squarish.\nOrientation", params.Orientation)
	params.Query = GetUserInputAsString(">>>>>>> Limit selection to photos matching a search term.\nQuery", params.Query)
	params.Size = GetUserInputAsString(">>>>>>> Photos size. Valid values: all, thumb, small, regular, full, raw.\nSize", params.Size)
	params.Username = GetUserInputAsString(">>>>>>> Limit selection to a single user.\nUsername", params.Username)
	return params
}

func getSensitiveUserInput(text string, info string) (string, error) {
	return getUserInput(text+" ["+maskString(info, 3)+"]", "")
}

func getSensitiveUserInputAsString(text string, info string) string {
	answer, err := getSensitiveUserInput(text, info)
	if err != nil {
		log.Fatalf("Unable to get user input about profile's name\n%v", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}

func getUserInput(text string, info string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if info == "" {
		fmt.Print(text + ": ")
	} else {
		fmt.Print(text + " [" + info + "]: ")
	}

	input, err := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		return input, fmt.Errorf("Unable to read user input \n%v", err)
	}
	return input, err
}

// GetUserInputAsBool prints `text` on console and return answer as `boolean`
func GetUserInputAsBool(text string, info bool) bool {
	answer, err := getUserInput(text, strconv.FormatBool(info))
	if err != nil {
		log.Fatalf("Unable to get user input as boolean\n%s", err)
	}

	if answer != "" && answer == "true" {
		return true
	}

	return false
}

// GetUserInputAsString prints `text` on console and return answer as `string`
func GetUserInputAsString(text string, info string) string {
	answer, err := getUserInput(text, info)
	if err != nil {
		log.Fatalf("Unable to get user input about profile's name\n%v", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}

func maskString(s string, showLastChars int) string {
	maskSize := len(s) - showLastChars
	if maskSize <= 0 {
		return s
	}

	return strings.Repeat("*", maskSize) + s[maskSize:]
}
