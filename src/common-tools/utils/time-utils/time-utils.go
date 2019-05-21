/*
time transfer tool, include time to string, string to time, timestamp to string and string to timestamp
*/
package time_utils

import "time"

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
