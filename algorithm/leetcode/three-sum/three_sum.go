package three_sum

import "sort"

/*
 * @author taohu
 * @date 2020/5/9
 * @time 下午8:30
 * @desc 三数之和
**/

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	// sort nums
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L := i + 1
		R := length - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[L], nums[R]})

				for L < R && nums[L] == nums[L+1] {
					L++
				}

				for L < R && nums[R] == nums[R-1] {
					R--
				}

				L++
				R--
			} else if sum < 0 {
				L++
			} else if sum > 0 {
				R--
			}
		}
	}

	return result
}
