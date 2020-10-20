package controller

import (
	"errors"
	"fmt"
	"log"

	cau "github.com/awslabs/clencli/cauldron"
	"github.com/spf13/cobra"
)

var initValidArgs = []string{"project"}

// InitCmd command to initialize projects
func InitCmd() *cobra.Command {
	man := cau.GetManual("init")
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
	name, typee, structure, onlyCustomizedStructure := initGetFlags(cmd)

	if args[0] == "project" {
		switch typee {
		case "basic":
			initCreateBasicProject(name, typee, structure, onlyCustomizedStructure)
		case "cloudformation":
			initCreateCloudFormationProject(name, typee, structure, onlyCustomizedStructure)
		case "terraform":
			initCreateTerraformProject(name, typee, structure, onlyCustomizedStructure)

		default:
			return errors.New("Unknow project type")
		}
	} else {
		return errors.New("invalid argument")
	}

	return nil
}

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

func initCreateBasicProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
	cau.InitializeProject(name)
	if !onlyCustomizedStructure {
		cau.InitBasic()
	}
	// cau.InitCustomProjectLayout(typee, "default")
	// cau.InitCustomProjectLayout(typee, structure)
	// cau.UpdateReadMe()
}

func initCreateCloudFormationProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
	cau.InitializeProject(name)
	if !onlyCustomizedStructure {
		cau.InitBasic()
		cau.InitHLD(name)
		cau.InitCloudFormation()
	}
	// cau.InitCustomProjectLayout(typee, "default")
	// cau.InitCustomProjectLayout(typee, structure)
	// cau.UpdateReadMe()
}

func initCreateTerraformProject(name string, typee string, structure string, onlyCustomizedStructure bool) {
	cau.InitializeProject(name)
	if !onlyCustomizedStructure {
		cau.InitBasic()
		cau.InitHLD(name)
		cau.InitTerraform()
	}
	// cau.InitCustomProjectLayout(typee, "default")
	// cau.InitCustomProjectLayout(typee, structure)
	// cau.UpdateReadMe()
}
