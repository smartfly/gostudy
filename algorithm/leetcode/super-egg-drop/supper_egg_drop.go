package super_egg_drop

var memoMap map[int]int

func superEggDrop(k int, n int) int {
	memoMap = make(map[int]int)
	return dp(k, n)
}

func dp(k, n int) int {
	// 备忘录
	if val, ok := memoMap[n*1000+k]; ok {
		return val
	}
	res := 0
	if k == 1 { // 1个鸡蛋
		return n
	} else if n == 0 { // 楼层为0
		return 0
	} else {
		low, high := 1, n
		for low+1 < high {
			x := (low + high) / 2
			t1 := dp(k-1, x-1)
			t2 := dp(k, n-x)
			if t1 < t2 {
				low = x
			} else if t1 > t2 {
				high = x
			} else {
				low, high = x, x
			}
		}
		res = 1 + findMin(findMax(dp(k-1, low-1), dp(k, n-low)), findMax(dp(k-1, high-1), dp(k, n-high)))
	}
	memoMap[n*1000+k] = res
	return res
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
