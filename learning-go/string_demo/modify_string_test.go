package string_demo

import (
	"fmt"
	"testing"
)

// 上面的实例并不是更新字符串正确姿势，因为一个UTF8编码的字符可能会占多个字节，
// 比如汉字就需要3~4个字节来存储, 此时更新其中的一个字节是错误的
func TestModifyStringByByte(t *testing.T) {
	x := "text"
	xBytes := []byte(x)
	xBytes[0] = 'T' // 注意此时的T是rune类型 T-->我 提示 constant 25105 overflows byte
	x = string(xBytes)
	fmt.Println(x) // Text
}

func TestModifyStringByRune(t *testing.T) {
	x := "text"
	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x) // 我ext
}
