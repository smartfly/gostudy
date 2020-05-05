package roman_to_integer

import "testing"

/*
 * @author taohu
 * @date 2020/5/5
 * @time 下午12:28
 * @desc please describe function
**/

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case III",
			args: args{s: "III"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
