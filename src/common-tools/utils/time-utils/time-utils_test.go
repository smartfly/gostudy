package time_utils

import (
	"fmt"
	"testing"
	"time"
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
