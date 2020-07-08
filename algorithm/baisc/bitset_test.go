package baisc

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2020/6/17
 * @time 上午10:35
 * @desc please describe function
**/

func Test_createBitSet(t *testing.T) {
	createBitSet(128)
	fmt.Println(bitSetValues)
	set(100)
	fmt.Println(bitSetValues)
	flag := get(101)
	fmt.Println(flag)
}

func Test_createBitSet1(t *testing.T) {
	createBitSet(32)
	fmt.Println(bitSetValues)
	set(10)
	fmt.Println(bitSetValues)
	flag := get(11)
	fmt.Println(flag)
}
