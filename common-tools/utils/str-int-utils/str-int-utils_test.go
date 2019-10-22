package str_int_utils

import (
	"fmt"
	"testing"
)

/*
 * @author: taohu
 * @date: 19-5-21
 * @time: 上午10:44
 * @desc: 包内测试
**/

func TestInt2Str(t *testing.T) {
	value := Int2Str(2000)
	if value == "2000" {
		fmt.Println(true)
	}
}

func TestInt642Str(t *testing.T) {
	value := Int642Str(1000000000000000000)
	if value == "1000000000000000000" {
		fmt.Println(true)
	}
}

func TestStr2Int(t *testing.T) {
	str, err := Str2Int("2000")
	if err != nil {
		fmt.Println(false)
	}
	fmt.Println(str)
}

func TestStr2Int64(t *testing.T) {
	str, err := Str2Int64("1000000000000000000")
	if err != nil {
		fmt.Println(false)
	}
	fmt.Println(str)
}
