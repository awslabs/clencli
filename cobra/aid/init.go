package aid

import (
	"fmt"
	"os"

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
		logrus.Errorf("unable to initialize cloudformation project")
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
	if err := CreateBasicProject(cmd, name); err != nil {
		return nil
	}

	if initialized := initCloudProject(); !initialized {
		logrus.Errorf("unable to initialize terraform project")
		return fmt.Errorf("error: unable to initalize project \"%s\"", name)
	}

	if initialized := initTerraformProject(); !initialized {
		logrus.Errorf("unable to initialize cloud project")
		return fmt.Errorf("error: unable to initalize project \"%s\"", name)
	}

	return nil
}

// InitTerraform initialize a project with Terraform structure
func initTerraformProject() bool {
	a := h.WriteFileFromBox("/init/project/type/terraform/Makefile", "Makefile")
	b := h.WriteFileFromBox("/init/project/type/terraform/LICENSE", "LICENSE")

	c := h.MkDirsIfNotExist("environments")
	d := h.WriteFileFromBox("/init/project/type/terraform/environments/dev.tf", "environments/dev.tf")
	e := h.WriteFileFromBox("/init/project/type/terraform/environments/prod.tf", "environments/prod.tf")

	f := h.WriteFileFromBox("/init/project/type/terraform/main.tf", "main.tf")
	g := h.WriteFileFromBox("/init/project/type/terraform/variables.tf", "variables.tf")
	h := h.WriteFileFromBox("/init/project/type/terraform/outputs.tf", "outputs.tf")

	return (a && b && c && d && e && f && g && h)

}

// TODO: allow users to inform additional files to be added to their project initialization
