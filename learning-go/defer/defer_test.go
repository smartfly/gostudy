package _defer

import (
	"fmt"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2020/3/20
 * @time 下午8:16
 * @desc [defer](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/)
**/

// 多个defer,其执行顺序如同栈，后进先出
func TestDeferZone(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// defer 在当前函数和方法返回之前被调用
func TestDeferRunOrder(t *testing.T) {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")
}

// 预计算：调用defer关键字会立刻对函数中引用的外部参数进行拷贝，所以time.Since(startedAt)的结果不是在main函数退出之前计算的
func TestDeferPreCalculated(t *testing.T) {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt).Seconds())

	time.Sleep(time.Second)
}

// 预计算改进：虽然调用defer关键字时也使用值传递，但是因为拷贝的是函数指针, 所以time.Since(startedAt)会在main函数执行前被调用
// 并打印出符合预期的结果
func TestDeferPreCalculatedAdvanced(t *testing.T) {
	startedAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startedAt).Seconds())
	}()

	time.Sleep(time.Second)
}
