package map_api

import (
	"fmt"
	"testing"
)

func TestMapKeyFloat(t *testing.T) {
	m := map[float64]int{}
	m[1.0] = 1
	m[1.1000000000000000001] = 2
	m[1.1000000000000000002] = 3
	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
	}
}
