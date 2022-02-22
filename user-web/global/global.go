package global

import (
	ut "github.com/go-playground/universal-translator"
	"micro-shop-api/user-web/config"
	"micro-shop-api/user-web/proto"
)

// 全局变量

var (
	Trans       ut.Translator
	Config      *config.Config   // 配置文件全局
	UserSrvConn proto.UserClient // 用户服务连接对象
)
