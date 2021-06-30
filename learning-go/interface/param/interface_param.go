package param

import "fmt"

type Animal interface {
	PrintAnimal(animal Animal)
}

type Dog struct {
	Name string
	claw string
}

type Cat struct {
	Name string
}

func (d *Dog) PrintAnimal(animal Animal) {
	fmt.Println(animal.(*Dog).Name + animal.(*Dog).claw)
}

func (d *Cat) PrintAnimal(animal Animal) {
	fmt.Println(animal.(*Cat).Name)
}
