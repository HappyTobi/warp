package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/chargeTracker"
	"github.com/spf13/cobra"
)

func ChargeTrackerCmd() *cobra.Command {
	chargeTracker := &cobra.Command{
		Use:   "charge-tracker",
		Short: "charge-tracker command to download the charge protocoll",
	}
	chargeTracker.AddCommand(ChargeLogCmd())

	return chargeTracker
}

func ChargeLogCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "log",
		Short: "Download complete charge-tracker log",
		RunE:  chargeTracker.ChargeLog,
	}
}
