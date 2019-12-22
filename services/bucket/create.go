package bucket

import (
	"context"
	"file_service/middleware/app"
	"file_service/middleware/client"
	"file_service/utils/logging"
	"github.com/gin-gonic/gin"
)

//创建腾讯Cos存储通
func CreateCosBucket(cont *gin.Context) {
	logging.Info("CreateBucket ...")
	this := app.Gin{Cont: cont}

	c, err := client.InitCosBucket("test")
	if err != nil {
		logging.Error("ERROR:%v", err.Error())
		this.ResponseJsonError(app.ERROR_DEFAULT)
	}
	logging.Debug("准备创建桶")
	res, err := c.Bucket.Put(context.Background(), nil)
	if err != nil {
		logging.Error("ERROR:%v", err.Error())
		this.ResponseJsonError(app.ERROR_DEFAULT)
	}
	logging.Debug("res:%+v", res)

	logging.Info("success")
	this.ResponseJsonSuccess()
}
