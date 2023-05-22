package utils

import (
	"time"
)

// 当前时间
func GetTime() time.Time {
	return time.Now()
}

// 当前时间戳
func GetUnixTime() int64 {
	return time.Now().Unix()
}
