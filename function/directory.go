package function

import (
	"log"
	"os"
)

// CreateDir create a directory, even if not existent, given its name
func CreateDir(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(name, os.ModePerm)
		if errDir != nil {
			log.Fatal(err)
			return false
		}
	}

	return true
}
