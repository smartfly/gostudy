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
 * @time 下午7:42
 * @desc please describe function
**/

func TestPbBuilderFunc(t *testing.T) {
	convey.Convey("TestPbBuilderFunc", t, func() {
		convey.Convey("first dsl", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patchBuilder := gomonkey.NewPatchBuilder(patches)

			patchBuilder.Func(fake.Belong).Stubs().With(gomonkey.Eq("zxl"),
				gomonkey.Any()).Will(gomonkey.Return(true)).Then(gomonkey.Repeat(gomonkey.Return(false), 2)).End()
			flag := fake.Belong("zxl", []string{})
			convey.So(flag, convey.ShouldBeTrue)
		})

		defer func() {
			if p := recover(); p != nil {
				str, ok := p.(string)
				convey.So(ok, convey.ShouldBeTrue)
				convey.So(str, convey.ShouldEqual, "input paras ddd is not matched")
			}
		}()
		fake.Belong("ddd", []string{"abc"})
	})
}
