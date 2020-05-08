package palindrome_number

import "testing"

/*
 * @author taohu
 * @date 2020/5/8
 * @time 下午3:11
 * @desc please describe function
**/

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 121",
			args: args{x: 121},
			want: true,
		},
		{
			name: "test case -121",
			args: args{x: -121},
			want: false,
		},
		{
			name: "test case 10",
			args: args{x: 10},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
