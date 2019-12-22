package logging

import (
	"file_service/utils/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

//启动服务
func Setup() {
	//初始化日志打印文件
	fmt.Println("init logging to file ...")
	//f, _ := os.Create("log/debug.log")
	f, _ := os.OpenFile("log/debug.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	gin.DefaultWriter = io.MultiWriter(f)
}

func BackLogFile() {
	filename := "log/debug.log"
	mdate := time.Now().Format("2006010215")
	err := os.Rename(filename, filename+"-"+mdate)
	if err != nil {
		Error("Move File ERROR[%v]", err.Error())
	}
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	gin.DefaultWriter = io.MultiWriter(f)
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

func Debug(format string, v ...interface{}) {

	//if global.Cfg.Debug == false { //false不打印日志
	//	return
	//}

	if setting.ServerSetting.RunMode == "false" {
		return
	}

	_, file, line, _ := runtime.Caller(1)
	filename := strings.Split(file, "/")
	tmp := fmt.Sprintf(formatLog(format, v...))
	fmt.Fprintln(gin.DefaultWriter, fmt.Sprintf("%v|%v|%v|%v|%v",
		time.Now().Format("2006-01-02 15:04:05"), "DEBUG", filename[len(filename)-1], line, tmp))
}

func Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	filename := strings.Split(file, "/")
	tmp := fmt.Sprintf(formatLog(format, v...))
	fmt.Fprintln(gin.DefaultWriter, fmt.Sprintf("%v|%v|%v|%v|%v",
		time.Now().Format("2006-01-02 15:04:05"), "INFO", filename[len(filename)-1], line, tmp))
}

func Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	filename := strings.Split(file, "/")
	tmp := fmt.Sprintf(formatLog(format, v...))
	fmt.Fprintln(gin.DefaultWriter, fmt.Sprintf("%v|%v|%v|%v|%v",
		time.Now().Format("2006-01-02 15:04:05"), "ERROR", filename[len(filename)-1], line, tmp))
}
