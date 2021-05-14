package helper

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	"io"
	"net/http"
	"os"

	"github.com/awslabs/clencli/box"
	"github.com/sirupsen/logrus"
)

// WriteFile writes a file and return true if successful
func WriteFile(filename string, data []byte) bool {
	err := ioutil.WriteFile(filename, data, os.ModePerm)

	if err != nil {
		logrus.Fatal(err)
		return false
	}

	return true
}

// WriteFileFromBox get the file from box's resources and write into the given destination, returns false if not able to.
func WriteFileFromBox(source string, dest string) bool {
	bytes, found := box.Get(source)

	if !found {
		logrus.Errorf("file \"%s\" not found under box/resources", source)
		return false
	}

	return WriteFile(dest, bytes)
}

// DownloadFileTo downloads a file and saves into the given directory with the given file name
func DownloadFileTo(url string, destination string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DownloadFile downloads a file and saves into the given directory with the given file name
func DownloadFile(url string, dirPath string, filename string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	sep := string(os.PathSeparator)
	out, err := os.Create(dirPath + sep + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(path string) bool {
	info, err := os.Stat(path)
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
func CopyFile(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		os.Exit(1)
	}
}

// CopyFileTo copy a file from source to destination
func CopyFileTo(source string, dest string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		logrus.Errorf("unable to read file\n%v", err)
		return err
	}

	err = ioutil.WriteFile(dest, input, os.ModePerm)
	if err != nil {
		logrus.Errorf("unable to write file\n%v", err)
		return err
	}

	return nil
}

// FileSize return the size of the give file path.
// Gives an error if files does not exist
func FileSize(path string) (int64, error) {
	var size int64 = -1
	if FileExists(path) {
		info, err := os.Stat(path)
		if err != nil {
			if err != nil {
				return size, fmt.Errorf("unable to obtain information about file: %s\n%s", path, err)
			}
			return size, err
		}
		size = info.Size()
	} else {
		return size, fmt.Errorf("file does not exist")
	}
	return size, nil
}

// ListFiles list of all file names in the given directory. Pass "." if you want to list at the current directory.
func ListFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		logrus.Fatal(err)
	}

	return files
}

// DeleteFile removes the named file or (empty) directory.
// If there is an error, it will be of type *PathError.
func DeleteFile(name string) error {
	return os.Remove(name)
}

// TrimRightFile open file in `path`, read line by line removing all trailing characters.
// Overwrite original file in `path` if `true`, it creates a new file `path`.tmp otherwise.
func TrimRightFile(path string, overwrite bool) error {
	f, err := os.Open(path)
	if err != nil {
		logrus.Errorf("unable to open file %s\n%v", f, err)
	}
	defer f.Close()

	tPath := path + ".tmp"
	tf, err := os.OpenFile(tPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		logrus.Error(err)
	}
	defer tf.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tr := strings.TrimRight(scanner.Text(), " ")
		if _, err := tf.WriteString(tr + "\n"); err != nil {
			logrus.Errorf("unable write trimmed string to temporary file\n%v", err)
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		logrus.Errorf("unable to use scanner\n%v", err)
		return err
	}

	if overwrite {
		if err := CopyFileTo(tPath, path); err != nil {
			logrus.Errorf("unable to replace original file with temporary\n%v", err)
			return err
		}
		defer os.Remove(tPath)
	}

	return nil
}
