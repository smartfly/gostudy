package pipeline

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range sum(sq(odd(echo(nums)))) {
		fmt.Println(n)
	}
}

func TestProxy(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println(n)
	}
}
