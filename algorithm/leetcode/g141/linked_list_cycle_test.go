package g141

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	a := []int{100, 5}
	ret := Min(a...)
	fmt.Println(ret)
}
