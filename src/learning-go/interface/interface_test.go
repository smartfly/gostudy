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

func TestShow(t *testing.T) {
	person := Person{
		Name: "路飞",
		age:  10,
	}
	fmt.Println(person)
	Show(&person)
	fmt.Println(person)
}

func TestInterfaceFunc(t *testing.T) {
	var a Integer = 1
	var b1 LessAdder = &a
	fmt.Printf("%v\n", b1.Less(1))
	var b2 LessAdder = a
	fmt.Println(b2)
}

type Integer int

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a Integer) Add(b Integer) {
	a += b
}
