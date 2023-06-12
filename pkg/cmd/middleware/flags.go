package middleware

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
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
		return fmt.Errorf("could not find warp charger address")
	}

	req(charger, user, pass, output)

	return nil
}
