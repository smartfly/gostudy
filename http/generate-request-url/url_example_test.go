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

func ExampleValue() {
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
	ExampleValue()
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

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(js), ",", ", ", -1)
}
