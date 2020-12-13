package tester

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/awslabs/clencli/helper"
)

// Setup create the necessary artifacts for a test, return the current working and the new working directory
func Setup(t *testing.T) (pwd string, nwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get the current directory path")
	}

	nwd, err = helper.CreateDirectoryNamedPath(t.Name())
	os.Chdir(nwd)

	return cwd, nwd
}

// SetupAll does TODO
func SetupAll() (pwd string, nwd string) {
	format := "2006-01-02-15-04-05.000000000"
	dt := time.Now().Format(format)

	cwd, err := os.Getwd()
	path := cwd + "/../../test/" + dt

	if err != nil {
		log.Fatal("Unable to get the current directory path")
	}

	nwd, err = helper.CreateDirectoryNamedPath(path)
	if err == nil {
		os.Chdir(nwd)
	}

	return cwd, nwd
}

// Teardown does TODO
func Teardown(pwd string, nwd string) {
	os.Chdir(pwd)
	err := os.RemoveAll(nwd)
	if err != nil {
		log.Fatal(err)
	}
}

// TeardownAll does TODO
func TeardownAll(pwd string, nwd string) {
	err := os.RemoveAll(nwd)
	if err != nil {
		log.Fatal(err)
	}
}
