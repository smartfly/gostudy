package letter_title_possibilities

import (
	"fmt"
	"testing"
)

func Test_numTilePossibilities(t *testing.T) {
	a := numTilePossibilities("AAABBC")
	fmt.Println(a)
}
