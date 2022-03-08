package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/satori/go.uuid"
	"go.uber.org/zap"

	"micro-shop-api/order-web/global"
	"micro-shop-api/order-web/initialize"
	"micro-shop-api/order-web/utils"
	"micro-shop-api/order-web/utils/register/consul"
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

	// 初始化服务连接
	initialize.InitSrvConnect()

	// 加载证书
	initialize.InitCert()

	// 生产环境动态获取端口
	if !utils.IsDev() {
		port, err := utils.GetFreePort()
		if err == nil {
			global.Config.Port = port
		}
	}

	// 将服务启动放入协程中，当接收到终止信号后，主进程销毁，协称也会随着销毁
	go func() {
		zap.S().Infof("%s 服务启动！ http://%s:%d", global.Config.Name, global.Config.Host, global.Config.Port)
		if err := r.Run(fmt.Sprintf(":%d", global.Config.Port)); err != nil {
			zap.S().Panicf("服务启动失败, port:%d,err:%s", global.Config.Port, err.Error())
		}
	}()

	// 注册服务
	registerClient := consul.NewRegistryClient(global.Config.Consul.Host, global.Config.Consul.Port)
	sc := global.Config
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(sc.Host, sc.Port, sc.Name, serviceId, sc.Consul.Tags)
	if err != nil {
		zap.S().Fatalw("服务注册失败", err.Error())
	}
	zap.S().Infow("服务注册成功")

	// 优雅的退出程序
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Infow("服务注销中...")
	// 等到信号，如果接受到信号执行以下内容
	err = registerClient.DeRegister(serviceId)
	if err != nil {
		zap.S().Fatalw("服务注销失败", err.Error())
	}
	zap.S().Infow("服务注销成功")
}
