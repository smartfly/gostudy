package array

import "testing"

func Test_traverseByCommon(t *testing.T) {
	arr := [][]int{{1, 4, 3}, {7, 5, 6}}
	traverseByCommon(arr)
}

func Test_traverseByRange(t *testing.T) {
	arr := [][]int{{1, 4, 3}, {7, 5, 6}}
	traverseByRange(arr)
}
