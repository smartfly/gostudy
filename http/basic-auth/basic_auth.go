package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
 * @author: taohu
 * @date: 19-7-17
 * @time: 下午4:58
 * @desc: http basic Auth
**/

func BasicAuth(w http.ResponseWriter, r *http.Request, user []byte, passwd []byte) bool {
	basicAuthPrefix := "Basic "
	auth := r.Header.Get("Authorization") // 获取 request header
	// 如果是 http basic auth
	if strings.HasPrefix(auth, basicAuthPrefix) {
		// 解码认证信息
		payload, err := base64.StdEncoding.DecodeString(
			auth[len(basicAuthPrefix):],
		)
		if err == nil {
			pair := bytes.SplitN(payload, []byte(":"), 2)
			if len(pair) == 2 && bytes.Equal(pair[0], user) &&
				bytes.Equal(pair[1], passwd) {
				// 执行被装饰的函数
				return true
			}
		}
	}
	return false
}

// 需要被保护的内容
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	user := []byte("foo")
	passwd := []byte("bar")

	// 装饰需要保护的 handler
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if BasicAuth(w, r, user, passwd) {
			HelloServer(w, r)
		}
		// 认证失败，提示 401 Unauthorized
		// Restricted 可以改成其他的值
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		// 401 状态码
		w.WriteHeader(http.StatusUnauthorized)
	})

	log.Println("Listen :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
