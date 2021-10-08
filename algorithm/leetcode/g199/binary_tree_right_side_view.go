package g199

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	levelResult := make([][]int, 0)
	var traversal func(root *TreeNode, level int)
	traversal = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(levelResult) {
			levelResult = append(levelResult, []int{})
		}
		levelResult[level] = append(levelResult[level], root.Val)
		traversal(root.Left, level+1)
		traversal(root.Right, level+1)
	}
	traversal(root, 0)
	result := make([]int, len(levelResult))
	for i, array := range levelResult {
		result[i] = array[len(array)-1]
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
