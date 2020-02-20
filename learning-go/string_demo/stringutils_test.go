package string_demo

import (
	"fmt"
	"strings"
	"testing"
)

/*
 * @author taohu
 * @date 2020/1/16
 * @time 下午3:53
 * @desc please describe function
**/

func Benchmark_fmt(t *testing.B) {
	a := "test"
	b := "demo"
	result := fmt.Sprintf("%s_%s", a, b)
	fmt.Println(result)
}

func Benchmark_strings(t *testing.B) {
	a := "test"
	b := "demo"
	builder := strings.Builder{}
	builder.WriteString(a)
	builder.WriteString("_")
	builder.WriteString(b)
	fmt.Println(builder.String())
}
