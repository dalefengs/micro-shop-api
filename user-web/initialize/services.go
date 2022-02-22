package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"micro-shop-api/user-web/global"
	"micro-shop-api/user-web/proto"
)

func InitSrvConnect() {

}

func InitUserSrvConnect() {
	cfg := api.DefaultConfig()
	consulCfg := global.Config.Consul
	cfg.Address = fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)
	cClient, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [注册中心失败]")
		panic(err)
	}
	filter, err := cClient.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.Config.Services.UserSrvName))
	if err != nil || len(filter) == 0 {
		zap.S().Errorw("[GetUserList] 查找 [用户服务失败]",
			"msg", err)
		panic(err)
	}
	userSrvHost := ""
	userSrvPort := 0

	for _, item := range filter {
		userSrvHost = item.Address
		userSrvPort = item.Port
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]",
			"msg", err.Error())
		panic(err)
	}
	global.UserSrvConn = proto.NewUserClient(userConn)
}
