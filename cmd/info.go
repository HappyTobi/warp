package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/info"
	"github.com/spf13/cobra"
)

func InfoCmd() *cobra.Command {
	info := &cobra.Command{
		Use:   "info",
		Short: "Info command",
		Long:  "Info command print out information about WARP Charger",
	}
	info.AddCommand(versionCmd())

	return info
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print Warp Charger version",
		RunE:  info.Version,
	}
}
