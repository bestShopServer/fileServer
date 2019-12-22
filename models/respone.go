package models

//公共响应信息
type Response struct {
	Code   int    `json:"code"`
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
	Params []BucketList `json:"params"`
}

//文件
type UploadFile struct {
	FileName string `json:"file_name"`
}

//上传文件响应信息
type UploadFileResp struct {
	Response
	Params UploadFile `json:"params"`
}
