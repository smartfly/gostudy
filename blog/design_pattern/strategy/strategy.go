package strategy

import "fmt"

type cache struct{}

type evictionAlgo interface {
	evict(c *cache)
}

// FIFO 先进先去
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

// LRU最少最近使用
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

// LFU最少使用
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}
