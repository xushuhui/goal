package main

import (
	"go-web/core"
	"go-web/routes"
)

func main() {
	core.InitEnv()

	_ = routes.Router().Run(":8000")

}
