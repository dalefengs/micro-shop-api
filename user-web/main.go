package main

import (
	"fmt"
	"go.uber.org/zap"
	"micro-shop-api/user-web/global"
	"micro-shop-api/user-web/initialize"
)

func main() {
	// 初始化 zap logger
	initialize.InitZapLogger()
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化路由
	r := initialize.Routers()
	// 初始化验证器
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	zap.S().Infof("服务启动！ http://localhost:%d", global.Config.Port)
	if err := r.Run(fmt.Sprintf(":%d", global.Config.Port)); err != nil {
		zap.S().Infof("服务启动失败, port:%d,err:%s", global.Config.Port, err.Error())
	}

}
