package main

import (
	"log"
	"time"
)

func Logging() HandlerFunc {
	return func(c *Context) error {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v ms", c.StatusCode, c.Req.RequestURI, time.Since(t).Milliseconds())
		return nil
	}
}
