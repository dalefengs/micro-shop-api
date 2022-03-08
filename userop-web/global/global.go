package global

import (
	ut "github.com/go-playground/universal-translator"
	"micro-shop-api/userop-web/config"
	"micro-shop-api/userop-web/proto"
)

// 全局变量

var (
	Trans          ut.Translator
	Config         *config.Config    // 配置文件全局从Nacos中获取
	NacosConfig    *config.Nacos     // Nacos 配置
	GoodsSrvClient proto.GoodsClient // 商品服务连接对象
	MessageClient  proto.MessageClient
	AddressClient  proto.AddressClient
	UserFavClient  proto.UserFavClient
)
