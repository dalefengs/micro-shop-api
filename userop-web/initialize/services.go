package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"micro-shop-api/userop-web/global"
	"micro-shop-api/userop-web/proto"
)

func InitSrvConnect() {
	initGoodsSrvConnect()
	initUseropSrvConnect()
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

//  初始化用户操作服务
func initUseropSrvConnect() {
	consulInfo := global.Config.Consul
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.Config.Services.UserOpSrvName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[initUseropPSrvConnect] 连接 [用户操作服务失败]")
		panic(err)
	}
	global.UserFavClient = proto.NewUserFavClient(conn)
	global.AddressClient = proto.NewAddressClient(conn)
	global.MessageClient = proto.NewMessageClient(conn)
}
