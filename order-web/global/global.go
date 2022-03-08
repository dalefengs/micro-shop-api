package global

import (
	ut "github.com/go-playground/universal-translator"
	"micro-shop-api/order-web/config"
	"micro-shop-api/order-web/proto"
)

// 全局变量

var (
	Trans              ut.Translator
	Config             *config.Config        // 配置文件全局从Nacos中获取
	NacosConfig        *config.Nacos         // Nacos 配置
	GoodsSrvClient     proto.GoodsClient     // 商品服务连接对象
	OrderSrvClient     proto.OrderClient     // 订单服务连接对象
	InventorySrvClient proto.InventoryClient // 库存服务连接对象
)
