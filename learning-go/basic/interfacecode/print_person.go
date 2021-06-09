package interfacecode

import "fmt"

// 成员函数&函数
type Person struct {
	Name   string
	Sexual string
	Age    int32
}

// 函数
func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

// 成员函数
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

// Program to an interface not an implementation
type Country struct {
	Name string
}

type City struct {
	Name string
}

type StringAble interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}

func (c City) ToString() string {
	return "City = " + c.Name
}

func PrintStr(p StringAble) {
	fmt.Println(p.ToString())
}

// interface integrity
type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}
