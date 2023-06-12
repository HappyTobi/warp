package users

import (
	"bytes"
	"fmt"
	"io"
)

func (u *Users) AllUsernames() ([]*User, error) {
	u.request.Path = "users/all_usernames"

	data, err := u.request.Get()
	if err != nil {
		return nil, err
	}

	return deserialize(data)
}

func deserialize(data []byte) ([]*User, error) {
	users := make([]*User, 0)
	user := &User{}

	reader := bytes.NewReader(data)
	buf := make([]byte, 32)
	var i = 0
	var userId = 0
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
			user.Username = string(b)
			user.Id = userId
			i++
		} else {
			user.DisplayName = string(b)
			if len(user.DisplayName) > 0 && len(user.Username) > 0 {
				users = append(users, user)
			}
			user = &User{}
			i = 0
			userId++
		}

		if len(string(buf)) == 0 {
			break
		}
	}

	return users, nil
}
