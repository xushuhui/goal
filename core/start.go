package core

import (
	"flag"
	"goal/setting"
	"log"
	"strings"
	"time"
)

func StartModule() {
	//initRedis()
	InitValidate()
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = initSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "setting/", "config.yaml")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}
func initSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &setting.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &setting.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &setting.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &setting.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &setting.EmailSetting)
	if err != nil {
		return err
	}

	setting.AppSetting.DefaultContextTimeout *= time.Second
	setting.JWTSetting.Expire *= time.Second
	setting.ServerSetting.ReadTimeout *= time.Second
	setting.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		setting.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		setting.ServerSetting.RunMode = runMode
	}

	return nil
}
