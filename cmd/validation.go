package cmd

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/validation"
	"github.com/spf13/cobra"
)

func ValidateOutputformat(cmd *cobra.Command, args []string) error {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	if output != "json" && output != "yaml" && output != "yml" {
		return fmt.Errorf("invalid output format: %s", output)
	}

	return nil
}

// validate charge start command
func ValidateChargeStart(cmd *cobra.Command, args []string) error {
	chargePower, err := cmd.Flags().GetInt("charge-power")
	if err != nil {
		return err
	}
	//validate charge-power flag if it is between 6000 and 32000
	if chargePower != -1 && (chargePower < 6000 || chargePower > 32000) {
		return fmt.Errorf("invalid charge-power: %d, value has to be between 6000 and 32000", chargePower)
	}

	if err := validation.ValidateWarpFeatures(cmd, args, "nfc", "evse"); err != nil {
		return err
	}

	connectionState, err := validation.ValidateChargerConnected(cmd, args)
	if err != nil {
		return err

	}
	if connectionState == 0 {
		return fmt.Errorf("Charger is not connected to a car, charging can't be started")
	}

	return nil
}

// validate charge stop command
func ValidateChargeStop(cmd *cobra.Command, args []string) error {
	if err := validation.ValidateWarpFeatures(cmd, args, "nfc", "evse"); err != nil {
		return err
	}

	connectionState, err := validation.ValidateChargerConnected(cmd, args)
	if err != nil {
		return err
	}

	if connectionState > 1 {
		return nil
	}

	return fmt.Errorf("Charger is not in correct state to stop charging")
}
