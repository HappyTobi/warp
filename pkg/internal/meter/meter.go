package meter

import "github.com/HappyTobi/warp/pkg/internal/warp"

func NewMeterService(request warp.Request) *Meter {
	return &Meter{
		request: request,
	}
}

func (i *Meter) Values() (map[string]interface{}, error) {
	i.request.Path = "meter/values"
	return i.request.GetJson()
}
