package iterator

import (
	"fmt"
	"testing"
)

func TestChannelForIterator(t *testing.T) {
	count := 10
	ch := generate(count)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

// for range 会报错
func TestChannelForRange(t *testing.T) {
	count := 10
	ch := generate(count)
	for v := range ch {
		fmt.Println(v)
	}
}

func generate(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			ch <- i
		}
	}()
	return ch
}
