package main

import "fmt"

func main() {
	var totalLine int
	_, _ = fmt.Scanln(&totalLine)
	nums := make([]string, totalLine)
	for i := 0; i < totalLine; i++ {
		var cnt int
		_, _ = fmt.Scanln(&cnt)
		negativeAndZeroNum := 0
		var temp int
		for j := 0; j < cnt; j++ {
			_, _ = fmt.Scan(&temp)
			if temp <= 0 {
				negativeAndZeroNum++
			}
		}
		if negativeAndZeroNum >= cnt/2 {
			nums[i] = "Yes"
		} else {
			nums[i] = "No"
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
