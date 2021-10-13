package g257

import (
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 5}
	root.Left, root.Right = node1, node2
	node2.Right = node3
	binaryTreePaths(root)
}
