package cauldron

import (
	"fmt"
	"io/ioutil"

	"io"
	"log"
	"net/http"
	"os"
)

// WriteFile writes a file and return true if successful
func WriteFile(filename string, data []byte) bool {
	err := ioutil.WriteFile(filename, data, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// DownloadFile downloads a file and saves into downloads/ folder
// It creates the downloads/ folder if it doesn't exists
func DownloadFile(url string, dirPath string, fileName string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if CreateDir(dirPath) {
		// Create the file
		out, err := os.Create(dirPath + "/" + fileName)
		if err != nil {
			return err
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		return err
	}
	return err
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirOrFileExists an error is known to report that a file or directory does not exist.
// It is satisfied by ErrNotExist as well as some syscall errors.
func DirOrFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// CopyFile copies a file from source into a given destination path
// https://github.com/mactsouk/opensource.com/blob/master/cp2.go
func CopyFile(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}

// FileSize return the size of the give file path.
// Gives an error if files does not exist
func FileSize(path string) (int64, error) {
	var size int64 = -1
	if FileExists(path) {
		info, err := os.Stat(path)
		if err != nil {
			if err != nil {
				return size, fmt.Errorf("Unable to obtain information about file: %s\n%s", path, err)
			}
			return size, err
		}
		size = info.Size()
	} else {
		return size, fmt.Errorf("File does not exist")
	}
	return size, nil
}
