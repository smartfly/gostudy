package reverse_integer

import "testing"

/*
 * @author taohu
 * @date 2020/5/8
 * @time 下午1:20
 * @desc please describe function
**/

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 123",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "test case -123",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "test case 120",
			args: args{x: 120},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
