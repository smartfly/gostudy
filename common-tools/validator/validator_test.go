package validator

import (
	"fmt"
	"testing"
	"time"
	"unsafe"

	"github.com/go-playground/validator/v10"
)

/*
 * @author taohu
 * @date 2020/9/23
 * @time 下午8:13
 * @desc https://github.com/go-playground/validator 验证struct结构体
**/

type User struct {
	Name string `validate:"min=6,max=10"`
	Age  int    `validate:"min=1,max=100"`
}

func TestUsageValidator(t *testing.T) {
	validate := validator.New()

	u1 := User{Name: "smart-fly", Age: 18}
	err := validate.Struct(u1)
	fmt.Println(err)

	u2 := User{Name: "lucas", Age: 101}
	err = validate.Struct(u2)
	fmt.Println(err)
}

func TestPtr(t *testing.T) {
	a := 10
	aPtr := &a
	c := 11

	fmt.Println(a)
	fmt.Println(aPtr)
	fmt.Println(*aPtr)
	var b *int
	b = &c
	fmt.Println(b)
	fmt.Println(*b)
	b = &(*aPtr)
	fmt.Println(b)
	fmt.Println(*b)
}

func TestSlice(t *testing.T) {
	a := make([]interface{}, 0, 10)
	fmt.Println(a)
	a = append(a, 1)
	fmt.Println(a)
}

func TestBit(t *testing.T) {
	a := int(1531289130965395)
	fmt.Println(unsafe.Sizeof(a))
}

func TestTime(t *testing.T) {
	a := time.Duration(1000) * time.Millisecond
	fmt.Println(a)
}
