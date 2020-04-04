package add_two_numbers

import (
	"fmt"
	"reflect"
	"testing"
)

/*
 * @author taohu
 * @date 2020/4/3
 * @time 下午8:58
 * @desc please describe function
**/

// [2, 4, 3] + [5, 6, 4] = [7, 0, 8]
func Test_addTwoNumbers(t *testing.T) {

	l11 := new(ListNode)
	l11.Val = 2

	ll2 := new(ListNode)
	ll2.Val = 4
	l11.Next = ll2

	l13 := new(ListNode)
	l13.Val = 3
	ll2.Next = l13

	fmt.Println(l11.Val, l11.Next.Val, l11.Next.Next.Val)

	l21 := new(ListNode)
	l21.Val = 5

	l22 := new(ListNode)
	l22.Val = 6
	l21.Next = l22

	l23 := new(ListNode)
	l23.Val = 4
	l22.Next = l23
	fmt.Println(l21.Val, l21.Next.Val, l21.Next.Next.Val)

	l31 := new(ListNode)
	l31.Val = 7

	l32 := new(ListNode)
	l32.Val = 0
	l31.Next = l32

	l33 := new(ListNode)
	l33.Val = 8
	l32.Next = l33
	fmt.Println(l31.Val, l31.Next.Val, l31.Next.Next.Val)

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: l11,
				l2: l21,
			},
			want: l31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// [1] + [9, 9] = [0, 0, 1]
func Test_addTwoNumbersByCase1(t *testing.T) {

	l11 := new(ListNode)
	l11.Val = 1

	l21 := new(ListNode)
	l21.Val = 9
	l22 := new(ListNode)
	l22.Val = 9
	l21.Next = l22

	l31 := new(ListNode)
	l31.Val = 0
	l32 := new(ListNode)
	l32.Val = 0
	l31.Next = l32
	l33 := new(ListNode)
	l33.Val = 1
	l32.Next = l33
	fmt.Println(l31.Val, l31.Next.Val, l31.Next.Next.Val)

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: l11,
				l2: l21,
			},
			want: l31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
