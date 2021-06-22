package url

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	fmt.Println("Hello World")
	HttpGet()
}

func HttpGet() {
	client := &http.Client{Timeout: 100 * time.Millisecond} // 30ms时候会超时, 100ms时候不超时
	response, err := client.Get("https://www.baidu.com/")
	if err != nil {
		log.Println(err)
		return
	}
	status := response.StatusCode //获取返回状态码，正常是200
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
	log.Println(status)
}
