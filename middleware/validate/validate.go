package validate

import (
	"context"
	"github.com/xushuhui/goal/errors"
	"github.com/xushuhui/goal/middleware"
)

type validator interface {
	Validate() error
}

// Validator is a validator middleware.
func Validator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if v, ok := req.(validator); ok {
				if err := v.Validate(); err != nil {
					return nil, errors.InvalidArgument("Validator", err.Error())
				}
			}
			return handler(ctx, req)
		}
	}
}
