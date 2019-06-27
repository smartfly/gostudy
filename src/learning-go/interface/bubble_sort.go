package _interface

/*
 * @author: taohu
 * @date: 19-6-21
 * @time: 上午9:58
 * @desc: 冒泡排序
**/

func bubbleSort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}
