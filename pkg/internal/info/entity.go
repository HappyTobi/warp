package info

import "github.com/HappyTobi/warp/pkg/internal/warp"

type InfoUpdate struct {
	request         *warp.Request
	warpVersion     string
	firmwareVersion string
}

type Update struct {
	Available      bool   `json:"update_available" yaml:"update_available"`
	UpdateVersion  string `json:"update_version" yaml:"update_version"`
	CurrentVersion string `json:"current_version" yaml:"current_version"`
}

type Tags struct {
	Name   string `json:"name"`
	ZipUrl string `json:"zipball_url"`
	TarUrl string `json:"tarball_url"`
}
