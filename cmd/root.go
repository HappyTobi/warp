package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HappyTobi/warp/pkg/cmd/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	oldConfig      = ".warp.yaml"
	configFileName = "warp.yaml"
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

// TODO move to settings package
func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configPath := filepath.Join(home, ".config", "warp")
	configFilePath := filepath.Join(configPath, configFileName)
	configFilePathLegacy := filepath.Join(home, oldConfig)

	migrated := false
	if _, err := os.Stat(configFilePathLegacy); !os.IsNotExist(err) {
		fmt.Println("Configurations file will me moved and updated to new format")
		_ = os.MkdirAll(configPath, os.ModePerm)
		if err := os.Rename(configFilePathLegacy, configFilePath); err != nil {
			fmt.Println("Error while moving config file")
			os.Exit(1)
		}
		os.Remove(configFilePathLegacy)
		migrated = true
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("warp")

	if err := viper.ReadInConfig(); err == nil && !migrated {
		return
	}

	viper.SetDefault("settings.user.firstname", "happy")
	viper.SetDefault("settings.user.lastname", "tobi")
	viper.SetDefault("settings.user.street", "githubroad")
	viper.SetDefault("settings.user.postcode", "0000")
	viper.SetDefault("settings.user.city", "internet")

	viper.SetDefault("settings.date_time.time_zone", "Europe/Berlin")
	viper.SetDefault("settings.date_time.time_format", "15:04:05 02-01-2006")
	viper.SetDefault("settings.power.price", "0.35")

	//csv settings
	viper.SetDefault("csv.comma", ";")
	viper.SetDefault("csv.header", true)

	//pdf settings
	viper.SetDefault("pdf.print_header", false)
	imagePath := filepath.Join(configPath, "logo.png")
	viper.SetDefault("pdf.image_path", imagePath)

	if err := settings.StoreImage(imagePath); err != nil {
		fmt.Print("Error while storing image")
		os.Exit(1)
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		os.WriteFile(configFilePath, []byte{}, 0644)
	}

	if err := viper.WriteConfig(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println("New warp config file created at:", viper.ConfigFileUsed())
}
