package map_api

import "testing"

// go test -bench=. -test.benchmem ./map_advance_test.go
// -test.benchmem 是一个可选的标志，用于显示内存分配
func test(m map[int]int) {
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := make(map[int]int)
		b.StartTimer()
		test(m)
	}
}

func BenchmarkCapMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := make(map[int]int, 10000)
		b.StartTimer()
		test(m)
	}
}

/*
goos: linux
goarch: amd64
BenchmarkMap-4              1624            725604 ns/op          687156 B/op        275 allocs/op
BenchmarkCapMap-4           3618            309102 ns/op            2699 B/op          9 allocs/op
		|		|			 |						|					|					|
	   名称	 CPU数量		  总运行次数            每次操作的纳秒数        每次操作的字节数    每次操作的内存分配数
PASS
ok      command-line-arguments  3.376s
*/
