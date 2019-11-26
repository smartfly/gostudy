package gomonkey

import (
	"encoding/json"
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"go/uint-test/mock/gomonkey/fake"
	"testing"
)

/*
 * @author taohu
 * @date 2019/11/23
 * @time 下午6:28
 * @desc please describe function
**/

var (
	outputExpect = "xxx-vethName100-yyy"
)

func TestApplyFunc(t *testing.T) {
	convey.Convey("TestApplyFunc", t, func() {
		patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
			return outputExpect, nil
		})
		defer patches.Reset()
		output, err := fake.Exec("", "")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, outputExpect)
	})

	convey.Convey("one func for fail", t, func() {
		patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
			return "", fake.ErrActual
		})
		defer patches.Reset()

		output, err := fake.Exec("", "")
		convey.So(err, convey.ShouldEqual, fake.ErrActual)
		convey.So(output, convey.ShouldEqual, "")
	})

	convey.Convey("two funcs", t, func() {
		patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
			return outputExpect, nil
		})
		defer patches.Reset()
		patches.ApplyFunc(fake.Belong, func(_ string, _ []string) bool {
			return true
		})
		output, err := fake.Exec("", "")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, outputExpect)
		flag := fake.Belong("", nil)
		convey.So(flag, convey.ShouldBeTrue)
	})

	convey.Convey("input and output param", t, func() {
		patches := gomonkey.ApplyFunc(json.Unmarshal, func(_ []byte, v interface{}) error {
			p := v.(*map[int]int)
			*p = make(map[int]int)
			(*p)[1] = 2
			(*p)[2] = 4
			return nil
		})
		defer patches.Reset()
		var m map[int]int
		err := json.Unmarshal(nil, &m)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(m[1], convey.ShouldEqual, 2)
		convey.So(m[2], convey.ShouldEqual, 4)
	})

}
