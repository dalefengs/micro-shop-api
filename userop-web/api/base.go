package api

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"micro-shop-api/userop-web/global"
	"micro-shop-api/userop-web/global/response"
	status2 "micro-shop-api/userop-web/global/status"
	"micro-shop-api/userop-web/utils"
)

// HandlerGrpcErrorToHttp grpc异常转换为http异常
func HandlerGrpcErrorToHttp(err error, c *gin.Context) {
	if err == nil {
		zap.S().Infof("Grpc Error : %s", err.Error())
		response.Fail(c, status2.Fail)
		return
	}
	s, ok := status.FromError(err)
	if !ok {
		response.FailCodeMsg(c, status2.Fail, s.Message())
		return
	}
	var (
		msg      string
		httpCode int
		code     = status2.ServerError
	)
	switch s.Code() {
	case codes.NotFound:
		httpCode = http.StatusNotFound
		msg = s.Message()
	case codes.Internal:
		httpCode = http.StatusInternalServerError
		msg = "内部错误"
	case codes.InvalidArgument:
		httpCode = http.StatusBadRequest
		code = status2.InvalidParameter
	case codes.Unavailable:
		httpCode = http.StatusInternalServerError
		code = status2.UnavailableServer
	default:
		httpCode = http.StatusInternalServerError
	}
	zap.S().Infof("Grpc Error : %s", err.Error())
	response.FailResponse(c, httpCode, code, msg, "")
}

func HandleValidatorError(err error, c *gin.Context) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		response.FailCodeMsg(c, status2.InvalidParameter, err.Error())
		return
	}
	response.FailCodeMsg(c, status2.InvalidParameter, utils.GetFirstError(errs.Translate(global.Trans)))
	return
}
