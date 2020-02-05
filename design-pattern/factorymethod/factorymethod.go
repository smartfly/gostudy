package factorymethod

/*
 * @author taohu
 * @date 2020/1/9
 * @time 下午7:35
 * @desc 工厂方法模式
**/

// Operator 被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// operatorBase 是Operator接口实现的基类，封装公用方法
type operatorBase struct {
	a, b int
}

// SetA 设置A
func (o *operatorBase) SetA(a int) {
	o.a = a
}

// SetB 设置B
func (o *operatorBase) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		operatorBase: &operatorBase{},
	}
}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*operatorBase
}

//Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		operatorBase: &operatorBase{},
	}
}

//MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*operatorBase
}

//Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}
