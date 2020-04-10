package _defer

import (
	"log"
	"strconv"
	"testing"
)

/*
 * @author taohu
 * @date 2020/4/10
 * @time 下午5:11
 * @desc please describe function
**/

// defer 延迟调用
func Test_Delay(t *testing.T) {
	defer log.Println("smartfly")
	log.Println("end.")
}

// defer 先进先出
func Test_LIFO(t *testing.T) {
	for i := 0; i < 6; i++ {
		defer log.Println("smartfly" + strconv.Itoa(i) + ".")
	}
	log.Println("end.")
}

// defer 运行时间点
func Test_RunTime(t *testing.T) {
	func() {
		defer log.Println("defer.smartfly.")
	}()
	log.Println("main.smartfly")
}

// defer 异常处理
func Test_Exception(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("smartfly.")
		}
	}()
	panic("end.")
}
