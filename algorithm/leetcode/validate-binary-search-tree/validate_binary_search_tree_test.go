package validate_binary_search_tree

import (
	"fmt"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}
	a := isValidBST(root)
	fmt.Println(a)
}
