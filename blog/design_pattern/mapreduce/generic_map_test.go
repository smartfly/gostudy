package mapreduce

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}
	squaredArr := Map(nums, square)
	fmt.Println(squaredArr)

	upCase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, upCase)
	fmt.Println(upstrs)
	//[HAO CHEN MEGAEASE]

	x := Map(5, 5)
	fmt.Println(x)
}
