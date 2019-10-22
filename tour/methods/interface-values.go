package main

import (
	"fmt"
	"math"
)

/**
 * @author taohu
 * @date 19-4-28
 * @time 上午9:45
 * @desc 接口值
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	t := &T{"hello"}
	describe(t)
	t.M()

	f := F(math.Pi)
	describe(f)
	f.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
