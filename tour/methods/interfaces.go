package main

import (
	"fmt"
	"math"
)

/**
 * @author taohu
 * @date 19-4-23
 * @time 上午9:50
 * @desc 接口
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	//v := Vertex2{3, 4}

	a = f // a MyFloat 实现了Abser
	//a = &v // a *Vertex 实现了Abser

	//a = v // v是一个Vertex(而不是*Vertex)
	fmt.Println(a.Abs())
}
