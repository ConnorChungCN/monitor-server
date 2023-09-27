package utils

import "time"

func Timestamp2Time(st int64) time.Time {
	return time.Unix(st, 0)
}

func NowTime() time.Time {
	return time.Now()
}

func NowTimestamp() int64 {
	return NowTime().Unix()
}
