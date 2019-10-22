package main

import "fmt"

/**
 * @author taohu
 * @date 19-4-28
 * @time 上午9:58
 * @desc 类型推断
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
