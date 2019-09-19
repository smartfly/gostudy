package rate_limiting

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2019/9/17
 * @time 下午3:54
 * @desc 限流demo
**/
// r：每秒发生多少次事件, b：其缓存最大可存多少个事件
var limiter = rate.NewLimiter(2, 5)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func TestRateLimitByTenQPS(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	// warp the server mux with the limit middleware
	http.ListenAndServe(":4000", limit(mux))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func TestLimiter(t *testing.T) {
	l := rate.NewLimiter(2, 5)
	ctx := context.Background()
	start := time.Now()
	// 要处理二十个事件
	for i := 0; i < 20; i++ {
		l.Wait(ctx)
		fmt.Println(l.Limit(), l.Burst())
	}
	fmt.Println(time.Since(start))
}

func TestReserve(t *testing.T) {
	l := rate.NewLimiter(1, 3)
	for {
		r := l.ReserveN(time.Now(), 3)
		fmt.Println(r.Delay())
		fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
	}
}

func TestChannelImplementRateLimit(t *testing.T) {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now().UnixNano())
	}

	fmt.Println("---------------------")

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now().UnixNano())
	}
}
