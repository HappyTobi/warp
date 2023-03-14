package info

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/tools"
	"github.com/HappyTobi/warp/pkg/internal/info"
	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func Update(cmd *cobra.Command, args []string) error {
	requests := make([]*warp.Request, 0, 2)

	nameRequest := &warp.Request{
		Path:        "info/name",
		ContentType: warp.JSON,
	}
	versionRequest := &warp.Request{
		Path:        "info/version",
		ContentType: warp.JSON,
	}

	githubRequest := &warp.Request{
		Path:        "tags",
		ContentType: warp.JSON,
		Warp:        "https://api.github.com/repos/Tinkerforge/esp32-firmware",
	}

	requests = append(requests, nameRequest, versionRequest)

	if err := tools.LoadGlobalParams(cmd, func(charger, username, password, output string) {
		for _, req := range requests {
			req.Warp = charger

			if len(username) > 0 && len(password) > 0 {
				req.Username = username
				req.Password = password
			}

			req.OutputRenderer = renderer.NewRenderer(output)
		}
	}); err != nil {
		return err
	}

	nameResp, err := nameRequest.GetJson()
	if err != nil {
		return err
	}

	versionResp, err := versionRequest.GetJson()
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

	updateInfo := info.NewUpdate(githubRequest, warpVersion, firmwareVersion)
	resp, err := updateInfo.UpdateAvailable()
	if err != nil {
		return err
	}

	fmt.Print(nameRequest.OutputRenderer.Render(resp))

	return nil
}
