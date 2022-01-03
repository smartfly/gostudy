package rand

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMathRand(t *testing.T) {

}

func TestCryptoRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		ret := Key()
		fmt.Println(ret)
	}
}

func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf)
}
