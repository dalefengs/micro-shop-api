package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"micro-shop-api/goods-web/global"
	"micro-shop-api/goods-web/proto"
)

func InitSrvConnect() {
	initUserSrvConnect()
}

//  初始化用户服务
func initUserSrvConnect() {
	consulInfo := global.Config.Consul
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.Config.Services.GoodsSrvName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [商品服务失败]")
		panic(err)
	}
	global.GoodsSrvClient = proto.NewGoodsClient(conn)
}
