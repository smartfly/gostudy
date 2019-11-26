package gomonkey

import (
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
 * @author taohu
 * @date 2019/11/23
 * @time 下午6:57
 * @desc please describe function
**/

var num = 10

func TestApplyGlobalVar(t *testing.T) {
	convey.Convey("TestApplyGlobalVar", t, func() {
		convey.Convey("change", func() {
			patches := gomonkey.ApplyGlobalVar(&num, 150)
			defer patches.Reset()
			convey.So(num, convey.ShouldEqual, 150)
		})

		convey.Convey("recover", func() {
			convey.So(num, convey.ShouldEqual, 10)
		})
	})

}
