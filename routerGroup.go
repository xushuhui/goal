package main

import "net/http"

type RouterGroup struct {
	Prefix      string
	Middlewares []HandlerFunc
	parent      *RouterGroup
	Engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.Engine
	newGroup := &RouterGroup{
		Prefix: group.Prefix + prefix,
		parent: group,
		Engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.Prefix + comp
	group.Engine.router.AddRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPost, pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.Middlewares = append(group.Middlewares, middlewares...)
}
