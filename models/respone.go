package models

//公共响应信息
type Response struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Notice string `json:"notice"`
}

//桶列表
type BucketList struct {
	//cos.
	Name         string `json:"name"`
	Region       string `json:"region"`
	CreationDate string `json:"creation_date"`
}

//桶列表响应信息
type BucketListResp struct {
	Response
	Data []BucketList `json:"data"`
}

//文件
type UploadFile struct {
	FileName string `json:"url"`
}

//上传文件响应信息
type UploadFileResp struct {
	Response
	Data UploadFile `json:"data"`
}
