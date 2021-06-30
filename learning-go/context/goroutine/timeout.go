package main

import (
	"context"
	"fmt"
	"time"
)

/**
同一个context可以控制多个goroutine, 确保线程可控, 而不是每新建一个goroutine就要有一个chan去通知他关闭
有了他代码更加简洁
*/

func main() {
	fmt.Println("run demo \n\n\n")
	demo()
}

func demo() {
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	go watch(ctx, "[线程1]", 5)
	go watch(ctx, "[线程2]", 2)
	//go watch(ctx, "[线程3]", 3)

	index := 0
	for {
		index++
		//fmt.Printf("%d 秒过去了 \n", index)
		time.Sleep(5 * time.Second)
		if index > 10 {
			break
		}
	}

	fmt.Println("通知停止监控")
	// 其实此时已经超时, 协程已经提前退出
	cancel()

	// 防止主进程提前退出
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}

func watch(ctx context.Context, name string, num int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v %s  监控退出, 停止了...\n", time.Now(), name)
			return
		default:
			//fmt.Printf("%s goroutine监控中... \n", name)
			fmt.Println(time.Now(), name, ctx)
			fmt.Printf("%p\n", &ctx)
			a := time.Duration(num)
			time.Sleep(a * time.Second)
		}
	}
}
