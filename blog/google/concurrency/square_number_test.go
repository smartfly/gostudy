package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPipeline(t *testing.T) {
	// Set up the pipeline
	c := gen(2, 3)

	// Consume the output
	for n := range sq(c) {
		fmt.Println(n) // 4 then 9
	}
}

func TestPipelineTimes(t *testing.T) {
	// Set up the pipeline and consume the output
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

func TestMerge(t *testing.T) {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func TestAssign(t *testing.T) {
	type Person struct {
		Name string
	}
	a := &Person{}
	r := rand.Int()
	fmt.Println(r)
	if r%2 == 0 {
		a = &Person{
			Name: "test",
		}
	}
	fmt.Println(a)
}

func TestRand(t *testing.T) {
	a := rand.Intn(10)
	fmt.Println(a)
	rand.Seed(100)
	b := rand.Intn(100)
	fmt.Println(b)

}
