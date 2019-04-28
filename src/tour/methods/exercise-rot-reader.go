package main

import (
	"io"
	"os"
	"strings"
)

/**
 * @author taohu
 * @date 19-4-28
 * @time 下午3:23
 * @desc 练习：rot13Reader
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type rot13Reader struct {
	r io.Reader
}

// 实现加密函数
func rot13(b byte) byte {
	var a byte
	switch {
	case b >= 'a' && b <= 'z':
		a = 'a'
	case b >= 'A' && b <= 'Z':
		a = 'A'
	default:
		return b
	}
	return (b-a+13)%26 + a
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	// 先从元数据上读取内容
	n, err = r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
