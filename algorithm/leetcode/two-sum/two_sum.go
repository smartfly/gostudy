package two_sum

/*
 * @author taohu
 * @date 2020/3/29
 * @time 下午6:33
 * @desc please describe function
**/

// 暴力破解，时间复杂度：O(n^2)
func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}
	return result
}

// hashMap实现-两遍hsh, 时间复杂度：O(n) 空间复杂度：O(n)
func twoSumAdvanced2Iter(nums []int, target int) []int {
	sourceMap := make(map[int]int, len(nums))
	for i, value := range nums {
		sourceMap[value] = i
	}

	for i, value := range nums {
		dst := target - value
		if index, ok := sourceMap[dst]; ok {
			if index != i {
				return []int{i, index}
			}
		}
	}
	return []int{}
}

// hashMap实现-一遍hash，时间复杂度：O(n) 空间复杂度：O(n)
func twoSumAdvanced(nums []int, target int) []int {
	sourceMap := make(map[int]int, 0)
	for i, value := range nums {
		if index, ok := sourceMap[target-value]; ok {
			return []int{index, i}
		}
		sourceMap[value] = i
	}
	return []int{}
}
