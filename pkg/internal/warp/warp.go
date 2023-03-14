package warp

import "strings"

func WarpType(warpType string) string {
	switch warpType {
	case "warp":
		return "warp"
	case "warp2":
		return "warp2"
	}
	return "unknown"
}

func WarpFirmware(firmware string) string {
	return strings.Split(firmware, "-")[0]
}
