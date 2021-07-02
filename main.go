package main

import (
	. "github.com/xushuhui/goal/internal"
	"github.com/xushuhui/goal/middleware"
	"net/http"
)

func main() {
	r := Default()

	r.GET("/", func(c *Context) error {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
		return nil
	})
	r.GET("/hello", func(c *Context) error {

		panic("is error")
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		return nil
	})

	r.GET("/hello/:name", func(c *Context) error {

		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		return nil
	})
	v1 := r.Group("/v1")
	v1.Use(middleware.Logger())
	{
		v1.GET("/", func(c *Context) error {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
			return nil
		})

		v1.GET("/hello", func(c *Context) error {
			c.JSON(http.StatusOK, H{"x": "x"})

			return nil
		})
		v1.GET("/hello/:name", func(c *Context) error {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
			return nil
		})
	}
	r.Run(":8000")
}
