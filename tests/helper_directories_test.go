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
		"empty":              {dir: "", want: false},
		"single backslack":   {dir: "\\", want: true},
		"multiple backslack": {dir: "\\\\\\", want: true},
		"dollar sign":        {dir: "$", want: true},
		"illegal char":       {dir: "/", want: false},
		"illegal char \\":    {dir: "\\", want: true},
		"illegal char /":     {dir: "/", want: false},
		"illegal char :":     {dir: ":", want: true},
		"illegal char *":     {dir: "*", want: true},
		"illegal char ?":     {dir: "?", want: true},
		"illegal char \" ":   {dir: "\"", want: true},
		"illegal char <":     {dir: "<", want: true},
		"illegal char >":     {dir: ">", want: true},
		"illegal char |":     {dir: "|", want: true},
		"single letter":      {dir: "a", want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			wd := createAndEnterTestDirectory(t)
			got := helper.MkDirsIfNotExist(tc.dir)
			os.Chdir(wd)
			assert.Equal(t, tc.want, got)

		})
	}
}
