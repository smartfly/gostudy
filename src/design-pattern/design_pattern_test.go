package design_pattern

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 19-8-21
 * @time 上午10:06
 * @desc please describe function
**/

func TestUser_String(t *testing.T) {
	saveUser(&user{"smartfly"})
}

func TestNewOperateFactory(t *testing.T) {
	operator := NewOperateFactory().CreateOperate("+")
	fmt.Printf("add result is %d\n", operator.Operate(1, 2))
	operator = NewOperateFactory().CreateOperate("*")
	fmt.Printf("multi result is %d\n", operator.Operate(2, 4))
}
