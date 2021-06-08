package main

/*
 * @author taohu
 * @date 2020/8/11
 * @time 下午8:31
 * @desc please describe function
**/

//func main() {
//	var cnt int
//	_, _ = fmt.Scanln(&cnt)
//	for i := 0; i < cnt; i++ {
//		var n, m int
//		_, _ = fmt.Scanln(&n, &m)
//		records := make([][]int, n)
//		for i := 0; i < n; i++ {
//			nums := make([]int, n)
//			var temp int
//			for i := 0; i < n; i++ {
//				_, _ = fmt.Scan(&temp)
//				nums[i] = temp
//			}
//			records[i] = nums
//		}
//
//		maxValue := 0
//		row := 0
//		for i := 0; i < n; i++ {
//			if records[i][3] > maxValue {
//				row = i
//				maxValue = records[i][3]
//			}
//		}
//		min := 20
//		rowMin := 0
//		for i := 0; i < n; i++ {
//			if records[i][1] < min {
//				rowMin = i
//				min = records[i][1]
//			}
//		}
//		v := 0
//		c := 0
//		for i := 0; i < n; i++ {
//			if rowMin == i {
//				continue
//			}
//			if i == row {
//				fmt.Println(row)
//				c += records[i][0] + records[i][2]
//				v += records[i][1] + records[i][3]
//			}
//			c += records[i][0]
//			v += records[i][1]
//		}
//		fmt.Println(v, c)
//	}
//}
