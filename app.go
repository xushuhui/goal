package main

import (
	. "github.com/xushuhui/goal/internal"
	"github.com/xushuhui/goal/middleware"
)

func Default() *Engine {
	engine := New()
	engine.Use(middleware.Logging(), middleware.Recovery())
	return engine
}
