package main

import (
	"fmt"
	"time"
)

/**
 * @author taohu
 * @date 19-4-28
 * @time 上午10:33
 * @desc 错误
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
