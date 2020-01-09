package factorymethod

import "testing"

/*
 * @author taohu
 * @date 2020/1/9
 * @time 下午7:46
 * @desc please describe function
**/

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)
	factory = PlusOperatorFactory{}

	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}

	factory = MinusOperatorFactory{}
	if compute(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
