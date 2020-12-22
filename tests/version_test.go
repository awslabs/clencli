package tests

import (
	"testing"

	"github.com/awslabs/clencli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func TestVersionCmd(t *testing.T) {
	args := []string{"version"}
	cmd := controller.VersionCmd()
	out, err := executeCommandR(t, cmd, args)
	assert.Nil(t, err)
	assert.Contains(t, out, "CLENCLI v")
}
