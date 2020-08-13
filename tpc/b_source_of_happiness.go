package main

/*
 * @author taohu
 * @date 2020/8/4
 * @time 下午7:16
 * @desc please describe function
**/

//func main() {
//	var n int
//	_, _ = fmt.Scanln(&n)
//	for i := 0; i < n; i++ {
//		var total, sum int
//		_, _ = fmt.Scanln(&total, &sum)
//		nums := make([]int, total)
//		var temp int
//		for i := 0; i < total; i++ {
//			_, _ = fmt.Scan(&temp)
//			nums[i] = temp
//		}
//
//		if total <= 2 {
//			fmt.Println("Impossible")
//			continue
//		} else {
//			fixNums := make([]int, total)
//			for i := 2; i < total; i++ {
//				tempSum := nums[i-2] + nums[i-1] + nums[i]
//				if tempSum < sum {
//					if nums[i-2] == -1 && nums[i-1] == -1 && nums[i] == -1 { // -1 -1 -1
//						fmt.Println("Impossible")
//						break
//					} else if nums[i-2] == -1 && nums[i-1] != -1 && nums[i] != -1 { // -1 0 -1
//						temp = sum - nums[i-1] - nums[i]
//						if temp > fixNums[i-2] {
//							fixNums[i-2] = temp
//						}
//					} else if nums[i-2] != -1 && nums[i-1] == -1 && nums[i] != -1 { // 0 -1 0
//						temp = sum - nums[i-2] - nums[i]
//						if temp > fixNums[i-1] {
//							fixNums[i-1] = temp
//						}
//					} else if nums[i-2] != -1 && nums[i-1] != -1 && nums[i] == -1 { // 0 0 -1
//						temp = sum - nums[i-2] - nums[i-1]
//						if temp > fixNums[i] {
//							fixNums[i] = temp
//						}
//					} else if nums[i-2] == -1 && nums[i-1] == -1 && nums[i] != -1 { // -1 -1 0
//						temp = sum - fixNums[i-2] - nums[i]
//						if temp > fixNums[i-1] {
//							fixNums[i-1] = temp
//						}
//					} else if nums[i-2] == -1 && nums[i-1] != -1 && nums[i] == -1 { // -1 0 -1
//						temp = sum - fixNums[i-2] - nums[i-1]
//						if temp > fixNums[i] {
//							fixNums[i] = temp
//						}
//					} else if nums[i-2] != -1 && nums[i-1] == -1 && nums[i] == -1 { // 0 0 -1
//						temp = sum - fixNums[i-2] - nums[i-1]
//						if temp > fixNums[i] {
//							fixNums[i] = temp
//						}
//					} else if nums[i-2] != -1 && nums[i-1] != -1 && nums[i] != -1 { // 0 0 0
//						if tempSum == 0 {
//							fmt.Println("Impossible")
//							break
//						}
//					}
//				}
//			}
//			var result int
//			for i := 0; i < total; i++ {
//				if nums[i] == -1 {
//					result += fixNums[i]
//				} else {
//					result += nums[i]
//				}
//			}
//			fmt.Println(result)
//		}
//	}
//}
