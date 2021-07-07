package err

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	p := Person{}
	p.ReadName().ReadAge().ReadWeight().Print()
	fmt.Println(p.err) // EOF错误
}
