package enum

import (
	"fmt"
	"testing"
)

type ArticleState int

const (
	Draft ArticleState = iota + 1
	Published
	Deleted
)

func checkArticleState(state ArticleState) {
	// ...
	fmt.Println(state)
}

func TestEnum(t *testing.T) {

	checkArticleState(Deleted)
	checkArticleState(Published)
	//a := 0
	//checkArticleState(a)  // 编译报错
	checkArticleState(0) //
}

func TestEnumString(t *testing.T) {
	var f FishType = C
	fmt.Println(f)
	switch f {
	case A:
		fmt.Println("A 测试完成了")
	case B:
		fmt.Println("B 也来了")
	case C:
		fmt.Println("C 也来了")
	case D:
		fmt.Println("D 也来了")
	default:
		fmt.Println("不知道怎么来了")
	}
}
