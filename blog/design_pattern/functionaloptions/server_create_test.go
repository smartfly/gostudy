package functionaloptions

import (
	"fmt"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s1, _ := NewServer("localhost", 1024)
	s2, _ := NewServer("localhost", 2048, Protocol("trpc"))
	s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
