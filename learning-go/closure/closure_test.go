package closure

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2020/3/19
 * @time 下午8:15
 * @desc go 闭包demo
**/

// getSequence 不带参数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func Test_getSequence(t *testing.T) {
	// nextNumber为一个函数, 函数i为0
	nextNumber := getSequence()
	// 调用nextNumber函数, i变量自增1并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 创建新的函数nextNumber1并查看结果
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

// add 带参数闭包使用
func add(a, b int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, a + b
	}
}

func Test_add(t *testing.T) {
	addFunc := add(1, 2)
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	fmt.Println(addFunc())
}
