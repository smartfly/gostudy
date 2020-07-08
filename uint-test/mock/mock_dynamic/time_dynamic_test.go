package mock_dynamic

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"unsafe"
)

/*
 * @author taohu
 * @date 2020/3/30
 * @time 下午3:32
 * @desc [编写可测试 Go 代码的一种模式](https://blog.betacat.io/post/2020/03/a-pattern-for-writing-testable-go-code/)
**/

func TestGetUnix(t *testing.T) {

	nowFn = func() time.Time {
		return time.Now().Add(-1 * time.Hour)
	}

	type args struct {
		createAt time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{createAt: time.Now().Add(-23 * time.Hour)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTime24Inner(tt.args.createAt); got != tt.want {
				t.Errorf("GetUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTimeUnix(now func() time.Time) int64 {
	return now().Unix()
}

func Test(t *testing.T) {
	fmt.Println(getTimeUnix(time.Now))
	var a uint64 = 4
	var b uint64 = 65535
	var c uint64 = 3
	fmt.Printf("%d\n", c)
	var d uint64 = 788
	f := (a << uint64(48)) | (b << uint64(32)) | (c << uint64(16)) | d
	fmt.Printf("%b\n", f)
	fmt.Println(f)
}

func Test1(t *testing.T) {
	fmt.Println(time.Now().Unix() + int64(time.Hour*24*90))
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)
	remainSecond := time.Duration(todayLastTime.Unix()-time.Now().Local().Unix()) * time.Second
	fmt.Println(remainSecond)
}

func Test2(t *testing.T) {
	var x int = 2309000000
	//i := int(x)     // i is -1 on 32-bit systems, 0xffffffff on 64-bit
	fmt.Printf("%d\n", x)
	fmt.Println(unsafe.Sizeof(x))
}

func test() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func Test3(t *testing.T) {
	fmt.Println(test())
	a := make(map[string]string)
	_ = json.Unmarshal([]byte("eyJhcHBfY2FsbF90eXBlIjoiMSIsInNjZW5lIjoidXNlcl9sb2dpbiJ9"), &a)
	fmt.Println(a)
}
