package dynamic_method

import (
	"fmt"
	"reflect"
	"testing"
)

/*
 * @author taohu
 * @date 2020/8/13
 * @time 下午4:24
 * @desc golang动态调用方法 https://studygolang.com/articles/377
**/

type YourT1 struct {
}

func (y *YourT1) MethodBar() {
	fmt.Println("MethodBar called")
}

type YourT2 struct {
}

func (y *YourT2) MethodFoo(i int, v string) {
	fmt.Println("MethodFoo called", i, v)
}

func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}

func TestDynamicMethod(t *testing.T) {
	InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")
	InvokeObjectMethod(new(YourT1), "MethodBar")
}

func TestSlice(t *testing.T) {
	var s []int64
	s = append(s, 1)
	fmt.Println(s)
}

func TestMap(t *testing.T) {
	var s map[int]string
	s[1] = "test"
	fmt.Println(s)
}
