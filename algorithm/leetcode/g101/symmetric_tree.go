package g101

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var traversal func(l *TreeNode, r *TreeNode) bool
	traversal = func(l *TreeNode, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}

		if l == nil || r == nil {
			return false
		}
		return l.Val == r.Val && traversal(l.Left, r.Right) && traversal(l.Right, r.Left)
	}
	return traversal(root, root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
