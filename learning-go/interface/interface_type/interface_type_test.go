package interface_type_test

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2020/3/17
 * @time 下午2:27
 * @desc Go语言的接口类型不是任意类型
**/

type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	return v == nil
}

func TestInterface(t *testing.T) {
	var s *TestStruct
	fmt.Println(s == nil)
	fmt.Println(NilOrNot(s))
}

// 1. 将上述变量与nil比较会返回true
// 2. 将上述变量传入NilOrNot方法并与nil比较会返回false
// 3. 调用NilOrNot函数时发生了隐式的类型转换、除了向方法传入参数之外，变量的赋值也会触发隐式类型转换。
// 在类型转换时，*TestStruct 类型会转换成 interface{} 类型，转换后的变量不仅包含转换前的变量，还包含变量的类型信息 TestStruct，
// 所以转换后的变量与 nil 不相等。
