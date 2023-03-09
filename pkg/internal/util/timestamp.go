package util

import (
	"time"
)

/*
Returns a date based on the timestamp min
*/
func TimestampMinutesToDate(timestampMin uint32) time.Time {
	if timestampMin == 0 {
		return time.Now()
	}

	datetime := (timestampMin * 60)
	return time.Unix(int64(datetime), 0)
}

// 27962627
// 	time := uint32(27962627)
// 	dt := timestampMinutesToDate(time)
// 	fmt.Print(dt)
