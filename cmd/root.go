package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/configuration"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use: "warp",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			//load the configuration file
			configPath, _ := cmd.Flags().GetString("config")
			return configuration.ReadConfig(configPath)
		},
	}

	root.PersistentFlags().StringP("charger", "c", "", "Url to Warp charger like http://192.168.1.2 or http://warp.local")
	root.PersistentFlags().StringP("output", "o", "json", "Output format (json, yaml)")

	root.PersistentFlags().String("config", "", "Path to the warp configuration file (default $HOME/.config/warp/warp.yaml)")

	root.PersistentFlags().StringP("username", "u", "", "Username to authenticate (required if password is set)")
	root.PersistentFlags().StringP("password", "p", "", "Password to authenticate (required if username is set)")
	root.MarkFlagsRequiredTogether("username", "password")

	root.AddCommand(ChargeCmd())
	root.AddCommand(InfoCmd())
	root.AddCommand(UserCmd())
	root.AddCommand(ChargeTrackerCmd())
	root.AddCommand(MeterCmd())
	root.AddCommand(EvccCmd())
	root.AddCommand(VersionCmd())
	root.AddCommand(ConfigurationCmd())

	return root
}

func init() {
	// enable all pre functions
	cobra.EnableTraverseRunHooks = true
}
