package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 虚拟头节点
	newHead := &ListNode{
		Next: head,
	}
	pre := newHead
	left := head
	if left == nil {
		return head
	}
	right := head.Next
	for left != nil && right != nil {
		pre.Next = right
		left.Next = right.Next
		right.Next = left
		if left.Next == nil {
			break
		}
		pre, left, right = left, left.Next, left.Next.Next
	}
	return newHead.Next
}
