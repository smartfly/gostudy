package main

/*
 * @author taohu
 * @date 2020/8/11
 * @time 下午7:13
 * @desc please describe function
**/

//func main() {
//	var cnt int
//	_, _ = fmt.Scanln(&cnt)
//	outputs := make([][]int, cnt)
//	for i := 0; i < cnt; i++ {
//		var temp string
//		outputs[i] = make([]int, 2)
//		_, _ = fmt.Scan(&temp)
//		length := len(temp)
//		records := make(map[int]string, 0)
//		zeroNum := 0
//		oneNum := 0
//		for i := 0; i < length; i++ {
//			records[i] = string(temp[i])
//			if temp[i] == '0' {
//				zeroNum++
//			} else if temp[i] == '1' {
//				oneNum++
//			}
//		}
//		sub := zeroNum - oneNum
//		if length%2 == 1 {
//			if sub < 0 {
//				startIndex, endIndex := flip(sub, temp)
//				outputs[i] = []int{startIndex, endIndex}
//			} else if sub > 0 {
//				startIndex, endIndex := flip(sub, temp)
//				outputs[i] = []int{startIndex, endIndex}
//			}
//		} else {
//			if sub == 0 {
//				outputs[i] = []int{1, length}
//			} else if sub > 0 {
//				startIndex, endIndex := flip(sub, temp)
//				outputs[i] = []int{startIndex, endIndex}
//			} else if sub < 0 {
//				startIndex, endIndex := flip(sub, temp)
//				outputs[i] = []int{startIndex, endIndex}
//			}
//		}
//
//	}
//	for i := 0; i < cnt; i++ {
//		fmt.Println(outputs[i][0], outputs[i][1])
//	}
//}
//
//func flip(num int, temp string) (int, int) {
//	var start, end int
//	if int(math.Abs(float64(num))) >= len(temp) {
//		return 1, 1
//	}
//
//	if num > 0 {
//		startIndex := 0
//		maxZeroLen := 0
//		for i := 0; i < len(temp); i++ {
//			if temp[i] == '0' {
//				continue
//			} else {
//				temp := i - startIndex
//				if temp > maxZeroLen {
//					maxZeroLen = temp
//					start = startIndex
//					end = i
//				}
//				if maxZeroLen == -num {
//					return start + 1, end
//				}
//				startIndex = i
//			}
//		}
//		return start + 1, end
//	} else {
//		startIndex := 0
//		maxOneLen := 0
//		for i := 0; i < len(temp); i++ {
//			if temp[i] == '1' {
//				continue
//			} else {
//				temp := i - startIndex
//				if temp > maxOneLen {
//					maxOneLen = temp
//					start = startIndex
//					end = i
//				}
//
//				if maxOneLen == -num {
//					return start + 1, end
//				}
//
//				startIndex = i
//			}
//		}
//
//		return start + 1, end
//	}
//}
