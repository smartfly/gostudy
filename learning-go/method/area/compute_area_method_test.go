package area

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
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

func TestUintBoxList(t *testing.T) {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	assert.Equal(t, 5, len(boxes))
	assert.Equal(t, float64(64), boxes[0].Volume())
	assert.Equal(t, "YELLOW", boxes[len(boxes)-1].color.String())
	assert.Equal(t, "YELLOW", boxes.BiggestColor().String())
	boxes.PaintItBlack()
	assert.Equal(t, "BLACK", boxes[1].color.String())
	assert.Equal(t, "BLACK", boxes.BiggestColor().String())
	// 注释
}
