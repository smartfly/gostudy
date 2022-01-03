package g142

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			pre := head
			for pre != slow {
				pre = pre.Next
				slow = slow.Next
			}
			return pre
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
