package main

import (
	"github.com/xushuhui/goal/internal"
	"log"
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*internal.Node
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*internal.Node),
		handlers: make(map[string]HandlerFunc)}
}
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &internal.Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) getRoute(method string, path string) (*internal.Node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.Search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.Pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
func (r *Router) Handle(c *Context) {

	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.Pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) error {
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound), c.Path)
			return nil
		})
	}
	c.Next()
}
