package simplefilter

import (
	"fmt"
	"go/filter/simplefilter/account"
)

/*
 * @author taohu
 * @date 2020/1/4
 * @time 下午3:01
 * @desc please describe function
**/

type Proxy struct {
	Account account.Account
}

func (p *Proxy) Query(id string) int {
	fmt.Println("Proxy.Query begin")
	value := p.Account.Query(id)
	fmt.Println("Proxy.Query end")
	return value
}

func (p *Proxy) Update(id string, value int) {
	fmt.Println("Proxy.Update begin")
	p.Account.Update(id, value)
	fmt.Println("Proxy.Update end")
}
