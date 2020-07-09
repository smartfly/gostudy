package string

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
 * @author taohu
 * @date 2020/7/9
 * @time 下午2:25
 * @desc please describe function
**/

type Student struct {
	NumberID string `json:"number_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type Chemistry struct {
	TeacherName string     `json:"teacher_name"`
	Students    []*Student `json:"students"`
}

func TestStruct2Str(t *testing.T) {
	s1 := Student{
		NumberID: "1",
		Name:     "小明",
		Age:      15,
	}

	s2 := Student{
		NumberID: "2",
		Name:     "小花",
		Age:      15,
	}

	students := make([]*Student, 0)
	students = append(students, &s1)
	students = append(students, &s2)

	chemistry := Chemistry{
		TeacherName: "小可",
		Students:    students,
	}

	fmt.Println(chemistry)
	values, _ := json.Marshal(chemistry)
	fmt.Println(string(values))
}
