package info

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/tools"
	"github.com/HappyTobi/warp/pkg/internal/info"
	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func Name(cmd *cobra.Command, args []string) error {
	request := &warp.Request{
		Path:        "info/name",
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

	info := info.NewInfo(request)
	name, err := info.LoadName()
	if err != nil {
		return err
	}

	fmt.Print(renderer.JsonInterface(name))

	return nil
}
