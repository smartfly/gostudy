package errgroup

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestMD5All(t *testing.T) {
	m, err := MD5All(context.Background(), ".")
	if err != nil {
		log.Fatal(err)
	}
	for k, sum := range m {
		fmt.Printf("%s:\t%x\n", k, sum)
	}
}
