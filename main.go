package main

import (
	"fmt"
	"go-web/core"
	"go-web/routes"
)

func main() {
	core.InitEnv()
	fmt.Println(core.Md5Encode("123456"))
	_ = routes.Router().Run(":8000")

}
