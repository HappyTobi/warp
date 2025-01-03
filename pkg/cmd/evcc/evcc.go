package evcc

import (
	"fmt"
	"strings"

	"github.com/HappyTobi/warp/pkg/cmd/middleware"
	"github.com/HappyTobi/warp/pkg/internal/evse"
	"github.com/HappyTobi/warp/pkg/internal/nfc"
	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/spf13/cobra"
)

// https://docs.warp-charger.com/docs/mqtt_http/api_reference/evse#evse_state_warp1
func Status(cmd *cobra.Command, args []string) error {
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	evseService := evse.NewEvseService(request)
	state, err := evseService.State()
	if err != nil {
		return err
	}

	//we have to return state A...F from field: iec61851_state
	iecState := state["iec61851_state"]
	if v, ok := iecState.(float64); ok {
		switch v {
		case 0:
			fmt.Print("A")
		case 1:
			fmt.Print("B")
		case 2:
			fmt.Print("C")
		case 3:
			fmt.Print("D")
		case 4:
			fmt.Print("F")
		}

		return nil
	}

	return fmt.Errorf("error while checking charger state")
}

func Enabled(cmd *cobra.Command, args []string) error {
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	evseService := evse.NewEvseService(request)
	current, err := evseService.ReadExternalCurrent()
	if err != nil {
		return err
	}

	enabled := "false"

	if current >= 6000 {
		enabled = "true"
	}

	fmt.Print(enabled)

	return nil
}

func Enable(cmd *cobra.Command, args []string) error {
	enable, err := cmd.Flags().GetString("enable")
	if err != nil {
		return err
	}

	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	nfcTagService := nfc.NewNfcTagsService(request)
	nfcTags, err := nfcTagService.Config()
	if err != nil {
		return err
	}

	userService := users.NewUsersService(request)
	allUsers, err := userService.AllUsernames()
	if err != nil {
		return err
	}

	//validate passed user
	userFilter, _ := cmd.Flags().GetString("user")
	if len(userFilter) == 0 {
		userFilter, _ = cmd.InheritedFlags().GetString("username")
	}

	userTag := nfc.UserTag{}
	for _, user := range allUsers {
		if strings.EqualFold(user.Username, userFilter) {
			for _, tag := range nfcTags.AuthorizedTags {
				if tag.UserId == user.Id {
					userTag = tag
					break
				}
			}
		}
	}

	if len(userTag.Id) == 0 {
		return fmt.Errorf("The passed user is not valid or has no valid nfc tag")
	}

	evseService := evse.NewEvseService(request)

	current := 0
	if strings.EqualFold(enable, "true") {
		current = 6000
	}

	if err := evseService.SetExternalCurrent(current); err != nil {
		return err
	}

	if err := nfcTagService.StartCharging(userTag); err != nil {
		return err
	}

	fmt.Print(enable)
	return nil
}

func MaxCurrent(cmd *cobra.Command, args []string) error {
	currentAmpere, err := cmd.Flags().GetInt("current")
	if err != nil {
		return err
	}

	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	//change 16 -> 16000
	currentAmpere = (currentAmpere * 1e3)

	fmt.Print(currentAmpere)

	evseService := evse.NewEvseService(request)
	return evseService.SetExternalCurrent(currentAmpere)
}