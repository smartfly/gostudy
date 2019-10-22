package main

import "fmt"

/**
 * @author taohu
 * @date 19-4-20
 * @time 下午6:13
 * @desc range和close
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 发送者可以通过close关闭一个信道来表示没有需要发送值了
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
