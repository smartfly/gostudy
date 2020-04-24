package longest_palindromic_substring

/*
 * @author taohu
 * @date 2020/4/12
 * @time 下午8:26
 * @desc please describe function
**/

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(i, i, s)
		len2 := expandAroundCenter(i, i+1, s)
		length := len1
		if length < len2 {
			length = len2
		}
		if length > (end - start) {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(left, right int, s string) int {
	L := left
	R := right
	for (L >= 0 && R < len(s)) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
