package pipeline

import (
	"fmt"
	"testing"
)

func Test_makeRange(t *testing.T) {
	nums := makeRange(1, 10000)
	in := echo(nums)
	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum(prime(in))
	}
	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
}
