package binary_tree_maximum_path_sum

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

var maxVal int

func maxPathSum(root *TreeNode) int {
	maxVal = math.MinInt64
	traverse(root)
	return maxVal
}

// traverse 后序遍历
func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftVal := findMax(traverse(root.Left), 0)
	rightVal := findMax(traverse(root.Right), 0)
	pathVal := leftVal + rightVal + root.Val
	maxVal = findMax(maxVal, pathVal)
	return root.Val + findMax(leftVal, rightVal)
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
