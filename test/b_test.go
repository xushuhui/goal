package test

import (
	"context"
	"testing"
)

func TestB(t *testing.T) {
	ctx := context.TODO()
	f(ctx)
	t.Log(ctx.Value("foo"))

}

func f(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}
