package utils

import (
	"time"
)

const (
	TIMETEMPLATE1 = "2006-01-02 15:04:05" //常规类型
)

// 当前时间
func GetTime() time.Time {
	return time.Now()
}

// 当前时间戳
func GetUnixTime() int64 {
	return time.Now().Unix()
}

// 格式化为:2006-01-02 15:04:05
func GetNormalTimeString(t time.Time) string {
	return t.Format(TIMETEMPLATE1)
}

// 转为时间戳->秒数
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

// 转为时间戳->毫秒数
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
