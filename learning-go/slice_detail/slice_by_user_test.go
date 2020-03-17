package slice_detail

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2020/3/6
 * @time 下午7:48
 * @desc please describe function
**/

// slice的替换操作
func SliceReplace(s []int) {
	s[0] = 10
	fmt.Printf("inner: %v, &addr: %p\n", s, &s)
}

func TestSliceReplace(t *testing.T) {
	s := []int{2, 3, 4, 5, 6}
	fmt.Printf("init: %v, &addr: %p\n", s, &s)
	SliceReplace(s)
	fmt.Printf("out: %v, &addr: %p\n", s, &s)
}

// slice的append操作
func sliceAppend(a []int) {
	a = append(a, 5)
	fmt.Printf("inner: %v, &init: %p\n", a, &a)
}

func TestSliceAppend(t *testing.T) {
	s := []int{2, 3}
	fmt.Printf("init: %v, &addr: %p\n", s, &s)
	sliceAppend(s)
	fmt.Printf("out: %v, &addr: %p\n", s, &s)
}
