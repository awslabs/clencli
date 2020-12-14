package aid

import (
	"os"

	"github.com/awslabs/clencli/box"
	"github.com/awslabs/clencli/helper"
)

/* BASIC PROJECT */

// CreateBasicProject creates a basic project
func CreateBasicProject(name string) {
	createProjectDir(name)
	initBasicProject()
	// helper.InitCustomProjectLayout(typee, "default")
	// helper.InitCustomProjectLayout(typee, structure)
	// helper.UpdateReadMe()
}

func createProjectDir(name string) {

	// Create the project directory, only if doesn't exist
	helper.CreateDir(name)

	// Change current directory to the project directory
	os.Chdir(name)
}

// create the basic configuration files
func initBasicProject() {

	// Create a directory for CLENCLI
	helper.CreateDir("clencli")

	initReadme := "clencli/readme.yaml"
	blobReadme, _ := box.Get("/init/clencli/readme.yaml")
	helper.WriteFile(initReadme, blobReadme)

	// Gomplate
	initReadMeTmpl := "clencli/readme.tmpl"
	blobReadMeTmpl, _ := box.Get("/init/clencli/readme.tmpl")
	helper.WriteFile(initReadMeTmpl, blobReadMeTmpl)

	// Gitignore
	initGitIgnore := ".gitignore"
	blobGitIgnore, _ := box.Get("/init/.gitignore")
	helper.WriteFile(initGitIgnore, blobGitIgnore)
}

/* CLOUD PROJECT */

// CreateCloudProject copies the necessary templates for cloud projects
func CreateCloudProject(name string) {
	CreateBasicProject(name)
	initCloudProject()
}

// copies the High Level Design template file
func initCloudProject() {
	initHLD := "clencli/hld.yaml"
	blobHLD, _ := box.Get("/init/clencli/hld.yaml")
	helper.WriteFile(initHLD, blobHLD)

	initHLDTmpl := "clencli/hld.tmpl"
	blobHLDTmpl, _ := box.Get("/init/clencli/hld.tmpl")
	helper.WriteFile(initHLDTmpl, blobHLDTmpl)
}

/* CLOUDFORMATION PROJECT */

// CreateCloudFormationProject creates an AWS CloudFormation project
func CreateCloudFormationProject(name string) {
	CreateBasicProject(name)
	initCloudProject()
	initCloudFormationProject()
}

// initialize a project with CloudFormation structure and copies template files
func initCloudFormationProject() {

	helper.CreateDir("environments")
	helper.CreateDir("environments/dev")
	helper.CreateDir("environments/prod")

	initCFSkeleton := "skeleton.yaml"
	blobCFSkeleton, _ := box.Get("/init/type/clouformation/skeleton.yaml")
	helper.WriteFile(initCFSkeleton, blobCFSkeleton)

	initCFSkeleton = "skeleton.json"
	blobCFSkeleton, _ = box.Get("/init/type/clouformation/skeleton.json")
	helper.WriteFile(initCFSkeleton, blobCFSkeleton)

	/* TODO: copy a template to create standard tags for the entire stack easily
	https://docs.aws.amazon.com/cli/latest/reference/cloudformation/create-stack.html
	example aws cloudformation create-stack ... --tags */

	/* TODO: copy Makefile */
	/* TODO: copy LICENSE */
}

/* TERRAFORM PROJECT */

// CreateTerraformProject creates a HashiCorp Terraform project
func CreateTerraformProject(name string) {
	CreateBasicProject(name)
	initCloudProject()
	initTerraformProject()
}

// InitTerraform initialize a project with Terraform structure
func initTerraformProject() {
	initMakefile := "Makefile"
	blobMakefile, _ := box.Get("/init/type/terraform/Makefile")
	helper.WriteFile(initMakefile, blobMakefile)

	initLicense := "LICENSE"
	blobLicense, _ := box.Get("/init/type/terraform/LICENSE")
	helper.WriteFile(initLicense, blobLicense)

	helper.CreateDir("environments")

	initDevEnvironment := "environments/dev.tf"
	blobDevEnvironment, _ := box.Get("/init/type/terraform/environments/dev.tf")
	helper.WriteFile(initDevEnvironment, blobDevEnvironment)

	initProdEnvironment := "environments/prod.tf"
	blobProdEnvironment, _ := box.Get("/init/type/terraform/environments/prod.tf")
	helper.WriteFile(initProdEnvironment, blobProdEnvironment)

	initMainTF := "main.tf"
	blobMainTF, _ := box.Get("/init/type/terraform/main.tf")
	helper.WriteFile(initMainTF, blobMainTF)

	initVariablesTF := "variables.tf"
	blobVariablesTF, _ := box.Get("/init/type/terraform/variables.tf")
	helper.WriteFile(initVariablesTF, blobVariablesTF)

	initOutputsTF := "outputs.tf"
	blobOutputsTF, _ := box.Get("/init/type/terraform/outputs.tf")
	helper.WriteFile(initOutputsTF, blobOutputsTF)
}

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

// func initCreateTerraformProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
// 	helper.createProjectDir(name)
// 	if !onlyCustomizedStructure {
// 		helper.InitBasicProject()
// 		helper.InitHLD(name)
// 		helper.InitTerraform()
// 	}
// 	// helper.InitCustomProjectLayout(typee, "default")
// 	// helper.InitCustomProjectLayout(typee, structure)
// 	// helper.UpdateReadMe()
// }
