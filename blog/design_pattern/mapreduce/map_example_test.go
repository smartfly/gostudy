package mapreduce

import (
	"fmt"
	"strings"
	"testing"
)

func TestMapStr(t *testing.T) {
	list := []string{"Hao", "Chen", "MegaEase"}
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x) // [HAO CHEN MEGAEASE]

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y) // [3 4 8]
}

func TestReduce(t *testing.T) {
	list := []string{"Hao", "Chen", "MegaEase"}
	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x) // 15
}

func TestFilter(t *testing.T) {
	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n", out) // [1 3 5 7 9]

	out = Filter(intset, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out) // [6 7 8 9 10]
}
