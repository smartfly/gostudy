package fake

import "fmt"

/*
 * @author taohu
 * @date 2019/11/23
 * @time 下午3:39
 * @desc please describe function
**/

func TestFunc(num1 int, num2 int) int {
	res := compareInt(num1, num2)
	fmt.Println("CompareInt return", res)
	return res
}

func compareInt(num1 int, num2 int) int {
	fmt.Println(num2)
	if num1 == num2 {
		return 0
	} else if num1 > num2 {
		return 1
	} else {
		return -1
	}
}
