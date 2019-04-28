package main

import (
	"fmt"
	"math"
)

/**
 * @author taohu
 * @date 19-4-22
 * @time 下午1:09
 * @desc 方法：一类带特殊的接收者参数的函数，
 * 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
