package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/meter"
	"github.com/spf13/cobra"
)

func MeterCmd() *cobra.Command {
	meter := &cobra.Command{
		Use:   "meter",
		Short: "Meter command",
	}
	meter.AddCommand(valuesCmd())

	return meter
}

func valuesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "values",
		Short:   "Return measurement of meter",
		RunE:    meter.Values,
		PreRunE: ValidateOutputformat,
	}
}
