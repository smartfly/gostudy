package roman_to_integer

import "testing"

/*
 * @author taohu
 * @date 2020/5/5
 * @time 下午5:33
 * @desc please describe function
**/

func Test_romanToIntAdvance(t *testing.T) {
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
		{
			name: "test case IV",
			args: args{s: "IV"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToIntAdvance(tt.args.s); got != tt.want {
				t.Errorf("romanToIntAdvance() = %v, want %v", got, tt.want)
			}
		})
	}
}
