/*
time transfer tool, include time to string, string to time, timestamp to string and string to timestamp
*/
package time_utils

import (
	"github.com/jinzhu/now"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

// Time 时间转化为string
func Time2Str(time time.Time) string {
	timeString := time.Format(timeLayout)
	return timeString
}

// 时间戳timestamp转string
func Timestamp2Str(timestamp int64) string {
	timeValue := time.Unix(timestamp, 0)
	timeString := Time2Str(timeValue)
	return timeString
}

// string转时间time
func Str2Time(timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	// 使用模板在对应时区转化为time.time类型
	theTime, err := time.ParseInLocation(timeLayout, timeStr, loc)
	if err != nil {
		return time.Time{}, err
	}
	return theTime, nil
}

// string转时间戳timestamp
func Str2Timestamp(timeStr string) (int64, error) {
	theTime, err := Str2Time(timeStr)
	if err != nil {
		return 0, err
	}
	timestamp := theTime.Unix()
	return timestamp, nil
}

// 获取上周周一起始时间
func GetFirstDayOfLastWeek() time.Time {
	time := now.Monday()
	time = time.AddDate(0, 0, -7)
	return time
}

// 获取上周周日结束时间
func GetLastDayOfLastWeek() time.Time {
	time := now.EndOfSunday()
	time = time.AddDate(0, 0, -7)
	return time
}

// 获取上月第一天起始时间
func GetFirstDayOfLastMonth() time.Time {
	time := now.BeginningOfMonth()
	time = time.AddDate(0, -1, 0)
	return time
}

// 获取上月最后一天结束时间
func GetLastDayOfLastMonth() time.Time {
	time := now.BeginningOfMonth()
	time = time.AddDate(0, -1, 0)
	time = now.New(time).EndOfMonth()

	return time
}
