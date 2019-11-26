package area

import (
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2019/11/26
 * @time 下午7:19
 * @desc please describe function
**/

func TestUnitByRectangle(t *testing.T) {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))

}

func TestUnitByArea(t *testing.T) {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
