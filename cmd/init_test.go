package cmd

import (
	"os"
	"testing"
	"time"

	fun "github.com/awslabs/clencli/function"
)

func TestMain(m *testing.M) {
	dt := time.Now().Format("2006-01-02-15-04-05.000000000")

	fun.CreateDir("../test/" + dt)
	code := m.Run()
	os.Exit(code)
}

func TestInitWithNoArgsAndNoFlags(t *testing.T) {

	cmd := InitCmd()
	cmd.Execute()
}
