package area

import "math"

/*
 * @author taohu
 * @date 2019/11/26
 * @time 下午7:17
 * @desc 计算面积
**/

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
