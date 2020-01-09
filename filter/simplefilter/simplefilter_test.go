package simplefilter

import (
	"go/filter/simplefilter/account"
	"testing"
)

/*
 * @author taohu
 * @date 2020/1/4
 * @time 下午3:04
 * @desc please describe function
**/

func TestProxy_Query(t *testing.T) {

	account.New = func(id, name string, value int) account.Account {
		a := &account.AccountImpl{id, name, value}
		p := &Proxy{a}
		return p
	}

	id := "100111"
	a := account.New(id, "ZhangSan", 100)
	a.Query(id)
	a.Update(id, 500)
}
