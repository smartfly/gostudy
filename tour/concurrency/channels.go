package main

import (
	"fmt"
)

/**
 * @author taohu
 * @date 19-4-19
 * @time 下午4:32
 * @desc 信道
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c中接收

	fmt.Println(x, y, x+y)

}
