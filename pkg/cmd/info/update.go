package info

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/middleware"
	"github.com/HappyTobi/warp/pkg/internal/info"
	"github.com/HappyTobi/warp/pkg/internal/update"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func Update(cmd *cobra.Command, args []string) error {
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	githubRequest, _ := middleware.LoadWarpRequest(cmd)
	githubRequest.Path = "tags"
	githubRequest.Warp = "https://api.github.com/repos/Tinkerforge/esp32-firmware"

	infoService := info.NewInfoService(request)
	nameResp, err := infoService.Name()
	if err != nil {
		return err
	}

	versionResp, err := infoService.Version()
	if err != nil {
		return err
	}

	warpType, ok := nameResp["type"].(string)
	if !ok {
		return fmt.Errorf("could not get warp type")
	}

	firmware, ok := versionResp["firmware"].(string)
	if !ok {
		return fmt.Errorf("could not get warp firmware version")
	}

	warpVersion := warp.WarpType(warpType)
	firmwareVersion := warp.WarpFirmware(firmware)

	updateInfo := update.NewUpdate(githubRequest, warpVersion, firmwareVersion)
	resp, err := updateInfo.UpdateAvailable()
	if err != nil {
		return err
	}

	fmt.Print(request.OutputRenderer.Render(resp))

	return nil
}
