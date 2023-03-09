package info

import (
	"github.com/HappyTobi/warp/pkg/internal/warp"
)

type Info struct {
	request *warp.Request
}

type Version struct {
	Firmware string `json:"firmware"`
	Config   string `json:"config"`
}

type Name struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	DisplayType string `json:"display_type"`
	Uid         string `json:"uid"`
}
