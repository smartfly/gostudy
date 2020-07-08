package generate_request_url

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"testing"
)

/*
 * @author: taohu
 * @date: 19-7-11
 * @time: 上午9:36
 * @desc: please describe function
**/

func ExampleURLEncodeValue() {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	//v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Encode())
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
}

func TestExampleValue(t *testing.T) {
	ExampleURLEncodeValue()
}

func TestUrlDecode(t *testing.T) {
	result := "result := app_id=2016073100129537&" +
		"biz_content=%7B%22body%22%3A%22body%22%2C%22subject%22%3A%22%E5%95%86%E5%93%81" +
		"%E6%A0%87%E9%A2%98%22%2C%22out_trade_no%22%3A%2201010101%22%2C%22timeout_express%22%3A%22%22%2C%22" +
		"total_amount%22%3A%22100.00%22%2C%22seller_id%22%3A%22%22%2C%22product_code%22%3A%22p_1010101%22%7D&" +
		"charset=utf-8&format=JSON&method=alipay.trade.app.pay&notify_url=http%3A%2F%2F203.86.24.181%3A3000%2Falipay&" +
		"sign=i0lEkNZjw%2FaBafm1HGH7wQ2KW5RWE7UUAmymwOZvbmHxofpIcGZnBBkOd%2BZolAb2XjjcPxJQ%2FrRx0Ts%2Bs58%2FpCsHR%2" +
		"BOZ8jsUFX6OQweqeq4uNnugBc0bBkZBosQxLazjsn9JVnqe0IW%2BxtXzju7g7kCcXZG9l8IEgG%2BanSGpYu0%2B6Dr30f7UoZPcT3VCw%" +
		"2FBNCNyh8BS0G7j%2FcCariZ3bZVLu2kUu4aHqhfzG58b2LnhAUkuZdxr6vsNZmbuC2G74yH1W9tOzBQPgKhDFagi0L3WAdL65r" +
		"1wiw7RIpHVp8n41hY5YkGr%2FVbtMQoHfEtOD2kzu9%2FL8Qb2ICoBq4C9lVw%3D%3D&sign_type=RSA2&" +
		"timestamp=2017-11-10+17%3A54%3A34&version=1.0 <nil>"
	u, err := url.ParseQuery(result)
	fmt.Println(u, err)
}

func ExampleURL() {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)
}

func TestExampleUrl(t *testing.T) {
	ExampleURL()

}

func ExampleURL_roundtrip() {
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())
}

func TestExampleURL_roundtrip(t *testing.T) {
	ExampleURL_roundtrip()
}

func ExampleURL_ResolveReference() {
	u, err := url.Parse("../../..//search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	base, err := url.Parse("http://example.com/directory/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base.ResolveReference(u))
	// Output:
	// http://example.com/search?q=dotnet
}

func TestExampleURL_ResolveReference(t *testing.T) {
	ExampleURL_ResolveReference()
}

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(js), ",", ", ", -1)
}

func ExampleParseQuery() {
	m, err := url.ParseQuery(`x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toJSON(m))
	// Output:
	// {"x":["1"], "y":["2", "3"], "z":[""]}
}

func TestExampleParseQuery(t *testing.T) {
	ExampleParseQuery()
}

func TestExample(t *testing.T) {
	s := "https://isee.weishi.qq.com/ws/app-pages/heroSDK_dapian/index.html?" +
		"offlineMode=1&showloading=0&navstyle=2&_wv=4096#/?upload_from=10052&autoAuth=1&sourceId=8"

	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	n, _ := url.ParseQuery(u.Fragment)
	fmt.Println(n)
	fmt.Println(u.Query())
}

func Test(t *testing.T) {
	fmt.Println(0x1F)
}

func TestA(t *testing.T) {

}
