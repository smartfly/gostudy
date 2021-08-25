package decoration

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	decorator(Hello)("Hello, World!")
	hello := decorator(Hello)
	hello("Hello")
}
