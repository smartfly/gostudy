package errgroup

import (
	"context"
	"fmt"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(ctx context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}
