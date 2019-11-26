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
 * @time 下午7:14
 * @desc 接口
**/

func TestApplyInterfaceReused(t *testing.T) {
	e := &fake.Etcd{}

	convey.Convey("TestApplyInterfaceReused", t, func() {
		patches := gomonkey.ApplyFunc(fake.NewDb, func(_ string) fake.Db {
			return e
		})
		defer patches.Reset()

		db := fake.NewDb("mysql")

		convey.Convey("TestApplyInterface", func() {
			info := "hello interface"
			patches.ApplyMethod(reflect.TypeOf(e), "Retrieve", func(_ *fake.Etcd, _ string) (string, error) {
				return info, nil
			})
			output, err := db.Retrieve("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info)
		})

	})

}
