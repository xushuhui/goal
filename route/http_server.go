package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"goal/config"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(config.ServerSetting.RunMode)
	r := InitRouter()
	port := config.ServerSetting.HttpPort
	HttpSrvHandler = &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  config.ServerSetting.ReadTimeout,
		WriteTimeout: config.ServerSetting.WriteTimeout,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", port)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
