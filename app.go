package main

func Default() *Engine {
	engine := New()
	engine.Use(Logging(), Recovery())
	return engine
}
