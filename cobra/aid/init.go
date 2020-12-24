package aid

import (
	"fmt"
	"os"

	"github.com/awslabs/clencli/box"
	"github.com/awslabs/clencli/helper"
	h "github.com/awslabs/clencli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/* BASIC PROJECT */

// CreateBasicProject creates a basic project
func CreateBasicProject(cmd *cobra.Command, name string) error {
	err := createAndEnterProjectDir(name)
	if err != nil {
		return err
	}

	if initalized := initProject(); !initalized {
		logrus.Errorf("unable to initialize basic project")
		return fmt.Errorf("error: unable to initalize project \"%s\"", name)
	}

	return nil
}

func createAndEnterProjectDir(name string) error {

	if !helper.MkDirsIfNotExist(name) {
		return fmt.Errorf("error: unable to create directory %s", name)
	}

	err := os.Chdir(name)
	if err != nil {
		return fmt.Errorf("error: unable to enter directory %s", name)
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error: unable to returns a rooted path name corresponding to the current directory:\n%s", err.Error())
	}
	logrus.Infof("current working directory changed to %s", wd)

	return nil
}

// create the basic configuration files
func initProject() bool {

	// Create a directory for CLENCLI
	a := h.MkDirsIfNotExist("clencli")
	b := h.WriteFileFromBox("/init/clencli/readme.yaml", "clencli/readme.yaml")
	c := h.WriteFileFromBox("/init/clencli/readme.tmpl", "clencli/readme.tmpl")
	d := h.WriteFileFromBox("/init/.gitignore", ".gitignore")

	return (a && b && c && d)

}

/* CLOUD PROJECT */

// CreateCloudProject copies the necessary templates for cloud projects
func CreateCloudProject(cmd *cobra.Command, name string) error {
	if err := CreateBasicProject(cmd, name); err != nil {
		return nil
	}

	if initialized := initCloudProject(); !initialized {
		logrus.Errorf("unable to initialize cloud project")
		return fmt.Errorf("error: unable to initalize project \"%s\"", name)
	}

	return nil
}

// copies the High Level Design template file
func initCloudProject() bool {
	a := h.WriteFileFromBox("/init/clencli/hld.yaml", "clencli/hld.yaml")
	b := h.WriteFileFromBox("/init/clencli/hld.tmpl", "clencli/hld.tmpl")

	return (a && b)
}

/* CLOUDFORMATION PROJECT */

// CreateCloudFormationProject creates an AWS CloudFormation project
func CreateCloudFormationProject(cmd *cobra.Command, name string) error {
	if err := CreateBasicProject(cmd, name); err != nil {
		return nil
	}

	if initialized := initCloudProject(); !initialized {
		logrus.Errorf("unable to initialize cloud project")
		return fmt.Errorf("error: unable to initalize project \"%s\"", name)
	}

	if initialized := initCloudFormationProject(); !initialized {
		logrus.Errorf("unable to initialize cloud project")
		return fmt.Errorf("error: unable to initalize project \"%s\"", name)
	}

	return nil
}

// initialize a project with CloudFormation structure and copies template files
func initCloudFormationProject() bool {

	a := h.MkDirsIfNotExist("environments")
	b := h.MkDirsIfNotExist("environments/dev")
	c := h.MkDirsIfNotExist("environments/prod")
	d := h.WriteFileFromBox("/init/project/type/clouformation/skeleton.yaml", "skeleton.yaml")
	e := h.WriteFileFromBox("/init/project/type/clouformation/skeleton.json", "skeleton.json")

	/* TODO: copy a template to create standard tags for the entire stack easily
	https://docs.aws.amazon.com/cli/latest/reference/cloudformation/create-stack.html
	example aws cloudformation create-stack ... --tags */

	/* TODO: copy Makefile */
	/* TODO: copy LICENSE */

	return (a && b && c && d && e)
}

/* TERRAFORM PROJECT */

// CreateTerraformProject creates a HashiCorp Terraform project
func CreateTerraformProject(cmd *cobra.Command, name string) error {
	CreateBasicProject(cmd, name)
	initCloudProject()
	initTerraformProject()
	return nil
}

// InitTerraform initialize a project with Terraform structure
func initTerraformProject() {
	initMakefile := "Makefile"
	blobMakefile, _ := box.Get("/init/type/terraform/Makefile")
	h.WriteFile(initMakefile, blobMakefile)

	initLicense := "LICENSE"
	blobLicense, _ := box.Get("/init/type/terraform/LICENSE")
	h.WriteFile(initLicense, blobLicense)

	h.MkDirsIfNotExist("environments")

	initDevEnvironment := "environments/dev.tf"
	blobDevEnvironment, _ := box.Get("/init/type/terraform/environments/dev.tf")
	h.WriteFile(initDevEnvironment, blobDevEnvironment)

	initProdEnvironment := "environments/prod.tf"
	blobProdEnvironment, _ := box.Get("/init/type/terraform/environments/prod.tf")
	h.WriteFile(initProdEnvironment, blobProdEnvironment)

	initMainTF := "main.tf"
	blobMainTF, _ := box.Get("/init/type/terraform/main.tf")
	h.WriteFile(initMainTF, blobMainTF)

	initVariablesTF := "variables.tf"
	blobVariablesTF, _ := box.Get("/init/type/terraform/variables.tf")
	h.WriteFile(initVariablesTF, blobVariablesTF)

	initOutputsTF := "outputs.tf"
	blobOutputsTF, _ := box.Get("/init/type/terraform/outputs.tf")
	h.WriteFile(initOutputsTF, blobOutputsTF)
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
// 						MkDirsIfNotExist(f.File.Path)
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
// 	h.createAndEnter(name)
// 	if !onlyCustomizedStructure {
// 		h.initProject()
// 		h.InitHLD(name)
// 		h.InitTerraform()
// 	}
// 	// h.InitCustomProjectLayout(typee, "default")
// 	// h.InitCustomProjectLayout(typee, structure)
// 	// h.UpdateReadMe()
// }
