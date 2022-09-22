package function

import (
	"fmt"
	"testing"
)

// 回调函数不一样
func TestAnonymousCallback(t *testing.T) {
	sumWorker([]int{1, 2, 3, 4}, func(i int) {
		fmt.Println(i)
	})

	sumWorker([]int{1, 2, 3, 4}, func(i int) {
		if i > 100 {
			fmt.Println("sum > 100")
		} else {
			fmt.Println("sum <= 100")
		}
	})
}

func sumWorker(data []int, callback func(int)) {
	sum := 0
	for _, num := range data {
		sum += num
	}
	callback(sum)
}
