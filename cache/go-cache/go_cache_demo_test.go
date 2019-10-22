package go_cache

import (
	"github.com/patrickmn/go-cache"
	"log"
	"testing"
	"time"
)

/*
 * @author: taohu
 * @date: 19-7-3
 * @time: 上午10:58
 * @desc: go-cache测试
**/

func TestGoCache(t *testing.T) {
	c := cache.New(30*time.Second, 10*time.Second)
	c.Set("Title", "Spring Festival", cache.DefaultExpiration)

	value, found := c.Get("Title")
	if found {
		log.Println("found:", value)
	} else {
		log.Println("not found")
	}

	time.Sleep(60 * time.Second)
	log.Println("sleep 60s")
	value, found = c.Get("Title")
	if found {
		log.Println("found:", value)
	} else {
		log.Println("not found")
	}
}
