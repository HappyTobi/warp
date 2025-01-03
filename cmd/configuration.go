package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/configuration"
	"github.com/spf13/cobra"
)

func ConfigurationCmd() *cobra.Command {
	configuration := &cobra.Command{
		Use:   "configuration",
		Short: "Create warp cli configuration file",
		RunE:  configuration.Create,
	}

	return configuration
}
