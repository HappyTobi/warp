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
	info.AddCommand(nameCmd())
	info.AddCommand(displayNameCmd())
	info.AddCommand(modulesCmd())
	info.AddCommand(featuresCmd())

	return info
}

func versionCmd() *cobra.Command {
	version := &cobra.Command{
		Use:   "version",
		Short: "Show Warp Charger version or check for updates",
		Long:  "Show Warp Charger version that is currently installed at the Warp Charger. The other option is to check for an update",
	}

	version.AddCommand(updateCmd())
	version.AddCommand(warpmVersionCmd())

	return version
}

func warpmVersionCmd() *cobra.Command {
	warpVersion := &cobra.Command{
		Use:     "warp",
		Short:   "Print Warp Charger version",
		RunE:    info.Version,
		PreRunE: ValidateOutputformat,
	}
	return warpVersion
}

func updateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "update",
		Short:   "Check if a new firmware is available",
		RunE:    info.Update,
		PreRunE: ValidateOutputformat,
	}
}

func nameCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "name",
		Short:   "Print Warp Charger name and type",
		RunE:    info.Name,
		PreRunE: ValidateOutputformat,
	}
}

func displayNameCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "display-name",
		Short:   "Print Warp Charger display name",
		RunE:    info.DisplayName,
		PreRunE: ValidateOutputformat,
	}
}

func modulesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "modules",
		Short:   "Print Warp Charger firmware modules",
		RunE:    info.Modules,
		PreRunE: ValidateOutputformat,
	}
}

func featuresCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "features",
		Short:   "Print Warp Charger hardware features",
		RunE:    info.Features,
		PreRunE: ValidateOutputformat,
	}
}
