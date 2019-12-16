package simple_factory

import "fmt"

/*
 * @author taohu
 * @date 2019/12/16
 * @time 上午10:03
 * @desc 简单工厂模式
**/

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI return Api instance by type
func NewAPI(t int) API {
	switch t {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	}
	return nil
}

// hiAPI is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// HelloAPI is another API implement
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
