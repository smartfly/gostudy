package main

import (
	"fmt"
	"io"
	"strings"
)

/**
 * @author taohu
 * @date 19-4-28
 * @time 下午2:37
 * @desc Reader
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
