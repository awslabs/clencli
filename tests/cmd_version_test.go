package tests

import (
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func TestVersionCmd(t *testing.T) {
	args := []string{"version"}
	cmd := controller.VersionCmd()
	sout, serr, err := executeCommand(t, cmd, args)
	assert.Nil(t, err)
	assert.Empty(t, serr)
	assert.Contains(t, sout, "CLENCLI v")
}
