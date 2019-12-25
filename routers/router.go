package routers

import (
	"file_service/middleware/cors"
	"file_service/services/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//配置腾讯cos
	//apiConf := r.Group("/v1/conf/cos")
	//{
	//	//创建桶
	//	apiConf.POST("/bucket/create", bucket.CreateCosBucket)
	//	apiConf.POST("/bucket/list", bucket.ListBucket)
	//
	//}

	//上传文件
	apiUp := r.Group("/baseshop/upload")
	{
		//上传文件到本地服务器
		apiUp.POST("/local", upload.LocalCreateFile)
		//上传文件到本地服务器
		apiUp.OPTIONS("/local", upload.LocalCreateFile)
	}
	return r
}
