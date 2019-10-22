package main

import (
	"fmt"
	"go/tour/tree"
)

/**
 * @author taohu
 * @date 19-4-20
 * @time 下午6:46
 * @desc 等价二叉查找树
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

// walk 步进tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, c chan int) {
	treeTraverseLeft(t, c)
	close(c)
}

// 树的左序遍历
func treeTraverseLeft(t *tree.Tree, ch chan int) {
	if t != nil {
		treeTraverseLeft(t.Left, ch)
		ch <- t.Value
		treeTraverseLeft(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
