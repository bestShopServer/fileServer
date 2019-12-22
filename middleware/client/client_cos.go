package client

import (
	"file_service/utils/logging"
	"file_service/utils/setting"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"net/http"
	"net/url"
)

func InitCosBucket(bucketName string) (client *cos.Client, err error) {
	logging.Info("bucketName:%v", bucketName)

	if len(bucketName) == 0 {
		client = cos.NewClient(nil, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  setting.TencentSetting.SecretID,
				SecretKey: setting.TencentSetting.SecretKey,
				Transport: &debug.DebugRequestTransport{
					RequestHeader:  true,
					RequestBody:    true,
					ResponseHeader: true,
					ResponseBody:   true,
					Writer:         nil,
					Transport:      nil,
				},
			},
		})
	} else {
		//region := "beijing"
		//请求域名:<名称>-1257865098.cos.ap-beijing.myqcloud.com
		urlAddr := fmt.Sprintf("https://%v-%v.cos.%v.myqcloud.com",
			bucketName, setting.TencentSetting.AppID, setting.TencentSetting.Region)
		logging.Info("桶域名名称:%v", urlAddr)
		u, err := url.Parse(urlAddr)
		if err != nil {
			logging.Error("ERROR:%v", err.Error())
			return client, err
		}
		b := &cos.BaseURL{BucketURL: u}
		client = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  setting.TencentSetting.SecretID,
				SecretKey: setting.TencentSetting.SecretKey,
				Transport: &debug.DebugRequestTransport{
					RequestHeader:  true,
					RequestBody:    true,
					ResponseHeader: true,
					ResponseBody:   true,
					Writer:         nil,
					Transport:      nil,
				},
			},
		})
	}

	return client, err
}
