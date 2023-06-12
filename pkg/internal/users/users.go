package users

import "github.com/HappyTobi/warp/pkg/internal/warp"

func NewUsersService(request warp.Request) *Users {
	return &Users{
		request: request,
	}
}
