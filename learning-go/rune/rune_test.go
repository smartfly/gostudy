package rune

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(len(m)) // 100
}
