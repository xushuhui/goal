package lib

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	//"sao.cn/configs"
)

var (
	logPath string
	lastDay int
)

func init() {
	logPath = "./log"
	_, err := os.Stat(logPath)
	if err != nil {
		_ = os.Mkdir(logPath, 0755)
	}
}

func defaultWriter() io.Writer {
	writerCheck()
	return gin.DefaultWriter
}

func defaultErrorWriter() io.Writer {
	writerCheck()
	return gin.DefaultErrorWriter
}

func writerCheck() {
	nowDay := time.Now().Day()
	if nowDay != lastDay {
		var file *os.File
		filename := time.Now().Format("2006-01-02")
		logFile := fmt.Sprintf("%s/%s-%s.log", logPath, "gosapi", filename)

		file, _ = os.Create(logFile)
		if file != nil {
			gin.DefaultWriter = file
			gin.DefaultErrorWriter = file
		}
	}

	lastDay = nowDay
}
