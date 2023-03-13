package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	root := &cobra.Command{
		Use: "warp",
	}

	root.PersistentFlags().StringP("charger", "c", "", "Url to Warp charger like http://192.168.1.2 or http://warp.local")
	root.PersistentFlags().StringP("output", "o", "json", "Output format (json, yaml, csv)")

	root.PersistentFlags().StringP("username", "u", "", "Username to authenticate (required if password is set)")
	root.PersistentFlags().StringP("password", "p", "", "Password to authenticate (required if username is set)")
	root.MarkFlagsRequiredTogether("username", "password")

	root.AddCommand(InfoCmd())
	root.AddCommand(UserCmd())
	root.AddCommand(ChargeTrackerCmd())

	return root
}
