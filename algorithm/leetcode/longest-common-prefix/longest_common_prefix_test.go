package longest_common_prefix

import "testing"

/*
 * @author taohu
 * @date 2020/5/9
 * @time 下午4:40
 * @desc please describe function
**/

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "test case 2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefixAdvance(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "test case 2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "test case 3",
			args: args{strs: []string{"aaa", "aa", "aaa"}},
			want: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefixAdvance(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefixAdvance() = %v, want %v", got, tt.want)
			}
		})
	}
}
