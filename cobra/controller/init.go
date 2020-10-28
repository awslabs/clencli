/*
Copyright © 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"errors"
	"fmt"
	"log"

	helper "github.com/awslabs/clencli/helper"
	"github.com/spf13/cobra"
)

var initValidArgs = []string{"project"}

// InitCmd command to initialize projects
func InitCmd() *cobra.Command {
	man := helper.GetManual("init")
	return &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		ValidArgs: initValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   initPreRun,
		RunE:      initRun,
	}
}

func initPreRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("one the following arguments are required: %s", initValidArgs)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil || len(name) == 0 {
		// flag accessed but not defined
		return errors.New("required flag name not set")
	}

	// ensure the project types
	if args[0] == "project" {
		t, err := cmd.Flags().GetString("type")
		// flag accessed but not defined
		if err != nil {
			return errors.New("Project type must be defined")
		}
		if t == "" {
			return errors.New("Project type must be provided")
		}

		if t != "basic" && t != "cloudformation" && t != "terraform" {
			return fmt.Errorf("Unknown project type provided: %s", t)
		}

	}

	return nil
}

func initRun(cmd *cobra.Command, args []string) error {
	return nil
}

// func initRun(cmd *cobra.Command, args []string) error {
// 	name, typee, structure, onlyCustomizedStructure := initGetFlags(cmd)

// 	if args[0] == "project" {
// 		switch typee {
// 		case "basic":
// 			initCreateBasicProject(name, typee, structure, onlyCustomizedStructure)
// 		case "cloudformation":
// 			initCreateCloudFormationProject(name, typee, structure, onlyCustomizedStructure)
// 		case "terraform":
// 			initCreateTerraformProject(name, typee, structure, onlyCustomizedStructure)

// 		default:
// 			return errors.New("Unknow project type")
// 		}
// 	} else {
// 		return errors.New("invalid argument")
// 	}

// 	return nil
// }

func initGetFlags(cmd *cobra.Command) (name string, typee string, structure string, onlyCustomizedStructure bool) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatal("required flag name not set")
	}

	typee, _ = cmd.Flags().GetString("type")
	structure, _ = cmd.Flags().GetString("structure")
	onlyCustomizedStructure, _ = cmd.Flags().GetBool("only-customized-structure")
	return name, typee, structure, onlyCustomizedStructure
}

// func initCreateBasicProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
// 	helper.InitializeProject(name)
// 	if !onlyCustomizedStructure {
// 		helper.InitBasic()
// 	}
// 	// helper.InitCustomProjectLayout(typee, "default")
// 	// helper.InitCustomProjectLayout(typee, structure)
// 	// helper.UpdateReadMe()
// }

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

// // InitializeProject creates the project directory, change the current directory and places basic configuration files
// func InitializeProject(name string) {

// 	// Create the name directory
// 	CreateDir(name)

// 	// Change current directory to name directory
// 	os.Chdir(name)
// }

// // InitBasic create the basic configuration files
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

// InitCustomProjectLayout generates
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

// InitCloudFormation initialize a project with CloudFormation structure
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
