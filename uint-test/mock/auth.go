package mock

import "errors"

/*
 * @author taohu
 * @date 2019/11/8
 * @time 上午9:40
 * @desc please describe function
**/

type AuthService interface {
	Login(userName, password string) (token string, err error)
	Logout(token string) error
}

type authService struct{}

func (auth *authService) Login() (string, error) {
	return "token", nil
}

func (auth *authService) Logout(token string) error {
	return nil
}

// 模拟登录失败
type authLoginErr struct {
	auth AuthService // 可以使用组合的特性，Logout方法我们不关心，只用"覆盖"Login方法即可
}

func (auth *authLoginErr) Login(userName string, password string) (string, error) {
	return "", errors.New("用户密码错误")
}
