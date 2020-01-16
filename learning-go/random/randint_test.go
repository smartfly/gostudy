package random

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

/*
 * @author taohu
 * @date 2020/1/16
 * @time 上午10:19
 * @desc please describe function
**/

func getRand(num int) int {
	var mu sync.Mutex
	mu.Lock()
	// 伪随机
	v := rand.Intn(num)
	mu.Unlock()
	return v
}

func TestRandInt(t *testing.T) {
	num := 10
	for j := 0; j < num; j++ {
		res := getRand(num)
		fmt.Println(res)
	}
}
