package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use: "warp",
	}

	root.PersistentFlags().StringP("charger", "c", "", "Url to Warp charger like http://192.168.1.2 or http://warp.local")
	root.PersistentFlags().StringP("output", "o", "json", "Output format (json, yaml)")

	root.PersistentFlags().StringP("username", "u", "", "Username to authenticate (required if password is set)")
	root.PersistentFlags().StringP("password", "p", "", "Password to authenticate (required if username is set)")
	root.MarkFlagsRequiredTogether("username", "password")

	root.AddCommand(InfoCmd())
	root.AddCommand(UserCmd())
	root.AddCommand(ChargeTrackerCmd())
	root.AddCommand(MeterCmd())
	root.AddCommand(VersionCmd())

	return root
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".warp")

	if err := viper.ReadInConfig(); err == nil {
		return
	}

	viper.SetDefault("date_time.time_zone", "Europe/Berlin")
	viper.SetDefault("date_time.time_format", "15:04:05 02-01-2006")
	viper.SetDefault("power.price", "0.35")
	viper.SetDefault("csv.comma", ";")
	viper.SetDefault("csv.header", true)
	if err := viper.SafeWriteConfig(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println("New warp config file created at:", viper.ConfigFileUsed())
}
