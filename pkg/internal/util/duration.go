package util

import (
	"fmt"
	"math"
)

/*
Return duration in format HH:MM:SS
*/
func ChargeDuration(duration uint32) string {
	tmp := math.Floor(float64(duration) / 3600)
	hours := durationFormat(tmp)
	tmp = math.Floor((float64(duration) - tmp*3600) / 60)
	minutes := durationFormat(tmp)
	tmpSec := (duration % 60)
	seconds := durationFormat(float64(tmpSec))

	return fmt.Sprintf("%s:%s:%s", hours, minutes, seconds)
}

func durationFormat(time float64) string {
	if time < 9 {
		return fmt.Sprintf("0%d", int(time))
	}
	return fmt.Sprintf("%d", int(time))
}
