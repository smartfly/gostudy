package container_with_most_water

/*
 * @author taohu
 * @date 2020/5/4
 * @time 下午12:00
 * @desc 盛最多水的容器
**/

func maxArea(height []int) int {
	var startIndex, endIndex, result, temp int
	startIndex = 0
	endIndex = len(height) - 1

	for startIndex < endIndex {

		if height[startIndex] >= height[endIndex] {
			temp = (endIndex - startIndex) * height[endIndex]
			endIndex = endIndex - 1
		} else {
			temp = (endIndex - startIndex) * height[startIndex]
			startIndex = startIndex + 1
		}

		if temp > result {
			result = temp
		}
	}

	return result
}
