package g16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var result int
	n := len(nums)
	gap := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		// a重复,继续下一个
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target { // 相等是最小的
				return target
			}
			if gapVal := abs(sum - target); gapVal < gap {
				gap = gapVal
				result = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}

	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
