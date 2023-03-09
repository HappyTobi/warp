package info

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/tools"
	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func Version(cmd *cobra.Command, args []string) error {
	request := &warp.Request{
		Path:        "info/version",
		ContentType: warp.JSON,
	}

	if err := tools.LoadGlobalParams(cmd, func(charger, username, password string) {
		request.Warp = charger

		if len(username) > 0 && len(password) > 0 {
			request.Username = username
			request.Password = password
		}

	}); err != nil {
		return err
	}

	js, err := request.GetJson()
	if err != nil {
		return err
	}

	fmt.Print(renderer.JsonInterface(js))

	return nil
}
