package meter

import "github.com/HappyTobi/warp/pkg/internal/warp"

type Meter struct {
	request warp.Request `json:"-"`
}
