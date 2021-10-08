package g102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var traversal func(root *TreeNode, level int)
	traversal = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(result) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], root.Val)
		traversal(root.Left, level+1)
		traversal(root.Right, level+1)
	}
	traversal(root, 0)
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
