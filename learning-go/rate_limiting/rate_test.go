package rate_limiting

import (
	"bytes"
	"fmt"
	"github.com/juju/ratelimit"
	"golang.org/x/time/rate"
	"io"
	"sync"
	"sync/atomic"
	"testing"
	"time"
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

// 使用time.Ticker实现 rate limiting 官方建议10QPS使用time.Ticker, 超过10QPS,
// 使用[golang.org/x/time/rate](https://godoc.org/golang.org/x/time/rate)
func TestCompleteRateLimitByTimeTicker(t *testing.T) {
	rate := time.Second / 10
	throttle := time.Tick(rate)
	for {
		select {
		case <-throttle:
			fmt.Println("test", time.Now())
		}
	}
}

func TestJuJuRateLimit(t *testing.T) {
	src := bytes.NewBuffer(make([]byte, 1024*1024))
	dst := &bytes.Buffer{}
	bucket := ratelimit.NewBucketWithRate(100*1024, 100*1024)
	start := time.Now()
	io.Copy(dst, ratelimit.Reader(src, bucket))
	fmt.Printf("Copied %d bytes in %s\n", dst.Len(), time.Since(start))
}
