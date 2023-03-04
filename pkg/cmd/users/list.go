package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) error {
	request := &warp.Request{
		Path:        "users/all_usernames",
		ContentType: warp.JSON,
	}

	// move out
	charger, _ := cmd.Root().Flags().GetString("charger")
	request.Warp = charger

	user, _ := cmd.Root().Flags().GetString("username")
	pass, _ := cmd.Root().Flags().GetString("password")

	if len(user) > 0 && len(pass) > 0 {
		request.Username = user
		request.Password = pass
	}

	data, err := request.Get()
	if err != nil {
		return err
	}

	jsonData := &UserListResponse{Users: make([]User, 0)}
	userData := &User{}

	reader := bytes.NewReader(data)
	buf := make([]byte, 32)
	var i = 0
	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		// remove all null chars from buffer
		b := bytes.Trim(buf, "\x00")
		if i == 0 {
			userData.Username = string(b)
			i++
		} else {
			userData.DisplayName = string(b)
			if len(userData.DisplayName) > 0 && len(userData.Username) > 0 {
				jsonData.Users = append(jsonData.Users, *userData)
			}
			userData = &User{}
			i = 0
		}

		if len(string(buf)) == 0 {
			break
		}
	}

	jm, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	fmt.Print(renderer.PrettyJson(jm))

	return nil
}
