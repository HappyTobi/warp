package middleware

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadWarpRequest(cmd *cobra.Command) (warp.Request, error) {
	request := warp.Request{
		ContentType: warp.JSON,
	}

	return request, LoadGlobalParams(cmd, func(charger, username, password, output string) {
		request.Warp = charger

		if len(username) > 0 && len(password) > 0 {
			request.Username = username
			request.Password = password
		}

		request.OutputRenderer = renderer.NewRenderer(output)
	})
}

func LoadGlobalParams(cmd *cobra.Command, req func(charger, username, password, output string)) error {
	charger, _ := cmd.Root().Flags().GetString("charger")

	user, _ := cmd.Root().Flags().GetString("username")
	pass, _ := cmd.Root().Flags().GetString("password")

	output, _ := cmd.Root().Flags().GetString("output")

	if len(charger) == 0 {
		//load from config
		charger = viper.GetString("charger.url")
		if len(charger) == 0 {
			return fmt.Errorf("could not find warp charger address as argument or part of the configuration file %s", viper.ConfigFileUsed())
		}
	}

	// load user and pass from config when empty, empty values are fine
	if len(user) == 0 {
		user = viper.GetString("charger.username")
	}

	if len(pass) == 0 {
		pass = viper.GetString("charger.password")
	}

	req(charger, user, pass, output)

	return nil
}
