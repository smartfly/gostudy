package gomonkey

import (
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"go/uint-test/mock/gomonkey/fake"
	"testing"
)

/*
 * @author taohu
 * @date 2019/11/23
 * @time 下午6:48
 * @desc please describe function
**/

func TestApplyFuncVar(t *testing.T) {
	convey.Convey("TestApplyFuncVar", t, func() {
		convey.Convey("for succ", func() {
			str := "hello"
			patches := gomonkey.ApplyFunc(fake.Marshal, func(_ interface{}) ([]byte, error) {
				return []byte(str), nil
			})
			defer patches.Reset()
			bytes, err := fake.Marshal(nil)

			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, str)
		})

		convey.Convey("for fail", func() {
			patches := gomonkey.ApplyFunc(fake.Marshal, func(_ interface{}) ([]byte, error) {
				return nil, fake.ErrActual
			})
			defer patches.Reset()
			_, err := fake.Marshal(nil)
			convey.So(err, convey.ShouldEqual, fake.ErrActual)
		})
	})

}
