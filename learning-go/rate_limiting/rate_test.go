package rate_limiting

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync"
	"sync/atomic"
	"testing"
)

/*
 * @author taohu
 * @date 2019/9/25
 * @time 上午10:25
 * @desc rate测试
**/

func TestSimultaneousRequests(t *testing.T) {
	const (
		limit      = 1
		burst      = 5
		numRequest = 15
	)

	var (
		wg    sync.WaitGroup
		numOk = uint32(0)
	)

	lim := rate.NewLimiter(limit, burst)

	f := func() {
		defer wg.Done()
		if ok := lim.Allow(); ok {
			atomic.AddUint32(&numOk, 1)
		}
	}

	wg.Add(numRequest)
	for i := 0; i < numRequest; i++ {
		go f()
	}
	wg.Wait()
	if numOk != burst {
		fmt.Printf("numOK = %d, want %d\n", numOk, burst)
	}
}
