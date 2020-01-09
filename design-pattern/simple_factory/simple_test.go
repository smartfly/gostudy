package simple_factory

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"os"
	"testing"
)

/*
 * @author taohu
 * @date 2019/12/16
 * @time 上午10:38
 * @desc 简单工厂模式测试
**/

// TestType1 test get hiApi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	assert.Equal(t, "Hi, Tom", s)
}

// TestType2 test get helloApi with factory
func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	assert.Equal(t, "Hello, Tom", s)
	fmt.Println(os.Getpid())
}
