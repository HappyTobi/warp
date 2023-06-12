package users

import "github.com/HappyTobi/warp/pkg/internal/warp"

type Users struct {
	request warp.Request `json:"-"`
	Users   []*User      `json:"Users"`
}

type User struct {
	Username    string `json:"Username"`
	DisplayName string `json:"Displayname"`
	Id          int    `json:"UserId"`
}
