package _select

import (
	"fmt"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2020/1/27
 * @time 下午8:04
 * @desc select https://colobu.com/2017/07/07/select-vs-switch-in-golang/
**/

// TestSelectByCase 正常执行
func TestSelectByCase(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// TestSelectByDefault 不存在收发的channel, 执行default中的语句
func TestSelectByDefault(t *testing.T) {

	var c1 chan int
	select {
	case i := <-c1:
		fmt.Println(i)
	default:
		fmt.Println("terminal")
	}

}

// TestSelectByRandom 随机执行
func TestSelectByRandom(t *testing.T) {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}
