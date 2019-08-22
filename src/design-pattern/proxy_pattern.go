package design_pattern

import "fmt"

/*
 * @author taohu
 * @date 19-8-21
 * @time 上午10:01
 * @desc go 代理模式
**/

type user struct {
	Name string
}

func (u *user) String() string {
	return u.Name
}

func saveUser(user *user) {
	withTx(func() {
		fmt.Printf("保存用户 %s\n", user.Name)
	})
}

func withTx(fn func()) {
	fmt.Println("开启事务")
	fn()
	fmt.Println("结束事务")
}
