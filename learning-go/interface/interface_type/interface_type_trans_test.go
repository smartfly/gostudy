package interface_type

import "testing"

/*
 * @author taohu
 * @date 2020/3/17
 * @time 下午7:51
 * @desc please describe function
**/

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c *Cat) Quack() {
	println(c.Name + " meow")
}

func Test(t *testing.T) {
	var c Duck = &Cat{Name: "grooming"}
	c.Quack()
}
