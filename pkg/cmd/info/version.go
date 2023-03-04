package info

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func Version(cmd *cobra.Command, args []string) error {
	request := &warp.Request{
		Path:        "info/name",
		ContentType: warp.JSON,
	}

	charger, _ := cmd.Root().Flags().GetString("charger")
	request.Warp = charger

	//move out and wrap in inject
	user, _ := cmd.Root().Flags().GetString("username")
	pass, _ := cmd.Root().Flags().GetString("password")

	if len(user) > 0 && len(pass) > 0 {
		request.Username = user
		request.Password = pass
	}

	data, err := request.Get()
	if err != nil {
		return err
	}

	// TODO add more renderers
	fmt.Print(renderer.PrettyJson(data))

	return nil
}
