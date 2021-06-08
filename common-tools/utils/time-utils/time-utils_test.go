package time_utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/now"
)

func TestTime2Str(t *testing.T) {
	result := Time2Str(time.Now())
	fmt.Printf("time to time string: %s\n", result)
}

func TestTimestamp2Str(t *testing.T) {
	result := Timestamp2Str(time.Now().Unix())
	fmt.Printf("timestamp to time string: %s\n", result)
}

func TestStr2Time(t *testing.T) {
	result, _ := Str2Time("2019-05-10 10:26:36")
	fmt.Printf(" time string to time: %s\n", result)
}

func TestStr2Timestamp(t *testing.T) {
	result, _ := Str2Timestamp("2019-05-10 10:26:36")
	fmt.Printf(" time string to timestamp: %d\n", result)
}

func TestTimeZero(t *testing.T) {
	fmt.Printf("time zero value: %s\n", time.Time{})
}

func TestGetFirstDayOfLastWeek(t *testing.T) {
	fmt.Println(GetFirstDayOfLastWeek())
}

func TestGetLastDayOfLastWeek(t *testing.T) {
	fmt.Println(GetLastDayOfLastWeek())
}

func TestGetFirstDayOfLastMonth(t *testing.T) {
	fmt.Println(GetFirstDayOfLastMonth())
}

func TestGetLastDayOfLastMonth(t *testing.T) {
	fmt.Println(GetLastDayOfLastMonth())
}

func TestEndOfWeek(t *testing.T) {
	now.WeekStartDay = time.Monday // set Monday as first day, default is Sunday
	fmt.Println(now.BeginningOfWeek())
	fmt.Println(now.EndOfWeek())
}

func TestCalculatingBasedOnAnotherTime(t *testing.T) {
	t1 := time.Date(2019, 7, 1, 17, 13, 49, 123456789, time.Now().Location())
	fmt.Println(now.New(t1).EndOfMonth())
}

func TestTimeFormats(t *testing.T) {
	now.TimeFormats = append(now.TimeFormats, "02 Jan 2006 15:04")
	fmt.Println(now.TimeFormats[len(now.TimeFormats)-1])
}

func TestTime2Str1(t *testing.T) {
	type args struct {
		time time.Time
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time2Str(tt.args.time); got != tt.want {
				t.Errorf("Time2Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
