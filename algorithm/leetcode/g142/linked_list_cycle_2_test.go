package g142

import (
	"fmt"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	a := &ListNode{Val: 2}
	head.Next = a
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = a
	ret := detectCycle(head)
	fmt.Println(ret.Val)
}
