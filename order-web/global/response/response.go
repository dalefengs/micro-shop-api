package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"micro-shop-api/order-web/global/status"
)

// Data 统一响应结构体
type Data struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time int64       `json:"time"`
}

// NewData 初始化统一响应结构体
func NewData(code status.Code, msg string, data interface{}) *Data {
	if msg == "" {
		msg = code.Msg
	}
	if data == "" {
		data = []string{}
	}
	return &Data{
		Code: code.Code,
		Data: data,
		Msg:  msg,
		Time: time.Now().Unix(),
	}
}

// region SuccessResponse

// SuccessResponse 成功返回状态码 数据 信息
func SuccessResponse(ctx *gin.Context, httpCode int, code status.Code, msg string, data interface{}) {
	ctx.JSON(httpCode, NewData(code, msg, data))
}

// Success 返回空的成功信息
func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewData(status.OK, "", ""))
}

// SuccessData 成功返回数据 其他默认
func SuccessData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewData(status.OK, "", data))
}

// SuccessMsg 成功返回信息
func SuccessMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, NewData(status.OK, msg, ""))
}

// SuccessMsgData 成功返回信息和数据
func SuccessMsgData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, NewData(status.OK, msg, data))
}

// endregion

// region FailResponse

// FailResponse 失败功返回状态码 数据 信息
func FailResponse(ctx *gin.Context, httpCode int, code status.Code, msg string, data interface{}) {
	ctx.JSON(httpCode, NewData(code, msg, data))
}

// Fail 失败返回状态码 数据 信息
func Fail(ctx *gin.Context, code status.Code) {
	ctx.JSON(http.StatusInternalServerError, NewData(code, "", ""))
}

// FailCodeMsg 失败返回状态码 信息
func FailCodeMsg(ctx *gin.Context, code status.Code, msg string) {
	ctx.JSON(http.StatusInternalServerError, NewData(code, msg, ""))
}

// endregion
