package mock_dynamic

import (
	"testing"
	"time"
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
