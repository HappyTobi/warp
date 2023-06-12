package info

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/middleware"
	"github.com/HappyTobi/warp/pkg/internal/info"
	"github.com/spf13/cobra"
)

func Name(cmd *cobra.Command, args []string) error {
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	infoService := info.NewInfoService(request)
	js, err := infoService.Name()
	if err != nil {
		return err
	}

	fmt.Print(request.OutputRenderer.Render(js))

	return nil
}
