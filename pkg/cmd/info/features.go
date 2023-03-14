package info

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/tools"
	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func Features(cmd *cobra.Command, args []string) error {
	request := &warp.Request{
		Path:        "info/features",
		ContentType: warp.JSON,
	}

	if err := tools.LoadGlobalParams(cmd, func(charger, username, password, output string) {
		request.Warp = charger

		if len(username) > 0 && len(password) > 0 {
			request.Username = username
			request.Password = password
		}

		request.OutputRenderer = renderer.NewRenderer(output)

	}); err != nil {
		return err
	}

	data, err := request.Get()
	if err != nil {
		return err
	}

	fmt.Print(request.OutputRenderer.RenderBytes(data))

	return nil
}
