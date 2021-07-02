package middleware

import (
	"github.com/xushuhui/goal/internal"
	"log"
	"time"
)

func Logging() internal.HandlerFunc {
	return func(c *internal.Context) error {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v ms", c.StatusCode, c.Req.RequestURI, time.Since(t).Milliseconds())
		return nil
	}
}
