package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goal/global"
	"goal/pkg/utils"
	"os"
)

var (
	Formatter = map[string]logrus.Formatter{
		"json": &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		"text": &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		//"test": &inLog.TestFormatter{TimestampFormat: "2006-01-02 15:04:05"},
	}

	//ApiLog    = apiLog()
	//MysqlLog  = mysqlLog()
	//AccessLog = accessLog()
)

type F = logrus.Fields

func SetReportCaller(b bool) {
	logrus.SetReportCaller(b)
}

func SetLevel(level string) {
	if level, err := logrus.ParseLevel(level); err != nil {
		panic(fmt.Errorf("Log Level %s", err))
	} else {
		logrus.SetLevel(level)
	}
}
func SetFileOut() {
	e := os.MkdirAll(LogDir(), 777)
	if e != nil {
		return
	}
	//level := logrus.Level.String()
	//logFile := LogDir() + level+ global.AppSetting.LogFileExt
	//f,e := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	//if e != nil {
	//	return
	//}
	//
	//log.SetOutput(f)

	return
}

func NewLogger() (l *logrus.Logger, e error) {
	l = logrus.StandardLogger()

	l.SetFormatter(Formatter[global.LogSetting.Formatter])
	level, e := logrus.ParseLevel(global.LogSetting.Level)
	if e != nil {
		return
	}
	l.SetLevel(level)
	l.SetReportCaller(global.LogSetting.ReportCaller)
	SetFileOut()
	return
}

//func apiLog() *log.Entry {
//	return log.WithField("topic", "api")
//}
//
//func mysqlLog() *log.Entry {
//	return log.WithField("topic", "mysql")
//}
//
//func accessLog() *log.Entry {
//	return log.WithField("topic", "access")
//}
func LogDir() string {
	return global.AppSetting.LogSavePath + "/" + utils.GetCurrentDate() + "/"
}
