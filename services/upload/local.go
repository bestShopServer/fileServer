package upload

import (
	"file_service/middleware/app"
	"file_service/models"
	"file_service/utils/logging"
	"file_service/utils/setting"
	"fmt"
	"gitee.com/sunki/gutils/encryption"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//存储文件到服务器本地
func LocalCreateFile(cont *gin.Context) {
	logging.Info("CosCreateFile ...")
	//var jparam models.UploadFileReq
	this := app.Gin{Cont: cont}

	file, err := this.Cont.FormFile("file_name")
	if err != nil {
		logging.Error("File Image ERR[%v]", err.Error())
		this.ResponseJsonError(app.ERROR_UPLOAD_FILE)
		return
	}

	//重新文件名
	Suffix := strings.Split(file.Filename, ".")
	fname := fmt.Sprintf("%v/%v%v.%v", setting.ServerSetting.FilePath,
		time.Now().Unix(), encryption.RandomString(10), Suffix[len(Suffix)-1])

	//保存文件
	if err := this.Cont.SaveUploadedFile(file, fname); err != nil {
		logging.Error("File Image ERR[%v]", err.Error())
		this.ResponseJsonError(app.ERROR_UPLOAD_FILE)
		return
	}

	resp := models.UploadFileResp{}
	resp.Params.FileName = fname
	this.ResponseJsonMessage(resp)
}
