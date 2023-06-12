package evse

import "github.com/HappyTobi/warp/pkg/internal/warp"

func NewEvseService(request warp.Request) *Evse {
	return &Evse{
		request: request,
	}
}
