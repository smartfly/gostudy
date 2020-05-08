package palindrome_number

/*
 * @author taohu
 * @date 2020/5/4
 * @time 上午11:52
 * @desc 回文数
**/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	var res, temp int
	temp = x
	for temp > 0 {
		i := temp % 10
		res = res*10 + i // 可能存在int32溢出 overflow
		temp = temp / 10
	}

	if res != x {
		return false
	}

	return true
}
