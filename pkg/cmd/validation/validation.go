package validation

import (
	"encoding/json"
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/middleware"
	"github.com/HappyTobi/warp/pkg/internal/evse"
	"github.com/HappyTobi/warp/pkg/internal/info"
	"github.com/spf13/cobra"
)

// ValidateWarpFeatures validates if the charger has the given features
func ValidateWarpFeatures(cmd *cobra.Command, args []string, features ...string) error {
	//validate if warp charger has an nfc bricklet
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	infoService := info.NewInfoService(request)
	data, err := infoService.Features()
	if err != nil {
		return err
	}

	var featuresData interface{}
	if err := json.Unmarshal(data, &featuresData); err != nil {
		return err
	}

	foundFeatures := 0
	for _, f := range featuresData.([]interface{}) {
		for _, feature := range features {
			if f == feature {
				foundFeatures++
			}
		}
	}

	if foundFeatures == len(features) {
		return nil
	}

	return fmt.Errorf("some feature is not supported by this charger")
}

// ValidateChargerConnected validates if the charger is connected to a car
func ValidateChargerConnected(cmd *cobra.Command, args []string, features ...string) (int, error) {
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return -1, err
	}

	evseService := evse.NewEvseService(request)
	state, err := evseService.State()
	if err != nil {
		return -1, err
	}

	chargerState := state["charger_state"].(float64)
	return int(chargerState), nil
}
