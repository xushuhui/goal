package main

import (
	"goal/internal/router"
	"goal/pkg/core"
	"goal/pkg/msg"

	"os"
	"os/signal"
	"syscall"
)

func main() {

	core.StartModule()

	router.HttpServerRun()
	msg.StartMsg()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()

}
