package function

import (
	"io/ioutil"
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

// CreateDirectoryNamedPath creates a directory named path, along with any necessary parents, and returns nil, or else returns an error.
// The permission bits perm (before umask) are used for all directories. If path is already a directory does nothing and returns nil.
func CreateDirectoryNamedPath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create directory (and its parents)", err)
	}

	return err
}

// CreateTempDir creates a new temporary directory in the directory dir.
// The directory name is generated by taking pattern and applying a random string to the end
func CreateTempDir(dir string, pattern string) (string, error) {
	dir, err := ioutil.TempDir(dir, pattern)
	if err != nil {
		log.Fatal(err)
	}
	return dir, err
}
