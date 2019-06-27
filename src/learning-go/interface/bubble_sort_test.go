package _interface

import (
	"fmt"
	"testing"
)

/*
 * @author: taohu
 * @date: 19-6-21
 * @time: 上午10:00
 * @desc: please describe function
**/

func TestBubble(t *testing.T) {
	a := []int{5, 8, 3, 9, 0}
	bubbleSort(a)
	fmt.Println(a)
}

func TestBubbleEnhance(t *testing.T) {
	a := Xi{5, 8, 3, 9, 0}
	Sort(a)
	fmt.Println(a)
	s := Xs{"nut", "ape", "elephant", "zoo", "go"}
	Sort(s)
	fmt.Println(s)
}
