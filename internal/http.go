package internal

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context) error

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) middlewares(path string) (middlewares []HandlerFunc) {
	for _, group := range engine.groups {
		if strings.HasPrefix(path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	return
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	c := newContext(w, req)
	c.handlers = engine.middlewares(req.URL.Path)
	engine.router.handle(c)
}
