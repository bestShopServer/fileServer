package main

import (
	"file_service/routers"
	"file_service/utils/logging"
	"file_service/utils/setting"
	"fmt"
	"github.com/robfig/cron"
	"net/http"
)

//准备初始化数据
func init() {
	setting.Setup()
	logging.Setup()

	//添加定时任务
	c := cron.New()
	//second min hour day month week
	err := c.AddFunc("0 0 0 * * *", logging.BackLogFile)
	if err != nil {
		logging.Error("定时任务执行失败[%v]", err.Error())
	}

	c.Start()
}

//主函数
func main() {

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf("%s:%d",
		setting.ServerSetting.HttpAddr, setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	fmt.Printf("[info] start http server listening %s\n", endPoint)
	logging.Info("服务正常启动 listening[%v]\n", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("启动服务ERROR[%v]", err.Error())
	}
	logging.Info("服务正常启动[%v]", endPoint)
}
