package configuration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/HappyTobi/warp/pkg/cmd/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Create(cmd *cobra.Command, args []string) error {
	fmt.Println("New warp config file created at:", viper.ConfigFileUsed())

	return nil
}

func CreateConfigFile(configPath string) error {
	viper.AddConfigPath(filepath.Dir(configPath))

	viper.SetConfigType(strings.ReplaceAll(filepath.Ext(configPath), ".", ""))
	viper.SetConfigFile(filepath.Base(configPath))

	if err := viper.ReadInConfig(); err == nil {
		// file already exist
		return nil
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
	dirName := filepath.Dir(configPath)
	imagePath := filepath.Join(dirName, "logo.png")
	viper.SetDefault("pdf.image_path", imagePath)

	//charger settings
	viper.SetDefault("charger.url", "")
	viper.SetDefault("charger.username", "")
	viper.SetDefault("charger.password", "")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if _, derr := os.Stat(dirName); derr != nil {
			_ = os.MkdirAll(configPath, os.ModePerm)
		}
	}

	if err := settings.StoreImage(imagePath); err != nil {
		fmt.Print("Error while storing image")
		return err
	}

	if err := viper.WriteConfig(); err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
