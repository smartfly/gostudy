package param

import "testing"

func TestDog(t *testing.T) {
	d := Dog{
		Name: "哈士奇",
	}
	d.PrintAnimal(&d)
}

func TestCat(t *testing.T) {
	c := Cat{
		Name: "蓝猫",
	}
	c.PrintAnimal(&c)
}
