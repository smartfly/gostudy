package main

import (
	"fmt"
	"time"
)

/**
 * @author taohu
 * @date 19-4-19
 * @time 下午4:25
 * @desc Go 程
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
