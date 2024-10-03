package helper

import (
	"strconv"
	"time"
)

// MsToTime преобразует время в мс с начало эпохи в структуру time.Time
func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	seconds := msInt / 1000
	milliseconds := msInt % 1000

	return time.Unix(seconds, milliseconds*int64(time.Millisecond)), nil
}
