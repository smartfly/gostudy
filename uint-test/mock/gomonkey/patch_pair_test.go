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
 * @time 下午7:39
 * @desc please describe function
**/

func TestPatchPair(t *testing.T) {
	convey.Convey("TestPatchPair", t, func() {
		patchPairs := [][2]interface{}{
			{
				fake.Exec,
				func(_ string, _ ...string) (string, error) {
					return outputExpect, nil
				},
			},
			{
				json.Unmarshal,
				func(_ []byte, v interface{}) error {
					p := v.(*map[int]int)
					*p = make(map[int]int)
					(*p)[1] = 2
					(*p)[2] = 4
					return nil
				},
			},
		}

		patches := gomonkey.NewPatches()
		defer patches.Reset()
		for _, pair := range patchPairs {
			patches.ApplyFunc(pair[0], pair[1])
		}
		output, err := fake.Exec("", "")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, outputExpect)
	})
}
