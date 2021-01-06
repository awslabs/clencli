package tests

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/helper"
	"github.com/stretchr/testify/assert"
)

func TestMkDir(t *testing.T) {
	tests := map[string]struct {
		dir  string
		want bool
	}{
		"foo": {dir: "foo", want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			wd := createAndEnterTestDirectory(t)
			os.Chdir(wd)
			got := helper.MkDirsIfNotExist(tc.dir)
			assert.Equal(t, tc.want, got)

		})
	}
}
