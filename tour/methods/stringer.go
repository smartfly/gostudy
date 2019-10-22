package main

import "fmt"

/**
 * @author taohu
 * @date 19-4-28
 * @time 上午10:07
 * @desc stringer 接口
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
