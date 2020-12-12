package aid

import (
	"os"

	"github.com/awslabs/clencli/helper"
)

// CreateBasicProject creates a basic project
func CreateBasicProject(name string) {
	initializeProject(name)
	// helper.InitCustomProjectLayout(typee, "default")
	// helper.InitCustomProjectLayout(typee, structure)
	// helper.UpdateReadMe()
}

func initializeProject(name string) {

	// Create the name directory
	helper.CreateDir(name)

	// Change current directory to name directory
	os.Chdir(name)
}

// func initCreateCloudFormationProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
// 	helper.InitializeProject(name)
// 	if !onlyCustomizedStructure {
// 		helper.InitBasic()
// 		helper.InitHLD(name)
// 		helper.InitCloudFormation()
// 	}
// 	// helper.InitCustomProjectLayout(typee, "default")
// 	// helper.InitCustomProjectLayout(typee, structure)
// 	// helper.UpdateReadMe()
// }

// func initCreateTerraformProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
// 	helper.InitializeProject(name)
// 	if !onlyCustomizedStructure {
// 		helper.InitBasic()
// 		helper.InitHLD(name)
// 		helper.InitTerraform()
// 	}
// 	// helper.InitCustomProjectLayout(typee, "default")
// 	// helper.InitCustomProjectLayout(typee, structure)
// 	// helper.UpdateReadMe()
// }

// InitBasic create the basic configuration files
// func InitBasic() {

// 	// Create a directory for CLENCLI
// 	CreateDir("clencli")

// 	initReadme := "clencli/readme.yaml"
// 	blobReadme, _ := box.Get("/init/clencli/readme.yaml")
// 	WriteFile(initReadme, blobReadme)

// 	// Gomplate
// 	initReadMeTmpl := "clencli/readme.tmpl"
// 	blobReadMeTmpl, _ := box.Get("/init/clencli/readme.tmpl")
// 	WriteFile(initReadMeTmpl, blobReadMeTmpl)

// 	// Gitignore
// 	initGitIgnore := ".gitignore"
// 	blobGitIgnore, _ := box.Get("/init/.gitignore")
// 	WriteFile(initGitIgnore, blobGitIgnore)
// }

// // InitHLD copies the High Level Design template file
// func InitHLD(project string) {
// 	initHLD := "clencli/hld.yaml"
// 	blobHLD, _ := box.Get("/init/clencli/hld.yaml")
// 	WriteFile(initHLD, blobHLD)

// 	initHLDTmpl := "clencli/hld.tmpl"
// 	blobHLDTmpl, _ := box.Get("/init/clencli/hld.tmpl")
// 	WriteFile(initHLDTmpl, blobHLDTmpl)
// }

// // InitCustomProjectLayout generates
// func InitCustomProjectLayout(projectType string, projectStructureName string) error {
// 	var g GlobalConfig
// 	err := viper.Unmarshal(&g)
// 	if err != nil {
// 		log.Fatalf("Unable to decode into struct, %v", err)
// 	}

// 	for i := 0; i < len(g.Config.Init.Types); i++ {
// 		t := g.Config.Init.Types[i]
// 		// match the project type
// 		if projectType == t.Type {
// 			// only create if project structure is enabled
// 			if t.Enabled {
// 				log.Println("Creating customized project structure")
// 				if projectStructureName == t.Name {
// 					log.Printf("Using project structure: %s\n", t.Name)
// 				}

// 				for _, f := range t.Files {
// 					if f.File.State == "directory" {
// 						CreateDir(f.File.Path)
// 					} else if f.File.State == "file" {
// 						dir, file := filepath.Split(f.File.Dest)
// 						// in case it's the current directory
// 						if dir == "" {
// 							dir = "."
// 						}
// 						if strings.Contains(f.File.Src, "http") {
// 							DownloadFile(f.File.Src, dir, file)
// 						} else {
// 							CopyFile(f.File.Src, f.File.Dest)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return nil
// }

// // InitCloudFormation initialize a project with CloudFormation structure
// func InitCloudFormation() {

// 	CreateDir("environments")
// 	CreateDir("environments/dev")
// 	CreateDir("environments/prod")

// 	initCFStack := "stack.yaml"
// 	blobCFStack, _ := box.Get("/init/type/clouformation/stack.yaml")
// 	WriteFile(initCFStack, blobCFStack)

// 	initCFNested := "nested.yaml"
// 	blobCFNested, _ := box.Get("/init/type/clouformation/nested.yaml")
// 	WriteFile(initCFNested, blobCFNested)

// }

// // InitTerraform initialize a project with Terraform structure
// func InitTerraform() {
// 	initMakefile := "Makefile"
// 	blobMakefile, _ := box.Get("/init/type/terraform/Makefile")
// 	WriteFile(initMakefile, blobMakefile)

// 	initLicense := "LICENSE"
// 	blobLicense, _ := box.Get("/init/type/terraform/LICENSE")
// 	WriteFile(initLicense, blobLicense)

// 	CreateDir("environments")

// 	initDevEnvironment := "environments/dev.tf"
// 	blobDevEnvironment, _ := box.Get("/init/type/terraform/environments/dev.tf")
// 	WriteFile(initDevEnvironment, blobDevEnvironment)

// 	initProdEnvironment := "environments/prod.tf"
// 	blobProdEnvironment, _ := box.Get("/init/type/terraform/environments/prod.tf")
// 	WriteFile(initProdEnvironment, blobProdEnvironment)

// 	initMainTF := "main.tf"
// 	blobMainTF, _ := box.Get("/init/type/terraform/main.tf")
// 	WriteFile(initMainTF, blobMainTF)

// 	initVariablesTF := "variables.tf"
// 	blobVariablesTF, _ := box.Get("/init/type/terraform/variables.tf")
// 	WriteFile(initVariablesTF, blobVariablesTF)

// 	initOutputsTF := "outputs.tf"
// 	blobOutputsTF, _ := box.Get("/init/type/terraform/outputs.tf")
// 	WriteFile(initOutputsTF, blobOutputsTF)
// }