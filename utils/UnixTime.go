package utils

import (
	"strconv"
	"time"
)

func ParseUnixTimeStr(unixTimeStr string) (time.Time, error) {
	timeInt, err := strconv.ParseInt(unixTimeStr, 10, 64)
	timeUnix := time.Unix(timeInt, 0)
	return timeUnix, err
}

func ParseUnitTimeInt(unixTimeInt int64) time.Time {
	return time.Unix(unixTimeInt, 0)
}
