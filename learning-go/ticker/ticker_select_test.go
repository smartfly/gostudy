package ticker

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2020/1/27
 * @time 下午3:42
 * @desc select-case & time.Ticker https://studygolang.com/articles/5224
**/
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 原始版 在使用select-case时候，当两个case条件都满足的时候，运行时系统会通过一个伪随机的算法决定那个Case将会被执行
func TestOrigin(t *testing.T) {
	ch := make(chan int, 1024)
	// 定时取出数据
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val: %d\n", val)
		}
	}(ch)
	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}

// 改进版
func TestImprovement(t *testing.T) {
	ch := make(chan int, 1024)
	// 定时取出数据
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val: %d\n", val)
		}
	}(ch)
	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		}

		select {
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		default:
		}
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}
