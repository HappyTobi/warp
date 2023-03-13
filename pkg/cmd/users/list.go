package users

import (
	"encoding/json"
	"fmt"

	"github.com/HappyTobi/warp/pkg/cmd/tools"
	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) error {
	request := &warp.Request{
		Path:        "users/all_usernames",
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

	usersList := users.NewUsersList(request)
	list, err := usersList.Load()
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
