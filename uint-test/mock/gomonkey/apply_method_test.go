package gomonkey

import (
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"go/uint-test/mock/gomonkey/fake"
	"reflect"
	"testing"
)

/*
 * @author taohu
 * @date 2019/11/23
 * @time 下午7:24
 * @desc method
**/

func TestApplyMethod(t *testing.T) {
	slice := fake.NewSlice()
	var s *fake.Slice
	convey.Convey("TestApplyMethod", t, func() {
		convey.Convey("for succ", func() {
			err := slice.Add(1)
			convey.So(err, convey.ShouldEqual, nil)
			err = slice.Remove(1)
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(len(slice), convey.ShouldEqual, 0)
		})

		convey.Convey("for already exist", func() {
			err := slice.Add(2)
			convey.So(err, convey.ShouldEqual, nil)
			patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
				return fake.ErrElemExsit
			})
			defer patches.Reset()

			err = slice.Add(1)
			convey.So(err, convey.ShouldEqual, fake.ErrElemExsit)
			err = slice.Remove(2)
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(len(slice), convey.ShouldEqual, 0)
		})

	})
}
