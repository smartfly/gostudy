package _struct

import "testing"

/*
 * @author taohu
 * @date 2020/5/13
 * @time 下午8:46
 * @desc 因为Go 白皮书明确规定比较两个结构体值时空字段(名为_字段将被忽略) https://golang.org/ref/spec#Comparison_operators
**/

type T struct {
	_ int
	_ bool
}

func TestStructEqual(t *testing.T) {
	var t1 = T{123, true}
	var t2 = T{789, false}
	println(t1 == t2)
}
