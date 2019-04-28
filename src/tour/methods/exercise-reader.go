package main

import "golang.org/x/tour/reader"

/**
 * @author taohu
 * @date 19-4-28
 * @time 下午2:53
 * @desc 练习：实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
