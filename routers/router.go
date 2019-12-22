package routers

import (
	"file_service/services/bucket"
	"file_service/services/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//获取token
	//r.POST("/auth", services.PostAuth )

	//apiv1 := r.Group("/api/dm")
	//apiv1.Use(middleware.JWT())
	//{
	//	//获取标签列表
	//	apiv1.POST("/list",  services.PostDramaList )
	//}

	//配置腾讯cos
	apiConf := r.Group("/v1/conf/cos")
	{
		//创建桶
		apiConf.POST("/bucket/create", bucket.CreateCosBucket)
		apiConf.POST("/bucket/list", bucket.ListBucket)

	}

	//上传文件
	apiUp := r.Group("/v1/upload")
	{
		//上传文件到本地服务器
		apiUp.POST("/local", upload.LocalCreateFile)
		//上传文件到腾讯云
		apiUp.POST("/cos", upload.CosCreateFile)
	}

	return r
}
