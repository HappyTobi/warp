package update

import (
	"encoding/json"
	"strings"

	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func NewUpdate(request warp.Request, warpVersion string, firmwareVersion string) *InfoUpdate {
	return &InfoUpdate{
		request:         request,
		warpVersion:     warpVersion,
		firmwareVersion: firmwareVersion,
	}
}

func (i *InfoUpdate) UpdateAvailable() (*Update, error) {
	data, err := i.request.Get()
	if err != nil {
		return nil, err
	}

	tags, err := deserialize(data)
	if err != nil {
		return nil, err
	}

	for j := range tags {
		warpVersion := strings.Split(tags[j].Name, "-")
		if warpVersion[0] == i.warpVersion {
			if strings.Compare(warpVersion[1], i.firmwareVersion) > 0 {
				return &Update{
					Available:      true,
					UpdateVersion:  warpVersion[1],
					CurrentVersion: i.firmwareVersion,
				}, nil
			}
		}
	}

	return &Update{
		Available:      false,
		UpdateVersion:  i.firmwareVersion,
		CurrentVersion: i.firmwareVersion,
	}, nil
}

func deserialize(data []byte) ([]Tags, error) {
	var githubTags []Tags
	if err := json.Unmarshal(data, &githubTags); err != nil {
		return nil, err
	}

	return githubTags, nil
}
