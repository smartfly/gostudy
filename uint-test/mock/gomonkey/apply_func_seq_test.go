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
 * @time 下午6:13
 * @desc please describe function
**/

func TestApplyFuncSeq(t *testing.T) {
	convey.Convey("TestApplyFuncSeq", t, func() {
		info1 := "hello cpp"
		info2 := "hello golang"
		info3 := "hello gomonkey"
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{info1, nil}},
			{Values: gomonkey.Params{info2, nil}},
			{Values: gomonkey.Params{info3, nil}},
		}
		patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
		defer patches.Reset()

		output, err := fake.ReadLeaf("")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, info1)

		output, err = fake.ReadLeaf("")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, info2)

		output, err = fake.ReadLeaf("")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, info3)
	})

	convey.Convey("retry succ util the third times", func() {
		info1 := "hello cpp"
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{"", fake.ErrActual}, Times: 2},
			{Values: gomonkey.Params{info1, nil}},
		}

		patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
		defer patches.Reset()
		output, err := fake.ReadLeaf("")
		convey.So(err, convey.ShouldEqual, fake.ErrActual)
		output, err = fake.ReadLeaf("")
		convey.So(err, convey.ShouldEqual, fake.ErrActual)
		output, err = fake.ReadLeaf("")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, info1)
	})
}
