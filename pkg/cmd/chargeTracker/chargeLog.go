package chargeTracker

import (
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/tools"
	"github.com/HappyTobi/warp/pkg/internal/chargeTracker"
	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func ChargeLog(cmd *cobra.Command, args []string) error {
	requests := make([]*warp.Request, 0, 2)

	request := &warp.Request{
		Path:        "charge_tracker/charge_log",
		ContentType: warp.JSON,
	}

	userRequest := &warp.Request{
		Path:        "users/all_usernames",
		ContentType: warp.JSON,
	}

	requests = append(requests, request, userRequest)

	if err := tools.LoadGlobalParams(cmd, func(charger, username, password string) {
		for _, req := range requests {
			req.Warp = charger

			if len(username) > 0 && len(password) > 0 {
				req.Username = username
				req.Password = password
			}
		}
	}); err != nil {
		return err
	}

	chargeLog := chargeTracker.NewChargeLog(request)
	user := users.NewUsersList(userRequest)
	users, _ := user.Load()

	charges, err := chargeLog.Load(users)
	if err != nil {
		return err
	}

	fmt.Print(renderer.JsonInterface(charges))
	return nil
}
