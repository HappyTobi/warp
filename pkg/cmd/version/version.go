package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
)

func Print(cmd *cobra.Command, args []string) error {
	if len(version) == 0 {
		version = "dev"
	}
	fmt.Printf("Version: \t %s \n", version)
	return nil
}
