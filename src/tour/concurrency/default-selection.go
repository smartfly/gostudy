package main

import (
	"fmt"
	"time"
)

/**
 * @author taohu
 * @date 19-4-20
 * @time 下午6:33
 * @desc 默认选择
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
