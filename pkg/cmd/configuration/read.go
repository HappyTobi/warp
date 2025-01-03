package configuration

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func ReadConfig(configPathArg string) error {

	configPath := configPathArg
	if len(configPathArg) == 0 {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configPath = filepath.Join(home, ".config", "warp", "warp.yaml")
	}

	// TODO add validation that file base and extension exists

	absConfigFilePath, _ := filepath.Abs(configPath)

	if err := CreateConfigFile(absConfigFilePath); err != nil {
		return err
	}

	return nil
}
