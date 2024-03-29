package cmd

import (
	"fmt"

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
	chargeLogCmd := &cobra.Command{
		Use:     "log",
		Short:   "Download complete charge-tracker log",
		RunE:    chargeTracker.ChargeLog,
		PreRunE: validateCustomOutputFormat,
	}

	chargeLogCmd.Flags().StringP("output", "o", "json", "Output format (json, yaml,csv)")
	chargeLogCmd.Flags().StringP("file", "f", "", "Output file (default: stdout)")
	chargeLogCmd.Flags().StringP("user", "r", "", "Filter by username (case-insensitive)")
	chargeLogCmd.Flags().IntP("month", "m", -1, "Filter by month (format: 1...12)")
	chargeLogCmd.Flags().IntP("year", "y", -1, "Filter by year (format: YYYY)")

	return chargeLogCmd

}

func validateCustomOutputFormat(cmd *cobra.Command, args []string) error {
	output, _ := cmd.Flags().GetString("output")
	file, _ := cmd.Flags().GetString("file")

	if output != "json" && output != "yaml" && output != "csv" && output != "pdf" {
		return fmt.Errorf("invalid output format: %s", output)
	}

	if output == "csv" && file == "" {
		return fmt.Errorf("output format csv requires a file to be set")
	}

	if output == "pdf" && file == "" {
		return fmt.Errorf("output format pdf requires a file to be set")
	}

	return nil
}
