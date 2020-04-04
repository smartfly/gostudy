package length_of_longest_substring

/*
 * @author taohu
 * @date 2020/4/4
 * @time 上午10:25
 * @desc please describe function
**/

func lengthOfLongestSubstring(s string) int {
	longestSubstringMap := make(map[int32]int, 0)
	maxLength := 0
	startIndex, endIndex := 0, 0
	for i, val := range s {
		if index, ok := longestSubstringMap[val]; ok {
			if startIndex < index+1 {
				startIndex = index + 1
			}
		}
		longestSubstringMap[val] = i
		value := endIndex - startIndex + 1
		if value > maxLength {
			maxLength = value
		}
		endIndex = i + 1
	}
	return maxLength
}
