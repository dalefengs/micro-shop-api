package initialize

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"micro-shop-api/user-web/global"
	"micro-shop-api/user-web/utils"
)

// 项目名称
var projectName = "user-web"

func InitConfig() {
	filepath := "config%s"
	if utils.IsDev() { // 开发环境
		filepath = fmt.Sprintf(filepath, "-dev")
	} else { // 生产环境
		filepath = fmt.Sprintf(filepath, "")
	}
	v := viper.New()
	getwd, _ := os.Getwd()
	getwd = strings.TrimRight(getwd, "/")
	// 获取到项目根路径 + 项目名称
	rootPath := getwd + "/" + projectName
	if index := strings.LastIndex(getwd, projectName); index != -1 {
		rootPath = getwd[:index] + "/" + projectName
	}
	v.AddConfigPath(rootPath) // 设置项目路径
	v.SetConfigType("yaml")   // 配置配置文件类型
	v.SetConfigName(filepath) // 设置配置文件路径
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析配置文件
	if err := v.Unmarshal(&global.Config); err != nil {
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
}
