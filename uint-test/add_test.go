package uint_test

import "testing"

/*
 * @author taohu
 * @date 2019/10/21
 * @time 下午9:54
 * @desc please describe function
**/

func Add(a, b int) int {
	return a + b
}

func TestAdd1(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Errorf("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

//func TestAdd2(t *testing.T) {
//	if Add(2, 3) != 6 {
//		t.Fatal("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N 测试循环次数
		Add(4, 5)
	}
}
