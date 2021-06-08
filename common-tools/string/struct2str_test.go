package string

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
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

func TestPrint(t *testing.T) {
	toBeCharge := "2015-01-01"                                      //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02"                                      //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)
}
