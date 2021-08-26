package strategy

import "fmt"

var strategiesFunc map[int]func(c *cache)

func init() {
	strategiesFunc = make(map[int]func(c *cache))
	strategiesFunc[0] = func(c *cache) {
		fmt.Println("Evicting by fifo strategy")
	}
	strategiesFunc[1] = func(c *cache) {
		fmt.Println("Evicting by lru strategy")
	}
	strategiesFunc[2] = func(c *cache) {
		fmt.Println("Evicting by lfu strategy")
	}
}
