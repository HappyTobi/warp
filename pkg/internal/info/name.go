package info

import (
	"github.com/HappyTobi/warp/pkg/internal/util"
)

func (i *Info) LoadName() (*Name, error) {
	data, err := i.request.Get()
	if err != nil {
		return nil, err
	}

	name := &Name{}
	if err := util.DeserializeJsonData(data, name); err != nil {
		return nil, err
	}

	return name, nil
}
