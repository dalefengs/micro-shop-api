package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"micro-shop-api/userop-web/global"
	"micro-shop-api/userop-web/utils"
)

// 项目名称
var projectName = "userop-web"

func InitConfig() {
	filepath := "nacos%s"
	if utils.IsDev() { // 开发环境
		filepath = fmt.Sprintf(filepath, "-dev")
	} else { // 生产环境
		filepath = fmt.Sprintf(filepath, "")
	}
	v := viper.New()
	rootPath := utils.GetProjectPath(projectName)
	v.AddConfigPath(rootPath) // 设置项目路径
	v.SetConfigType("yaml")   // 配置配置文件类型
	v.SetConfigName(filepath) // 设置配置文件路径
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析配置文件
	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic(err)
	}
	// 动态监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := v.Unmarshal(&global.Config); err != nil {
			panic(err)
		}
		zap.S().Infof("配置文件发生变化:%v", global.Config)
	})

	// 从 Nacos 中获取配置信息
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.GoodsApiDataId,
		Group:  global.NacosConfig.Group,
	})
	if err != nil {
		zap.S().Fatalw("获取 Nacos 配置信息失败", err.Error())
	}
	err = json.Unmarshal([]byte(content), &global.Config)
	if err != nil {
		zap.S().Fatalf("解析 Nacos 配置信息失败:%s", err.Error())
	}

	//Listen config change,key=dataId+group+namespaceId.
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConfig.GoodsApiDataId,
		Group:  global.NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(global.Config.Services.UserOpSrvName)
			err = json.Unmarshal([]byte(data), &global.Config)
			if err != nil {
				zap.S().Fatalw("文件发生变化，解析 Nacos 配置信息失败", err.Error())
			}
			fmt.Println(global.Config.Services.UserOpSrvName)

		},
	})
	if err != nil {
		zap.S().Fatalw("监听配置文件出现异常", err.Error())
	}
}
