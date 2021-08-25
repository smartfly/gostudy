package ioc

import (
	"fmt"
	"testing"
)

func TestEmbedding(t *testing.T) {
	label := Label{Widget{10, 10}, "State:"}
	label.X = 11
	label.Y = 12
	fmt.Println(label)
}

func TestLabel_Paint(t *testing.T) {
	l := Label{Text: "testLabel"}
	l.Paint()
}
