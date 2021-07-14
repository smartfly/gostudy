package closure

import (
	"fmt"
	"testing"
)

// range 与 go 协程引起的错误。loop variable v captured by func literal
func TestClosureAndGoroutine(t *testing.T) {
	done := make(chan bool)
	values := []string{"a", "b", "c"}
	for _, v := range values {
		// v := v		// create a new 'v'
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
	// wait for all goroutines to complete before existing
	for _ = range values {
		<-done
	}
}
