package add_two_numbers

/*
 * @author taohu
 * @date 2020/4/2
 * @time 下午4:25
 * @desc please describe function
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val: 0,
	}
	now := result
	var carry int
	for true {
		value1, value2 := 0, 0
		if l1 == nil && l2 == nil {
			break
		}

		if l1 != nil {
			value1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			value2 = l2.Val
			l2 = l2.Next
		}

		val := (carry + value1 + value2) % 10
		carry = (carry + value1 + value2) / 10

		now.Next = &ListNode{
			Val: val,
		}
		now = now.Next
	}
	if carry == 1 {
		now.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return result.Next
}
