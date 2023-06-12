package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/charge"
	"github.com/spf13/cobra"
)

func ChargeCmd() *cobra.Command {
	users := &cobra.Command{
		Use:   "charge",
		Short: "Charge command",
	}
	users.AddCommand(StartCmd())
	users.AddCommand(StopCmd())

	return users
}

func StartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Use:     "start",
		Short:   "Start charging for a specified user",
		RunE:    charge.Start,
		PreRunE: ValidateChargeStart,
	}

	startCmd.Flags().StringP("user", "r", "", "User to start charging")
	startCmd.Flags().IntP("charge-power", "a", -1, "Ampere to charge with")
	_ = startCmd.MarkFlagRequired("user")

	return startCmd
}

func StopCmd() *cobra.Command {
	stopCmd := &cobra.Command{
		Use:     "stop",
		Short:   "Stop charging",
		RunE:    charge.Stop,
		PreRunE: ValidateChargeStop,
	}

	return stopCmd
}
