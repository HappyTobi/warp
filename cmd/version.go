package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/version"
	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print warp cli version",
		RunE:  version.Print,
	}
}
