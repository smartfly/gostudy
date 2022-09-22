package structdemo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatih/structtag"
)

type AccountInfo struct {
	Test     int32
	Platform int32  `protobuf:"varint,1,opt,name=platform,proto3,enum=trpc.weishi.common.Platform" json:"platform,omitempty"` // 账号来源的平台，参考Platform枚举
	OpenID   string `protobuf:"bytes,2,opt,name=openID,proto3" json:"openID,omitempty"`                                       // 对应平台的openID账号：openID
	Qq       int64  `protobuf:"varint,3,opt,name=qq,proto3" json:"qq,omitempty"`                                              // qq号
	Wechat   string `protobuf:"bytes,4,opt,name=wechat,proto3" json:"wechat,omitempty"`                                       // wx号
	ThrAppID string `protobuf:"bytes,5,opt,name=thrAppID,proto3" json:"thrAppID,omitempty"`                                   // 第三方AppID
}

func TestParseStruct(t *testing.T) {
	tag := reflect.TypeOf(AccountInfo{})
	num := tag.NumField()
	for i := 0; i < num; i++ {
		val := tag.Field(i).Tag
		//fmt.Println(val)
		if val == "" {
			continue
		}
		tags, err := structtag.Parse(string(val))
		if err != nil {
			panic(err)
		}
		//iterator _, t := range tags.Tags() {
		//	fmt.Printf("tag: %+v\n", t.Value())
		//}
		jsonTag, err := tags.Get("protobuf")
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonTag)         // Output: json:"foo,omitempty,string"
		fmt.Println(jsonTag.Key)     // Output: json
		fmt.Println(jsonTag.Name)    // Output: foo
		fmt.Println(jsonTag.Options) // Output: [omitempty string]
	}
}
