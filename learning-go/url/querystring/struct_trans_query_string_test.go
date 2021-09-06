package querystring

import (
	"fmt"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestStruct2QueryString(t *testing.T) {
	type Options struct {
		Query   string `url:"q"`
		ShowAll bool   `url:"all"`
		Page    int    `url:"page"`
	}
	opt := Options{"foo", true, 2}
	v, _ := query.Values(opt)
	fmt.Println(v.Encode()) // will output: "q=foo&all=true&page=2"
}
