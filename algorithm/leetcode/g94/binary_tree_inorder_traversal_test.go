package g94

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	root.Right = node2
	node2.Left = node3
	ret := inorderTraversal(root)
	fmt.Println(ret)
}
