package info

import "github.com/HappyTobi/warp/pkg/internal/util"

func (i *Info) LoadVersion() (*Version, error) {
	data, err := i.request.Get()
	if err != nil {
		return nil, err
	}

	version := &Version{}
	if err := util.DeserializeJsonData(data, version); err != nil {
		return nil, err
	}

	return version, nil
}
