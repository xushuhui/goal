package core

import (
	"flag"
	"goal/global"
	"goal/internal/model"
	"goal/pkg/logger"
	"goal/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"strings"
	"time"
)

func StartModule() {
	//initRedis()
	var err error
	err = initValidate()
	if err != nil {
		log.Fatalf("initValidate err: %v", err)
	}
	err = initFlag()
	if err != nil {
		log.Fatalf("initFlag err: %v", err)
	}
	err = initSetting()
	if err != nil {
		log.Fatalf("initSetting err: %v", err)
	}
	err = initLogger()
	if err != nil {
		log.Fatalf("initLogger err: %v", err)
	}
	//err = initDBEngine()
	//if err != nil {
	//	log.Fatalf("initDBEngine err: %v", err)
	//}
}

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func initFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "config.yaml")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}
func initSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}
func initLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func initDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
