package design_pattern

/*
 * @author taohu
 * @date 19-8-23
 * @time 上午10:16
 * @desc 简单工厂模式
**/

type OperateR interface {
	Operate(int, int) int
}

type AddOperate struct {
}

func (o *AddOperate) Operate(a int, b int) int {
	return a + b
}

type MultipleOperate struct {
}

func (o *MultipleOperate) Operate(a int, b int) int {
	return a * b
}

type OperateFactory struct {
}

func NewOperateFactory() *OperateFactory {
	return &OperateFactory{}
}

func (f *OperateFactory) CreateOperate(operateName string) OperateR {
	switch operateName {
	case "+":
		return &AddOperate{}
	case "*":
		return &MultipleOperate{}
	default:
		return nil
	}
}
