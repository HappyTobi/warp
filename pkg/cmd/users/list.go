package users

import (
	"encoding/json"
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/middleware"
	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) error {
	request, err := middleware.LoadWarpRequest(cmd)
	if err != nil {
		return err
	}

	usersList := users.NewUsersService(request)
	list, err := usersList.AllUsernames()
	if err != nil {
		return err
	}

	data, err := json.Marshal(list)
	if err != nil {
		return err
	}

	fmt.Print(request.OutputRenderer.RenderBytes(data))

	return nil
}
