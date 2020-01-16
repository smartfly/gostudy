package prototype

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

/*
 * @author taohu
 * @date 2020/1/14
 * @time 上午10:18
 * @desc please describe function
**/

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	t2 := t1.Clone()
	assert.Equal(t, t1, t2)
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1")
	t1 := c.(*Type1)
	assert.Equal(t, "type1", t1.name)
}
