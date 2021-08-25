package ioc

import (
	"fmt"
	"testing"
)

func TestIntSetAdd(t *testing.T) {
	intSet := NewIntSet()
	intSet.Add(5)
	intSet.Add(4)
	fmt.Println(intSet.data)
	_ = intSet.undo.Undo()
	fmt.Println(intSet.data)
	_ = intSet.undo.Undo()
	fmt.Println(intSet.data)
}
