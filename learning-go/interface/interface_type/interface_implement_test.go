package interface_type

import "testing"

/*
 * @author taohu
 * @date 2020/3/17
 * @time 下午8:21
 * @desc 接口实现两种形式: 结构体指针实现和结构体实现
 * go test -gcflags=-N -benchmem -test.count=3 -test.cpu=1 -test.benchtime=1s -bench=.
**/

type Animal interface {
	CanFly() bool
}

type Bird struct {
	Name string
}

// 结构体指针实现
func (b *Bird) CanFly() bool {
	return true
}

func BenchmarkDynamicDispatch(b *testing.B) {
	bird := &Bird{Name: "test"}
	animal := Animal(bird)
	for n := 0; n < b.N; n++ {
		animal.CanFly()
	}
}

func BenchmarkDirectCall(b *testing.B) {
	bird := &Bird{Name: "test"}
	for n := 0; n < b.N; n++ {
		bird.CanFly()
	}
}

// 结构体实现
//func (b Bird) CanFly() bool {
//	return true
//}
//
//func BenchmarkDynamicDispatch(b *testing.B) {
//	bird := Bird{Name: "test"}
//	animal := Animal(bird)
//	for n := 0; n < b.N; n++ {
//		animal.CanFly()
//	}
//}
//
//func BenchmarkDirectCall(b *testing.B) {
//	bird := Bird{Name: "test"}
//	for n := 0; n < b.N; n++ {
//		bird.CanFly()
//	}
//}
