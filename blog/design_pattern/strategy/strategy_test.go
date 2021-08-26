package strategy

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestEvict(t *testing.T) {
	strategies := make(map[int]evictionAlgo)
	strategies[0] = &fifo{}
	strategies[1] = &lru{}
	strategies[2] = &lfu{}
	val, _ := rand.Int(rand.Reader, big.NewInt(100))
	condition := int(val.Int64()) % 3
	strategies[condition].evict(&cache{})
}

func TestFunc(t *testing.T) {
	val, _ := rand.Int(rand.Reader, big.NewInt(100))
	condition := int(val.Int64()) % 3
	f := strategiesFunc[condition]
	f(&cache{})
}
