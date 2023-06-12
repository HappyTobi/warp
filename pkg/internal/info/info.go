package info

import "github.com/HappyTobi/warp/pkg/internal/warp"

func NewInfoService(request warp.Request) *Info {
	return &Info{
		request: request,
	}
}

func (i *Info) DisplayName() (map[string]interface{}, error) {
	i.request.Path = "info/display_name"
	return i.request.GetJson()
}

func (i *Info) Features() ([]byte, error) {
	i.request.Path = "info/features"
	return i.request.Get()
}

func (i *Info) Modules() (map[string]interface{}, error) {
	i.request.Path = "info/modules"
	return i.request.GetJson()
}

func (i *Info) Name() (map[string]interface{}, error) {
	i.request.Path = "info/name"
	return i.request.GetJson()
}

func (i *Info) Version() (map[string]interface{}, error) {
	i.request.Path = "info/version"
	return i.request.GetJson()
}
