package interfacecode

import (
	"fmt"
	"testing"
)

func TestPrintPerson(t *testing.T) {
	p := Person{
		Name:   "woody",
		Sexual: "Male",
		Age:    28,
	}
	PrintPerson(&p)
	p.Print()
}

func TestPrintStr(t *testing.T) {
	a := Country{"China"}
	b := City{"Shenzhen"}
	PrintStr(a)
	PrintStr(b)
}

func TestShape(t *testing.T) {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}

// 校验square接口完整性
func TestCheckSquareIntegrity(t *testing.T) {
	//var _ Shape = (*Square) (nil)
	// Cannot use '(*Square) (nil)' (type *Square) as type Shape Type does not
	// implement 'Shape' as some methods are missing: Area() int
}
