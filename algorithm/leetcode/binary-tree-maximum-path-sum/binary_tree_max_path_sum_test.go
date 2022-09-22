package binary_tree_maximum_path_sum

import (
	"fmt"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	root := TreeNode{Val: 0}
	a := maxPathSum(&root)
	fmt.Println(a)
}
