package nfc

import "github.com/HappyTobi/warp/pkg/internal/warp"

type Nfc struct {
	request warp.Request `json:"-"`
}

type Tag struct {
	Type     int    `json:"tag_type"`
	Id       string `json:"tag_id"`
	LastSeen int    `json:"last_seen"`
}

type AuthorizedTags struct {
	AuthorizedTags []UserTag `json:"authorized_tags"`
}

type UserTag struct {
	UserId int    `json:"user_id"`
	Type   int    `json:"tag_type"`
	Id     string `json:"tag_id"`
}

type userTagCharge struct {
	TagType int    `json:"tag_type"`
	Id      string `json:"tag_id"`
}
