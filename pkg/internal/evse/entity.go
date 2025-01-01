package evse

import (
	"github.com/HappyTobi/warp/pkg/internal/warp"
)

type Evse struct {
	request warp.Request `json:"-"`
}

type ChargePower struct {
	Current int `json:"current"`
}

type ExternelCurrent struct {
	Current int `json:"current"`
}
