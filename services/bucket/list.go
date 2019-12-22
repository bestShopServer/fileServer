package bucket

import (
	"context"
	"file_service/middleware/app"
	"file_service/middleware/client"
	"file_service/models"
	"file_service/utils/logging"
	"github.com/gin-gonic/gin"
)

//查看桶列表
func ListBucket(cont *gin.Context) {
	logging.Info("CreateBucket ...")
	var bucket models.BucketList
	this := app.Gin{Cont: cont}

	c, err := client.InitCosBucket("")
	if err != nil {
		logging.Error("ERROR:%v", err.Error())
		this.ResponseJsonError(app.ERROR_DEFAULT)
	}
	list, _, err := c.Service.Get(context.Background())
	if err != nil {
		panic(err)
	}
	logging.Debug("桶列表:%+v", list.Buckets)

	resp := models.BucketListResp{}

	for _, tmp := range list.Buckets {
		bucket.Name = tmp.Name
		bucket.Region = tmp.Region
		bucket.CreationDate = tmp.CreationDate
		resp.Params = append(resp.Params, bucket)
	}
	//resp.Params = list.Buckets
	this.ResponseJsonMessage(resp)
}
