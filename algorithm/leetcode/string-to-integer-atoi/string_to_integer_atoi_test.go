package string_to_integer_atoi

import "testing"

/*
 * @author taohu
 * @date 2020/5/5
 * @time 下午10:36
 * @desc please describe function
**/

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 42",
			args: args{"42"},
			want: 42,
		},
		{
			name: "test case    -42",
			args: args{"   -42"},
			want: -42,
		},
		{
			name: "test case words and 987",
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: "test case -91283472332",
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			name: "test case +-2",
			args: args{str: "+-2"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
