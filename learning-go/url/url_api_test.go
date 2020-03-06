package url

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

/*
 * @author taohu
 * @date 2020/2/20
 * @time 下午4:56
 * @desc please describe function
**/

func TestQueryEscape(t *testing.T) {
	webpage := "http://mywebpage.com/thumbify"
	image := "http://images.com/cat.png"
	fmt.Println(webpage + "?image=" + url.QueryEscape(image))
}

func TestQueryUnEscape(t *testing.T) {
	urlValue := "http://mywebpage.com/thumbify?image=http%3A%2F%2Fimages.com%2Fcat.png"
	rawurl, err := url.QueryUnescape(urlValue)
	if err != nil {
		fmt.Println(err)
	}
	u, _ := url.Parse(rawurl)
	// 直接访问 scheme
	scheme := u.Scheme
	// User 包含了所有的认证信息，这里调用username和password来获取独立值
	username := u.User.Username()
	password, _ := u.User.Password()
	fmt.Printf("scheme: [%s], username: [%s], password:[%s]\n", scheme, username, password)
	// Host 同时包括主机名和端口信息，如果端口存在的话，使用 strings.Split() 从 Host 中手动提取端口。
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	if len(h) > 1 {
		fmt.Printf("ip:[%s], port:[%s]\n", h[0], h[1])
	}
	//这里我们提出路径和查询片段信息。
	fmt.Printf("path:[%s], fragment:[%s]\n", u.Path, u.Fragment)
	//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。你也可以将查询参数解析为一个map。
	// 已解析的查询参数 map_api 以查询字符串为键，对应值字符串切片为值，所以如何只想得到一个键对应的第一个值，将索引位置设置为 [0] 就行了。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Printf("iamge: [%s]\n", m["image"][0])
}
