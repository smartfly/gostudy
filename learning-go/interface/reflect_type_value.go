package _interface

import (
	"reflect"
)

/*
 * @author: taohu
 * @date: 19-7-2
 * @time: 上午9:13
 * @desc: 自省和反射
**/

type Person struct {
	Name string "test"
	age  int
}

func Show(i interface{}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
	}
}
