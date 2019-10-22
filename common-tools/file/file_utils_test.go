package file

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2019/9/12
 * @time 下午7:26
 * @desc flie测试
**/

func TestReadFile(t *testing.T) {
	linesList, err := ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(linesList)
}
