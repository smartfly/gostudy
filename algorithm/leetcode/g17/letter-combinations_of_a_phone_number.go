package g17

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	pairs := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var backtrack func(s string, index int, temp string)
	backtrack = func(s string, index int, temp string) {
		if index == len(s) {
			result = append(result, temp)
		} else {
			letters := pairs[digits[index]]
			lettersCount := len(letters)
			for j := 0; j < lettersCount; j++ {
				backtrack(s, index+1, temp+string(letters[j])) // 使用临时变量，不要改变temp
			}
		}
	}
	backtrack(digits, 0, "")
	return result
}
