package slice

import (
	"bytes"
	"fmt"
)

func dataShare() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	fmt.Println(foo, bar)
}

// append()这个函数在cap不够用的时候就会重新分配内存以扩大容量, 而如果够用的时候不会重新分配内存
func appendOp() {
	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42
	fmt.Println(a, b)
}

func appendOpByCapEnough() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}
