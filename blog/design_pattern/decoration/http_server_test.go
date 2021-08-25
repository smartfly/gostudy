package decoration

import (
	"log"
	"net/http"
	"testing"
)

func TestWithServerHeader(t *testing.T) {
	http.HandleFunc("/v1/hello", WithServerHeader(hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func TestCollection(t *testing.T) {
	http.HandleFunc("/v1/hello", WithServerHeader(WithAuthCookie(hello)))
	http.HandleFunc("/v2/hello", WithServerHeader(WithBasicAuth(hello)))
	http.HandleFunc("/v3/hello", WithServerHeader(WithBasicAuth(WithDebugLog(hello))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func TestHandle(t *testing.T) {
	http.HandleFunc("/v4/hello", Handler(hello, WithServerHeader, WithAuthCookie, WithDebugLog))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
