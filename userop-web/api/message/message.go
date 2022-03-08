package message

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"micro-shop-api/userop-web/api"
	"micro-shop-api/userop-web/global"
	"micro-shop-api/userop-web/global/response"
	"micro-shop-api/userop-web/proto"
	"micro-shop-api/userop-web/validator/forms"
)

func List(ctx *gin.Context) {
	request := &proto.MessageRequest{}

	//userId, _ := ctx.Get("user_id")
	//claims, _ := ctx.Get("claims")
	//model := claims.(*models.CustomClaims)
	//if model.AuthorityId == 1 {
	//	request.UserId = int32(userId.(uint))
	//}

	rsp, err := global.MessageClient.MessageList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("获取留言失败")
		api.HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["user_id"] = value.UserId
		reMap["type"] = value.MessageType
		reMap["subject"] = value.Subject
		reMap["message"] = value.Message
		reMap["file"] = value.File

		result = append(result, reMap)
	}
	reMap["data"] = result

	response.SuccessData(ctx, reMap)
}

func New(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	messageForm := forms.MessageForm{}
	if err := ctx.ShouldBindJSON(&messageForm); err != nil {
		api.HandleValidatorError(err, ctx)
		return
	}

	rsp, err := global.MessageClient.CreateMessage(context.Background(), &proto.MessageRequest{
		UserId:      int32(userId.(uint)),
		MessageType: messageForm.MessageType,
		Subject:     messageForm.Subject,
		Message:     messageForm.Message,
		File:        messageForm.File,
	})

	if err != nil {
		zap.S().Errorw("添加留言失败")
		api.HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	response.SuccessData(ctx, gin.H{
		"id": rsp.Id,
	})
}
