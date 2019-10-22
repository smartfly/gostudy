package main

import (
	"fmt"
	"math"
)

/**
 * @author taohu
 * @date 19-4-28
 * @time 上午10:39
 * @desc 练习：错误
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type ErrNegativeSqrt float64

// fmt.Sprint(e)会优先调用Error()方法，即本身，这会导致死循环。
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	// 当它是一个负数时，返回一个错误
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := x / 2
	for i, d, t := 0, 0.0, 0.0; i < 10; i++ {
		t, z = z, z-(z*z-x)/(2*z)
		d = math.Abs(z - t)
		if d < 1e-12 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
