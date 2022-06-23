package value

import (
	"context"
	"fmt"
	"testing"
)

type contextKey string

const (
	k1 = contextKey("k1")
	k2 = contextKey("k2")
)

func Func1(ctx context.Context) {
	ctx = context.WithValue(ctx, k1, "v1")
	Func2(ctx)
	v1 := ctx.Value(k1)
	v2 := ctx.Value(k2)
	// context只能自上而下携带值
	fmt.Println("Func1", v1, v2) // Func1 v1 <nil>
}

func Func2(ctx context.Context) {
	ctx = context.WithValue(ctx, k2, "v2")
	v1 := ctx.Value(k1)
	v2 := ctx.Value(k2)
	fmt.Println("Func2", v1, v2) // Func2 v1 v2
}

func TestWithValue(t *testing.T) {
	ctx := context.Background()
	Func1(ctx)
}
