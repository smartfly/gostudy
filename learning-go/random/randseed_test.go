package random

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2020/1/16
 * @time 上午10:22
 * @desc please describe function
**/

func TestRandSeed(t *testing.T) {
	num := 10
	for j := 0; j < num; j++ {
		res := getRandSeed(num)
		fmt.Println(res)
	}
}

func getRandSeed(num int) int {
	rand.Seed(time.Now().UnixNano())
	var mu sync.Mutex
	mu.Lock()
	v := rand.Intn(num)
	mu.Unlock()
	return v
}
