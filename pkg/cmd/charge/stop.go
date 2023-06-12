package charge

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/middleware"
	"github.com/HappyTobi/warp/pkg/internal/nfc"
	"github.com/spf13/cobra"
)

func Stop(cmd *cobra.Command, args []string) error {
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	nfcTagService := nfc.NewNfcTagsService(request)
	if err := nfcTagService.StopCharging(); err != nil {
		return err
	}

	fmt.Printf("Charge stopped")
	return nil
}
