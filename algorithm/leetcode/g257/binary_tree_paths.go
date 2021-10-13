package g257

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	var tmp []string
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			tmp = append(tmp, strconv.Itoa(root.Val))
			result = append(result, strings.Join(tmp, "->"))
			tmp = tmp[:len(tmp)-1]
			return
		}
		tmp = append(tmp, strconv.Itoa(root.Val))
		traversal(root.Left)
		traversal(root.Right)
		tmp = tmp[:len(tmp)-1]
	}
	traversal(root)
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
