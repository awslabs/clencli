package controller

import (
	"fmt"
	"runtime"

	"github.com/awslabs/clencli/box"
	cau "github.com/awslabs/clencli/cauldron"
	"github.com/spf13/cobra"
)

func versionRun(cmd *cobra.Command, args []string) error {
	// Get the version defined in the VERSION file
	version, status := box.Get("/VERSION")
	if status {
		goOS := runtime.GOOS
		goVersion := runtime.Version()
		goArch := runtime.GOARCH

		fmt.Printf("CLENCLI v%s %s %s %s\n", version, goVersion, goOS, goArch)
	} else {
		return fmt.Errorf("Version not available")
	}
	return nil
}

// VersionCmd command to display CLENCLI current version
func VersionCmd() *cobra.Command {
	man := cau.GetManual("version")

	return &cobra.Command{
		Use:   man.Use,
		Short: man.Short,
		Long:  man.Long,
		RunE:  versionRun,
	}
}