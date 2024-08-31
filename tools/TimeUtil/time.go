package TimeUtil

import (
	"fmt"
	"time"
)

const (
	CHINESE_LAYOUT = "2006-01-02 15:04:05"
)

var (
	loc, _ = time.LoadLocation("Local")
)

// 时间戳转为日期格式
func UnixToDatetime(t string) string {
	// 使用`time.Unix`将时间戳转换为时间类型
	timestamp, _ := time.Parse(CHINESE_LAYOUT, time.Unix(0, 0).In(loc).Format(CHINESE_LAYOUT))
	fmt.Sscan(t, &timestamp)

	// 将时间类型转换为日期格式
	date := timestamp.In(loc)
	return date.Format(CHINESE_LAYOUT)
}

// 获取某日期的时间戳
func DatetimeToUnix(date string) int64 {
	t, _ := time.ParseInLocation(CHINESE_LAYOUT, date, loc)
	return t.Unix()
}


// 获取当前时间
func GetCurrentDatetime() string {
	return time.Now().In(loc).Format(CHINESE_LAYOUT)
}


