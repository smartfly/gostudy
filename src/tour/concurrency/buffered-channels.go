package main

import "fmt"

/**
 * @author taohu
 * @date 19-4-19
 * @time 下午4:59
 * @desc 带缓冲的信道
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
