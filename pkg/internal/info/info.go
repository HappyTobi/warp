package info

import "github.com/HappyTobi/warp/pkg/internal/warp"

func NewInfo(request *warp.Request) *Info {
	return &Info{request: request}
}
