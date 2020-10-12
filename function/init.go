package function

import (
	"os"

	"github.com/awslabs/clencli/box"
)

// Init creates the project directory, change the current directory and places basic configuration files
func Init(project string) {
	// Create the project directory
	CreateDir(project)

	// Change current directory to project directory
	os.Chdir(project)

	// Create a directory for CLENCLI
	CreateDir("clencli")

	initReadme := "clencli/readme.yaml"
	blobReadme, _ := box.Get("/init/clencli/readme.yaml")
	WriteFile(initReadme, blobReadme)

	// Gomplate
	initReadMeTmpl := "clencli/readme.tmpl"
	blobReadMeTmpl, _ := box.Get("/init/clencli/readme.tmpl")
	WriteFile(initReadMeTmpl, blobReadMeTmpl)

	// Gitignore
	initGitIgnore := ".gitignore"
	blobGitIgnore, _ := box.Get("/init/.gitignore")
	WriteFile(initGitIgnore, blobGitIgnore)
}

// InitHLD copies the High Level Design template file
func InitHLD(project string) {
	initHLD := "clencli/hld.yaml"
	blobHLD, _ := box.Get("/init/clencli/hld.yaml")
	WriteFile(initHLD, blobHLD)

	initHLDTmpl := "clencli/hld.tmpl"
	blobHLDTmpl, _ := box.Get("/init/clencli/hld.tmpl")
	WriteFile(initHLDTmpl, blobHLDTmpl)
}
