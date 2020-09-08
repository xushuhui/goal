package logger

import (
	"github.com/sirupsen/logrus"
	"goal/pkg/utils"

	"io"
	"os"
)

var (
	Formatter = map[string]logrus.Formatter{
		"json": &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		"text": &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		//"test": &inLog.TestFormatter{TimestampFormat: "2006-01-02 15:04:05"},
	}
	logDir string
	//ApiLog    = apiLog()

)

func setFileOut() (f io.Writer, e error) {
	e = os.MkdirAll(logDir, 777)
	if e != nil {
		return
	}
	logFile := logDir + "app.log"
	f, e = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if e != nil {
		return
	}

	return
}

func LogDir(savePath string) string {
	return savePath + "/" + utils.GetCurrentDate() + "/"
}

type Logger struct {
	Logrus *logrus.Logger
}

func NewLogger(format, lvl string, reportCaller bool, savePath string) (l *Logger, e error) {
	lrus := logrus.StandardLogger()

	lrus.SetFormatter(Formatter[format])
	level, e := logrus.ParseLevel(lvl)
	if e != nil {
		return
	}
	lrus.SetLevel(level)
	lrus.SetReportCaller(reportCaller)
	logDir = LogDir(savePath)
	l = &Logger{
		lrus,
	}
	l.setOutput()
	return
}
func (l *Logger) setOutput() {
	writer, e := setFileOut()
	if e != nil {

	}
	l.Logrus.SetOutput(io.MultiWriter(l.Logrus.Out, writer))
}

//func apiLog() *log.Entry {
//	return log.WithField("topic", "api")
//}
//
