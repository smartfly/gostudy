package two_sum

import (
	"reflect"
	"testing"
)

/*
 * @author taohu
 * @date 2020/3/29
 * @time 下午6:42
 * @desc please describe function
**/

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{nums: []int{2, 7, 11, 15}, target: 9},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumAdvanced(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success-case1",
			args: args{nums: []int{3, 2, 4}, target: 6},
			want: []int{1, 2},
		},
		{
			name: "success-case2",
			args: args{nums: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumAdvanced(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumAdvanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumAdvanced2Iter(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success-case1",
			args: args{nums: []int{3, 2, 4}, target: 6},
			want: []int{1, 2},
		},
		{
			name: "success-case2",
			args: args{nums: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumAdvanced2Iter(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumAdvanced2Iter() = %v, want %v", got, tt.want)
			}
		})
	}
}
