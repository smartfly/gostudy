package baisc

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2020/3/19
 * @time 下午7:57
 * @desc please describe function
**/

// fibonacciByRecursion 递归实现
func fibonacciByRecursion(num int) int {
	if num < 2 {
		return 1
	}
	return fibonacciByRecursion(num-1) + fibonacciByRecursion(num-2)
}

func TestFibonacciByRecursion(t *testing.T) {
	for i := 0; i < 10; i++ {
		nums := fibonacciByRecursion(i)
		fmt.Println(nums)
	}
}

// fibonacciByClosure 闭包实现
func fibonacciByClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestFibonacciByClosure(t *testing.T) {
	f := fibonacciByClosure()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacciByChannel 通道channel实现
func fibonacciByChannel(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		c <- a
	}
	close(c)
}

func TestFibonacciByChannel(t *testing.T) {
	c := make(chan int, 10)
	fibonacciByChannel(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
