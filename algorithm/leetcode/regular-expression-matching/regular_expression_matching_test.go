package regular_expression_matching

import "testing"

/*
 * @author taohu
 * @date 2020/5/8
 * @time 下午6:34
 * @desc please describe function
**/

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case aa-a*",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "test case ab-.*",
			args: args{
				s: "aa",
				p: ".*",
			},
			want: true,
		},
		{
			name: "test case aa-a",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
