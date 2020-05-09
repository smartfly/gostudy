package reverse_integer

import "math"

/*
 * @author taohu
 * @date 2020/5/4
 * @time 上午11:47
 * @desc 整数反转
**/

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	sign := 1
	init := x
	value := 0
	if x < 0 {
		sign = -1
		init = sign * x
	}
	for init > 0 {
		temp := init % 10
		value = value*10 + temp // int32存在overflow问题
		init = init / 10
	}

	ret := sign * value

	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}

	return ret
}
