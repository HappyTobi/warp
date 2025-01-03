package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/evcc"
	"github.com/spf13/cobra"
)

func EvccCmd() *cobra.Command {
	evcc := &cobra.Command{
		Use:   "evcc",
		Short: "Evcc command",
		Long:  "Evcc command to integration the warp chager with user authentication into evcc",
	}

	evcc.AddCommand(statusCmd())
	evcc.AddCommand(enabledCmd())
	evcc.AddCommand(enableCmd())
	evcc.AddCommand(maxcurrentCmd())

	return evcc
}

func statusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Return status of the charger for evcc",
		RunE:  evcc.Status,
	}
}

func enabledCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "enabled",
		Short: "Return enabled status of the charger for evcc",
		RunE:  evcc.Enabled,
	}
}

func enableCmd() *cobra.Command {
	enableCmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable the charger for evcc",
		RunE:  evcc.Enable,
	}

	enableCmd.Flags().StringP("user", "r", "", "User to start charging (default is provided username)")
	enableCmd.Flags().String("enable", "false", "Enable charging")

	return enableCmd
}

func maxcurrentCmd() *cobra.Command {
	maxcurrentCmd := &cobra.Command{
		Use:   "maxcurrent",
		Short: "Set the maxcurrent apaire for the charger from evcc",
		RunE:  evcc.MaxCurrent,
	}

	maxcurrentCmd.Flags().Int("current", -1, "Currente ampere")

	return maxcurrentCmd
}
