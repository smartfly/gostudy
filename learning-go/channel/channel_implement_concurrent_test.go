package channel

import "testing"

/*
 * @author taohu
 * @date 2020/4/3
 * @time 下午8:25
 * @desc please describe function
**/

func Test_currentComputer(t *testing.T) {
	array := make([]int, 0)
	for i := 0; i < 1000000; i++ {
		array = append(array, i)
	}

	type args struct {
		numArray []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{numArray: array},
			want: 499999500000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := currentComputer(tt.args.numArray); got != tt.want {
				t.Errorf("currentComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}
