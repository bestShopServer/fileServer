package upload

import (
	"context"
	"file_service/middleware/app"
	"file_service/middleware/client"
	"file_service/models"
	"file_service/utils/logging"
	"fmt"
	"gitee.com/sunki/gutils/encryption"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//存储文件到腾讯云服务器COS
func CosCreateFile(cont *gin.Context) {
	logging.Info("CosCreateFile ...")
	//var jparam models.UploadFileReq
	this := app.Gin{Cont: cont}
	//err := this.Cont.ShouldBindJSON(&jparam)
	//if err != nil {
	//	logging.Error("ERROR:%+v", err.Error())
	//	this.ResponseJsonError(app.ERROR_PARAMS_PARSE)
	//	return
	//}
	//logging.Info("参数:%+v", jparam)

	c, err := client.InitCosBucket("test")
	if err != nil {
		logging.Error("ERROR:%v", err.Error())
		this.ResponseJsonError(app.ERROR_DEFAULT)
		return
	}

	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	//name := "test/objectPut.go"
	// 1.Normal put string content
	//f := strings.NewReader("test")

	fil, err := this.Cont.FormFile("file_name")
	if err != nil {
		logging.Error("ERROR:%v", err.Error())
		this.ResponseJsonError(app.ERROR_DEFAULT)
		return
	}
	logging.Debug("file name:%v", fil.Filename)
	//f, err := os.Open(fil.Filename)
	f, err := fil.Open()
	//os.Create()
	if err != nil {
		logging.Error("ERROR:%v", err.Error())
		this.ResponseJsonError(app.ERROR_DEFAULT)
		return
	}
	defer f.Close()
	logging.Debug("f:%v", f)

	//重命名文件
	path := "test"
	Suffix := strings.Split(fil.Filename, ".")
	fname := fmt.Sprintf("%v/%v%v.%v", path, time.Now().Unix(),
		encryption.RandomString(8), Suffix[len(Suffix)-1])
	logging.Info("文件名称是:%v", fname)

	res, err := c.Object.Put(context.Background(), fname, f, nil)
	if err != nil {
		logging.Error("ERROR:%v", err.Error())
		this.ResponseJsonError(app.ERROR_DEFAULT)
		//panic(err)
		return
	}
	logging.Debug("res:%+v", res.Status)
	//// 2.Put object by local file path
	//filePath := "tmpfile" + time.Now().Format(time.RFC3339)
	//
	//res, err = c.Object.PutFromFile(context.Background(), fil.Filename, "./test", nil)
	//if err != nil {
	//	logging.Error("ERROR:%v", err.Error())
	//	this.ResponseJsonError(app.ERROR_DEFAULT)
	//	//panic(err)
	//	return
	//}
	logging.Debug("res:%+v", res)

	resp := models.UploadFileResp{}
	resp.Data.FileName = fname
	this.ResponseJsonMessage(resp)
}
