package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"micro-shop-api/order-web/global"
	"micro-shop-api/order-web/proto"
)

func InitSrvConnect() {
	initGoodsSrvConnect()
	initOrderSrvConnect()
	initInventorySrvConnect()
}

//  初始化商品服务
func initGoodsSrvConnect() {
	consulInfo := global.Config.Consul
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.Config.Services.GoodsSrvName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[initGoodsSrvConnect] 连接 [商品服务失败]")
		panic(err)
	}
	global.GoodsSrvClient = proto.NewGoodsClient(conn)
}

//  初始化订单服务
func initOrderSrvConnect() {
	consulInfo := global.Config.Consul
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.Config.Services.OrderSrvName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[initOrderSrvConnect] 连接 [商品服务失败]")
		panic(err)
	}
	global.OrderSrvClient = proto.NewOrderClient(conn)
}

//  初始化库存服务
func initInventorySrvConnect() {
	consulInfo := global.Config.Consul
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.Config.Services.InventorySrvName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[initInventorySrvConnect] 连接 [库存服务失败]")
		panic(err)
	}
	global.InventorySrvClient = proto.NewInventoryClient(conn)
}
