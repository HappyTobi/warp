package charge

import (
	"fmt"
	"strings"

	"github.com/HappyTobi/warp/pkg/cmd/middleware"
	"github.com/HappyTobi/warp/pkg/internal/evse"
	"github.com/HappyTobi/warp/pkg/internal/nfc"
	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/spf13/cobra"
)

func Start(cmd *cobra.Command, args []string) error {
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

	evseService := evse.NewEvseService(request)
	chargePower, err := evseService.CurrentChargePower()
	if err != nil {
		return err
	}

	//validate passed user
	userTag := nfc.UserTag{}
	userFilter, _ := cmd.Flags().GetString("user")
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

	//check if chargepower is set and changed
	cPowerArg, _ := cmd.Flags().GetInt("charge-power")
	if cPowerArg > 0 && cPowerArg != chargePower.Current {
		//update chargepower at warp charger
		if err := evseService.UpdateChargePower(cPowerArg); err != nil {
			return err
		}

		chargePower.Current = cPowerArg
	}

	if err := nfcTagService.StartCharging(userTag); err != nil {
		return err
	}

	//fmt.Printf(request.OutputRenderer.Render()
	fmt.Printf("Charge started for user %s with %d Ampere", userFilter, chargePower.Current)
	return nil
}
