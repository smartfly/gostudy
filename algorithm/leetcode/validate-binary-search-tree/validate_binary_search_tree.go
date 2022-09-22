package validate_binary_search_tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// 前序遍历
	var f func(root *TreeNode, lower, upper int) bool
	f = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if lower >= root.Val || upper <= root.Val {
			return false
		}
		left := f(root.Left, lower, root.Val)
		right := f(root.Right, root.Val, upper)
		return left && right
	}
	return f(root, math.MinInt64, math.MaxInt64)
}
