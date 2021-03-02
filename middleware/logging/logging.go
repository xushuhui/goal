package logging

import (
	"github.com/gin-gonic/gin"

	"github.com/xushuhui/goal/log"
)

// Option is HTTP logging option.
type Option func(*options)

type options struct {
	logger log.Logger
}

// WithLogger with middleware logger.
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// Server is an server logging middleware.
func Server(opts ...Option) gin.HandlerFunc {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}
	log := log.NewHelper("middleware/logging", options.logger)
	return func(ctx *gin.Context) {

		var (
			path      string
			method    string
			params    string
			component string
		)

		component = "HTTP"
		path = ctx.Request.RequestURI
		method = ctx.Request.Method
		params = ctx.Request.Form.Encode()
		//if err != nil {
		//	log.Errorw(
		//		"kind", "server",
		//		"component", component,
		//		"path", path,
		//		"method", method,
		//		"params", params,
		//		"code", errors.Code(err),
		//		"error", err.Error(),
		//	)
		//	return
		//}
		log.Infow(
			"kind", "server",
			"component", component,
			"path", path,
			"method", method,
			"params", params,
			"code", 0,
		)
		ctx.Next()
		return
	}
}

// Client is an client logging middleware.
func Client(opts ...Option) gin.HandlerFunc {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}
	log := log.NewHelper("middleware/logging", options.logger)
	return func(ctx *gin.Context) {

		var (
			path      string
			method    string
			params    string
			component string
		)
		component = "HTTP"
		path = ctx.Request.RequestURI
		method = ctx.Request.Method
		params = ctx.Request.Form.Encode()

		log.Infow(
			"kind", "client",
			"component", component,
			"path", path,
			"method", method,
			"params", params,
			"code", 0,
		)
		return
	}
}
