package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectType(t *testing.T) {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) // type: float64
}

func TestReflectValue(t *testing.T) {
	var x float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(x).String()) // value: <float64 Value>
}

func TestReflectValueMethod(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                               // type: float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // kind is float64: true
	fmt.Println("value:", v.Float())                             // value: 3.4
}

func TestReflectValueLargestType(t *testing.T) {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Uint())                                       // v.Uint returns a uint64.
	fmt.Println(x)                                            // 120
}

func TestMyInt(t *testing.T) {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                        // type: reflect.MyInt
	fmt.Println("kind is int: ", v.Kind() == reflect.Int) // true.
	fmt.Println(v.Int())                                  // 7
}

func TestVInterface(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println(v.Interface())
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(y)
}
