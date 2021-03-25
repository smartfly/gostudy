package nil

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2021/3/1
 * @time 上午11:39
 * @desc please describe function
**/

func TestNilValuesNotEqual(t *testing.T) {
	ret := (interface{})(nil) == (*int)(nil)
	fmt.Println(ret) // false
}

func TestRetrieveElementsForNilMaps(t *testing.T) {
	retInt := (map[string]int)(nil)["key"]
	fmt.Println(retInt) // 0
	retBool := (map[int]bool)(nil)[123]
	fmt.Println(retBool) // false
	retPtr := (map[int]*int64)(nil)[123]
	fmt.Println(retPtr) // <nil>
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func TestExample(t *testing.T) {
	example(5)
}

func TestDouble(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}
