package closure

import (
	"fmt"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2020/7/8
 * @time 上午11:42
 * @desc please describe function
**/

// for range中使用闭包
func TestNotClosure(t *testing.T) {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	d := time.Duration(2 * time.Second)
	t1 := time.NewTicker(d)
	defer t1.Stop()
	select {
	case <-t1.C:
		fmt.Println("timeout...")
	}
}

func TestClosure(t *testing.T) {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(value string) {
			fmt.Println(value)
		}(v)
	}
	d := time.Duration(1 * time.Second)
	t1 := time.NewTicker(d)
	defer t1.Stop()
	select {
	case <-t1.C:
		fmt.Println("timeout...")
	}
}

// 函数列表使用不当
func fly() []func() {
	var fs []func()

	for i := 0; i < 3; i++ {
		fs = append(fs, func() { //将多个匿名函数添加到列表
			fmt.Println(&i, i)
		})
	}
	return fs
}

func fixFly() []func() {
	var fs []func()

	for i := 0; i < 3; i++ {
		x := i
		fs = append(fs, func() { //将多个匿名函数添加到列表
			fmt.Println(&x, x)
		})
	}
	return fs
}

func TestFly(t *testing.T) {
	for _, f := range fly() {
		f()
	}
}

func TestFixFly(t *testing.T) {
	for _, f := range fixFly() {
		f()
	}
}

func TestModifyGlobalVariable(t *testing.T) {
	var x int = 1
	y := func() int {
		x += 1
		return x
	}
	fmt.Println("result: ", x, y()) // 先执行函数, 然后执行打印
}

//defer延迟调用
func TestDeferDelayCall(t *testing.T) {
	x, y := 1, 2
	defer func(a int) {
		fmt.Printf("x: %d, y:%d\n", a, y) //y为闭包引用
	}(x)

	x += 100
	y += 100
	fmt.Println(x, y)
}
