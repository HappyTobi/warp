package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ValidateOutputformat(cmd *cobra.Command, args []string) error {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	if output != "json" && output != "yaml" {
		return fmt.Errorf("invalid output format: %s", output)
	}

	return nil
}
