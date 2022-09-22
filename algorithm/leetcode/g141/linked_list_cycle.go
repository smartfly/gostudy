package g141

import "fmt"

/**
 * Definition iterator singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // 最大的 int
	fmt.Printf("%b\n", min)
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
