package median_of_two_sorted_array

import (
	"math"
)

/*
 * @author taohu
 * @date 2020/4/12
 * @time 下午12:50
 * @desc 两个数组的中位数
**/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		leftValue := findKth(nums1, 0, nums2, 0, length/2)
		rightValue := findKth(nums1, 0, nums2, 0, length/2+1)
		return (float64(leftValue) + float64(rightValue)) / 2
	} else {
		value := float64(findKth(nums1, 0, nums2, 0, (length+1)/2))
		return value
	}
}

func findKth(num1 []int, index1 int, num2 []int, index2 int, k int) int {
	if index1 >= len(num1) {
		return num2[index2+k-1]
	}

	if index2 >= len(num2) {
		return num1[index1+k-1]
	}

	if k == 1 {
		a := num1[index1]
		b := num2[index2]
		var value int
		if a >= b {
			value = b
		} else {
			value = a
		}
		return value
	}

	var result, mid1, mid2 int
	if index1+k/2-1 < len(num1) {
		mid1 = num1[index1+k/2-1]
	} else {
		mid1 = math.MaxInt64
	}

	if index2+k/2-1 < len(num2) {
		mid2 = num2[index2+k/2-1]
	} else {
		mid2 = math.MaxInt64
	}

	if mid1 >= mid2 {
		result = findKth(num1, index1, num2, index2+k/2, k-k/2)
	} else {
		result = findKth(num1, index1+k/2, num2, index2, k-k/2)
	}
	return result
}
