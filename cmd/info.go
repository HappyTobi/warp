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
		Use:     "version",
		Short:   "Print Warp Charger version",
		RunE:    info.Version,
		PreRunE: ValidateOutputformat,
	}

	version.AddCommand(updateCmd())

	return version
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
