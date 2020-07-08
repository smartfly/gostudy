package rate

import (
	"fmt"
	"go.uber.org/ratelimit"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2020/7/7
 * @time 下午9:12
 * @desc please describe function
**/

func TestUberRateLimit(t *testing.T) {
	rl := ratelimit.New(100)

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
