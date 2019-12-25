package app

import (
	"encoding/json"
	"encoding/xml"
	"file_service/utils/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	Cont *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (this *Gin) Response(httpCode, errCode int, data interface{}) {
	logging.Info("Response Code: %v  Msg: %v", errCode, GetMsg(errCode))
	jtmp, err := json.Marshal(data)
	if err == nil {
		logging.Debug("Response Data: %v", string(jtmp))
	}

	this.Cont.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (this *Gin) ResponseJsonSuccess() {
	logging.Info("%v", http.StatusOK)
	this.Cont.JSON(http.StatusOK, Response{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
	})
}

// Response setting gin.JSON
func (this *Gin) ResponseJsonError(code int) {
	logging.Error("code:%v", code)
	this.Cont.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  GetMsg(code),
	})
	this.Cont.Abort()
}

// Response setting gin.JSON
func (this *Gin) ResponseJsonMessage(obj interface{}) {
	jtmp, err := json.Marshal(obj)
	if err == nil {
		logging.Debug("Response Data: %v", string(jtmp))
	}
	this.Cont.JSON(http.StatusOK, obj)
	println(1)
}

func (this *Gin) ResponseJsonData(obj interface{}) {

	this.Cont.JSON(http.StatusOK, Response{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: obj,
	})
}

// Response setting gin.XML
func (this *Gin) ResponseXmlSuccess() {
	this.Cont.JSON(http.StatusOK, Response{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
	})
	logging.Info("%v", http.StatusOK)
	return
}

// Response setting gin.XML
func (this *Gin) ResponseXmlMessage(obj interface{}) {
	by, _ := xml.Marshal(obj)
	logging.Info("响应信息[%v]", string(by))
	this.Cont.XML(http.StatusOK, obj)
	return
}
