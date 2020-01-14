package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * @author taohu
 * @date 2020/1/14
 * @time 上午9:46
 * @desc please describe function
**/

func TestBuilder1(t *testing.T) {
	asserts := assert.New(t)
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	asserts.Equal("123", res, "generator")
}

func TestBuilder2(t *testing.T) {
	asserts := assert.New(t)
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	asserts.Equal(6, res, "plus")
}
