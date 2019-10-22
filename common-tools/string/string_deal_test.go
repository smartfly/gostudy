package string

import (
	"fmt"
	"strings"
	"testing"
)

/*
 * @author: taohu
 * @date: 19-7-5
 * @time: 上午9:17
 * @desc: 字符串处理demo
**/

// 去除字符串收尾空格
func TestStringTrimSpace(t *testing.T) {
	str := " Hello World! "
	str = strings.TrimSpace(str)
	fmt.Println(str)
}
