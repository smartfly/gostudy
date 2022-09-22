package g279

import "math"

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		minNum := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minNum = findMin(minNum, f[i-j*j])
		}
		f[i] = minNum + 1
	}
	return f[n]
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
