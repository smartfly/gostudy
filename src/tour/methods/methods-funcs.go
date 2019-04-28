package main

import (
	"fmt"
	"math"
)

/**
 * @author taohu
 * @date 19-4-22
 * @time 下午1:17
 * @desc 方法即函数：方法只是个带接收者参数的函数
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type Vertex1 struct {
	X, Y float64
}

func Abs(v Vertex1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex1{3, 4}
	fmt.Println(Abs(v))
}
