package app

const (
	SUCCESS            = 0
	ERROR              = 500
	INVALID_PARAMS     = 400
	ERROR_DEFAULT      = 10000
	ERROR_PARAMS_PARSE = 10001
	ERROR_UPLOAD_FILE  = 10002

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

var MsgFlags = map[int]string{
	SUCCESS:            "success",
	ERROR:              "网络错误",
	INVALID_PARAMS:     "无效参数",
	ERROR_DEFAULT:      "操作出错",
	ERROR_PARAMS_PARSE: "参数解析错误",
	ERROR_UPLOAD_FILE:  "上传文件出错",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {

	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
