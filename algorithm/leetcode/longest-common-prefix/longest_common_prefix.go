package longest_common_prefix

import "math"

/*
 * @author taohu
 * @date 2020/5/9
 * @time 下午4:29
 * @desc 最长公共前缀
**/

func longestCommonPrefix(strs []string) string {
	minLen := math.MaxInt64
	var minStr string
	for _, val := range strs {
		temp := len(val)
		if temp < minLen {
			minLen = temp
			minStr = val
		}
	}

	var result string
	for i, ch := range minStr {
		flag := true
		for _, str := range strs {
			if str == minStr {
				continue
			}
			if str[i:i+1] != string(ch) {
				flag = false
				break
			}
		}
		if flag {
			result += string(ch)
		} else {
			break
		}
	}
	return result
}

func longestCommonPrefixAdvance(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(result) && (j < len(strs[i])) {
			if result[j:j+1] != strs[i][j:j+1] {
				break
			}
			j++
		}
		result = result[0:j]
		if result == "" {
			break
		}
	}

	return result
}
