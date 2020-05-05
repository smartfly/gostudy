package integer_to_roman

import "testing"

/*
 * @author taohu
 * @date 2020/5/5
 * @time 上午11:53
 * @desc please describe function
**/

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case: 3",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "test case: 1994",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
