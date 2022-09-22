package tag

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Student struct {
	Name string
	Age  int `json:"age"`
	sex  int // (1-man, 2-women)
}

func TestJsonToTag(t *testing.T) {
	s := Student{
		Name: "smartfly",
		Age:  18,
		sex:  1,
	}
	result, _ := json.Marshal(s)
	fmt.Println(string(result)) // {"Name":"smartfly","age":18}
}
