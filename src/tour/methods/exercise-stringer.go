package main

import "fmt"

/**
 * @author taohu
 * @date 19-4-28
 * @time 上午10:10
 * @desc 通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。 例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ipAddr IPAddr) String() string {
	var result string
	for _, value := range ipAddr {
		result += fmt.Sprintf("%d.", value)
	}
	return result[:len(result)-1]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
