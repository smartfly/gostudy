package fontscolorutils

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2020/2/19
 * @time 下午6:04
 * @desc please describe function
**/

func TestBlack(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "black",
			args: args{msg: "test"},
			want: SetColor("test", 0, 0, TextYellow),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yellow(tt.args.msg); got != tt.want {
				t.Errorf("Black() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func TestSetColor(t *testing.T) {
	a := SetColor("test=2 a = 1", DisplayHighLight, BackgroundBlack, TextRed)
	fmt.Println(a)
}
