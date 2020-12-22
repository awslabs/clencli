package tester

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var tDir string

func TestMain(m *testing.M) {
	format := "2006-01-02-15-04-05.000000000"
	dt := time.Now().Format(format)

	tDir, err := ioutil.TempDir(os.TempDir(), "clencli-"+dt)
	if err != nil {
		log.Fatal(err)
	}

	// enter the new directory
	err = os.Chdir(tDir)
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	// change to the previous directory
	// err = os.Chdir(tDir)
	// if err != nil {
	// log.Fatal(err)
	// }

	os.Exit(code)
}

// func TestConfigureEmpty(t *testing.T) {
// 	aid.DeleteConfigurationsDirectory()
// 	err := tester.ExecuteCommand(controller.ConfigureCmd(), "configure")
// 	assert.Contains(t, err.Error(), "invalid argument")
// }

// TODO: test configure command and provide the input via the test block
