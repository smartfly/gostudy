package fake

import (
	"fmt"
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * @author taohu
 * @date 2019/11/23
 * @time 下午3:42
 * @desc please describe function
**/

func TestCompareInt(t *testing.T) {
	//convey.Convey("CompareInttest", t, func() {
	//	// ApplyFunc 的第一个参数是待测试的函数名字；
	//	// 第二个参数是待测试函数的函数原型
	//	// 这里模拟返回值是0 的情况， 我们还可以返回1或者-1
	//	patch := gomonkey.ApplyFunc(CompareInt, func(num1 int, num2 int) int {
	//		fmt.Println(num1+num2)
	//		return -1
	//	})
	//	defer patch.Reset()
	//	result := TestFunc(1,1)
	//	convey.So(result, convey.ShouldEqual, -1)
	//})
	//fmt.Println("CompareInt test success")

	p := gomonkey.NewPatches()
	p.ApplyFunc(compareInt, func(_ int, _ int) int {
		return -8
	})
	defer p.Reset()
	result := TestFunc(1, 1)
	assert.Equal(t, -8, result)
	fmt.Println("CompareInt test success")
}
