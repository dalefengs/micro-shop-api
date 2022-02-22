package api

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"micro-shop-api/user-web/extends/aliyun/dysmsapi"
	"micro-shop-api/user-web/extends/jwtauth"
	"micro-shop-api/user-web/global"
	"micro-shop-api/user-web/global/response"
	status2 "micro-shop-api/user-web/global/status"
	"micro-shop-api/user-web/proto"
	"micro-shop-api/user-web/utils"
	"micro-shop-api/user-web/validator/forms"
)

// 生成UserResponse 实例
func returnUserResponse(id uint32, mobile, nickname, orangeKey string, birthday uint64, gender int32) *response.UserResponse {
	return &response.UserResponse{
		Id:        id,
		Mobile:    mobile,
		NickName:  nickname,
		OrangeKey: orangeKey,
		Birthday:  response.JsonTime(time.Unix(int64(birthday), 0)),
		Gender:    gender,
	}
}

// HandlerGrpcErrorToHttp grpc异常转换为http异常
func HandlerGrpcErrorToHttp(err error, c *gin.Context) {
	if err == nil {
		return
	}
	s, ok := status.FromError(err)
	if !ok {
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
	response.FailResponse(c, httpCode, code, msg, "")
}

// GetUserList 获取用户列表
func GetUserList(ctx *gin.Context) {
	pn := ctx.DefaultQuery("pn", "1")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	pageInfo := proto.PageInfo{
		Page:  uint32(pnInt),
		Limit: uint32(pSizeInt),
	}
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &pageInfo)
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表] 失败",
			"msg", err.Error())
		HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]response.UserResponse, 0)
	for _, user := range rsp.Data {
		u := response.UserResponse{
			Id:        user.Id,
			Mobile:    user.Mobile,
			NickName:  user.Nickname,
			OrangeKey: user.OrangeKey,
			Birthday:  response.JsonTime(time.Unix(int64(user.Birthday), 0)),
			Gender:    user.Gender,
		}
		result = append(result, u)
	}
	response.SuccessData(ctx, result)
}

// PasswordLogin 密码登录
func PasswordLogin(c *gin.Context) {
	pwdLoginForm := forms.PasswordLoginForm{}
	// 绑定验证 Json
	if err := c.ShouldBindJSON(&pwdLoginForm); err != nil {
		if err == io.EOF {
			response.Fail(c, status2.InvalidParameter)
			return
		}
		errs := err.(validator.ValidationErrors)
		response.FailCodeMsg(c, status2.InvalidParameter, utils.GetFirstError(errs.Translate(global.Trans)))
		return
	}

	uInfo, err := global.UserSrvClient.GetUserInfoByMobile(c, &proto.MobileRequest{Mobile: pwdLoginForm.Mobile})
	if err != nil {
		// 没有查找到用户信息
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				response.FailCodeMsg(c, status2.NotFoundUser, s.Message()) // 用户不存在
			default:
				response.FailCodeMsg(c, status2.Fail, s.Message()) // 操作失败
			}
			return
		}
		response.Fail(c, status2.Fail) // 操作失败
		return
	}
	// 验证密码
	if err := CheckPassword(c, uInfo.Mobile, pwdLoginForm.Password); err != nil {
		zap.S().Infof("验证密码异常%v", err)
		// 没有查找到信息
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				response.FailCodeMsg(c, status2.NotFoundUser, s.Message()) // 用户不存在
			case codes.InvalidArgument:
				response.FailCodeMsg(c, status2.PasswordError, s.Message())
			default:
				response.FailCodeMsg(c, status2.Fail, s.Message()) // 操作失败
			}
			return
		}
		response.Fail(c, status2.Fail) // 操作失败
		return
	}
	j := jwtauth.NewJWT()
	claims := jwtauth.NewCustomClaimsDefault(uint(uInfo.Id), uInfo.Mobile, uInfo.Nickname)
	token, err := j.CreateToken(*claims)
	if err != nil {
		response.FailCodeMsg(c, status2.Fail, err.Error())
		return
	}
	rsp := response.UserTokenResponse{
		Id:       claims.Id,
		Mobile:   claims.Mobile,
		NickName: claims.Nickname,
		Token:    token,
	}

	response.SuccessMsgData(c, "登录成功", rsp)
}

// CheckPassword 验证密码
func CheckPassword(ctx *gin.Context, mobile string, password string) (err error) {
	rsp, err := global.UserSrvClient.CheckPassword(ctx, &proto.CheckPasswordInfo{Mobile: mobile, Password: password})
	if err == nil && rsp.Success {
		return nil
	}
	return err
}

// RegisterUser 用户注册
func RegisterUser(c *gin.Context) {
	// 表单验证
	createUserInfo := forms.RegisterUserForm{}
	if err := c.ShouldBindJSON(&createUserInfo); err != nil {
		if err == io.EOF {
			response.Fail(c, status2.InvalidParameter)
			return
		}
		e, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.S().Infof("error : %s", e.Error())
			response.FailCodeMsg(c, status2.InvalidParameter, utils.GetFirstError(e.Translate(global.Trans)))
		}
		response.FailCodeMsg(c, status2.InvalidParameter, utils.GetFirstError(e.Translate(global.Trans)))
		return
	}
	if !dysmsapi.VerifySmsCode(1, createUserInfo.Mobile, createUserInfo.SmsCode) {
		response.Fail(c, status2.CodeIncorrect)
		return
	}

	cuInfo := proto.CreateUserInfo{
		Mobile:   createUserInfo.Mobile,
		Nickname: createUserInfo.Nickname,
		Password: createUserInfo.Password,
	}

	uRsp, err := global.UserSrvClient.CreateUser(context.Background(), &cuInfo)
	if err != nil {
		if gerr, ok := status.FromError(err); ok {
			switch gerr.Code() {
			case codes.AlreadyExists:
				response.FailCodeMsg(c, status2.AlreadyExists, gerr.Message()) // 用户已存在
			case codes.InvalidArgument:
				response.FailCodeMsg(c, status2.Fail, gerr.Message())
			default:
				response.FailCodeMsg(c, status2.Fail, gerr.Message())
			}
			return
		}
		response.Fail(c, status2.Fail) // 操作失败
		return
	}
	rsp := returnUserResponse(uRsp.Id, uRsp.Mobile, uRsp.Nickname, uRsp.OrangeKey, uRsp.Birthday, uRsp.Gender)
	response.SuccessMsgData(c, "注册成功", rsp)

}
